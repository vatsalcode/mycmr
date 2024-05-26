package services

import (
    "context"
    "mycrm/models"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type CustomerService struct {
    db *mongo.Client
}

func NewCustomerService(db *mongo.Client) *CustomerService {
    return &CustomerService{db: db}
}

func (cs *CustomerService) CreateCustomer(customer *models.Customer) error {
    collection := cs.db.Database("mycrm").Collection("customers")
    _, err := collection.InsertOne(context.TODO(), customer)
    return err
}

func (cs *CustomerService) GetCustomerByID(id string) (*models.Customer, error) {
    collection := cs.db.Database("mycrm").Collection("customers")
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return nil, err
    }
    var customer models.Customer
    err = collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&customer)
    if err != nil {
        return nil, err
    }
    return &customer, nil
}

func (cs *CustomerService) UpdateCustomer(id string, customer *models.Customer) error {
    collection := cs.db.Database("mycrm").Collection("customers")
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return err
    }
    _, err = collection.UpdateOne(
        context.TODO(),
        bson.M{"_id": objID},
        bson.M{"$set": customer},
    )
    return err
}

func (cs *CustomerService) DeleteCustomer(id string) error {
    collection := cs.db.Database("mycrm").Collection("customers")
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return err
    }
    _, err = collection.DeleteOne(context.TODO(), bson.M{"_id": objID})
    return err
}
