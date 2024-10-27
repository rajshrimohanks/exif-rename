package main

import (
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/jessevdk/go-flags"
	"github.com/rwcarlsen/goexif/exif"
)

func main() {

	var opts struct {
		// Supply the directory with pictures as an argument to the program.
		// e.g. go run main.go -d /mnt/f/zfest-2024/Unedited\ Photos/
		Directory string `short:"d" long:"dir" description:"The directory path which contains the pictures." required:"true"`
		DryRun    bool   `long:"dry-run" description:"Performs a dry run without actually executing the changes."`
	}
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal("Exiting...")
	}

	// Grab the files within the directory
	dirEntries, err := os.ReadDir(opts.Directory)
	if err != nil {
		log.Fatal(err)
	}

	if opts.DryRun {
		log.Print("This is a dry run. No changes will be performed.")
	}
	// Iterate through the files in the directory
	for idx, de := range dirEntries {
		// We want the full path of the image. So we join the directory path with the file name.
		fullFileName := filepath.Join(opts.Directory, de.Name())

		// Open the file.
		file, err := os.Open(fullFileName)
		if err != nil {
			log.Fatal(err)
		}

		// Decode EXIF data from the file.
		exif, err := exif.Decode(file)
		if err != nil {
			log.Fatal(err)
		}

		// Grab the time when the image was taken.
		taken, err := exif.DateTime()
		if err != nil {
			log.Fatal(err)
		}
		// Format the taken time to a filename friendly string.
		dateTime := taken.Format("2006-01-02_15-04-05")

		//Our original file name can be any of the following patterns:
		// AVK00001.jpg
		// AVK00002 (2).jpg
		// AVK00002 (2) (1).jpg
		// We want only the AVKxxxxx.jpg portion and nothing else.
		// So we create a RegEx to split the file name to groups.
		re := regexp.MustCompile(`((AVK|DSC)[\d]+)(\s*)(\([\d]\)[\s]*)*.jpg`)
		// FindStringSubmatch() gets us the captured groups. In our case, we want the first
		// group which has the core name.
		// We append this to the created date time and join this to the dir name to get the
		// new file path.
		newFileName := filepath.Join(opts.Directory, dateTime+"_"+re.FindStringSubmatch(de.Name())[1]+".jpg")

		// Finally we call os.Rename to rename the file.
		if !opts.DryRun {
			err = os.Rename(fullFileName, newFileName)
			if err != nil {
				log.Fatal(err)
			}
		}
		log.Printf("%5v of %5v: %60v -> %60v", (idx + 1), len(dirEntries), fullFileName, newFileName)
	}
}
