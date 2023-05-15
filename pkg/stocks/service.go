package stockBroker

import (
	"api.ainvest.com/controller/models"
)

type Service interface {
GetAllStockBrokers() (*[]models.StockBrokerModel, error)
InsertNewBroker(broker *models.StockBrokerModel) error
UpdateStockBroker(id string, update map[string]interface{}) (bool,error)
DeleteStockBroker(id string) (bool,error)
}

type service struct {
	repository Repository
}


func  NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAllStockBrokers() (*[]models.StockBrokerModel,error) {
return s.repository.ReadAllStockBrokers()
}


func (s *service) InsertNewBroker(broker *models.StockBrokerModel) error{
return s.repository.NewStockBroker(broker)
}


func (s *service) UpdateStockBroker(id string, update map[string]interface{}) (bool,error){
return s.repository.EditStockBroker(id, update)


}


func (s *service) DeleteStockBroker(id string) (bool,error){
return s.repository.RemoveStockBroker(id)
}