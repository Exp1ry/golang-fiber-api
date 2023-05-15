package nftBroker

import (
	"context"
	"errors"
"fmt"
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

realID,err:= primitive.ObjectIDFromHex(id)
if err!= nil {
	return false, err
}
    filter := bson.M{"_id": realID}

    for k, v := range updates {

if k == "_id" {
	return false, errors.New("cannot change an ID")
}
		update := bson.M{"$set": bson.M{k: v}}
        _, err := r.collection.UpdateOne(context.Background(), filter, update)
        if err != nil {
            fmt.Println(err)
            return false, err
        }
    }

    return true, nil
}


func (r *repository) RemoveNFTBroker(id string) (bool,error){

	ctx := context.Background()
	realID, err := primitive.ObjectIDFromHex(id)
	if err!= nil {
		return false, err
	}
	resp := r.collection.FindOneAndDelete(ctx, bson.M{"_id":realID})
	if resp.Err() == mongo.ErrNoDocuments {
		fmt.Println(err)
		return  false, errors.New("no document found")
	}
	return true,nil



}