package user

import "golang.org/x/crypto/bcrypt"

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}

	/* Ya, secara sederhana, fungsi NewService digunakan untuk membuat objek service baru yang dapat 
	mengakses metode Save dengan menggunakan argumen repository yang diberikan.*/
}


func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Occupation = input.Occupation
	user.Email = input.Email
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password),bcrypt.MinCost)
	if err != nil{
		return user,err
	}
	user.PasswordHash = string(passwordHash)
	user.Role = "user"
	
	newUser,err := s.repository.Save(user)
	if err != nil{
		return newUser,err
	}
	return newUser,nil
}