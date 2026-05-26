package cache

import (
	"testing"
	"time"
)

func TestAddAndGet(t *testing.T) {
	interval := time.Millisecond * 5
	mockCache := NewCache(interval)
	if mockCache.cache == nil {
		t.Error("cache is nil")
	}

	cases := []struct {
		inputKey     string
		inputedValue []byte
	}{
		{
			inputKey:     "key1",
			inputedValue: []byte("val1"),
		},
	}

	for _, cs := range cases {

		mockCache.Add(cs.inputKey, cs.inputedValue)

		actual, ok := mockCache.Get(cs.inputKey)

		if !ok {
			t.Error("Key does not exist")
		}

		if string(actual) != string(cs.inputedValue) {
			t.Error("Value doesnt match")
		}
	}

}

func TestDeleteLoop(t *testing.T) {
	interval := time.Millisecond * 5
	mockCache := NewCache(interval)
	if mockCache.cache == nil {
		t.Error("cache is nil")
	}

	cases := []struct {
		inputKey     string
		inputedValue []byte
	}{
		{inputKey: "key1", inputedValue: []byte("val1")},
	}

	for _, cs := range cases {
		mockCache.Add(cs.inputKey, cs.inputedValue)

		// wait for the reap loop to evict
		time.Sleep(interval + time.Millisecond)

		actual, ok := mockCache.Get(cs.inputKey)

		if ok {
			t.Error("key should have been evicted")
		}
		if actual != nil {
			t.Errorf("expected nil value after eviction, got %q", actual)
		}
	}
}
