package lp

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/jniltinho/golp/ini"
	"log"
	"os"
)

func GetIni(inifile, section, name string) string {
	cfg, err := ini.LoadFile(inifile)
	if err != nil {
		log.Fatal(err)
	}
	getname, ok := cfg.Get(section, name)
	if !ok {
		log.Fatal("app not found in file INI ", inifile)
	}
	return getname
}

func GoQueryGet(url, find1, find2 string) string {
	var fileName string
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(find1).Each(func(i int, s *goquery.Selection) {
		fileName = s.Find(find2).Text()
	})

	return fileName
}

func LogToFile(logfile, msg string, verbose bool) {
	if verbose {
		log.SetOutput(os.Stdout)
		log.Println(msg)
	}
	f, err := os.OpenFile(logfile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Println("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)
	log.Println(msg)
}
