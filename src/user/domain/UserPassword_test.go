package user_domain

import "testing"

// TestCreateUserPassword Testing for UserPassword Create Successful
func TestCreateUserPassword(t *testing.T) {
	userPasswordString := "password"
	_, err := UserPasswordCreate(userPasswordString)

	if err != nil {
		t.Fatal("Fail to Create UserPassword")
	}
}

// TestFailToCreateUserPasswordWhenUserPasswordStringIsEmpty Testing for UserPassword Create Fail When userPasswordString was Empty.
func TestFailToCreateUserPasswordWhenUserPasswordStringIsEmpty(t *testing.T) {
	userPasswordString := ""
	_, err := UserPasswordCreate(userPasswordString)

	if err == nil {
		t.Fatal("Fail to Create UserPassword")
	}
}

// TestFailToCreateUserPasswordUserPasswordLessThanSixWord Testing for UserPassword Create Fail When userPasswordString length less than six word.
func TestFailToCreateUserPasswordUserPasswordLessThanSixWord(t *testing.T) {
	userPasswordString := "abdd"
	userPasswordStringKoreanVer := "신영르르릅"
	_, err := UserPasswordCreate(userPasswordString)
	_, errKoreanVer := UserPasswordCreate(userPasswordStringKoreanVer)

	if err == nil || errKoreanVer == nil {
		t.Fatal("Fail to Create UserPassword")
	}
}