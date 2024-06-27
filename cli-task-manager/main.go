package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("**************************************\n Weclome to your daily task manager! \n**************************************\n")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Choose one of the following options:-\n1:- Create\n2:- View\n3:- Update\n4:- Delete\n5:- Quit\n")
		userChoice, err := reader.ReadString('\n')
		userChoice = strings.TrimSpace(strings.ToLower(userChoice))
		if err != nil {
			fmt.Println("An Unexpected Error has occured")
			break
		}
		if userChoice == "quit" {
			fmt.Println("Exiting the Program!")
			break
		} else if userChoice == "create" {

		} else if userChoice == "view" {

		} else if userChoice == "update" {

		} else if userChoice == "delete" {

		} else {

		}
	}
}
