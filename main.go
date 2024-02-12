package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var menus = []string{}

var filePath = "D:/GitWork/CLI-Golang"
var path = filepath.Join(filePath, "CLI.json")

type Menus struct {
	Menu []string
}

func main() {

	var menu string
	var listMenu bool
	var delMenu string
	var upMenu string
	var newValue string
	flag.StringVar(&menu, "add", "", "ini commmand untuk add menu")
	flag.StringVar(&upMenu, "update", "", "ini command untuk update menu")
	flag.BoolVar(&listMenu, "list", false, "ini command untuk get list menu")
	flag.StringVar(&delMenu, "delete", "", "ini command untuk delete menu dari list")
	flag.StringVar(&newValue, "value", "", "input value barunya dengan ini")
	flag.Parse()

	switch {
	case menu != "":
		addMenu(menu)

	case listMenu:
		getMenu()

	case delMenu != "":
		deleteMenu(delMenu)

	case upMenu != "":
		updateMenu(upMenu, newValue)
	}

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

func isMenuExist(menu []string, addMenu string) bool {

	if menu != nil {
		for _, c := range menu {
			if c == addMenu {
				return true
			}
		}
	}

	return false

}

func trimSpaceBehind(addMenu string) string {

	var blank = ""

	for index, s := range addMenu {

		if index == len(addMenu)-1 {
			if string(s) == " " {
				return blank
			}
		}
		blank += string(s)
	}

	return addMenu
}

func addMenu(menu string) {

	//readfile menu
	theMenu, err := readFile(path)
	if err != nil {
		fmt.Println("error :", err)
		return
	}

	menu = trimSpaceBehind(menu)

	validate := isMenuExist(theMenu, menu)
	if validate {
		fmt.Println("Menu already exist !!")
		return
	}

	if menu != "" {
		var addMenu = append(theMenu, menu)
		if err := writeFile(path, addMenu); err != nil {
			fmt.Println("Error writing file : ", err)
			return
		}
	}
}

func updateMenu(item string, newValue string) {

	var idxItemToBeEdited int = -1

	//readfile menu
	menu, err := readFile(path)
	if err != nil {
		fmt.Println("error :", err)
		return
	}

	for idx, menuItem := range menu {
		if menuItem == item {
			idxItemToBeEdited = idx
			break
		}
	}

	if idxItemToBeEdited == -1 {
		fmt.Println("Menu tidak ditemukan !!")
		return
	}

	if newValue == " " {
		fmt.Println("Menu tidak boleh kosong !!")
		return
	}

	menu[idxItemToBeEdited] = newValue

	if err := writeFile(path, menu); err != nil {
		fmt.Println("Error writing file : ", err)
		return
	}

	fmt.Printf("New Menu: %s\n", strings.Join(menu, ", "))

}

func getMenu() {

	//readfile menu
	theMenu, err := readFile(path)
	if err != nil {
		fmt.Println("error :", err)
		return
	}

	fmt.Printf("Menus: %s", strings.Join(theMenu, ", "))

}

func deleteMenu(item string) {

	var idxItemToBeDeleted int = -1

	//readfile menu
	menu, err := readFile(path)
	if err != nil {
		fmt.Println("error :", err)
		return
	}

	for idx, menuItem := range menu {
		if menuItem == item {
			idxItemToBeDeleted = idx
			break
		}
	}

	if idxItemToBeDeleted == -1 {
		fmt.Println("Menu tidak ditemukan !!")
		return
	}

	menu = append(menu[:idxItemToBeDeleted], menu[idxItemToBeDeleted+1:]...)

	if err := writeFile(path, menu); err != nil {
		fmt.Println("Error writing file : ", err)
		return
	}

	fmt.Printf("Menus: %s\n", strings.Join(menu, ", "))
}
