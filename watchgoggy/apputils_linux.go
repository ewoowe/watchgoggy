// +build linux

package watchgoggy

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

//GetPid get pid by name of app, if app exist, return pid; else return error
//ps -ef | grep omc-application.jar | grep -v "grep" | awk '{print $2}
//root      1265     1  2 Nov09 ?        04:14:06 /java/jdk1.8.0_221/bin/java -Dopennms.home=/omc/config/opennms -jar /omc/omc-application.jar
func GetPid(name string) (int, error) {
	cmdFull := exec.Command("bash", "-c", "ps -ef | grep \""+name+"\" | grep -v \"grep\"")
	var stdout bytes.Buffer
	cmdFull.Stdout = &stdout
	cmdFull.Run()

	fmt.Println(cmdFull.String())
	fmt.Println(string(stdout.Bytes()))
	lineFull := strings.Split(string(stdout.Bytes()), "\n")
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

func IsPidExist(pid int) bool {
	return false
}

func GetAppCmd(name string) (*exec.Cmd, error) {
	app := GetApp(name)
	if app != nil {
		return exec.Command(app.Cmd), nil
	}
	return nil, errors.New("has not app[" + name + "]'s config info")
}

func RunApp(cmd *exec.Cmd, name string) error {
	if cmd == nil {
		return errors.New("cmd cant be nil")
	}
	var out, err = os.OpenFile("/var/log/"+name+".log", os.O_RDWR|os.O_CREATE, 0)
	if err != nil {
		return err
	}
	cmd.Stdout = out
	cmd.Stderr = out
	return cmd.Run()
}
