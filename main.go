package main

import (
	"github.com/Amniversary/wedding-logic-info/server"
	"github.com/Amniversary/wedding-logic-info/config"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	server.NewServer(config.NewConfig()).Run()
}



