package main

type HttpGetter interface {
	Get(url string) ([]byte, error)
	GetSaveDir() string
	GetKey() string
}
