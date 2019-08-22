package main

import "fmt"

func iAmTheMan(name string) string {
	return fmt.Sprintf("%s is the man", name)
}

func main() {

	fmt.Println(iAmTheMan("Who"))

}

