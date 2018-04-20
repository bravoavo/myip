// Network project main.go
package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	listner, _ := net.Listen("tcp", ":5555")

	for {
		cn, err := listner.Accept()

		if err != nil {
			fmt.Println("EGOG")
			cn.Close()
			continue
		}

		fmt.Println("Connecnted")
		bufReader := bufio.NewReader(cn)
		//wr := bufio.NewWriter(cn)
		var data = "Your IP address is " + cn.RemoteAddr().String()
		cn.Write([]byte(data))
		fmt.Println("Start Reading")
		fmt.Println(data)
		for {
			rbyte, err := bufReader.ReadByte()

			if err != nil {
				fmt.Println("cant read", err)
				break
			}
			fmt.Print(string(rbyte))
		}

	}
}
