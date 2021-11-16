// +build linux

package watchgoggy

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

//GetPid get pid by name of app, if app exist, return pid; else return error
//ps -ef | grep omc-application.jar | grep -v "grep" | awk '{print $2}
//root      1265     1  2 Nov09 ?        04:14:06 /java/jdk1.8.0_221/bin/java -Dopennms.home=/omc/config/opennms -jar /omc/omc-application.jar
func GetPid(name string) (int, error) {

	cmdFull := exec.Command("bash", "-c", "ps -ef | grep \""+name+"\" | grep -v \"grep\"")
	var stdout1 bytes.Buffer
	cmdFull.Stdout = &stdout1
	cmdFull.Run()

	fmt.Println(cmdFull.String())
	fmt.Println(string(stdout1.Bytes()))
	lineFull := strings.Split(string(stdout1.Bytes()), "\n")
	if cap(lineFull) < 1 {
		return -1, errors.New("name of app not exist")
	}
	for _, s := range lineFull {
		tmps := strings.Fields(s)
		if cap(tmps) < 8 {
			continue
		}
		tmps2 := strings.Split(s, tmps[6])
		if cap(tmps2) == 2 {
			if strings.Trim(tmps2[1], " ") == name {
				ret, err := strconv.Atoi(tmps[1])
				if err == nil {
					return ret, nil
				}
			}
		}
	}
	return -1, errors.New("name of app not exist")
}
