package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"

	log "github.com/sirupsen/logrus"
)

var (
	SFiLES orbit
	tmpS   strings.Builder
)

func init() {
	tmpS.WriteString("all.bash ")
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	files, err := ioutil.ReadDir(pwd)
	if err != nil {
		log.Fatalln(err)
		return
	}
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".SAFE") {
			SFiLES = append(SFiLES, strings.Split(file.Name(), ".SAFE")[0])
		}
		if strings.HasSuffix(file.Name(), ".zip") {
			SFiLES = append(SFiLES, strings.Split(file.Name(), ".zip")[0])
		}
	}
	sort.Sort(SFiLES)
	log.Infoln(SFiLES)
}

func main() {
	for _, SFiLE := range SFiLES {
		s, err := NewSentinel(SFiLE)
		if err != nil {
			log.Error(err)
		}
		log.Info("Try to Download AUX_POEORB")
		if s.Download(AUX_POEORB) {
			log.Info("Downloading AUX_POEORB success")
		} else {
			log.Info("Try to Download AUX_RESORB")
			if s.Download(AUX_RESORB) {
				log.Info("Downloading AUX_RESORB success")
			} else {
				log.Info("Downloading err")
			}
		}
	}
	fmt.Println(tmpS.String())
}
