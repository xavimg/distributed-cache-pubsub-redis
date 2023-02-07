package main

import "fmt"

type Store struct {
	data  map[int]string
	cache Cacher
}

func NewStore(c Cacher) *Store {
	data := map[int]string{
		1: "Elon musk is the new owner of twitter !",
		2: "Argentina worldclass team",
		3: "Harry Potter legacy",
	}

	return &Store{
		data:  data,
		cache: c,
	}
}

func (s *Store) Get(key int) (string, error) {
	val, ok := s.cache.Get(key)
	if ok {
		// busting cache
		if err := s.cache.Remove(key); err != nil {
			fmt.Println(err)
		}
		fmt.Println("returning value from the cache")
		return val, nil
	}

	val, ok = s.data[key]
	if !ok {
		return "", fmt.Errorf("key-not-found %d", key)
	}

	if err := s.cache.Set(key, val); err != nil {
		return "", nil
	}

	fmt.Println("returning key from internal storage")

	return val, nil
}
