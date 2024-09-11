package models

import (
	"database/sql"
	"fmt"
	"log"
)

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

func GetAllCredentials(db *sql.DB) ([]Credential, error) {
	var credentials []Credential

	var rows, err = db.Query("SELECT " +
		"id, one_password_uuid, vault, name, url, username, email " +
		"FROM credentials " +
		"WHERE completed_at IS NULL " +
		"ORDER BY name ",
	)
	if err != nil {
		log.Fatal(err)
	}

	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	for rows.Next() {
		var c Credential
		if err := rows.Scan(&c.Id, &c.OnePasswordUuid, &c.Vault, &c.Name, &c.Url, &c.Username, &c.Email); err != nil {
			return nil, fmt.Errorf("GetAllCredentials: %v", err)
		}
		credentials = append(credentials, c)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetAllCredentials: %v", err)
	}

	return credentials, nil
}
