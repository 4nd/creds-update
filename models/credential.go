package models

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

type Credential struct {
	Id              int64      `json:"id"`
	OnePasswordUuid string     `json:"one_password_uuid"`
	Vault           string     `json:"vault"`
	Name            string     `json:"name"`
	Url             string     `json:"url"`
	Username        string     `json:"username"`
	Password        string     `json:"password"`
	Email           string     `json:"email"`
	Payload         string     `json:"payload"`
	CompletedAt     *time.Time `json:"completed_at"`
}

type CredentialId struct {
	Id int64 `json:id`
}

func GetAllCredentials(db *sql.DB) ([]Credential, error) {
	var credentials []Credential

	var rows, err = db.Query("SELECT " +
		"id, one_password_uuid, vault, name, url, username, email, completed_at " +
		"FROM credentials " +
		"WHERE vault='Personal' AND completed_at IS NULL " +
		"ORDER BY name LIMIT 5",
	)
	if err != nil {
		log.Fatal(err)
	}

	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	for rows.Next() {
		var c Credential
		if err := rows.Scan(&c.Id, &c.OnePasswordUuid, &c.Vault, &c.Name, &c.Url, &c.Username, &c.Email, &c.CompletedAt); err != nil {
			return nil, fmt.Errorf("GetAllCredentials: %v", err)
		}
		credentials = append(credentials, c)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetAllCredentials: %v", err)
	}

	return credentials, nil
}

func CompleteCredential(db *sql.DB, credentialId int64) (Credential, error) {
	var credential Credential

	_, err := db.Exec("UPDATE credentials SET completed_at=NOW() WHERE id=?", credentialId)
	if err != nil {
		return credential, fmt.Errorf("completeCredential Update: %v", err)
	}

	credential, err = GetCredentialById(db, credentialId)
	if err != nil {
		return credential, fmt.Errorf("completeCredential Get: %v", err)
	}

	return credential, nil
}

func GetCredentialById(db *sql.DB, id int64) (Credential, error) {
	var c Credential

	row := db.QueryRow("SELECT id, one_password_uuid, vault, name, url, username, email, completed_at FROM credentials WHERE id=?", id)
	if err := row.Scan(&c.Id, &c.OnePasswordUuid, &c.Vault, &c.Name, &c.Url, &c.Username, &c.Email, &c.CompletedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c, fmt.Errorf("GetCredentialById %d: no such credential", id)
		}
		return c, fmt.Errorf("GetCredentialById %d: %v", id, err)
	}

	return c, nil
}
