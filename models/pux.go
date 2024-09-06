package models

type Pux struct {
	Accounts []struct {
		Attrs struct {
			AccountName string `json:"accountName"`
			Name        string `json:"name"`
			Avatar      string `json:"avatar"`
			Email       string `json:"email"`
			Uuid        string `json:"uuid"`
			Domain      string `json:"domain"`
		} `json:"attrs"`
		Vaults []struct {
			Attrs struct {
				Uuid   string `json:"uuid"`
				Desc   string `json:"desc"`
				Avatar string `json:"avatar"`
				Name   string `json:"name"`
				Type   string `json:"type"`
			} `json:"attrs"`
			Items []struct {
				Uuid         string `json:"uuid"`
				FavIndex     int    `json:"favIndex"`
				CreatedAt    int    `json:"createdAt"`
				UpdatedAt    int    `json:"updatedAt"`
				State        string `json:"state"`
				CategoryUuid string `json:"categoryUuid"`
				Details      struct {
					LoginFields []struct {
						Value       string `json:"value"`
						Id          string `json:"id"`
						Name        string `json:"name"`
						FieldType   string `json:"fieldType"`
						Designation string `json:"designation"`
					} `json:"loginFields"`
					NotesPlain string `json:"notesPlain"`
					Sections   []struct {
						Title  string `json:"title"`
						Name   string `json:"name"`
						Fields []struct {
							Title string `json:"title"`
							Id    string `json:"id"`
							Value struct {
								Concealed string `json:"concealed"`
							} `json:"value"`
							IndexAtSource int  `json:"indexAtSource"`
							Guarded       bool `json:"guarded"`
							Multiline     bool `json:"multiline"`
							DontGenerate  bool `json:"dontGenerate"`
							InputTraits   struct {
								Keyboard       string `json:"keyboard"`
								Correction     string `json:"correction"`
								Capitalization string `json:"capitalization"`
							} `json:"inputTraits"`
						} `json:"fields"`
					} `json:"sections"`
					PasswordHistory []struct {
						Value string `json:"value"`
						Time  int    `json:"time"`
					} `json:"passwordHistory"`
					DocumentAttributes struct {
						FileName      string `json:"fileName"`
						DocumentId    string `json:"documentId"`
						DecryptedSize int    `json:"decryptedSize"`
					} `json:"documentAttributes"`
				} `json:"details"`
				Overview struct {
					Subtitle string `json:"subtitle"`
					Urls     []struct {
						Label string `json:"label"`
						Url   string `json:"url"`
					} `json:"urls"`
					Title string  `json:"title"`
					Url   string  `json:"url"`
					Ps    int     `json:"ps"`
					Pbe   float64 `json:"pbe"`
					Pgrng bool    `json:"pgrng"`
				} `json:"overview"`
			} `json:"items"`
		} `json:"vaults"`
	} `json:"accounts"`
}
