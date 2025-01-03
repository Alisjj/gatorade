package main

import (
	"fmt"

	"github.com/alisjj/gatorade/internal/config"
)

func main() {
	data, err := config.Read()
	if err != nil {
		fmt.Printf("Error")
	}
	if err = data.SetUser("aminu"); err != nil {
		fmt.Printf("Error")
	}

	fmt.Println(data.DbUrl)
	fmt.Println(data.CurrentUserName)

}
