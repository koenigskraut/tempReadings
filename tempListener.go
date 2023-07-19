package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"os"
)

func processReadings(data []byte) {
	var temps [3]float32
	if len(data) != 6 {
		log.Println("ERROR")
		return
	}
	for i := 0; i <= 4; i += 2 {
		temp := int16(binary.BigEndian.Uint16(data[i : i+2]))
		temps[i/2] = float32(temp) / 125
	}
	insert, err := db.Query(
		"INSERT INTO temperature (inside, radiator, outside) VALUES (?,?,?)",
		temps[0], temps[1], temps[2],
	)
	if err != nil {
		log.Println(err)
	}
	insert.Close()
}

func runUDPServer() {
	address := fmt.Sprintf(":%s", os.Getenv("UDP_PORT"))
	udpServer, err := net.ListenPacket("udp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer udpServer.Close()

	for {
		buf := make([]byte, 1024)
		n, _, err := udpServer.ReadFrom(buf)
		if err != nil {
			continue
		}
		go processReadings(buf[:n])
	}
}
