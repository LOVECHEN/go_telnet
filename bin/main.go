package main

import (
	"bufio"
	"fmt"
	"go_telnet/telnet"
	"os"
)

func show(recv string) {
	fmt.Println(recv)
}

func input(c *telnet.Client) {
	inputReader := bufio.NewReader(os.Stdin)

	for {
		in, err := inputReader.ReadString('\n')

		fmt.Println("Input:", in)

		if err != nil {
			fmt.Println(err.Error())
		}

		c.Write(in)
	}
}

func main() {

	c := telnet.NewClient("192.168.0.107", "23")

	err := c.Connect(show)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	go input(c)

	c.Process()

	defer c.Delete()
}
