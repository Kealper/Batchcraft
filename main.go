/*
Copyright 2016 Ken Piper

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"net"
	"os"
)

// Define RCON packet types
const (
	Response = iota
	_
	Command
	Login
)

// Structure for RCON requests and responses
type Packet struct {
	RequestID int
	Type      int
	Payload   []byte
}

// Log in to the specified RCON server and send the provided command, and return the server's response
// Many commands have no response, so response may be blank
func sendCommand(addr, pass, data string) (string, error) {
	raddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return "", fmt.Errorf("Failed to resolve hostname.")
	}
	c, err := net.DialTCP("tcp", nil, raddr)
	if err != nil {
		return "", fmt.Errorf("Failed to connect to server. Is the server running with RCON configured correctly?")
	}
	defer c.Close()

	// Attempt login
	c.Write(buildPacket(1, Login, []byte(pass)))
	buf := make([]byte, 2048)
	c.Read(buf)
	resp := readPacket(buf)
	if resp.RequestID != 1 { // Bad RCON password
		return "", fmt.Errorf("Failed to log in. Bad RCON password?")
	}

	// Send command if successful, and return the result as a string
	c.Write(buildPacket(2, Command, []byte(data)))
	buf = make([]byte, 2048)
	c.Read(buf)
	resp = readPacket(buf)
	if resp.RequestID != 2 {
		return "", fmt.Errorf("Response from server could not be verified.")
	}
	return string(resp.Payload), nil
}

// Formats data into a ready-to-send Minecraft RCON request
func buildPacket(id, t int, data []byte) []byte {
	buf := bytes.NewBuffer([]byte(""))
	binary.Write(buf, binary.LittleEndian, uint32(len(data) + 10))
	binary.Write(buf, binary.LittleEndian, uint32(id))
	binary.Write(buf, binary.LittleEndian, uint32(t))
	buf.Write(data)
	buf.Write([]byte{0, 0})
	return buf.Bytes()
}

// Parses out key information from a Minecraft RCON response
func readPacket(data []byte) Packet {
	buf := bytes.NewBuffer(data)
	var l uint32
	var id uint32
	var t uint32
	binary.Read(buf, binary.LittleEndian, &l)
	binary.Read(buf, binary.LittleEndian, &id)
	binary.Read(buf, binary.LittleEndian, &t)
	payload := make([]byte, l - 10)
	buf.Read(payload)
	resp := Packet{
		RequestID: int(id),
		Type: int(t),
		Payload: payload,
	}
	return resp
}

func main() {
	addr := flag.String("a", "127.0.0.1:25575", "Address for the server's RCON instance, in host:port format.")
	pass := flag.String("p", "", "Password for authenticating with the server. This is required.")
	cmd := flag.String("c", "", "Command string to send to the server. This should be wrapped in quotes if the command is longer than one word.")
	flag.Parse()
	resp, err := sendCommand(*addr, *pass, *cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1) // If there was an error, exit with a non-zero status to signal a problem
	}
	if resp != "" {
		fmt.Print(resp)
	}
}
