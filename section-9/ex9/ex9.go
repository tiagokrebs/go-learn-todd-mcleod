package main

import "fmt"

func main() {
	favSport := "basketball"
	switch {
	case favSport == "Soccer":
		fmt.Println("Hi! my favSport is soccer")
	case favSport == "basketball":
		fmt.Println("Hi! My favSport is basketball")
	}
}
