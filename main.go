package main

import (
	"fmt"
	"log"
	"os"

	"github.com/alisjj/gatorade/internal/config"
)

func main() {
	data, err := config.Read()
	if err != nil {
		fmt.Printf("Error")
	}

	s := &state{cfg: &data}
	// var cmds commands
	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
		return
	}

	var cmd command
	cmd.name = os.Args[1]
	cmd.args = os.Args[2:]

	err = cmds.run(s, cmd)
	if err != nil {
		log.Fatal(err)
		return
	}

}
