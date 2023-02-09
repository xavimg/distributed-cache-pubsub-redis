package main

type Cacher interface {
	Get(int) ([]byte, bool)
	Set(int, []byte) error
	Remove(int) error
}

type NOPCache struct{}

func (c NOPCache) Get(int) ([]byte, bool) {
	return nil, false
}

func (c NOPCache) Set(int, *User) error {
	return nil
}

func (c NOPCache) Remove(int) error {
	return nil
}
