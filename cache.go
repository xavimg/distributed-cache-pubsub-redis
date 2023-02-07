package main

type Cacher interface {
	Get(int) (string, bool)
	Set(int, string) error
	Remove(int) error
}

type NOPCache struct{}

func (c NOPCache) Get(int) (string, bool) {
	return "", false
}

func (c NOPCache) Set(int, string) error {
	return nil
}

func (c NOPCache) Remove(int) error {
	return nil
}
