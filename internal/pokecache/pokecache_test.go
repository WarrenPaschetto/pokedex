package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache(10 * time.Millisecond)
	if cache.entries == nil {
		t.Error("cache is nil")
	}
}

func TestAddGetCache(t *testing.T) {
	cache := NewCache(1 * time.Minute)

	cases := []struct {
		inputKey string
		inputVal []byte
	}{
		{
			inputKey: "key1",
			inputVal: []byte("val1"),
		},
		{
			inputKey: "key2",
			inputVal: []byte("val2"),
		},
		{
			inputKey: "",
			inputVal: []byte("val3"),
		},
	}

	for _, cas := range cases {
		cache.Add(cas.inputKey, cas.inputVal)
		actual, ok := cache.Get(cas.inputKey)
		if !ok {
			t.Errorf("%s not found", cas.inputKey)
			continue
		}
		if string(actual) != string(cas.inputVal) {
			t.Errorf("%s does not match %s", string(actual), string(cas.inputVal))
			continue
		}
	}

}

func TestReapCache(t *testing.T) {
	cache := NewCache(1 * time.Millisecond)

	cache.Add("key1", []byte("val1"))
	time.Sleep(50 * time.Millisecond)
	_, ok := cache.Get("key1")
	if ok {
		t.Error("cache not deleted")
	}
}

func TestReapFail(t *testing.T) {
	cache := NewCache(5 * time.Millisecond)

	cache.Add("key1", []byte("val1"))
	time.Sleep(time.Millisecond)
	_, ok := cache.Get("key1")
	if !ok {
		t.Error("key1 should have been reaped")
	}
}
