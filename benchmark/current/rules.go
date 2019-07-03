package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"strings"

	"github.com/diegoholiveira/jsonlogic"
)

type (
	User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	Users []User
)

func getData(size int) io.Reader {
	rand.Seed(1199)

	users := make(Users, size)
	for i := 0; i < size; i++ {
		user := User{
			Name: "Zeze",
			Age:  rand.Intn(78),
		}
		users[i] = user
	}

	data := map[string]interface{}{
		"users": users,
	}

	var b bytes.Buffer

	encoder := json.NewEncoder(&b)
	encoder.Encode(data)

	return &b
}

func ExecuteRule(size int) Users {
	data := getData(size)
	logic := strings.NewReader(`{
		"filter": [
			{"var": "users"},
			{">=": [
				{"var": ".age"},
				18
			]}
		]
	}`)

	var result bytes.Buffer

	err := jsonlogic.Apply(logic, data, &result)
	if err != nil {
		fmt.Println(err.Error())

		return Users{}
	}

	var users Users

	decoder := json.NewDecoder(&result)
	decoder.Decode(&users)

	return users
}
