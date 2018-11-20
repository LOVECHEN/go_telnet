package main

import (
	"bufio"
	"fmt"
	"github.com/lixiangyun/go_telnet/telnet"
	"os"
)

func UserOutput(recv []byte) {

TRY:

	begin, end := -1, -1
	for i, v := range recv {

		switch v {
		case 27:
			begin = i
		case 'm':
			end = i
		case 7:
			recv[i] = ' '
		}

		if begin != -1 && end != -1 && begin < end {

			if begin == 0 {
				recv = recv[end+1:]
			} else if end+1 >= len(recv) {
				recv = recv[0:begin]
			} else {
				recv = append(recv[0:begin], recv[end+1:]...)
			}

			//fmt.Println("len:", len(recv), recv)
			//fmt.Println(begin, end)

			goto TRY
		}
	}

	fmt.Print(string(recv))
}

func UserInput(c *telnet.Client) {
	input := bufio.NewReader(os.Stdin)

	for {
		in, err := input.ReadBytes('\r')
		//fmt.Println("Input:", in)
		if err != nil {
			fmt.Println(err.Error())
		}

		c.Write(in)
	}
}

func main() {

	args := os.Args

	if len(args) < 3 {
		fmt.Println("Usage: <IP> <PORT>")
		return
	}

	IP := args[1]
	PORT := args[2]

	c := telnet.NewClient(IP, PORT)

	err := c.Connect(UserOutput)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	go UserInput(c)

	c.Process()

	defer c.Delete()
}
