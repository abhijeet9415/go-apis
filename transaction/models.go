package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type Transaction struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID    string             `bson:"user_id,omitempty" json:"user_id"`
	Amount    float64            `json:"amount"`
	Type      string             `json:"type"` // credit/debit
	Timestamp primitive.DateTime `bson:"timestamp,omitempty" json:"timestamp"`
}
