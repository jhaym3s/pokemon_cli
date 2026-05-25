package cache

import (
	"testing"
)


func TestCreateChase(t *testing.T){
		mockCache := NewCache()
	if mockCache.cache == nil{
	t.Error("cache is nil")
	}

	cases := []struct{
		inputKey string
		inputedValue []byte
	}{
		{
			inputKey: "key1",
			inputedValue: []byte("val1"),
		},
	}

	for _, cs := range cases {
	
	mockCache.Add(cs.inputKey, cs.inputedValue)

	actual , ok := mockCache.Get(cs.inputKey)

	if !ok {
		t.Error("Key does not exist")
	}

	if string(actual) != string(cs.inputedValue){
		t.Error("Value doesnt match")
	}
	}

	
	

}