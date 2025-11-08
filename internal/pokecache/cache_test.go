package pokecache

import (
	"testing"
	"time"
)

func TestAddGetAndReap(t *testing.T) {
	interval := 50 * time.Millisecond
	c := NewCache(interval)

	key := "https://www.pokeapi.co/api/v2/location-era"
	val := []byte("hello")

	c.Add(key, val)

	got, ok := c.Get(key)
	if !ok {
		t.Fatalf("expected key to exist right after Add")
	}
	if string(got) != string(val) {
		t.Fatalf("expected %q, got %q", string(val), string(got))
	}

	got[0] = 'H'
	again, _ := c.Get(key)
	if string(again) != "hello" {
		t.Fatalf("cache value mutated via returned slice; want %q got %q", "hello", string(again))
	}

	time.Sleep(3 * interval)

	if _, ok := c.Get(key); ok {
		t.Fatalf("expected key to be reaped after interval")
	}
}
