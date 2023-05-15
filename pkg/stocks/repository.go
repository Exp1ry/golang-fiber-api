package stockBroker

import (
	"context"
	"errors"

	"api.ainvest.com/controller/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


type Repository interface{
ReadAllStockBrokers() (*[]models.StockBrokerModel, error)
NewStockBroker(broker *models.StockBrokerModel) error
EditStockBroker(id string, update map[string]interface{}) (bool,error)
RemoveStockBroker(id string) (bool, error)

}

type repository struct {
	collection *mongo.Collection
}


func NewRepo(Collection *mongo.Collection) Repository{
	return &repository{
		collection:Collection,
	}
}

func (r *repository) ReadAllStockBrokers() (*[]models.StockBrokerModel , error){

allStockBrokers := []models.StockBrokerModel{}

ctx := context.Background()

resp, err := r.collection.Find(ctx, bson.D{})
if err!= nil {
return nil, errors.New("Couldn't find brokers.")
}

for resp.Next(ctx){
StockBroker := models.StockBrokerModel{}

err := resp.Decode(&StockBroker)
if err!= nil {
	return nil, err
}
allStockBrokers = append(allStockBrokers, StockBroker)
}

defer resp.Close(ctx)

return &allStockBrokers, nil
}


func (r *repository) NewStockBroker(broker *models.StockBrokerModel) error {
	if broker.Name == "" || broker == nil {
		return errors.New("Please input a correct Stock broker.")
	}

	newID := primitive.NewObjectID()
	broker.ID = newID

	ctx := context.Background()

	_,err := r.collection.InsertOne(ctx, &broker)
	if err!= nil {
		return err
	}


	return nil

}

func (r *repository) EditStockBroker(id string, updates map[string]interface{}) (bool,error){
existBroker := models.StockBrokerModel{}

ctx := context.Background()

 resp := r.collection.FindOne(ctx, bson.M{"_id":id})
 if resp.Err()== mongo.ErrNoDocuments {
	return false,errors.New("No document found.")
 }

 err := resp.Decode(&existBroker)
 if err!= nil {
	return false,err
 }

 for k, v := range updates {
_, err := r.collection.UpdateByID(ctx, id, bson.M{k:v})
if err!= nil {
	return false,err
}
 }

 return true,nil
}

func (r *repository) RemoveStockBroker(id string) (bool,error){


ctx := context.Background()
resp := r.collection.FindOneAndDelete(ctx, id)
if resp.Err() == mongo.ErrNoDocuments {
	return  false, errors.New("No document found.")
}
return true,nil



}