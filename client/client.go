package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
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

	// Connexion au serveur
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", IP, PORT))
	checkError(err)

	for {
		// entrée utilisateur
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("client: ")
		text, err := reader.ReadString('\n')
		checkError(err)

		conn.Write([]byte(text))

		// On écoute tous les messages émis par le serveur et on rajoute un retour à la ligne
		message, err := bufio.NewReader(conn).ReadString('\n')

		checkError(err)

		f("serveur : " + message)
	}
}