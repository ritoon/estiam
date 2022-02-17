package model

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestUserMarshalJSON(t *testing.T) {
	pwd := Password("tst123")
	u := User{
		FirstName: "Steeve",
		LastName:  "Morin",
		Password:  &pwd,
	}

	data, err := json.Marshal(u)
	if err != nil {
		t.Error()
	}
	fmt.Println(string(data))
	t.Log(string(data))
}
