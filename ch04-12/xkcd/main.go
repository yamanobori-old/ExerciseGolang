package main

import (
	"encoding/json"
	"log"
	"os"
	"regexp"
	"strconv"

	"./format"
	"./index"
)

const indexdir = "info"

func init() {
	if _, err := os.Stat(indexdir); os.IsNotExist(err) {
		if err := os.Mkdir(indexdir, 0777); err != nil {
			log.Fatal(err)
		}
		index.CreateIndex(indexdir)
	}
}

// Info コミックのデータ
type Info struct {
	Month      string
	Num        int
	Link       string
	Year       string
	News       string
	SafeTitle  string `json:"safe_title"`
	Transcript string
	Alt        string
	Img        string
	Title      string
	Day        string
}

func main() {
	if len(os.Args) > 1 {
		keyword := os.Args[1]
		for i := 1; i < index.MaxIndex; i++ {
			indexFile := indexdir + "/" + strconv.Itoa(i)
			var (
				f   *os.File
				err error
			)
			if f, err = os.Open(indexFile); err != nil {
				continue
			}
			var info Info
			json.NewDecoder(f).Decode(&info)
			if match, _ := regexp.MatchString(keyword, info.Transcript); match {
				if err := format.Print(index.URLPrefix+strconv.Itoa(i), info.Transcript); err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}
