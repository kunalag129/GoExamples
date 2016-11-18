package main

import (
	"os"
	"io"
	"bytes"
	"fmt"
)

func check(e error) {
	if e!=nil {
		panic(e)
	}
}

func StreamToString(stream io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.String()
}

func main() {
	f, err  := os.Create("file location")
	check(err)
	defer f.Close()

	list := []string{""}

	fmt.Println(len(list))

	for _,p:= range list {
		r := apicall.GetProductDetails(p)
		f.WriteString(StreamToString(r))
	}
}

