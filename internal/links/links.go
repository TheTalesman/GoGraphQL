package links

import (
	"context"
	"log"
	"time"

	"github.com/thetalesman/apigraph/internal/db"
	"github.com/thetalesman/apigraph/internal/users"

	"gopkg.in/mgo.v2/bson"
)

//Link eh o struct dos links para dto
type Link struct {
	ID      string `bson:"_id" json:"_id"`
	Title   string
	Address string
	User    *users.User
}

//Save salva o link no banco
func (link Link) Save() string {

	collection := db.Client.Database("GRAPHQL").Collection("links")
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	res, err := collection.InsertOne(ctx, bson.M{"Title": link.Title, "Address": link.Address})
	if err != nil {
		log.Fatal("erro: ", err)
	}
	linkID := res.InsertedID.(bson.ObjectId)
	links := linkID.Hex()
	return links

}

func GetAll() []Link {
	collection := db.Client.Database("GRAPHQL").Collection("links")
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	res, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal("erro: ", err)
	}
	defer res.Close(ctx)
	var resultados []Link

	for res.Next(ctx) {
		var result Link
		err := res.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		resultados = append(resultados, result)
	}
	if err := res.Err(); err != nil {
		log.Fatal(err)
	}

	return resultados
}
