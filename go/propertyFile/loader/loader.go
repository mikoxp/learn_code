package loader

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// https://stackoverflow.com/questions/40022861/parsing-values-from-property-file-in-golang
type Config map[string]string

func ReadConfig(filename string) (Config, error) {
	// init empty
	config := Config{}
	if len(filename) == 0 {
		return config, nil
	}
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		//read to end line
		line, err := reader.ReadString('\n')
		//search =
		if equal := strings.Index(line, "="); equal >= 0 {
			if key := strings.TrimSpace(line[:equal]); len(key) > 0 {
				value := ""
				if len(line) > equal {
					value = strings.TrimSpace(line[equal+1:])
				}
				config[key] = value
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
	}
	return config, nil
}
