package main

import (
	"bytes"
	"log"
	"net"
)

const (
	incompleteMessageError = "Failed to send entire payload"
)

func main() {
	network := "tcp"
	port := ":8080"
	ln, err := net.Listen(network, port)

	if err != nil {
		log.Fatal("Failed to open", network, "connection at port", port)
	}

	log.Print("Listening on port", port)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal("Failed to accept connection")
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	msg := []byte("Successfully connected.\n")
	n, err := conn.Write(msg)

	if err != nil {
		log.Fatal("Failed to write to connection")
	}
	if n < len(msg) {
		log.Fatal(incompleteMessageError)
	}

	maximumMessageLength := 256
	bytes := make([]byte, maximumMessageLength)
	for {
		if conn == nil {
			log.Print("Connection is nil")
			break
		}

		n, err = conn.Read(bytes)

		if n > maximumMessageLength {
			log.Fatal("Message too long")
		}
		if err != nil {
			log.Fatal("Failed to read from connection")
		}

		msg = bytes[:n]
		log.Print("Received message: ", string(msg))

		cmd, msgWithoutCmd := parseMessage(msg)
		err = cmd.execute(conn, msgWithoutCmd)

		if err != nil {
			log.Fatal("Failed to execute command", cmd, msg)
			conn.Close()
		}
	}
}

func parseMessage(msg []byte) (Command, []byte) {
	indexOfFirstSpace := bytes.Index(msg, []byte(" "))

	var possibleCommand string
	if indexOfFirstSpace == -1 {
		possibleCommand = string(bytes.TrimRight(msg, "\n"))
	} else {
		possibleCommand = string(msg[:indexOfFirstSpace])
	}

	var remainingMessage []byte
	if indexOfFirstSpace == -1 {
		remainingMessage = []byte{}
	} else {
		remainingMessage = bytes.TrimLeft(msg[indexOfFirstSpace:], " ")
	}

	switch possibleCommand {
	case "send":
		return SendCommand{}, remainingMessage
	case "close":
		return CloseCommand{}, remainingMessage
	default:
		return InvalidCommand{}, remainingMessage
	}
}
