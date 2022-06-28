package generator

import (
	"github.com/jackc/fake"
	"go.mongodb.org/mongo-driver/bson"
)

func GenerateDummyProducts(numberOfContacts int) []interface{} {
	var users []interface{}
	for i := 0; i < numberOfContacts; i++ {
		users = append(users, bson.D{
			{"fullName", fake.FullName()},
			{"gender", fake.Gender()},
		})
	}

	return users
}
