package pushing

import (
	"context"
	"encoding/json"
	"sync"
	"sync/atomic"
	"time"

	"github.com/cxptek/vdex/models/pg"
	"github.com/cxptek/vdex/service"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

var id int64

type Client struct {
	id       int64
	conn     *websocket.Conn
	writeCh  chan interface{}
	sub      *subscription
	channels map[string]struct{}
	mu       sync.Mutex
}

func NewClient(conn *websocket.Conn, sub *subscription) *Client {
	return &Client{
		id:       atomic.AddInt64(&id, 1),
		conn:     conn,
		writeCh:  make(chan interface{}, 256),
		sub:      sub,
		channels: map[string]struct{}{},
	}
}

func (c *Client) startServe() {
	go c.runReader()
	go c.runWriter()
}

func (c *Client) runReader() {
	c.conn.SetReadLimit(maxMessageSize)
	err := c.conn.SetReadDeadline(time.Now().Add(pongWait))
	if err != nil {
		logrus.Error(err)
		return
	}
	c.conn.SetPongHandler(func(string) error {
		return c.conn.SetReadDeadline(time.Now().Add(pongWait))
	})
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			c.close()
			break
		}

		var req Request
		if err := json.Unmarshal(message, &req); err != nil {
			logrus.Errorf("bad message : %v %v", string(message), err)
			c.close()
			break
		}

		c.onMessage(&req)
	}
}

func (c *Client) runWriter() {
	_, cancel := context.WithCancel(context.Background())
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		cancel()
		ticker.Stop()
		_ = c.conn.Close()
	}()

	for {
		select {
		case message := <-c.writeCh:
			err := c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err != nil {
				_ = c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				c.close()
				return
			}

			buf, err := json.Marshal(message)
			if err != nil {
				continue
			}
			// logrus.Infoln("Write socket to:", c.userID)
			// spew.Dump(message)
			err = c.conn.WriteMessage(websocket.TextMessage, buf)
			if err != nil {
				c.close()
				return
			}

		case <-ticker.C:
			_ = c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			err := c.conn.WriteMessage(websocket.PingMessage, nil)
			if err != nil {
				c.close()
				return
			}
		}
	}
}

func (c *Client) onMessage(req *Request) {
	switch req.Type {
	case "subscribe":
		c.onSub(req.ProductIDs, req.Channels, req.Token)
	case "unsubscribe":
		c.onUnSub(req.ProductIDs, req.Channels, req.Token)
	default:
	}
}

func (c *Client) onSub(productIDs []string, channels []string, token string) {
	user, _ := service.CheckJWT(token)
	var userID int64 = 0
	if user != nil {
		u, _ := pg.SharedStore().GetUserByPublicID(context.Background(), user.PublicID)
		if u != nil {
			userID = u.ID
		}
	}
	logrus.Infoln("userID ", userID, " subscribing")
	// subscribe init channel
	for _, productID := range productIDs {
		for _, channel := range channels {
			switch Channel(channel) {
			case ChannelTrade:
				c.subscribe(ChannelTrade.FormatWithProductID(productID))
			case ChannelOrderbook:
				c.subscribe(ChannelOrderbook.FormatWithProductID(productID))
			case ChannelOrder:
				c.subscribe(ChannelOrder.Format(productID, userID))
			case ChannelOrderbookDepth:
				c.subscribe(ChannelOrderbookDepth.FormatWithClientID(c.id))
			default:
				continue
			}
		}

		asks, bids := service.GetProductDepth(productID, 1000)
		c.sub.publish(ChannelOrderbookDepth.FormatWithClientID(c.id), map[string]interface{}{
			"type": ChannelOrderbookDepth,
			"data": map[string]interface{}{
				"productId": productID,
				"asks":      asks,
				"bids":      bids,
			},
		})
	}
}

func (c *Client) onUnSub(productIDs []string, channels []string, token string) {
	user, _ := service.CheckJWT(token)
	var userID int64 = 0
	if user != nil {
		u, _ := pg.SharedStore().GetUserByPublicID(context.Background(), user.PublicID)
		if u != nil {
			userID = u.ID
		}
	}

	for _, productID := range productIDs {
		for _, channel := range channels {
			switch Channel(channel) {
			case ChannelTrade:
				c.unsubscribe(ChannelTrade.FormatWithProductID(productID))
			case ChannelOrderbook:
				c.unsubscribe(ChannelOrderbook.FormatWithProductID(productID))
			case ChannelOrder:
				c.unsubscribe(ChannelOrder.Format(productID, userID))
			default:
				continue
			}
		}
	}
}

func (c *Client) subscribe(channel string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	_, found := c.channels[channel]
	if found {
		return false
	}

	if c.sub.subscribe(channel, c) {
		c.channels[channel] = struct{}{}
		return true
	}
	return false
}

func (c *Client) unsubscribe(channel string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.sub.unsubscribe(channel, c) {
		delete(c.channels, channel)
	}
}

func (c *Client) close() {
	c.mu.Lock()
	defer c.mu.Unlock()

	for channel := range c.channels {
		c.sub.unsubscribe(channel, c)
	}
}
