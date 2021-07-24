package user_domain

import "testing"

// TestUserPasswordCreate Testing for UserPassword Create Successful
func TestUserPasswordCreate(t *testing.T) {
	userPasswordString := "password"
	_, err := UserPasswordCreate(userPasswordString)

	if err != nil {
		t.Fatal("Fail to Create UserPassword")
	}
}

// TestFailToUserPasswordCreateWhenUserPasswordStringIsEmpty Testing for UserPassword Create Fail When userPasswordString was Empty.
func TestFailToUserPasswordCreateWhenUserPasswordStringIsEmpty(t *testing.T) {
	userPasswordString := ""
	_, err := UserPasswordCreate(userPasswordString)

	if err == nil {
		t.Fatal("Fail to Create UserPassword")
	}
}

// TestFailToUserPasswordCreateUserPasswordLessThanSixWord Testing for UserPassword Create Fail When userPasswordString length less than six word.
func TestFailToUserPasswordCreateUserPasswordLessThanSixWord(t *testing.T) {
	userPasswordString := "abdd"
	userPasswordStringKoreanVer := "신영르르릅"
	_, err := UserPasswordCreate(userPasswordString)
	_, errKoreanVer := UserPasswordCreate(userPasswordStringKoreanVer)

	if err == nil || errKoreanVer == nil {
		t.Fatal("Fail to Create UserPassword")
	}
}