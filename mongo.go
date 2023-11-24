package sessions

import (
	"github.com/gorilla/sessions"
	"github.com/laziercoder/mongostore"
	"go.mongodb.org/mongo-driver/mongo"
)

var _ sessions.Store = (*store)(nil)

func NewStore(c *mongo.Collection, maxAge int, ensureTTL bool, keyPairs ...[]byte) sessions.Store {
	return &store{mongostore.NewMongoStore(c, maxAge, ensureTTL, keyPairs...)}
}

type store struct {
	*mongostore.MongoStore
}

func (c *store) Options(options sessions.Options) {
	c.MongoStore.Options = &sessions.Options{
		Path:     options.Path,
		Domain:   options.Domain,
		MaxAge:   options.MaxAge,
		Secure:   options.Secure,
		HttpOnly: options.HttpOnly,
	}
}
