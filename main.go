package main

import (
	"FloatingBooks/db"
	"FloatingBooks/router"
	"flag"
)

func main () {
	port := flag.String("port", "8080", "监听端口")
	flag.Parse()
	db.DBInit()
	r := router.RouterInit()
	_ = r.Run(":" + *port)
}