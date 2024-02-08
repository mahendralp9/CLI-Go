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

type Menus struct {
	Menu []string
}

func main() {

	var menu string
	flag.StringVar(&menu, "menunya", "", "ini commmand untuk add menu")
	flag.Parse()

	//readfile menu
	theMenu, err := readFile(path)
	if err != nil {
		fmt.Println("error :", err)
		return
	}

	if menu != "" {
		var addMenus = append(theMenu, menu)
		if err := writeFile(path, addMenus); err != nil {
			fmt.Println("Error writing file : ", err)
			return
		}
	}

	fmt.Printf("Menus: %s\n", theMenu)

}

func readFile(path string) ([]string, error) {

	var menus []string

	content, err := os.ReadFile(path)
	if err != nil {

		if os.IsNotExist(err) { //jika filenya tidak ketemu atau tidak ada maka akan menampilkan slice kosong
			return menus, nil
		}

		return nil, err

	}

	var m Menus
	json.Unmarshal(content, &m)

	return m.Menu, nil
}

func writeFile(path string, addMenus []string) error {

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
