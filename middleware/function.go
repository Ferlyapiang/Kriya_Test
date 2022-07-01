package middleware

import (
	"encoding/json"
	"fmt"
	"kriya_Test/util/db"
)

// True Or False IF Checking
func CekRoleByID(id string) bool {
	var (
		roles Roles
		data  DataJSONB
	)
	db := db.Connect()
	defer db.Close()

	sql := `SELECT data from roles
			WHERE id = $1`

	err := db.Get(&roles, sql, id)
	if err != nil {
		fmt.Println(err)
		return false
	}

	json.Unmarshal([]byte(*roles.Data), &data)

	fmt.Println(*data.Role_name == "admin")
	return *data.Role_name == "admin"

}
