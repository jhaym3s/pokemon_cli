package cache

import (
	"testing"
)


func TestCreateChase(t *testing.T){
		mockCache := NewCache()
	if mockCache.cache == nil{
	t.Error("cache is nil")
	}

	mockCache.Add("key1", []byte("Val1"))

	actual , ok := mockCache.Get("key1")

	if !ok {
		t.Error("Key does not exist")
	}

	if string(actual) != "Val1"{
		t.Error("Value doesnt exist")
	}
	

}