package repository

import (
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
)

type twoFactorRepositoryDB struct {
	db *memcache.Client
}

func NewTwoFactorRepositoryDB(db *memcache.Client) twoFactorRepositoryDB {
	return twoFactorRepositoryDB{db: db}
}

func (twoFactorRepo twoFactorRepositoryDB) Get(k string) (*Secret, error) {
	v, err := twoFactorRepo.db.Get(k)
	if err != nil {
		fmt.Println("Get secret key and value error ", err)
		return nil, err
	}
	s := Secret{
		Key:   v.Key,
		Value: string(v.Value),
	}
	return &s, nil
}
