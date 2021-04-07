package testing

import (
	"encoding/json"
	"go_redis/core/db"
)

func (a *Author) Testing() (interface{}, error) {
	db := db.Redis()
	var result Author
	jsons, err := json.Marshal(a)
	if err != nil {
		return nil, err
	}

	err = db.Set("insert", jsons, 0).Err()
	if err != nil {
		return nil, err
	}

	val, err := db.Get("insert").Result()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(val), &result)

	return result, nil
}

func Delete(key string) error {
	db := db.Redis()

	_, err := db.Del(key).Result()
	if err != nil {
		return err
	}

	return nil
}

func Find(key string) (interface{}, error) {
	db := db.Redis()
	var result Author

	val, err := db.Get(key).Result()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(val), &result)

	return result, nil
}
