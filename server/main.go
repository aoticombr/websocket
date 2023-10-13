package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handleConnection1(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	// Adicione a nova conexão ao mapa de clientes
	clients[conn] = struct{}{}

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			// Remova a conexão do mapa de clientes em caso de erro ou desconexão
			delete(clients, conn)
			return
		}
		fmt.Println("for...", string(p))
		if string(p) == "route1" {
			// Rota 1
			response := []byte("Resposta da Rota 1")
			conn.WriteMessage(messageType, response)
		}
	}
}

func handleConnection2(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	// Adicione a nova conexão ao mapa de clientes
	clients[conn] = struct{}{}

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			// Remova a conexão do mapa de clientes em caso de erro ou desconexão
			delete(clients, conn)
			return
		}
		fmt.Println("for...", string(p))
		if string(p) == "route2" {
			// Rota 2
			response := []byte("Resposta da Rota 2")
			conn.WriteMessage(messageType, response)
		}
	}
}

func sendOLA() {
	for {
		fmt.Println("sendOLA")
		time.Sleep(10 * time.Second)
		// Envie "OLA" para todos os clientes
		for client := range clients {
			fmt.Println("Enviando OLA para um cliente")
			client.WriteMessage(websocket.TextMessage, []byte("OLA"))
		}
	}
}

var clients = make(map[*websocket.Conn]struct{})

func main() {
	http.HandleFunc("/route1", handleConnection1)
	http.HandleFunc("/route2", handleConnection2)

	// Inicie a goroutine para enviar "OLA" a cada 10 segundos
	go sendOLA()

	log.Fatal(http.ListenAndServe(":3030", nil))
}
