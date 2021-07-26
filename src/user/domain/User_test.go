package user_domain

import "testing"

// TestNewUserCreate Testing for NewUserCreate Create Successful
func TestNewUserCreate(t *testing.T) {
	userProps := NewUserProps {
		UserEmail:    UserEmail { "den.shin.dev@gmail.com" },
		UserName:     UserName { "신영현" },
		UserPassword: UserPassword { "password" },
	}
	user, err := NewUserCreate(userProps)

	if err != nil || user.Email.Value != "den.shin.dev@gmail.com" {
		t.Fatal("Fail to Create UserEmail")
	}
}