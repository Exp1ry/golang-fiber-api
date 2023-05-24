package users

import (
	"context"
	"errors"
	"fmt"
	"time"

	"api.ainvest.com/controller/models"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)



type Repository interface {
GetAllUsers() ([]*models.UserModel, error)
CreateUser(u *models.UserModel) error
EditUser(id string, updates map[string]interface{}) (bool,error)
RemoveUser(id string) (bool, error)
AdminSignup(email,password,firstName,lastName string) (bool,error)
AdminSignin(email, password string)  (string,error)
}

type repository struct {
collection *mongo.Collection
adminCollection *mongo.Collection
}

func NewRepo(col, adminCol *mongo.Collection ) Repository {
	return &repository{
collection: col,
adminCollection: adminCol,
	}
}


func (r *repository) GetAllUsers() ([]*models.UserModel, error) {
	userSlice := []*models.UserModel{}

	ctx := context.Background()
	cursor,err := r.collection.Find(ctx, bson.M{})
	if err!= nil {
		fmt.Println(err)
		return nil, err
	}

	for cursor.Next(ctx){
		user := models.UserModel{}
err = cursor.Decode(&user)
if err!= nil {
	fmt.Println(err)

	return nil ,err
}

userSlice = append(userSlice, &user)
	}

	return userSlice, nil
}


func (r *repository) CreateUser(u *models.UserModel) error {

	existUser := models.UserModel{}

	ctx := context.Background()
	r.collection.FindOne(ctx, bson.M{"email":u.Email}).Decode(&existUser)

	if len(existUser.Email) > 1 || existUser.Email == u.Email {
return errors.New("user already exists.")
	}

	realID := primitive.NewObjectID()
u.ID = realID
	_, err := r.collection.InsertOne(ctx, u)
	if err!= nil {
		return  err
	}

	return  nil

} 


func (r *repository) EditUser(id string, updates map[string]interface{}) (bool,error){

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
	
	
	func (r *repository) RemoveUser(id string) (bool,error){
	
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

	func (r *repository) AdminSignup(email,password,firstName,lastName string) (bool, error) {
		ctx := context.Background()
		existAdmin := models.AdminModel{}

        r.adminCollection.FindOne(ctx, bson.M{"email":email}).Decode(&existAdmin)

		if existAdmin.Email == email || existAdmin.Email != "" {
return false, errors.New("existing account")
		}


		encryptedPW,err := bcrypt.GenerateFromPassword([]byte(password), 8)
		if err!= nil {

return false, err
		}

		newID := primitive.NewObjectID()

		newAdmin := models.AdminModel{
			ID: newID,
			Email: email,
			Password: string(encryptedPW),
			FirstName: firstName,
			LastName: lastName,
		}
		_, err = r.adminCollection.InsertOne(ctx, newAdmin)
		if err!= nil {

			return false, err
		}
return true, nil

	}


	func (r *repository) AdminSignin(email,password string) (string,error){
		ctx := context.Background()

		admin := models.AdminModel{}
		err := r.adminCollection.FindOne(ctx, bson.M{"email":email}).Decode(&admin)
		if err!= nil {
			return "", err
		}

		err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password))
		if err!= nil {
			return "", err
		}

// Create a new JWT token
token := jwt.New(jwt.SigningMethodHS256)

// Set the claims (payload) of the token
claims := token.Claims.(jwt.MapClaims)
claims["email"] = email
claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Set the expiration time (e.g., 24 hours from now)
signedToken, err := token.SignedString([]byte("secret"))
if err!= nil {
	return "", err
}
return signedToken, nil
	}