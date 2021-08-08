package user_domain

import "testing"

// TestUserNameCreate Testing for UserName Create Successful
func TestUserNameCreate(t *testing.T) {
	userNameString := "Shin"
	userName, err := UserNameCreate(userNameString)

	if err != nil || userName.Value() != "Shin" {
		t.Fatal("Fail to Create UserName")
	}
}

// TestFailToUserNameCreateWhenUserNameStringIsEmpty Testing for UserName Create Fail When userNameString was Empty.
func TestFailToUserNameCreateWhenUserNameStringIsEmpty(t *testing.T) {
	userNameString := ""
	_, err := UserNameCreate(userNameString)

	if err == nil {
		t.Fatal("Fail to Create UserName")
	}
}

// TestFailToUserNameCreateUserNameLessThanThreeWord Testing for UserName Create Fail When userNameString length less than three word.
func TestFailToUserNameCreateUserNameLessThanThreeWord(t *testing.T) {
	userNameString := "ab"
	userNameStringKoreanVer := "신영"
	_, err := UserNameCreate(userNameString)
	_, errKoreanVer := UserNameCreate(userNameStringKoreanVer)

	if err == nil || errKoreanVer == nil {
		t.Fatal("Fail to Create UserName")
	}
}