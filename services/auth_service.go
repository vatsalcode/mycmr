package services

import (
    "context"
    "mycrm/models"
    "mycrm/utils"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson"
)

type AuthService struct {
    db *mongo.Client
    jwtSecret string
}

func NewAuthService(db *mongo.Client, jwtSecret string) *AuthService {
    return &AuthService{db: db, jwtSecret: jwtSecret}
}

func (service *AuthService) RegisterUser(user *models.User) error {
    user.HashPassword()
    collection := service.db.Database("mycrm").Collection("users")
    _, err := collection.InsertOne(context.TODO(), user)
    return err
}

func (service *AuthService) LoginUser(email, password string) (string, error) {
    collection := service.db.Database("mycrm").Collection("users")
    var user models.User
    err := collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
    if err != nil {
        return "", err
    }
    if !user.CheckPassword(password) {
        return "", err
    }
    token, err := utils.GenerateJWT(user.ID, service.jwtSecret)
    if err != nil {
        return "", err
    }
    return token, nil
}
