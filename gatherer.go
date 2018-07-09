package main

import (
	"net/http"
	"io/ioutil"
)

type Gatherer struct {
	collection map[string]bool
}
/**
爬取一个url，返回byte和error
 */
func (self *Gatherer) gather(url string) ([]byte,error) {
	resp, err := http.Get(url)
	if err!=nil {
		return nil,err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
func (self *Gatherer) put(url string) bool {
	if url=="" || self.collection[url]==true{
		return false
	}
	if self.collection[url]==false {
		self.collection[url]=true
	}
	return true
}
func NewGatherer() *Gatherer{
	return &Gatherer{make(map[string]bool)}
}
func (self *Gatherer) GetCollection() map[string]bool {
	return self.collection
}