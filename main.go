package main

import "github.com/tsinghua-cel/attacker-client-go/client"

func main() {
	c, err := client.Dial("http://localhost:8545", 0)
	if err != nil {
		panic(err)
	}
	defer c.Close()
}
