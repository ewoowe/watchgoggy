//+build windows

package watchgoggy

import (
	"errors"
	"os/exec"
)

func GetPid(name string) (int, error) {
	return -1, errors.New("name of app not exist")
}

func IsPidExist(pid int) bool {
	return false
}

func GetAppCmd(name string) (*exec.Cmd, error) {
	return nil, errors.New("operation dont support")
}

func RunApp(cmd *exec.Cmd, name string) error {
	return errors.New("operation dont support")
}
