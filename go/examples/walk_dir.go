package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var depth int
var trackedFileTypes = []string{".go", ".js", ".md", ".py", ".es", ".java", ".cpp", ".c", ".png", ".jpg", ".jpeg"}

func init() {
	flag.IntVar(&depth, "depth", 0, "how deep is your walk...")
	flag.Parse()
}

func main() {
	if depth == 0 {
		walkCurDir("*")
	} else {
		walkRecursive(".")
	}
}

func walkCurDir(dir string) {
	for fi := range genDirFileInfo(dir) {
		if fi.IsDir() {
			printDir(fi)
		} else {
			fmt.Println(fi.Name())
		}
	}
}

func genDirFileInfo(dir string) chan os.FileInfo {
	// files, err := filepath.Glob(dir)
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Print(err)
	}

	fChan := make(chan os.FileInfo)
	// generator anon func
	go func() {
		for _, fi := range files {
			fChan <- fi
		}
		close(fChan)
	}()

	return fChan
}

func walkRecursive(dir string) {
	filepath.Walk(dir, walkProjFunc)
}

func walkProjFunc(fp string, fi os.FileInfo, err error) error {
	if err != nil {
		log.Print(err)
		return nil
	}

	// what are the 2 letter dirs?
	if fi.IsDir() && len(fi.Name()) > 2 {
		printDir(fi)
		return nil
	}

	for _, fType := range trackedFileTypes {
		match, err := filepath.Match("*"+fType, fi.Name())
		if err != nil {
			log.Print(err)
			return err
		}
		if match {
			// fmt.Println(fp)
			_, file := filepath.Split(fp)
			fmt.Println(file)
		}
	}

	return nil
}

func printDir(fi os.FileInfo) {
	fmt.Println()
	fmt.Println(fi.Name())
	fmt.Println("-----------------------")
}
