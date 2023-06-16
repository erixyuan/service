package utils

import (
	"fmt"
	"os"
)

func checkFileIsExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

func WriteFile(filename string, str string) {
	fmt.Printf("write to %s >> %s\n", filename, str)
	f, _ := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	defer f.Close()
	if _, err := f.WriteString(str); err != nil {
		fmt.Printf("WriteFile %v\n", err)
	}
	if err := f.Sync(); err != nil {
		fmt.Printf("WriteFile %v\n", err)
	}
}
