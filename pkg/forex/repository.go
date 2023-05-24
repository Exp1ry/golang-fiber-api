package forexBrokers

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"api.ainvest.com/controller/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)


type Repository interface{
ReadAllForexBrokers() (*[]models.ForexBrokerModel, error)
NewForexBroker(broker *models.ForexBrokerModel) error
EditForexBroker(id string, update map[string]interface{}) (bool,error)
RemoveForexBroker(id string) (bool, error)

}

type repository struct {
	collection *mongo.Collection
}


func NewRepo(Collection *mongo.Collection) Repository{
	return &repository{
		collection:Collection,
	}
}

func (r *repository) ReadAllForexBrokers() (*[]models.ForexBrokerModel , error){

allForexBrokers := []models.ForexBrokerModel{}

ctx := context.Background()

resp, err := r.collection.Find(ctx, bson.D{})
if err!= nil {
return nil, errors.New("couldnt find brokers")
}

for resp.Next(ctx){
ForexBroker := models.ForexBrokerModel{}

err := resp.Decode(&ForexBroker)
if err!= nil {
	return nil, err
}
allForexBrokers = append(allForexBrokers, ForexBroker)
}

defer resp.Close(ctx)

return &allForexBrokers, nil
}


func (r *repository) NewForexBroker(broker *models.ForexBrokerModel) error {
	if broker.Name == "" || broker == nil {
		return errors.New("please input a correct Forex broker")
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

func (r *repository) EditForexBroker(id string, updates map[string]interface{}) (bool,error){

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

func (r *repository) RemoveForexBroker(id string) (bool , error){


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