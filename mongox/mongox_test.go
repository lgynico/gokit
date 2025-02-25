package mongox

import (
	"fmt"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

type player struct {
	PlayerId string   `bson:"player_id,omitempty"`
	Company  *company `bson:"company,omitempty"`
	Hobbies  []string `bson:"hobbies,omitempty"`
}

type company struct {
	Name     string `bson:"name"`
	Employee int32  `bson:"employee"`
}

func TestMongo(t *testing.T) {

	client := NewClient(DefaultURI, "test_me")
	collection := client.Collection("player")

	player := player{
		PlayerId: "1",
		Company: &company{
			Name:     "Tencent",
			Employee: 100000,
		},
		Hobbies: []string{"Game", "Girl", "Cat"},
	}

	id, err := collection.Insert(&player)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(id)
}

func TestFindOne(t *testing.T) {
	client := NewClient(DefaultURI, "test_me")
	collection := client.Collection("player")

	var player player
	err := collection.FindOne("6772623eb43474a8c1266e78", &player)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v %+v\n", &player, player.Company)
}

func TestFind(t *testing.T) {
	client := NewClient(DefaultURI, "test_me")
	collection := client.Collection("player")
	var results []*player

	err := collection.Find(bson.M{"company.name": "Tencent"}, &results)
	if err != nil {
		t.Fatal(err)
	}

	for _, player := range results {
		fmt.Printf("%+v %+v\n", player, player.Company)
	}
}
