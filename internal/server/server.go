package server

import (
	"flag"
	"os"
	"time"

	"github.com/gerceboss/VChat/internal/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
	"github.com/gofiber/websocket/v2"
)
var (
	addr=flag.String("addr",":",os.Getenv("PORT"))
	cert=flag.String("cert","","")
	key=flag.String("key","","")
)
func Run()error{
	flag.Parse()
	if *addr==":"{
		*addr=":8080"
	}
	engine:=html.New("./../views",".html")
	app:=fiber.New(fiber.Config{Views : engine})
	app.Use(logger.New())
	app.Use(cors.New())

	app.Get("/",handlers.Welcome)
	app.Get("/room/create",handlers.CreateRoom)
	app.Get("/room/:uuid",handlers.Room)
	app.Get("/room/:uuid/chat",handlers.RoomChat)
	app.Get("/room/:uuid/websocket",websocket.New(handlers.RoomWebsocket,websocket.Config{
		HandshakeTimeout: 10*time.Second,
	}))
	app.Get("/room/:uuid/chat/websocket",websocket.New(handlers.RoomChatWebsocket))
	app.Get("/room/:uuid/viewer/websocket",websocket.New(handlers.RoomViewerWebsocket))
	app.Get("stream/:ssuid",handlers.Stream)
	app.Get("stream/:ssuid/websocket",)
	app.Get("stream/:ssuid/chat/websocket",websocket.New(handlers.StreamChatWebsocket))
	app.Get("stream/:ssuid/viewer/websocket",websocket.New(handlers.StreamViewerWebsocket))


}