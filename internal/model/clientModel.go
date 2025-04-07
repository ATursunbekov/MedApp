package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Client struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:name bson:"name"`
	Age      int                `json:age bson:"age"`
	Sex      string             `json:sex bson:"sex"`
	Phone    string             `json:phone bson:"phone"`
	Email    string             `json:email bson:"email"`
	Password string             `json:password bson:"password"`
}

type ClientInput struct {
	Email    string `json:email bson:"email"`
	Password string `json:password bson:"password"`
}
