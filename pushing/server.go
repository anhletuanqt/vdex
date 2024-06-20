package pushing

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type Server struct {
	addr string
	path string
	sub  *subscription
}

func NewServer(addr, path string, sub *subscription) *Server {
	return &Server{
		addr: addr,
		path: path,
		sub:  sub,
	}
}

func (s *Server) ws(c *gin.Context) {
	upGrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	conn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logrus.Errorln(err)
		return
	}

	NewClient(conn, s.sub).startServe()
}

func (s *Server) Run() {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard

	r := gin.Default()
	r.GET(s.path, s.ws)
	err := r.Run(s.addr)
	if err != nil {
		panic(err)
	}
}
