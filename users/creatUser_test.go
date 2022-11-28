package users

import (
	"testing"
	"user/users/controller"
)

func creatUser(t *testing.T, account string, keyword string, email string) {
	check, err := controller.CreatUser(account, keyword, email)
	if err != nil {
		t.Error("creatError : " + account + " , " + keyword + " , " + email)
	} else if !check {
		t.Error("creatExist : " + account + " , " + keyword + " , " + email)
	}
}

func TestCreatData(t *testing.T) {
	creatUser(t, "a", "b", "c")
}
