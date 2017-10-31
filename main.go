package main

import (
	"github.com/reechou/wedding-chatroom/config"
	"github.com/reechou/wedding-chatroom/controller"
)

func main() {
	controller.NewLogic(config.NewConfig()).Run()
}
