package server

import (
	"flag"
	"os"
	"time"
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
	app.Get()

}