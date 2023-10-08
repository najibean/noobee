package main

import "fmt"

func main() {
	var car = map[string]string{}
	car["name"] = "BMW"
	car["color"] = "Black"

	print(car["name"], car["color"])
}

func stringReturn(name string, color string) string {
	return "Mobil " + name + " berwarna " + color
}

func print(name string, color string) {
	fmt.Println(stringReturn(name, color))
}
