package main

import (
	"os"
	"os/exec"
)

func main() {
	c1 := exec.Command("netstat", "-aon")
	c2 := exec.Command("findstr", "1080")
	c2.Stdin, _ = c1.StdoutPipe()
	c2.Stdout = os.Stdout
	_ = c2.Start()
	_ = c1.Run()
	_ = c2.Wait()
}
