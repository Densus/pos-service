package service

import (
	"fmt"
	"github.com/densus/pos_service/pos/model/dto"
	"github.com/densus/pos_service/pos/model/entity"
	"github.com/densus/pos_service/pos/repository"
	"github.com/mashingan/smapping"
)

//UserService is a contract about what UserService can do
type UserService interface {
	Update(user dto.UpdateUserDTO) entity.User
	SingleUser(UserID string) entity.User
	AllUser() []entity.User
	Delete(UserID string, user entity.User)
}

type userService struct {
	userRepository repository.UserRepository
}

//NewUserService creates a new instance that represent UserService interface
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepository: userRepo}
}

//Update is a method that update data user with model dto.UpdateUserDTO
func (u *userService) Update(user dto.UpdateUserDTO) entity.User {
	mapped := smapping.MapFields(&user)
	userToUpdate := entity.User{}
	err := smapping.FillStruct(&userToUpdate, mapped)
	if err != nil {
		panic(err)
	}
	fmt.Println("userToUpdate: ", userToUpdate)

	res := u.userRepository.UpdateUser(userToUpdate)
	return res
}

//SingleUser is a method to get profile data
func (u *userService) SingleUser(UserID string) entity.User {
	return u.userRepository.GetUser(UserID)
}

//AllUser is a method to get all users
func (u *userService) AllUser() []entity.User {
	return u.userRepository.AllUser()
}

func (u *userService) Delete(UserID string, user entity.User)  {
	u.userRepository.DeleteUser(UserID, user)
}
