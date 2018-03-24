package conf

import (
	"encoding/xml"
	"io/ioutil"

	"github.com/ArakiTakaki/Context"
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
	Key  int
}

// GetMeta サイトの情報を返却する。
func GetMeta() Site {
	data, err := context.Get("meta")
	if err != nil {
		return ResetMeta()
	}
	return data.(Site)
}

// ResetMeta メタデータのセットし直し
func ResetMeta() Site {
	file, err := ioutil.ReadFile(`./conf/conf.xml`)
	if err != nil {
		panic(err)
	}
	var site Site
	xml.Unmarshal(file, &site)
	context.Put("meta", site)
	return site
}
