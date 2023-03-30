package domain

type City struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	CountryId   int64  `json:"country_id"`
	CountryName string `json:"country_name"`
	CountryCca2 string `json:"country_cca2"`
	Admin1Id    int64  `json:"admin1_id"`
	Admin1Name  string `json:"admin1_name"`
	Admin1Code  string `json:"admin1_code"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
	WikiDataId  string `json:"wiki_data_id"`
}
