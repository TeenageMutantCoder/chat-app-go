package main

import (
	"errors"
	"net"
)

type Command interface {
	execute(conn net.Conn, msg []byte) error
}

type SendCommand struct{}

func (c SendCommand) execute(conn net.Conn, msg []byte) error {
	n, err := conn.Write(msg)

	if n < len(msg) {
		return errors.New(incompleteMessageError)
	}

	return err
}

type InvalidCommand struct{}

func (c InvalidCommand) execute(conn net.Conn, msg []byte) error {
	msg = []byte("Invalid command. Please try again.\n")
	n, err := conn.Write(msg)

	if n < len(msg) {
		return errors.New(incompleteMessageError)
	}

	return err
}

type CloseCommand struct{}

func (c CloseCommand) execute(conn net.Conn, msg []byte) error {
	msg = []byte("Closing connection...\n")
	n, err := conn.Write(msg)
	if n < len(msg) {
		return errors.New(incompleteMessageError)
	}
	if err != nil {
		return err
	}

	err = conn.Close()
	return err
}
