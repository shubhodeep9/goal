package exercises

import (
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

func storeCurrent(val int) bool {
	usr, err := user.Current()
	if err != nil {
		return false
	}
	path := usr.HomeDir + "/.config/goal"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, os.ModePerm)
	}
	file, err := os.OpenFile(path+"/current.txt", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		return false
	}
	defer file.Close()
	if _, err = file.Write([]byte(strconv.Itoa(val))); err != nil {
		return false
	}
	file.Sync()
	return true
}

func ExerciseChecker() int {
	return 2
}
