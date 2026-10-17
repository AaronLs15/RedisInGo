package main

import (
	"log"
	"os"
	"redisgo/store"
)

type LoggingMiddleware struct {
	inner  store.Storer
	logger *log.Logger
}

func NewLoggingMiddleware(inner store.Storer) *LoggingMiddleware {
	return &LoggingMiddleware{
		inner:  inner,
		logger: log.New(os.Stdout, "[log]", log.Ltime|log.Lmicroseconds),
	}
}

func (l *LoggingMiddleware) Get(Key string) (string, error) {
	l.logger.Printf("Get %q", Key) //side effect
	val, err := l.inner.Get(Key)

	if err != nil {
		l.logger.Printf("Get %q error: %v", Key, err)
		return "", err
	} else {
		l.logger.Printf("Get %q value: %q", Key, val)
	}

	return val, err
}

func (l *LoggingMiddleware) Set(key string, value string) error {
	l.logger.Printf("Set %q", key) //side effect
	return l.inner.Set(key, value)
}

func (l *LoggingMiddleware) Delete(key string) {
	l.logger.Printf("delete %q", key) //side effect
	l.inner.Delete(key)
}

func (l *LoggingMiddleware) Len() int {
	got := l.inner.Len()
	l.logger.Printf("Len %d", got) //side effect
	return got
}

func (l *LoggingMiddleware) Keys() []string {
	got := l.inner.Keys()
	l.logger.Printf("Keys %v", got)
	return got
}
