package handler

import (
	"github.com/gofiber/contrib/websocket"
	"log"
)

type Client struct {
	UUID   string
	Conn   *websocket.Conn
	UserId int32
}

type Hub struct {
	Clients map[string]*Client
}

// Websockets godoc
// @Summary      Websockets for real-time communication
// @Description  Websockets connection to receive messages in real-time
// @Tags         Websockets
// @Accept       json
// @Produce      json
// @Param		 Authorization header string true "Bearer token"
// @Success      200  {object}  response.WebsocketMessage
// @Failure      401  {object}  response.Error
// @Failure      500  {object}  response.Error
// @Router       /ws [get]
func Websockets(c *websocket.Conn) {
	for {
		mt, msg, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", msg)

		if err = c.WriteMessage(mt, msg); err != nil {
			log.Println("write:", err)
			break
		}
	}
}
