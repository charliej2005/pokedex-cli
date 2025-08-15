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
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
		{
			key: "https://pokeapi.co/api/v2/pokemon/94",
			val: []byte("gengar"),
		},
		{
			key: "https://pokeapi.co/api/v2/pokemon/25",
			val: []byte("pikachu"),
		},
		{
			key: "https://pokeapi.co/api/v2/location/1",
			val: []byte("kanto"),
		},
		{
			key: "https://pokeapi.co/api/v2/location/8",
			val: []byte("galar"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}

func TestReapLoopMultipleKeys(t *testing.T) {
	const baseTime = 10 * time.Millisecond
	const waitTime = baseTime + 10*time.Millisecond
	cache := NewCache(baseTime)

	keys := []string{
		"https://pokeapi.co/api/v2/pokemon/1",
		"https://pokeapi.co/api/v2/pokemon/4",
		"https://pokeapi.co/api/v2/pokemon/7",
	}
	vals := [][]byte{
		[]byte("bulbasaur"),
		[]byte("charmander"),
		[]byte("squirtle"),
	}

	for i, key := range keys {
		cache.Add(key, vals[i])
	}

	for i, key := range keys {
		val, ok := cache.Get(key)
		if !ok || string(val) != string(vals[i]) {
			t.Errorf("expected to find key %s with value %s", key, vals[i])
		}
	}

	time.Sleep(waitTime)

	for _, key := range keys {
		_, ok := cache.Get(key)
		if ok {
			t.Errorf("expected key %s to be reaped", key)
		}
	}
}
