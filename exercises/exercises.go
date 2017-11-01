package exercises

import (
	"github.com/shubhodeep9/goal/exercises/helloworld"
	"os"
	"os/user"
	"strconv"
)

var exercises []interface{} = []interface{}{
	helloworld.Output,
}

func ExerciseGo(index int) bool {
	// execute function from the list
	if index >= len(exercises) {
		return false
	}
	exercises[index].(func())()
	return storeCurrent(index)
}

// A big function to store current exercise,
// simple task :P
func storeCurrent(val int) bool {
	// get HOME dir
	usr, err := user.Current()
	if err != nil {
		return false
	}

	// set path as $HOME/.config/goal
	path := usr.HomeDir + "/.config/goal"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, os.ModePerm)
	}

	// write to $HOME/.config/goal/current.txt
	// I know its just an int number being stored
	// but it is important :P
	file, err := os.OpenFile(path+"/current.txt", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		return false
	}
	defer file.Close()
	if _, err = file.Write([]byte(strconv.Itoa(val))); err != nil {
		return false
	}
	// Sync file, I don't know why :P
	if file.Sync() != nil {
		return false
	}
	return true
}

func ExerciseChecker() int {
	return 2
}
