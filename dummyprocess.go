package main

import (
	"log"
	"os/exec"
	"time"
)

func main() {
	numChildren := 5

	for i := 0; i < numChildren; i++ {
		time.Sleep(time.Second)
		cmd := exec.Command("./dummychild")
		err := cmd.Start()
		if err != nil {
			log.Fatal(err)
		}
	}
}
