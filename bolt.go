package sessions

import (
	"time"

	"github.com/gorilla/sessions"
	"github.com/uncle-gua/bolthold"
	"github.com/uncle-gua/boltstore"
)

func NewBoltStore(store *bolthold.Store, expiredSessionCleanup bool, keyPairs ...[]byte) *boltStore {
	s := boltstore.New(store, keyPairs...)
	if expiredSessionCleanup {
		quit := make(chan struct{})
		go s.PeriodicCleanup(1*time.Hour, quit)
	}
	return &boltStore{s}
}

type boltStore struct {
	*boltstore.BoltStore
}

func (s *boltStore) Options(options Options) {
	s.BoltStore.Options = &sessions.Options{
		Path:     options.Path,
		Domain:   options.Domain,
		MaxAge:   options.MaxAge,
		Secure:   options.Secure,
		HttpOnly: options.HttpOnly,
		SameSite: options.SameSite,
	}
}
