package domain

type Admin1 struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Admin1Code  string `json:"admin1_code"`
	CountryID   int    `json:"country_id"`
	CountryName string `json:"country_name"`
	CountryCCA2 string `json:"country_cca2"`
	Type        string `json:"type"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
}

type Admin1Filter struct {
	Name        *string
	CountryCCA2 *string
	Admin1Code  *string
}
