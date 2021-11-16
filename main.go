package main

import (
	"fmt"
	"watchgoggy/watchgoggy"
)

func main() {
	pid, err := watchgoggy.GetPid("/java/jdk1.8.0_221/bin/java -Dopennms.home=/omc/config/opennms -jar /omc/omc-application.jar")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(pid)
}
