package nftBroker

import (
	"context"
	"errors"

	"api.ainvest.com/controller/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


type Repository interface{
ReadAllNFTBrokers() (*[]models.NFTBrokerModel, error)
NewNFTBroker(broker *models.NFTBrokerModel) error
EditNFTBroker(id string, update map[string]interface{}) (bool,error)
RemoveNFTBroker(id string) (bool, error)

}

type repository struct {
	collection *mongo.Collection
}


func NewRepo(Collection *mongo.Collection) Repository{
	return &repository{
		collection:Collection,
	}
}

func (r *repository) ReadAllNFTBrokers() (*[]models.NFTBrokerModel , error){

allNFTBrokers := []models.NFTBrokerModel{}

ctx := context.Background()

resp, err := r.collection.Find(ctx, bson.D{})
if err!= nil {
return nil, errors.New("Couldn't find brokers.")
}

for resp.Next(ctx){
NFTBroker := models.NFTBrokerModel{}

err := resp.Decode(&NFTBroker)
if err!= nil {
	return nil, err
}
allNFTBrokers = append(allNFTBrokers, NFTBroker)
}

defer resp.Close(ctx)

return &allNFTBrokers, nil
}


func (r *repository) NewNFTBroker(broker *models.NFTBrokerModel) error {
	if broker.Name == "" || broker == nil {
		return errors.New("Please input a correct NFT broker.")
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

func (r *repository) EditNFTBroker(id string, updates map[string]interface{}) (bool,error){
existBroker := models.NFTBrokerModel{}

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

func (r *repository) RemoveNFTBroker(id string) (bool,error){


ctx := context.Background()
resp := r.collection.FindOneAndDelete(ctx, id)
if resp.Err() == mongo.ErrNoDocuments {
	return  false, errors.New("No document found.")
}
return true,nil



}