package conf

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

// Site Pageを取得
type Site struct {
	SiteTitle string
	Caption   string
	MaxPage   int
	Page      []Page
}

// Page URLと名前のマップ
type Page struct {
	Url  string
	Name string
}

func GetMeta() {
	const BUFSIZE = 1024
	file, err := ioutil.ReadFile(`./conf/conf.xml`)
	Error(err, "ファイルのオープンエラー:context.go")
	var pages Page
	xml.Unmarshal(file, &pages)
	fmt.Printf("%+v", pages)
}

func Error(err error, errCode string) {
	if err != nil {
		fmt.Println(errCode)
		panic(err)
	}
}
