package main

import (
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("ping", "127.0.0.1")
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}
