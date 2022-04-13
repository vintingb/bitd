package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/imroc/req"
	log "github.com/sirupsen/logrus"
)

const (
	sciHubUrl            = "https://scihub.copernicus.eu/gnss/odata/v1/Products"
	auth                 = "Basic Z25zc2d1ZXN0Omduc3NndWVzdA=="
	AUX_POEORB orbitType = "AUX_POEORB"
	AUX_RESORB orbitType = "AUX_RESORB"
)

var filterString = "startswith(Name,'%s') and substringof('%s',Name) and ContentDate/Start lt datetime'%s' and ContentDate/End gt datetime'%s'"

type orbitType string

type orbit []string

func (o orbit) Len() int { return len(o) }

func (o orbit) Swap(i, j int) { o[i], o[j] = o[j], o[i] }
func (o orbit) Less(i, j int) bool {
	x, err := NewSentinel(o[i])
	if err != nil {
		log.Error(err)
	}
	y, err := NewSentinel(o[j])
	if err != nil {
		log.Error(err)
	}
	return x.startTime.Before(y.startTime)
}

type sentinel struct {
	dirname      string
	sentinelType string
	startTime    time.Time
	endTime      time.Time
	orbitUrl     string
	orbitName    string
}

func NewSentinel(dirname string) (*sentinel, error) {
	s := new(sentinel)
	s.dirname = dirname
	err := s.parse()
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (s *sentinel) parse() error {
	var err error
	r := strings.Split(s.dirname, "_")
	s.sentinelType = r[0]
	t1 := r[5]
	t2 := r[6]
	s.startTime, err = time.Parse("20060102T150405", t1)
	if err != nil {
		return err
	}
	s.endTime, err = time.Parse("20060102T150405", t2)
	if err != nil {
		return err
	}
	return nil
}

func (s *sentinel) Download(ot orbitType) bool {
	param := req.Param{
		"$top":     1,
		"$orderby": "ContentDate/Start asc",
		"$filter":  fmt.Sprintf(filterString, s.sentinelType, ot, s.startTime.Format(`2006-01-02T15:04:05`), s.endTime.Format(`2006-01-02T15:04:05`)),
	}
	header := req.Header{
		"Authorization": auth,
	}
	r, err := req.Get(sciHubUrl, header, param)
	if err != nil {
		log.Fatal(err)
	}
	st := new(success)
	err = r.ToXML(st)
	if err != nil {
		return false
	}
	if st.Entry.Title.Text == "" || st.Entry.ID.Text == "" {
		return false
	}
	s.orbitName = st.Entry.Title.Text + ".EOF"
	s.orbitUrl = st.Entry.ID.Text + "/$value"
	tmpS.WriteString(s.dirname + ".SAFE ")
	tmpS.WriteString(st.Entry.Title.Text + ".EOF ")
	if _, err := os.Stat(s.orbitName); os.IsNotExist(err) {
		log.Infoln(s.orbitName)
		r, err = req.Get(s.orbitUrl, header)
		if err != nil {
			log.Error(err)
			return false
		}
		err = r.ToFile(s.orbitName)
		if err != nil {
			log.Error(err)
			return false
		}
	} else {
		log.Infoln(s.orbitName + " exist")
	}
	return true

}
