package helloworld

import (
	"fmt"
	"github.com/philgebhardt/mdcat"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func Output() {
	gopath := os.Getenv("GOPATH")
	problem, err := ioutil.ReadFile(gopath + "/src/github.com/shubhodeep9/goal/exercises/helloworld/problem.md")
	if err != nil {
		log.Fatal(err)
	}
	markdown_string := fmt.Sprintf(string(problem))
	mdcat.Print(strings.NewReader(markdown_string))
}
