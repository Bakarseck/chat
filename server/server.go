package main

import (
	"fmt"
	"net"
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

	//On accepte les connexions entrantes sur le port 3569
	conn, err := ln.Accept()
	checkError(err)

	f("Un client est connecté depuis", conn.RemoteAddr())
	
	for {

		// On ecoute les messages émis par les connexions entrantes (ctrl-c pour quitter)

		buffer := make([]byte, 4096) //Taille maximum du message qui sera envoyé par le client

		length, err := conn.Read(buffer) // lire le message envoyé par le client

		message := string(buffer[:length]) // supprimer les bits qui servent à rien et convertir les bytes en sring

		if err != nil {
			f("Le client s'est connecté")
		}

		// On affiche le message du client en le convertissant de byte à string
		fmt.Print("Client:", message)

		// On envoie le message au client pour qu'il l'affiche
		conn.Write([]byte(message + "\n"))
	}
}