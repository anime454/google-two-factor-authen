package repository

import (
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
)

type qrCodeRepositoryDB struct {
	db *memcache.Client
}

func NewQrCodeRepositoryDB(db *memcache.Client) qrCodeRepositoryDB {
	return qrCodeRepositoryDB{db: db}
}

func (qrCodeRepo qrCodeRepositoryDB) Set(k string, v []byte) error {
	err := qrCodeRepo.db.Set(&memcache.Item{Key: k, Value: v})
	if err != nil {
		fmt.Println("Set secret key error ", err)
		return err
	}
	return nil
}
