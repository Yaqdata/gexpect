package main

import (
    gexpect "github.com/Yaqdata/gexpect"
    "log"
)

func main() {
	log.Println("Testing Ping interact...")

	child, err := gexpect.Spawn("ssh root@127.0.0.1")
	if err != nil {
		panic(err)
	}
	child.Interact()
	log.Printf("Success\n")
}

