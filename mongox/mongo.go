package mongox

import (
	"context"
	"reflect"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
)

const DefaultURI = "mongodb://127.0.0.1:27017"

type Client struct {
	client   mongo.Client
	database mongo.Database
}

func NewClient(uri string, database string, opts ...func(opts *Options)) *Client {
	opt := defaultOptions()
	for _, o := range opts {
		o(&opt)
	}

	client, _ := mongo.Connect(context.TODO(), opt.opts.ApplyURI(uri))
	return &Client{
		client:   *client,
		database: *client.Database(database),
	}
}

func (p *Client) Collection(name string) *Collection {
	return &Collection{
		c: p.database.Collection(name),
	}
}

func (p *Client) Ping() error {
	return p.client.Ping(context.TODO(), nil)
}

func FieldTagName(field reflect.StructField) string {
	tag := field.Tag.Get("bson")
	if tag == "" {
		return ""
	}

	if strings.Contains(tag, ",") {
		tag, _, _ = strings.Cut(tag, ",")
	}

	return tag
}
