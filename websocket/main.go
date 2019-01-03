package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"strconv"
)

var clients = make(map[*websocket.Conn]bool) // connected clients
var broadcast = make(chan Message)           // broadcast channel

// 定义升级
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
//参考: https://github.com/scotch-io/go-realtime-chat
//https://scotch.io/bar-talk/build-a-realtime-chat-server-with-go-and-websockets
//websocket类库 https://github.com/gorilla/websocket

// 定义消息结构体
type Message struct {
	Data_type    string `json:"data_type"`
	Data_msg string `json:"data_msg"`
}

func main() {
	//起http服务
	log.Println("http server started on 127.0.0.1:7777")
	http.HandleFunc("/ssh", handleConnections)
	http.ListenAndServe("127.0.0.1:7777", nil)

}

func handleConnections(w http.ResponseWriter, r *http.Request) {

	log.Println("handle---connnect....")
	// http请求升级为websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Make sure we close the connection when the function returns
	defer ws.Close()

	//起广播协程
	go handleMessages()

	// Register our new client
	clients[ws] = true

	//首次读取上报数据
	for {
		var msg Message
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}
		// Send the newly received message to the broadcast channel
		broadcast <- msg
	}
}

func handleMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast

		var respMsg Message
		if msg.Data_type == "join" {
			respMsg = Message{"join-resp", strconv.Itoa(len(clients))}
		} else if msg.Data_type == "input" {
			respMsg = Message{"input-resp", msg.Data_msg}
		} else {
			respMsg = Message{"other", "1"}
		}
		// 将数据广播到每个客户端
		for client := range clients {

			err := client.WriteJSON(respMsg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}