package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var SAFEFiLES []string

func init() {
	log.SetFlags(log.Lshortfile)
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	files, err := ioutil.ReadDir(pwd)
	if err != nil {
		return
	}
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".SAFE") {
			SAFEFiLES = append(SAFEFiLES, file.Name())
		}
	}
}
func main() {
	for _, SAFEFiLE := range SAFEFiLES {
		s, err := newSentinel(SAFEFiLE)
		if err != nil {
			log.Println(err)
		}
		err = s.download()
		if err != nil {
			log.Fatalln(err)
		}
	}
}
