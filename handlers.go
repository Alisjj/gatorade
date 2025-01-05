package main

import (
	"context"
	"fmt"
	"time"

	"github.com/alisjj/gatorade/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("error: login handler expects a username")
	}
	ctx := context.Background()

	_, err := s.db.GetUser(ctx, cmd.args[0])
	if err != nil {
		return err
	}

	if err := s.cfg.SetUser(cmd.args[0]); err != nil {
		return err
	}

	fmt.Println("User has been set")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("error: register handler expects a username")
	}

	ctx := context.Background()
	uid := uuid.New()
	created_at := time.Now()
	updated_at := time.Now()
	name := cmd.args[0]
	c_user, err := s.db.CreateUser(ctx, database.CreateUserParams{ID: uid, CreatedAt: created_at, UpdatedAt: updated_at, Name: name})
	if err != nil {
		return err
	}
	if err := s.cfg.SetUser(c_user.Name); err != nil {
		return err
	}
	fmt.Println("User has been created!")
	fmt.Println(c_user)
	return nil
}

func handlerReset(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("error: reset expects zero args")
	}

	ctx := context.Background()
	if err := s.db.ReserDB(ctx); err != nil {
		return err
	}
	return nil
}

func handlerGetUsers(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("error: users command expects zero args")
	}

	ctx := context.Background()

	users, err := s.db.GetUsers(ctx)
	if err != nil {
		return err
	}
	for _, u := range users {
		if s.cfg.CurrentUserName == u.Name {
			fmt.Printf("* %s (current)\n", u.Name)
		} else {
			fmt.Printf("* %s\n", u.Name)
		}
	}
	return nil
}

func handlerAgg(_ *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("error: agg command expects zero args")
	}

	ctx := context.Background()
	url := "https://www.wagslane.dev/index.xml"
	feed, err := fetchFeed(ctx, url)
	if err != nil {
		return err
	}
	fmt.Println(feed)
	return nil
}
