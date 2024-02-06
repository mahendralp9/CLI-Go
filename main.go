package main

import (
	"flag"
	"fmt"
)

var menus = []string{}

func main() {

	menu := flag.String("menunya", "", "ini commmand untuk add menu")

	flag.Parse()

	var addMenus = append(menus, *menu)
	fmt.Println(addMenus)
	fmt.Printf("menu: %s", *menu)

}
