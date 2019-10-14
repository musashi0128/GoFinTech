package main

import (
	"GoFinTech/config"
	"GoFinTech/utils"
	"log"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	log.Println("test")
}
