package user

import (
	"fmt"
	"kriya_Test/util/util"
	"math"

	"github.com/jmoiron/sqlx"
)

func getListUser(db *sqlx.DB, param GetUserList) ([]UserData, error) {
	var (
		datas []UserData
		data  UserDetail
	)
	sql := `SELECT id, data FROM users limit 5 offset $1`
	offset := math.Abs(float64(param.Page-1)) * 5
	result, err := db.Queryx(sql, offset)
	if err != nil {
		return nil, err
	}

	for result.Next() {
		err := result.StructScan(&data)
		if err != nil {
			return nil, err
		}
		json := data.UnMarshal()
		fmt.Println(json)
		datas = append(datas, UserData{
			Username: json.Username,
			Email:    json.Email,
			Status:   json.Is_Active,
		})
	}

	return datas, nil
}

func getUser(db *sqlx.DB, param GetUser) (*UserData, error) {
	var (
		data     UserDetail
		roleData UserDetail
	)
	sql := `SELECT u.data, u.role_id 
			FROM users u
			JOIN roles ro on ro.id  = u.role_id 
			WHERE u.id = $1`

	user, err := db.Queryx(sql, param.ID)
	if err != nil {
		return nil, err
	}

	for user.Next() {
		err := user.StructScan(&data)
		if err != nil {
			return nil, err
		}

		query := `SELECT data FROM roles WHERE id = $1`
		role, err := db.Queryx(query, data.Role_ID)
		if err != nil {
			return nil, err
		}

		for role.Next() {
			err = role.StructScan(&roleData)
			if err != nil {
				return nil, err
			}
		}

	}

	roleJson := roleData.UnMarshal()
	json := data.UnMarshal()

	return &UserData{
		ID:        data.ID,
		Username:  json.Username,
		Email:     json.Email,
		Role_Name: roleJson.Role_Name,
	}, nil
}

func CreatedUser(db *sqlx.DB, id, data, role string) error {
	sql := `INSERT into users("id","data","role_id","created_at", "updated_at", "deleted_at") 
		values($1,$2,$3,$4,$5,6)`

	result, err := db.Exec(sql, id, data, role, util.GetCurrentDate(), util.GetCurrentDate(), 0)
	if err != nil {
		return err
	}

	if affected, _ := result.RowsAffected(); affected > 0 {
		return nil
	}
	return fmt.Errorf("error when Insert New User")
}

func UpdateUser(db *sqlx.DB, data, id string) error {
	query := `UPDATE users 
		SET data = $1
		WHERE id = $2;`
	result, err := db.Exec(query, data, id)
	if err != nil {
		return err
	}

	if affected, _ := result.RowsAffected(); affected > 0 {
		return nil
	}
	return fmt.Errorf("error when Update New User")
}

func DeleteUser(db *sqlx.DB, id string) error {
	query := `DELETE FROM users WHERE id = $1;
	`
	result, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	if affected, _ := result.RowsAffected(); affected > 0 {
		return nil
	}
	return fmt.Errorf("error when Delete User")
}
