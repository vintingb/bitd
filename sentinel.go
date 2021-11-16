package main

import (
	"archive/zip"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	PRECISEORBIT    = "POEORB"
	RESTITUTEDORBIT = "RESORB"
	DOWNLOADURL     = "http://step.esa.int/auxdata/orbits/Sentinel-1/%s/%s/%s/"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

type sentinel struct {
	dirname      string
	sentinelType string
	startTime    time.Time
	endTime      time.Time
	orbitUrl     string
	orbitName    string
}

func newSentinel(dirname string) (*sentinel, error) {
	s := new(sentinel)
	s.dirname = dirname
	err := s.parse()
	if err != nil {
		return nil, err
	}
	if !s.search(PRECISEORBIT) {
	} else {
		s.search(RESTITUTEDORBIT)
	}
	return s, nil
}

func (s *sentinel) parse() error {
	var err error
	r := strings.Split(s.dirname, "_")
	fmt.Println(s.dirname)
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

func (s *sentinel) search(sType string) bool {
	url := fmt.Sprintf(DOWNLOADURL, sType, s.sentinelType, s.startTime.Format("2006/01"))
	// Request the HTML page.
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find("tr td").Each(func(i int, selection *goquery.Selection) {
		// For each item found, get the title
		title := selection.Find("a").Text()
		if title != "" && title != "Parent Directory" {
			ss := strings.Split(title, "_")
			t1 := ss[6]
			t2 := ss[7]
			sTime, err := time.Parse("V20060102T150405", t1)
			if err != nil {
				log.Fatalln(err)
			}
			eTime, err := time.Parse("20060102T150405.EOF.zip", t2)
			if err != nil {
				log.Fatalln(err)
			}
			if sTime.Before(s.startTime) && eTime.After(s.endTime) {
				s.orbitUrl = url + title
				s.orbitName = title
				return
			}
		}
	})
	if s.orbitName == "" {
		return false
	}
	return false

}

func (s *sentinel) download() error {
	response, err := http.Get(s.orbitUrl)
	file, err := os.Create(s.orbitName)
	if err != nil {
		return err
	}
	res, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	_, err = file.Write(res)
	if err != nil {
		return err
	}
	err = unZip("", s.orbitName)
	if err != nil {
		return err
	}
	err = os.Remove(s.orbitName)
	if err != nil {
		return err
	}
	return nil
}

func unZip(dst, src string) (err error) {
	// 打开压缩文件，这个 zip 包有个方便的 ReadCloser 类型
	// 这个里面有个方便的 OpenReader 函数，可以比 tar 的时候省去一个打开文件的步骤
	zr, err := zip.OpenReader(src)
	defer zr.Close()
	if err != nil {
		log.Fatalf("err-->: %v  file--> %s", err, src)
	}

	// 如果解压后不是放在当前目录就按照保存目录去创建目录
	if dst != "" {
		if err := os.MkdirAll(dst, 0755); err != nil {
			return err
		}
	}

	// 遍历 zr ，将文件写入到磁盘
	for _, file := range zr.File {
		path := filepath.Join(dst, file.Name)

		// 如果是目录，就创建目录
		if file.FileInfo().IsDir() {
			if err := os.MkdirAll(path, file.Mode()); err != nil {
				return err
			}
			// 因为是目录，跳过当前循环，因为后面都是文件的处理
			continue
		}

		// 获取到 Reader
		fr, err := file.Open()
		if err != nil {
			return err
		}

		// 创建要写出的文件对应的 Write
		fw, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_TRUNC, file.Mode())
		fmt.Println(path)
		if err != nil {
			return err
		}

		_, err = io.Copy(fw, fr)
		if err != nil {
			return err
		}

		// 将解压的结果输出
		// 因为是在循环中，无法使用 defer ，直接放在最后
		// 不过这样也有问题，当出现 err 的时候就不会执行这个了，
		// 可以把它单独放在一个函数中，这里是个实验，就这样了
		fw.Close()
		fr.Close()
	}
	return nil
}
