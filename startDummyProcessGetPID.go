package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("./dummyprocess")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d", cmd.Process.Pid)
}
