package main

import "github.com/ThomasRooney/gexpect"
import "fmt"

func main() {
	fmt.Printf("Starting python.. \n")
	child, err := gexpect.Spawn("python")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Expecting >>>.. \n")
	child.Expect(">>>")

	fmt.Printf("Interacting.. \n")
	child.Interact()
	fmt.Printf("Done \n")
}