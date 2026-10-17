package main

import (
	"encoding/base64"
	"fmt"

	"redisgo/store"
	"redisgo/store/kv"
)

func main() {
	plain := kv.NewStore(20)
	logger := NewLoggingMiddleware(plain)
	metrics := NewMetricsMiddleware(logger)

	encrypted, err := SetKeyWithEncryption(metrics, "a", "hey you can see me")
	if err != nil {
		fmt.Println("Error setting key:", err)
		return
	}

	fmt.Println(encrypted)
	metrics.GetMetrics()

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
