package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var menus = []string{}

var filePath = "D:/GitWork/CLI-Golang"
var path = filepath.Join(filePath, "CLI.json")

func main() {

	var menu string
	flag.StringVar(&menu, "menunya", "", "ini commmand untuk add menu")
	flag.Parse()

	var addMenus = append(menus, menu)

	if err := writeFile(path, addMenus); err != nil {
		fmt.Println("Error writing file : ", err)
	}

	//readfile menu
	// menus := readFile(path)

}

func readFile(path string) ([]string, error) {

	var menus []string

	content, err := os.ReadFile(path)
	if err != nil {

		if os.IsNotExist(err) { //jika filenya tidak ketemu atau tidak ada maka akan menampilkan slice kosong
			return menus, nil
		}

		json.Unmarshal(content, &menus)

		return nil, err

	}

	return menus, nil
}

func writeFile(path string, addMenus []string) error {

	type Menus struct {
		Menu []string
	}

	menyu := Menus{
		Menu: addMenus,
	}

	m, err := json.Marshal(menyu)
	if err != nil {
		return err
	}

	os.MkdirAll(filePath, os.ModePerm)
	err = os.WriteFile("CLI.json", m, os.ModePerm)
	if err != nil {
		log.Println("Error Creating File")
		return nil
	}

	log.Println("File created successfully at: ", path)

	return nil

}
