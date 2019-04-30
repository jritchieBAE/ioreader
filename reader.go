package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

var thing io.Reader

func main() {

	printThing := func() func() {
		myThing := &thing
		return func() {
			byteThing, err := ioutil.ReadAll(*myThing)
			if err != nil {
				fmt.Println(err)
				return
			}
			newReader := strings.NewReader(string(byteThing))
			*myThing = newReader
			fmt.Println(string(byteThing))
		}
	}()

	thing = strings.NewReader("hello")
	printThing()
	printThing()
	thing, _ = os.Open("file.txt")
	printThing()
	printThing()
}
