package services

import (
    "context"
    "mycrm/models"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson"
)

type InteractionService struct {
    db *mongo.Client
}

func NewInteractionService(db *mongo.Client) *InteractionService {
    return &InteractionService{db: db}
}

func (service *InteractionService) CreateInteraction(interaction *models.Interaction) error {
    collection := service.db.Database("mycrm").Collection("interactions")
    _, err := collection.InsertOne(context.TODO(), interaction)
    return err
}

func (service *InteractionService) GetInteractionsByCustomer(customerID string) ([]*models.Interaction, error) {
    collection := service.db.Database("mycrm").Collection("interactions")
    filter := bson.M{"customerID": customerID}
    cursor, err := collection.Find(context.TODO(), filter)
    if err != nil {
        return nil, err
    }
    var interactions []*models.Interaction
    if err = cursor.All(context.TODO(), &interactions); err != nil {
        return nil, err
    }
    return interactions, nil
}

func (service *InteractionService) CalculateStats() (*InteractionStats, error) {
    collection := service.db.Database("mycrm").Collection("interactions")
    stats := &InteractionStats{
        TypesCount: make(map[string]int64),
    }

    // Calculate total number of interactions
    totalCount, err := collection.CountDocuments(context.TODO(), bson.D{})
    if err != nil {
        return nil, err
    }
    stats.TotalCount = totalCount

    // Calculate counts by type using aggregation
    pipeline := mongo.Pipeline{
        {{"$group", bson.D{{"_id", "$type"}, {"count", bson.D{{"$sum", 1}}}}}},
    }
    cursor, err := collection.Aggregate(context.TODO(), pipeline)
    if err != nil {
        return nil, err
    }
    var results []bson.M
    if err = cursor.All(context.TODO(), &results); err != nil {
        return nil, err
    }
    for _, result := range results {
        if result["_id"] != nil {
            stats.TypesCount[result["_id"].(string)] = result["count"].(int64)
        }
    }

    return stats, nil
}

type InteractionStats struct {
    TotalCount int64
    TypesCount map[string]int64
}
