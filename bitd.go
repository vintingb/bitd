package main

import (
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"strings"
)

var SFiLES []string

func init() {
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
	log.Infoln(SFiLES)
}

func main() {
	for _, SFiLE := range SFiLES {
		s, err := newSentinel(SFiLE)
		if err != nil {
			log.Error(err)
		}
		log.Info("Try to download AUX_POEORB")
		if s.download(AUX_POEORB) {
			log.Info("Downloading AUX_POEORB success")
		} else {
			log.Info("Try to download AUX_RESORB")
			s.download(AUX_RESORB)
			log.Info("Downloading AUX_RESORB success")
		}
	}

}
