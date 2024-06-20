package pushing

import "fmt"

type Channel string

func (t Channel) Format(productID string, userID int64) string {
	return fmt.Sprintf("%v:%v:%v", t, productID, userID)
}

func (t Channel) FormatWithUserID(userID int64) string {
	return fmt.Sprintf("%v:%v", t, userID)
}

func (t Channel) FormatWithProductID(productID string) string {
	return fmt.Sprintf("%v:%v", t, productID)
}
func (t Channel) FormatWithClientID(clientID int64) string {
	return fmt.Sprintf("%v:%v", t, clientID)
}

const (
	ChannelOrderbook      = Channel("orderbook")
	ChannelTrade          = Channel("trade")
	ChannelOrder          = Channel("order")
	ChannelOrderbookDepth = Channel("orderbook_depth")
)

type Request struct {
	Type       string   `json:"type"`
	ProductIDs []string `json:"productIds"`
	Channels   []string `json:"channels"`
	Token      string   `json:"token"`
}

type Response struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}
