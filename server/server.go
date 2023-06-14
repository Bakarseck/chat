package main

import (
	"fmt"
	"net"
	"bufio"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

const (
	IP = "127.0.0.1" // IP local
	PORT = "3569" // Port Used
)

var f = fmt.Println

func main() {

	f("Lancement du serveur ...")

	// on écoute sur le port 3569
	ln, err := net.Listen("tcp", fmt.Sprintf("%s:%s", IP, PORT))
	checkError(err)

	var clients []net.Conn // tableau de clients
	
	for {

		conn, err := ln.Accept()
		if err == nil {
			clients = append(clients, conn)
		}
		checkError(err)
		fmt.Println("Un client est connecté", conn.RemoteAddr())
		
		go func() {
			buf := bufio.NewReader(conn)

			for {
				name, err := buf.ReadString('\n')

				if err != nil {
					fmt.Printf("Client disconnected.\n")
					break
				}

				for _, c := range clients {
					c.Write([]byte(name))
				}
			}

		}()
	}
}