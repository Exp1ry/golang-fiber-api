package users

import "api.ainvest.com/controller/models"

type Service interface{
FetchAllUsers() ([]*models.UserModel,error)
UpdateUser(id string, updates map[string]interface{}) (bool,error)
AddUser(u *models.UserModel) error
DeleteUser(id string) (bool,error)
AddAdmin(email,password,firstName,lastName string) (bool,error)
EnterAdmin(email,password string)  (string,error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}

}


func (s *service) FetchAllUsers() ([]*models.UserModel, error) {
	return s.repository.GetAllUsers()
}

func (s *service) UpdateUser(id string, updates map[string]interface{}) (bool, error) {
return s.repository.EditUser(id, updates)
}


func (s *service) AddUser(u *models.UserModel) error {
	return s.repository.CreateUser(u)
}

func (s *service) DeleteUser(id string) (bool,error) {
	return s.repository.RemoveUser(id)
}

func (s *service) AddAdmin(email,password,firstName,lastName string) (bool,error){
	return s.repository.AdminSignup(email,password,firstName,lastName)
}

func (s *service) EnterAdmin(email,password string)  (string,error){
	return s.repository.AdminSignin(email,password)
}