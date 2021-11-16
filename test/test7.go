package main

import (
	"os"
	"os/exec"
)

func main() {
	c1 := exec.Command("grep", "ERROR", "/var/log/messages")
	c2 := exec.Command("wc", "-l")
	c2.Stdin, _ = c1.StdoutPipe()
	c2.Stdout = os.Stdout
	_ = c2.Start()
	_ = c1.Run()
	_ = c2.Wait()
}
