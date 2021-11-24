package watchgoggy

import (
	"fmt"
	"os/exec"
	"time"
)

func Monitor() {
	timer := time.NewTimer(1 * time.Second)
	for {
		<-timer.C
		doMonitor()
		timer.Reset(1 * time.Second)
	}
}

func doMonitor() {
	for name, _ := range Applications {
		appCmd, err := GetAppCmd(name)
		if err != nil {
			fmt.Errorf("cant get exec.cmd of app[%s]\n", name)
		}
		_, err = GetPid(appCmd.String())
		if err != nil {
			fmt.Printf("app[%s] not exist, run it now[%s]...\n", name, appCmd.String())
			go func(appCmd *exec.Cmd, name string) {
				err := RunApp(appCmd, name)
				if err != nil {
					fmt.Errorf("run app[%s] failed\n", name)
				}
			}(appCmd, name)
		}
	}
}
