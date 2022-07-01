package user

import "encoding/json"

type UserDetail struct {
	ID      string `db:"id"`
	Data    string `db:"data"`
	Role_ID string `db:"role_id"`
}

type Response struct {
	Message       string      `json:"message,omitempty"`
	Error_key     string      `json:"error_key,omitempty"`
	Error_message string      `json:"error_message,omitempty"`
	Data          interface{} `json:"data,omitempty"`
}

type UserData struct {
	Role_ID      string `json:"role_id,omitempty"`
	ID           string `json:"User_Id,omitempty"`
	Dob          string `json:"Dob,omitempty"`
	Email        string `json:"Email,omitempty"`
	Phone        string `json:"Phone,omitempty"`
	Fullname     string `json:"Fullname,omitempty"`
	Password     string `json:"Password,omitempty"`
	Username     string `json:"Username,omitempty"`
	Is_Active    bool   `json:"Is_Active,omitempty"`
	Status       bool   `json:"Status,omitempty"`
	Role_Name    string `json:"Role_Name,omitempty"`
	Display_Name string `json:"Display_Name,omitempty"`
}

type GetUserList struct {
	Page int `uri:"page" binding:"required" `
}

type GetUser struct {
	ID string `uri:"userid" binding:"required" `
}

func (u *UserDetail) UnMarshal() (data UserData) {
	json.Unmarshal([]byte(u.Data), &data)
	return data
}
