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
	URL  string
	Name string
}

// GetMeta サイトの情報を返却する。
func GetMeta() Site {
	file, err := ioutil.ReadFile(`./conf/conf.xml`)
	if err != nil {
		panic(err)
	}
	var site Site
	xml.Unmarshal(file, &site)
	fmt.Printf("%+v", site)
	fmt.Printf("\n")

	return site
}
