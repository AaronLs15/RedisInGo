package main

import (
	"encoding/base64"
	"fmt"
	"time"

	"redisgo/store"
	"redisgo/store/kv"
	"redisgo/store/ttl"
)

func main() {
	keyValueStore := kv.NewStore(20)

	encrypted, err := SetKeyWithEncryption(keyValueStore, "a", "hey you can see me")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(encrypted)

	ttlStore := ttl.NewTTLStore(time.Second * 2)

	encryptedTtlValue, err := SetKeyWithEncryption(ttlStore, "a", "hey you can see me")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(encryptedTtlValue)

}

func SetKeyWithEncryption(store store.Storer, key, val string) (string, error) {
	encoded := base64.StdEncoding.EncodeToString([]byte(val))

	if err := store.Set(key, encoded); err != nil {
		return "", err
	}

	return store.Get(key)
}
