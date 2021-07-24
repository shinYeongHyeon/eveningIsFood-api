package user_domain

import "testing"

// TestNewUserCreate Testing for NewUserCreate Create Successful
func TestNewUserCreate(t *testing.T) {
	userProps := NewUserProps {
		UserEmail:    UserEmail { "den.shin.dev@gmail.com" },
		UserName:     UserName { "신영현" },
		UserPassword: UserPassword { "password" },
	}
	_, err := NewUserCreate(userProps)

	if err != nil {
		t.Fatal("Fail to Create UserEmail")
	}
}