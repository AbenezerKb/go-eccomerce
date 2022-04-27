package service

import (
	"gin-exercise/db"
	"context"
)

type LoginService interface {
	Login(ctx context.Context, email string, password string) bool
}

func NewLogin() LoginService {
	return &UserLogin{}
}

type UserLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *UserLogin) Login(ctx context.Context,email string, password string) bool {

	if _, status := db.UserInfo(ctx, email, password); status {
	
		return true
	}

	return false
}
