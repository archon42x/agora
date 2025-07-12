package logic

import (
	"github.com/archon42x/agora/auth/dao"
	"github.com/archon42x/agora/common/model"
)

func CreateUser(user *model.User) (uint64, error) {
	return dao.CreateUser(user)
}
func FindUserByUsername(username string) (*model.User, error) {
	return dao.FindUserByUsername(username)
}
