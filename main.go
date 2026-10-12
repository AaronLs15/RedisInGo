package main

import (
	"encoding/base64"
	"fmt"

	"redisgo/store"
	"redisgo/store/kv"
)

func main() {
	s := CreateStore()

	encrypted, err := SetKeyWithEncryption(s, "a", "hey you can see me")
	if err != nil {
		fmt.Println("Error setting key:", err)
		return
	}

	fmt.Println(encrypted)

}

func CreateStore() store.Storer {
	plain := kv.NewStore(20)
	logger := NewLoggingMiddleware(plain)
	return logger
}

func SetKeyWithEncryption(store store.Storer, key, val string) (string, error) {
	encoded := base64.StdEncoding.EncodeToString([]byte(val))

	if err := store.Set(key, encoded); err != nil {
		return "", err
	}

	store.Len()
	store.Get("key")

	return store.Get(key)
}
