package coreUuid

import (
	"reflect"
	"testing"
)

func TestCreateUuid(t *testing.T) {
	createdUuid := CreateUuid()
	if createdUuid == "" || reflect.TypeOf(createdUuid).String() != "string" {
		t.Fatal("Fail to create uuid")
	}
}
