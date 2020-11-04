package main

import "fmt"

func main() {
	// var colors := map[string]string
	colors := map[string]string{

		"red":   "#FF0000",
		"green": "#008000",
	}
	// colors := make(map[string]string)
	// colors["red"] = "#FF0000"
	// colors["green"] = "#008000"
	colors["white"] = "#FFFFFF"

	fmt.Println(colors)
	// delete(colors, "white")
	// fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		println(color, hex)
	}
}

/*
:::::::::::::::::::: maps VS structs :::::::::::::::::::::::::::::::::::::::::

Map

All keys must be the same type
Used to reperesent collection of related properties
All values must be same type
Dont need to knoe all the keys at compile time
Keys are indexed we can iterate through them
Reference type

Struct

Values can be different type
You need to know all the different feilds at compile time
Keys dont support indexing
Used to represent a "thing"  with lot of properties
Value type
Its like Object in Javascript


:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
*/
