package models

type Customer struct {
    ID          string `json:"id" bson:"_id,omitempty"`
    Name        string `json:"name"`
    ContactInfo string `json:"contactInfo"`
    Company     string `json:"company"`
    Status      string `json:"status"`
    Notes       string `json:"notes"`
}
