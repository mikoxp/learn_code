package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	path := "test"
	m := make(map[string]int)
	c := countFiles(path, m)
	fmt.Println("--------------------------------")
	writeMap(m)
	fmt.Println("--------------------------------")
	fmt.Printf("Numbers of files %d\n", c)
	fmt.Scanln()

}

func countFiles(path string, m map[string]int) int {
	var result = 0
	var tmp = 0
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		filePath := path + "/" + file.Name()
		if file.IsDir() {
			result = result + countFiles(filePath, m)
		} else {
			fmt.Println(filePath)
			tmp = tmp + 1
			result = result + 1
		}
	}
	if tmp > 0 {
		m[path] = tmp
	}
	return result
}

func writeMap(m map[string]int) {
	for key, value := range m {
		fmt.Printf("%s = %d\n", key, value)
	}
}
