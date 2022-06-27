package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	val := "E:\\playlista"
	output := "C:\\Users\\pcoles\\Downloads\\out"
	ensureDir(output)
	packFiles(val, output)

}

func packFiles(name string, output string) {
	files, err := ioutil.ReadDir(name)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		p := name + "/" + file.Name()
		fmt.Println(name+"/"+file.Name(), file.IsDir())
		if file.IsDir() {
			packFiles(p, output)
		} else {
			o := output + "/" + file.Name()
			copyFile(p, o)
		}
	}
}

// https://stackoverflow.com/questions/37932551/mkdir-if-not-exists-using-golang
func ensureDir(dirName string) error {
	err := os.Mkdir(dirName, os.ModeDir)
	if err == nil {
		return nil
	}
	if os.IsExist(err) {
		// check that the existing path is a directory
		info, err := os.Stat(dirName)
		if err != nil {
			return err
		}
		if !info.IsDir() {
			return errors.New("path exists but is not a directory")
		}
		return nil
	}
	return err
}

// https://opensource.com/article/18/6/copying-files-go
func copyFile(sourceFile, destinationFile string) {
	input, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ioutil.WriteFile(destinationFile, input, 0644)
	if err != nil {
		fmt.Println("Error creating", destinationFile)
		fmt.Println(err)
		return
	}
}
