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

	if err := s.cfg.SetUser(cmd.args[0]); err != nil {
		return err
	}

	fmt.Println("User has been set")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("error: login handler expects a username")
	}

	ctx := context.Background()
	uid := uuid.New()
	created_at := time.Now()
	updated_at := time.Now()
	name := cmd.args[0]
	_, err := s.db.GetUser(ctx, name)
	if err != nil {
		return err
	}
	// if user != {
	// 	return
	// }
	_, err = s.db.CreateUser(ctx, database.CreateUserParams{ID: uid, CreatedAt: created_at, UpdatedAt: updated_at, Name: name})
	if err != nil {
		return err
	}
	return nil
}
