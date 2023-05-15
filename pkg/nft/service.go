package nftBroker

import (
	"api.ainvest.com/controller/models"
)

type Service interface {
GetAllNFTBrokers() (*[]models.NFTBrokerModel, error)
InsertNewBroker(broker *models.NFTBrokerModel) error
UpdateNFTBroker(id string, update map[string]interface{}) (bool,error)
DeleteNFTBroker(id string) (bool,error)
}

type service struct {
	repository Repository
}


func  NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAllNFTBrokers() (*[]models.NFTBrokerModel,error) {
return s.repository.ReadAllNFTBrokers()
}


func (s *service) InsertNewBroker(broker *models.NFTBrokerModel) error{
return s.repository.NewNFTBroker(broker)
}


func (s *service) UpdateNFTBroker(id string, update map[string]interface{}) (bool,error){
return s.repository.EditNFTBroker(id, update)


}


func (s *service) DeleteNFTBroker(id string) (bool,error){
return s.repository.RemoveNFTBroker(id)
}