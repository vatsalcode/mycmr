package services

import (
    "context"
    "mycrm/models"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
    db *mongo.Client
}

func NewUserService(db *mongo.Client) *UserService {
    return &UserService{db: db}
}

func (us *UserService) CreateUser(user *models.User) error {
    collection := us.db.Database("mycrm").Collection("users")
    user.HashPassword()
    _, err := collection.InsertOne(context.TODO(), user)
    return err
}

func (us *UserService) GetUserByID(id string) (*models.User, error) {
    collection := us.db.Database("mycrm").Collection("users")
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return nil, err
    }
    var user models.User
    err = collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&user)
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func (us *UserService) UpdateUser(id string, user *models.User) error {
    collection := us.db.Database("mycrm").Collection("users")
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return err
    }
    _, err = collection.UpdateOne(
        context.TODO(),
        bson.M{"_id": objID},
        bson.M{"$set": user},
    )
    return err
}

func (us *UserService) DeleteUser(id string) error {
    collection := us.db.Database("mycrm").Collection("users")
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return err
    }
    _, err = collection.DeleteOne(context.TODO(), bson.M{"_id": objID})
    return err
}
