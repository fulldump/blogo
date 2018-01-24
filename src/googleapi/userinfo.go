package googleapi

type GoogleUserInfo struct {
	Id            string `json:"id" bson:",omitempty"`
	Email         string `json:"email" bson:",omitempty"`
	VerifiedEmail bool   `json:"verified_email" bson:",omitempty"`
	Name          string `json:"name" bson:",omitempty"`
	GivenName     string `json:"given_name" bson:",omitempty"`
	FamilyName    string `json:"family_name" bson:",omitempty"`
	Picture       string `json:"picture" bson:",omitempty"`
	Locale        string `json:"locale" bson:",omitempty"`
}
