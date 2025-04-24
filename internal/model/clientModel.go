package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Client struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name" example:"John Doe" binding:"required" bson:"name"`
	Age      int                `json:"age" example:"22" binding:"required" bson:"age"`
	Sex      string             `json:"sex" example:"male" binding:"required" bson:"sex"`
	Phone    string             `json:"phone" example:"+996700000000" binding:"required" bson:"phone"`
	Email    string             `json:"email" example:"user@example.com" binding:"required" bson:"email"`
	Password string             `json:"password" example:"strongPassword123" binding:"required" bson:"password"`
}

type ClientInput struct {
	Email    string `json:"email" example:"user@example.com" bson:"email" binding:"required"`
	Password string `json:"password" example:"securePass123" bson:"password" binding:"required"`
}
