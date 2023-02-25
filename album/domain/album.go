package domain

import "time"

type Album struct{
	ID string `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
	ReleaseDate time.Time `json:"release_date" bson:"ralease_date"`
	Producer string `json:"producer" bson:"producer"`
	
}