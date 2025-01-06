package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/alisjj/gatorade/internal/config"
	"github.com/alisjj/gatorade/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	data, err := config.Read()
	if err != nil {
		fmt.Printf("Error")
	}

	db, err := sql.Open("postgres", data.DbUrl)
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)

	s := &state{cfg: &data, db: dbQueries}
	// var cmds commands
	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerGetUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", handlerAddFeed)
	cmds.register("feeds", handlerFeeds)
	cmds.register("follow", handlerFollow)
	cmds.register("following", handlerFollowing)

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
		return
	}

	var cmd command
	cmd.name = os.Args[1]
	cmd.args = os.Args[2:]

	err = cmds.run(s, cmd)
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}

}
