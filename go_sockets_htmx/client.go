package main

import (
"net/http"
"github.com/gorilla/websocket"
)

type Client struct {
	hub *Hub
	conn *websocket.Conn
	send chan []byte

func serveWs(hub *Hub, w http.ResponseWriter,r *http.Request) {
}

