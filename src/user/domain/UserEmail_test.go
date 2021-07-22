package user_domain

import "testing"

// TestCreateUserEmail Testing for UserEmail Create Successful
func TestCreateUserEmail(t *testing.T) {
	userEmailString := "den.shin.dev@gmail.com"
	_, err := UserEmailCreate(userEmailString)

	if err != nil {
		t.Fatal("Fail to Create UserEmail")
	}
}

// TestFailToCreateUserNameWhenUserEmailStringWasNotEmailForm Testing for UserEmail Create Fail When userEmailString was not email form.
func TestFailToCreateUserNameWhenUserEmailStringWasNotEmailForm(t *testing.T) {
	userEmailString := "a@b"
	_, err := UserEmailCreate(userEmailString)

	if err == nil {
		t.Fatal("Should be fail to Create UserEmail")
	}
}