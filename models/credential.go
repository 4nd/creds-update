package models

type Credential struct {
	Id              int64  `json:"id"`
	OnePasswordUuid string `json:"one_password_uuid"`
	Vault           string `json:"vault"`
	Name            string `json:"name"`
	Url             string `json:"url"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	Email           string `json:"email"`
	Payload         string `json:"payload"`
}
