package main

import (
	"flag"
	"fmt"
)

func main() {

	menu := flag.String("menunya", "", "ini commmand untuk add menu")

	flag.Parse()

	fmt.Printf("menu: %s", *menu)

}
