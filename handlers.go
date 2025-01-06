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
		return fmt.Errorf("login handler expects a username")
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
		return fmt.Errorf("register handler expects a username")
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
		return fmt.Errorf("reset expects zero args")
	}

	ctx := context.Background()
	if err := s.db.ResetUsers(ctx); err != nil {
		return err
	}
	if err := s.db.ResetFeeds(ctx); err != nil {
		return err
	}
	return nil
}

func handlerGetUsers(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("users command expects zero args")
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
		return fmt.Errorf("agg command expects zero args")
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

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("addfeed requries two arguments")
	}

	uid := uuid.New()
	ctx := context.Background()
	name := cmd.args[0]
	url := cmd.args[1]
	user, err := s.db.GetUser(ctx, s.cfg.CurrentUserName)
	if err != nil {
		return nil
	}
	user_id := user.ID
	feed, err := s.db.CreateFeed(ctx, database.CreateFeedParams{ID: uid, Name: name, UserID: user_id, Url: url, CreatedAt: time.Now(), UpdatedAt: time.Now()})
	if err != nil {
		return err
	}
	_, err = s.db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{ID: uuid.New(), UserID: user.ID, FeedID: feed.ID, CreatedAt: time.Now(), UpdatedAt: time.Now()})
	if err != nil {
		return err
	}
	fmt.Println(feed)
	return nil
}

func handlerFeeds(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("feeds doesn't require any arguments")
	}

	ctx := context.Background()
	feeds, err := s.db.GetFeeds(ctx)
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Printf("%v, %v, %v\n", feed.Name, feed.Url, feed.Username)
	}
	return nil
}

func handlerFollow(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("follow only required one argument, <url>")
	}
	ctx := context.Background()
	user, err := s.db.GetUser(ctx, s.cfg.CurrentUserName)
	if err != nil {
		return err
	}
	feed, err := s.db.GetFeed(ctx, cmd.args[0])
	if err != nil {
		return err
	}
	follows, err := s.db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{ID: uuid.New(), UserID: user.ID, FeedID: feed.ID, CreatedAt: time.Now(), UpdatedAt: time.Now()})
	if err != nil {
		return err
	}

	for _, follow := range follows {
		fmt.Println(follow)
	}
	return nil

}

func handlerFollows(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("follows doesnt require any arguments")
	}
	ctx := context.Background()
	user, err := s.db.GetUser(ctx, s.cfg.CurrentUserName)
	if err != nil {
		return err
	}
	follows, err := s.db.GetFeedFollowsForUser(ctx, user.ID)
	if err != nil {
		return err
	}

	fmt.Println(follows)
	return nil
}
