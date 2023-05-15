package forexBrokers

import (
	"api.ainvest.com/controller/models"
)


type Service interface {
	GetAllForexBrokers() (*[]models.ForexBrokerModel, error)
	InsertNewForexBroker(broker *models.ForexBrokerModel) error
	UpdateForexBroker(id string, update map[string]interface{}) (bool,error)
	DeleteForexBroker(id string) (bool,error)
	}
	
	type service struct {
		repository Repository
	}
	
	
	func  NewService(r Repository) Service {
		return &service{
			repository: r,
		}
	}
	
	func (s *service) GetAllForexBrokers() (*[]models.ForexBrokerModel,error) {
	return s.repository.ReadAllForexBrokers()
	}
	
	
	func (s *service) InsertNewForexBroker(broker *models.ForexBrokerModel) error{
	return s.repository.NewForexBroker(broker)
	}
	
	
	func (s *service) UpdateForexBroker(id string, update map[string]interface{}) (bool,error){
	return s.repository.EditForexBroker(id, update)
	
	
	}
	
	
	func (s *service) DeleteForexBroker(id string) (bool,error){
	return s.repository.RemoveForexBroker(id)
	}