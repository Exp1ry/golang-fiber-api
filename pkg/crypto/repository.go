package cryptoBrokers

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
ReadAllCryptoBrokers() (*[]models.CryptoBrokerModel, error)
NewCryptoBroker(broker *models.CryptoBrokerModel) error
EditCryptoBroker(id string, update map[string]interface{}) (bool,error)
RemoveCryptoBroker(id string) (bool, error)

}

type repository struct {
	collection *mongo.Collection
}


func NewRepo(Collection *mongo.Collection) Repository{
	return &repository{
		collection:Collection,
	}
}

func (r *repository) ReadAllCryptoBrokers() (*[]models.CryptoBrokerModel , error){

allCryptoBrokers := []models.CryptoBrokerModel{}

ctx := context.Background()

resp, err := r.collection.Find(ctx, bson.D{})
if err!= nil {
return nil, errors.New("Couldn't find brokers.")
}

for resp.Next(ctx){
cryptoBroker := models.CryptoBrokerModel{}

err := resp.Decode(&cryptoBroker)
if err!= nil {
	return nil, err
}
allCryptoBrokers = append(allCryptoBrokers, cryptoBroker)
}

defer resp.Close(ctx)

return &allCryptoBrokers, nil
}


func (r *repository) NewCryptoBroker(broker *models.CryptoBrokerModel) error {
	if broker.Name == "" || broker == nil {
		return errors.New("Please input a correct crypto broker.")
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

func (r *repository) EditCryptoBroker(id string, updates map[string]interface{}) (bool,error){
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

func (r *repository) RemoveCryptoBroker(id string) (bool,error){

	realID, err := primitive.ObjectIDFromHex(id)
	if err!= nil {
		return false, err
	}

ctx := context.Background()
resp := r.collection.FindOneAndDelete(ctx, bson.M{"_id":realID})
if resp.Err() == mongo.ErrNoDocuments {
	return  false, errors.New("No document found.")
}
return true,nil



}