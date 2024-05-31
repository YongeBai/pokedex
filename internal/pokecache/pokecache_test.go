package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		value []byte
	}{
		{"key1", []byte("value1")},
		{"key2", []byte("value1")},
		{"", []byte("value2")},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.value)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("%s not found in cache", c.key)
			} else if string(val) != string(c.value) {
				t.Errorf("Expected %v, got %v", c.value, val)
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5 * time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("key1", []byte("value1"))
	_, ok := cache.Get("key1")
	if !ok {
		t.Errorf("Expected value to be found")
	}
	time.Sleep(waitTime)

	_, ok = cache.Get("key1")
	if ok {
		t.Errorf("Expected value to be reaped")
	}
}