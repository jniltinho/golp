package dl

import (
	"github.com/jniltinho/golp/curl"
	"log"
	"os"
	"strings"
	"time"
)

func DownloadFile(url string, fileName string, verbose bool) {
	req := curl.New(url)

	req.Method("GET")
	req.SaveToFile(fileName)

	// Print progress status per one second
	req.Progress(func(p curl.ProgressStatus) {
		if verbose {
			log.SetOutput(os.Stdout)
			log.Println(
				"speed", curl.PrettySpeedString(p.Speed),
				"len", curl.PrettySizeString(p.ContentLength),
				"got", curl.PrettySizeString(p.Size),
			)
		}
	}, time.Second)

	req.Do()
}

func DownloadFromUrl(url string, verbose bool) {
	tokens := strings.Split(url, "/")
	fileName := tokens[len(tokens)-1]
	req := curl.New(url)

	req.Method("POST")
	req.SaveToFile(fileName)

	// Print progress status per one second
	req.Progress(func(p curl.ProgressStatus) {
		if verbose {
			log.SetOutput(os.Stdout)
			log.Println(
				"speed", curl.PrettySpeedString(p.Speed),
				"len", curl.PrettySizeString(p.ContentLength),
				"got", curl.PrettySizeString(p.Size),
			)
		}
	}, time.Second)

	req.Do()
}
