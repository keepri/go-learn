package main

import (
	"errors"
	"fmt"
)

type user struct {
	name string
}

func (u user) Error() string {
	return fmt.Sprint("User error: ", u.name)
}

func main() {
	if user, err := get_user(
		// should error
		true,
		// error with
		fmt.Errorf("some error %w", fmt.Errorf("some bubbled up error")),
		// user{name: "me"},
	); err == nil {
		fmt.Println("We have user", user)
	} else {
		fmt.Println(err.Error())
		if unwrapped := errors.Unwrap(err); unwrapped != nil {
			fmt.Println(unwrapped)
		}
	}
}

func get_user(should_error bool, with error) (user, error) {
	if should_error == true {
		return user{}, with
	}

	return user{name: "me"}, nil
}
