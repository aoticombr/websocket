package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

func connectAndSend(url string, message string) {
	for {
		dataHora := time.Now()
		fmt.Println("#######################################", dataHora)
		fmt.Println("connectAndSend..", url)
		fmt.Println("$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$")

		ws, _, err := websocket.DefaultDialer.Dial(url, nil)

		if err != nil {
			log.Println(err)
			time.Sleep(1 * time.Second)
			continue
		}
		defer ws.Close()

		// Iniciar uma goroutine para receber dados do servidor em segundo plano
		go func() {
			for {
				messageType, p, err := ws.ReadMessage()
				if err != nil {
					log.Println(err)
					return
				}
				fmt.Println("Received Message Type:", messageType)
				fmt.Println("Received Data:", string(p))
			}
		}()

		err = ws.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			log.Println(err)
			return
		}

		time.Sleep(10 * time.Second)
	}
}

func main() {
	fmt.Println("Cliente Iniciado...")
	go connectAndSend("ws://localhost:3030/route1", "route1")
	go connectAndSend("ws://localhost:3030/route2", "route2")

	select {}
	fmt.Println("Cliente finalizado...")
}
