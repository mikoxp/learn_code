package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	path := "test"
	fmt.Printf("Numbers of files %d\n", countFiles(path))
}

func countFiles(path string) int {
	var result = 0
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		filePath := path + "/" + file.Name()
		if file.IsDir() {
			result = result + countFiles(filePath)
		} else {
			fmt.Println(filePath)
			result = result + 1
		}
	}
	return result

}
