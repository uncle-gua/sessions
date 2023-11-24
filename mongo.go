package sessions

import (
	"github.com/gorilla/sessions"
	"github.com/laziercoder/mongostore"
	"go.mongodb.org/mongo-driver/mongo"
)

var _ sessions.Store = (*mongoStore)(nil)

func NewMongoStore(c *mongo.Collection, maxAge int, ensureTTL bool, keyPairs ...[]byte) Store {
	return &mongoStore{mongostore.NewMongoStore(c, maxAge, ensureTTL, keyPairs...)}
}

type mongoStore struct {
	*mongostore.MongoStore
}

func (c *mongoStore) Options(options Options) {
	c.MongoStore.Options = &sessions.Options{
		Path:     options.Path,
		Domain:   options.Domain,
		MaxAge:   options.MaxAge,
		Secure:   options.Secure,
		HttpOnly: options.HttpOnly,
	}
}
