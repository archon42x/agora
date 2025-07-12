package dao

import (
	"log"

	"github.com/archon42x/agora/auth/config"
	"github.com/archon42x/agora/common/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func init() {
	var err error
	db, err = gorm.Open(mysql.Open(config.MysqlDSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("connect database failed: %v\n", err)
}

func CreateUser(user *model.User) (uint64, error) {
	if err := db.Create(user).Error; err != nil {
		return 0, err
	}
	return user.ID, nil
}

func FindUserByUsername(username string) (*model.User, error) {
	user := &model.User{}
	err := db.Where("username = ?", username).First(user).Error
	return user, err
}
