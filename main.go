package main

import (
	"fmt"
	"github.com/Abdulsametileri/ingilizce-kelime-go/config"
	"github.com/Abdulsametileri/ingilizce-kelime-go/services"
)

func main() {
	fmt.Println("Bismillah")

	config.Setup()

	services.Setup(&services.Telegram{})
}
