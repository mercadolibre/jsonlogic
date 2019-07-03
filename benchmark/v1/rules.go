package main

import (
	"fmt"
	"math/rand"

	"github.com/diegoholiveira/jsonlogic"
)

type (
	User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	Users []User
)

func getData(size int) interface{} {
	rand.Seed(1199)

	users := make([]interface{}, size)
	for i := 0; i < size; i++ {
		user := map[string]interface{}{
			"name": "Zeze",
			"age":  float64(rand.Intn(78)),
		}
		users[i] = interface{}(user)
	}

	data := map[string]interface{}{
		"users": users,
	}

	return interface{}(data)
}

func ExecuteRule(size int) Users {
	data := getData(size)
	logic := interface{}(map[string]interface{}{
		"filter": []interface{}{
			map[string]interface{}{
				"var": "users",
			},
			map[string]interface{}{
				">=": []interface{}{
					map[string]interface{}{
						"var": ".age",
					},
					float64(18),
				},
			},
		},
	})

	var result interface{}

	err := jsonlogic.Apply(logic, data, &result)
	if err != nil {
		fmt.Println(err.Error())

		return Users{}
	}

	var users Users

	for _, row := range result.([]interface{}) {
		_row := row.(map[string]interface{})

		users = append(users, User{
			Name: _row["name"].(string),
			Age:  int(_row["age"].(float64)),
		})
	}

	return users
}
