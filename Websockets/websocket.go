package websockets

import "github.com/gorilla/websocket"

var Upgraderr = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func WebsocketServer() {
	//init a one time handshake

}
