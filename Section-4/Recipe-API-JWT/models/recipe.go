package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

//swagger:parameters recipes newRecipe
type Recipe struct {
	//swagger:ignore
	ID           primitive.ObjectID `json:"id" `
	Name         string             `json:"name" `
	Tags         []string           `json:"tags" `
	Ingredients  []string           `json:"ingredients" `
	Instructions []string           `json:"instructions" `
	PublishedAt  time.Time          `json:"publishedAt" `
}
