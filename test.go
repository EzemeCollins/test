package main

import "fmt"

func concatenateStrings(strings ...string) string {
	result := ""
	for _, s := range strings {
		result += s
	}
	return result
}
func main() {
	combined := concatenateStrings("my name is Ezeme collins")
	fmt.Println(combined)
}
