package cryptoBrokers

import (
	"api.ainvest.com/controller/models"
)

type Service interface {
GetAllCryptoBrokers() (*[]models.CryptoBrokerModel, error)
InsertNewBroker(broker *models.CryptoBrokerModel) error
UpdateCryptoBroker(id string, update map[string]interface{}) (bool,error)
DeleteBroker(id string) (bool,error)
}

type service struct {
	repository Repository
}


func  NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAllCryptoBrokers() (*[]models.CryptoBrokerModel,error) {
return s.repository.ReadAllCryptoBrokers()
}


func (s *service) InsertNewBroker(broker *models.CryptoBrokerModel) error{
return s.repository.NewCryptoBroker(broker)
}


func (s *service) UpdateCryptoBroker(id string, update map[string]interface{}) (bool,error){
return s.repository.EditCryptoBroker(id, update)


}


func (s *service) DeleteBroker(id string) (bool,error){
return s.repository.RemoveCryptoBroker(id)
}