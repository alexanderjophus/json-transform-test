package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

const inDir = "/pfs/feed-raw/"
const outDir = "/pfs/out/"

func main() {
	os.Mkdir(outDir, 0700)

	files, err := ioutil.ReadDir(inDir)
	if err != nil {
		log.Fatalf("read dir: %s", err)
	}
	log.Println("Read directory")

	for _, f := range files {
		log.Printf("File found: %s\n", f.Name())
		_, err := copyFile(inDir+f.Name(), outDir+f.Name())
		if err != nil {
			log.Fatal(err)
		}
	}
}

func copyFile(src, dst string) (int64, error) {
	log.Printf("Copying %s to %s\n", src, dst)
	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
