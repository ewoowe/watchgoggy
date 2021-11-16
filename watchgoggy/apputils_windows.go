//+build windows

package watchgoggy

import "errors"

func GetPid(name string) (int, error) {
	return -1, errors.New("name of app not exist")
}
