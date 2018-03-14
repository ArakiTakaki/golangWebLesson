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
	var site Site
	xml.Unmarshal(file, &site)
	fmt.Printf("%+v", site)
	fmt.Printf("\n")
}

func Error(err error, errCode string) {
	if err != nil {
		fmt.Println(errCode)
		panic(err)
	}
}
