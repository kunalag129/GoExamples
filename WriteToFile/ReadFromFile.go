package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	content, e  := ioutil.ReadFile("file location")
	if e!=nil {
		panic(e)
	}

	lines := strings.Split(string(content), "\n")

	fmt.Println(len(lines))

}
