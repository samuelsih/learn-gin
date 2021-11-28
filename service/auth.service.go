package service

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/samuelsih/learn-gofiber/dto"
	"github.com/samuelsih/learn-gofiber/entity"
	"github.com/samuelsih/learn-gofiber/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	VerifyCredential(email string, password string) interface{}
	CreateUser(user dto.UserCreateDTO) entity.User
	FindByEmail(email string) entity.User
	IsDuplicateEmail(email string) bool
}

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{
		userRepository: userRepo,
	}
}

func (service *authService) VerifyCredential(email string, password string) interface{} {
	response := service.userRepository.VerifyCredential(email, password)

	if value, ok := response.(entity.User); ok {
		comparedPassword := comparePassword(value.Password, []byte(password))

		if value.Email == email && comparedPassword {
			return response
		}
		return false
	}

	return false
}

func (service *authService) CreateUser(user dto.UserCreateDTO) entity.User {
	userToCreate := entity.User{}

	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))

	if err != nil {
		log.Fatal("Failed map : ", err)
	}

	response := service.userRepository.InsertUser(userToCreate)
	return response
}

func (service *authService) FindByEmail(email string) entity.User {
	return service.userRepository.FindByEmail(email)
}

func (service *authService) IsDuplicateEmail(email string) bool {
	response := service.userRepository.IsDuplicateEmail(email)

	return !(response.Error == nil)
}

func comparePassword(hashPassword string, plainPassword []byte) bool {
	byteHash := []byte(hashPassword)

	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)

	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
