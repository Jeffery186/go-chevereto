package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
	"go-chevereto/conf"
	"golang.design/x/clipboard"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
)

func main() {
	var sFlag = flag.String("f", "", "待上传的图片的路径")
	flag.Parse()
	err0 := clipboard.Init()
	if err0 != nil {
		log.Println(err0)
	}
	var api conf.ApiConf
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Println(err)
	}
	api = conf.CheveretoConf(path.Join(dir, "conf.ini"))
	client := resty.New()
	profileImgBytes, _ := ioutil.ReadFile(*sFlag)
	_, file := path.Split(*sFlag)
	resp, err1 := client.R().
		SetFileReader("source", file, bytes.NewReader(profileImgBytes)).
		SetFormData(map[string]string{
			"key":    api.Key,
			"action": "upload",
			"format": "json",
		}).
		Post(api.Url)
	if err1 != nil {
		log.Println(err1)
	}
	url := gjson.Get(resp.String(), "image.url").String()
	//name := gjson.Get(resp.String(), "image.name").String()
	//extension := gjson.Get(resp.String(), "image.extension").String()
	fmt.Println(url)
	//println(fmt.Sprintf("%s.%s", name, extension))
	clipboard.Write(clipboard.FmtText, []byte(url))
}
