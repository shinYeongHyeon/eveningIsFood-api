package user_domain

import (
	"testing"
)

// TestUserEmailCreate Testing for UserEmail Create Successful
func TestUserEmailCreate(t *testing.T) {
	userEmailString := "den.shin.dev@gmail.com"
	userEmail, err := UserEmailCreate(userEmailString)

	if err != nil || userEmail.Value() != "den.shin.dev@gmail.com" {
		t.Fatal("Fail to Create UserEmail")
	}
}

// TestFailToUserEmailCreateWhenUserEmailStringWasNotEmailForm Testing for UserEmail Create Fail When userEmailString was not email form.
func TestFailToUserEmailCreateWhenUserEmailStringWasNotEmailForm(t *testing.T) {
	userEmailString := "a@b"
	_, err := UserEmailCreate(userEmailString)

	if err == nil {
		t.Fatal("Should be fail to Create UserEmail")
	}
}