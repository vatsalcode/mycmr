package models

import "time"

type Interaction struct {
    ID              string    `json:"id" bson:"_id,omitempty"`
    CustomerID      string    `json:"customerID"`
    UserID          string    `json:"userID"`
    Type            string    `json:"type"`
    Details         string    `json:"details"`
    InteractionDate time.Time `json:"interactionDate"`
}
