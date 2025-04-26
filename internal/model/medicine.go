package model

type Medicine struct {
	ID          string `json:"id" bson:"_id"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
	Quantity    int    `json:"quantity" bson:"quantity"`
}
