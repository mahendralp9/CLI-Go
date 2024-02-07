package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var menus = []string{}

var filePath = "D:/"
var path = filepath.Join(filePath, "CLI.json")

func main() {

	var menu string
	flag.StringVar(&menu, "menunya", "", "ini commmand untuk add menu")
	flag.Parse()

	var addMenus = append(menus, menu)

	os.MkdirAll(filePath, os.ModePerm)
	err := os.WriteFile(path, []byte(strings.Join(addMenus, "\n")), os.ModePerm)
	if err != nil {
		log.Println("Error Creating File")
		return
	}

	log.Println("File created successfully at: ", path)

}
