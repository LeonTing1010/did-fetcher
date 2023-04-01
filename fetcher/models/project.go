package models

/*
Domain Project Collection
Project_Collection
{"_id":"ENS","createdAt":{"$date":{"$numberLong":"1674226465473"}},
"description":"","name":"ENS","updatedAt":{"$date":{"$numberLong":"1675073679027"}}}
*/
type Project struct {
	Id          string `bson:"_id"`
	Description string `bson:"description,omitempty"`
	Name        string `bson:"name,omitempty"`
}

// var collection *mongo.Collection
