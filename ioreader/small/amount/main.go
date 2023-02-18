package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file, err := os.Open("config.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Config file contents:", string(data))
}
