package leveldb

import (
	"encoding/json"
	"github.com/clearcodecn/dep/b2s"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

var (
	db *leveldb.DB
)

func Open(path string, options *opt.Options) error {
	var err error
	db, err = leveldb.OpenFile(path, options)
	if err != nil {
		return err
	}
	return nil
}

func GetString(s string) string {
	data, _ := db.Get(b2s.StringToByte(s), nil)
	return b2s.BytesToString(data)
}

func GetObj(s string, v interface{}) bool {
	data, _ := db.Get(b2s.StringToByte(s), nil)
	if len(data) == 0 {
		return false
	}
	err := json.Unmarshal(data, v)
	if err != nil {
		return false
	}
	return true
}

func Set(key string, val interface{}) (err error) {
	switch v := val.(type) {
	case []byte:
		err = db.Put(b2s.StringToByte(key), v, nil)
	case string:
		err = db.Put(b2s.StringToByte(key), b2s.StringToByte(v), nil)
	default:
		data, err := json.Marshal(val)
		if err != nil {
			return err
		}
		err = db.Put(b2s.StringToByte(key), data, nil)
	}
	return err
}
