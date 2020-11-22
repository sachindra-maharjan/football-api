package api

import (
	"context"
)

type UserService service

type User struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

//ListAll lists all the users
func (u *UserService) ListAll(ctx context.Context) ([]*User, *Response, error) {

	req, err := u.client.NewRequest("GET", "todos", nil)
	if err != nil {
		return nil, nil, err
	}

	users := []*User{}
	resp, err := u.client.Do(ctx, req, &users)
	if err != nil {
		return nil, nil, err
	}

	return users, resp, err

}
