package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {

	cmdSecSMS := "window.location=\"http:/www.tf1.fr/\""

	fmt.Println(cmdSecSMS)
	hostName := "localhost"
	portNum := "32000"

	doDial(cmdSecSMS, hostName, portNum)

}

func doDial(cmd, host, port string) {
	// connect to this socket
	conn, err := net.Dial("tcp", host+":"+port)

	if err != nil {
		fmt.Printf("Some error %v", err)
		return
	}

	defer conn.Close()
	fmt.Printf("Connection established between %s and localhost.\n", host)
	fmt.Printf("Local Address : %s \n", conn.LocalAddr().String())
	fmt.Printf("Remote Address : %s \n", conn.RemoteAddr().String())

	// send to socket
	fmt.Fprintf(conn, cmd+"\n")
	// listen for reply
	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Print("Message from server: " + message)

}
