package exercises

import (
	"fmt"
	"github.com/shubhodeep9/goal/exercises/helloworld"
	"os"
	"os/user"
	"strconv"
)

func ExerciseGo(index int) bool {
	switch index {
	case 0:
		helloworld.Output()
	}
	return storeCurrent(index)
}

func storeCurrent(val int) error {
	usr, err := user.Current()
	if err != nil {
		return err
	}
	file, err := os.OpenFile(usr.HomeDir+"/.config/goal/current.txt", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	if _, err = file.Write([]byte(strconv.Itoa(val))); err != nil {
		return err
	}
	file.Sync()
	return nil
}

func ExerciseChecker() int {
	return 2
}
