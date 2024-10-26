package main

import (
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/rwcarlsen/goexif/exif"
)

func main() {
	dir := os.Args[1] // /mnt/f/zfest-2024/Edited\ Photos/

	log.Printf("selected directory: %v", dir)

	dirEntries, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for idx, de := range dirEntries {
		fullFileName := filepath.Join(dir, de.Name())
		file, err := os.Open(fullFileName)
		if err != nil {
			log.Fatal(err)
		}

		exif, err := exif.Decode(file)
		if err != nil {
			log.Fatal(err)
		}

		taken, err := exif.DateTime()
		if err != nil {
			log.Fatal(err)
		}

		re := regexp.MustCompile(`(AVK[\d]+)(\s*)(\([\d]\)[\s]*)*.jpg`)
		dateTime := taken.Format("2006-01-02_15-04-05")
		newFileName := dateTime + "_" + re.FindStringSubmatch(de.Name())[1] + ".jpg"
		log.Printf("%5v of %5v: %20v -> %35v", idx, len(dirEntries), de.Name(), newFileName)
	}
}
