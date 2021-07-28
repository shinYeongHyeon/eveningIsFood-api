package user_domain

import (
	"testing"
)

// TestNewUserCreate Testing for NewUserCreate Create Successful
func TestNewUserCreate(t *testing.T) {
	userProps := NewUserProps {
		UserEmail:    UserEmail { "den.shin.dev@gmail.com" },
		UserName:     UserName { "신영현" },
		UserPassword: UserPassword { "password" },
	}
	user, err := NewUserCreate(userProps)

	if err != nil || user.Uuid == "" || user.Email.Value != "den.shin.dev@gmail.com" || user.Name.Value != "신영현" || user.Password.Value != "password" {
		t.Fatal("Fail to Create UserEmail")
	}
}