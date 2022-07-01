package middleware

type Roles struct {
	ID   *string `json:"id,omitempty" db:"id"`
	Data *string `db:"data"`
}

type DataJSONB struct {
	Role_name    *string `json:"role_name,omitempty" db:"role_name"`
	Display_name *string `json:"display_name,omitempty" db:"display_name"`
}

type User struct {
	ID   *string `json:"id,omitempty" db:"id"`
	Data *string `json:"data,omitempty" db:"data"`
}
