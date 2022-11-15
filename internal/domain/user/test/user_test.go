package test

import (
	"easy_im/internal/domain/user/model"
	"easy_im/internal/domain/user/repo"
	"easy_im/internal/domain/user/service"
	"fmt"
	"testing"
)

func TestUser(t *testing.T) {
	t.Run("repo", func(t *testing.T) {
		userRepo := repo.NewUserRepo()
		err := userRepo.Add(&model.User{
			Username: "123456",
			Password: "123456",
			Nickname: "test1",
		})
		if err != nil {
			fmt.Println("add", err.Error())
		}
		users, _ := userRepo.Search("123")
		fmt.Printf("%v \n", users)
	})
	t.Run("service", func(t *testing.T) {
		err, _ := service.UserService.Auth("123456", "123456")
		if err != nil {
			fmt.Println("auth", err.Error())
		}
		service.UserService.Register("1234444", "123456", "test2")
	})

}
