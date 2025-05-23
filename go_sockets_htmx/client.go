package main

import (
"net/http"
"github.com/gorilla/websocket"
)

type Message struct {
	ClientID string
	Text string
}

type Client struct {
	id string
	hub *Hub
	conn *websocket.Conn
	send chan []byte

var upgrader = websocket.Upgrader{
	ReadBufferSize:1024,
	WriteBufferSize:1024,

func serveWs(hub *Hub, w http.ResponseWriter,r *http.Request) {
	// upgrade the conn
	// create client
	// listening the hbu
	conn, err := upgrader.Upgrade(w,r,nil)
	if err != nil {
		log.Println(err)
		return 
	}
	client := &Client{hub:hub,conn:conn, send:make(chan[]byte)}



}

