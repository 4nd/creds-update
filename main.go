package main

import (
	"archive/zip"
	"database/sql"
	"encoding/json"
	"fmt"
	"importCreds/config"
	"importCreds/models"
	"log"
	"net/mail"
)

func main() {

	fileName := "/Users/andy/1PasswordExport-QB2HPMCOL5BB7G4V6COGZIF7MM-20240906-225532.1pux"
	zipArchive, zipError := zip.OpenReader(fileName)
	if zipError != nil {
		return
	}
	defer func(zipArchive *zip.ReadCloser) {
		err := zipArchive.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(zipArchive)

	//for _, file := range zipArchive.File {
	//	fmt.Println(file.Name)
	//}

	reader, openError := zipArchive.Open("export.data")
	if openError != nil {
		log.Fatal(openError)
	}

	pux := models.Pux{}
	jsonError := json.NewDecoder(reader).Decode(&pux)
	if jsonError != nil {
		log.Fatal(jsonError)
	}
	db := config.GetDatabase()

	credentials := make(map[string]models.Credential)

	for _, vault := range pux.Accounts[0].Vaults {
		vaultName := vault.Attrs.Name
		for _, cred := range vault.Items {
			credential := models.Credential{}
			credential.OnePasswordUuid = cred.Uuid
			credential.Vault = vaultName
			credential.Name = cred.Overview.Title
			credential.Url = cred.Overview.Url
			credential.Username = findInLoginFields(&cred.Details.LoginFields, "username")
			credential.Password = findInLoginFields(&cred.Details.LoginFields, "password")
			if credential.Username != "" && isEmail(credential.Username) {
				credential.Email = credential.Username
			}
			if credential.Username == "" && credential.Password == "" {
				continue
			}

			credentials[cred.Uuid] = credential
			storeCredential(db, &credential)
		}
	}

	fmt.Printf("%d credentials found", len(credentials))
}

func storeCredential(db *sql.DB, credential *models.Credential) {
	var _, err = db.Exec("INSERT INTO credentials "+
		"(one_password_uuid, vault, name, url, username, password, email, created_at, updated_at)"+
		"VALUES"+
		"(?, ?, ?, ?, ?, ?, ?, NOW(), NOW())",
		credential.OnePasswordUuid,
		credential.Vault,
		credential.Name,
		credential.Url,
		credential.Username,
		credential.Password,
		credential.Email,
	)
	if err != nil {
		log.Fatal(err)
	}
}

func findInLoginFields(fields *[]struct {
	Value       string `json:"value"`
	Id          string `json:"id"`
	Name        string `json:"name"`
	FieldType   string `json:"fieldType"`
	Designation string `json:"designation"`
}, s string) string {
	for _, field := range *fields {
		if field.Designation == s {
			return field.Value
		}
	}
	return ""
}

func isEmail(value string) bool {
	email, _ := mail.ParseAddress(value)
	return email != nil
}
