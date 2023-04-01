package domain

type City struct {
	ID          int    `json:"id" bson:"id"`
	Name        string `json:"name" bson:"name"`
	CountryId   int64  `json:"country_id" bson:"country_id"`
	CountryName string `json:"country_name" bson:"country_name"`
	CountryCca2 string `json:"country_cca2" bson:"country_cca2"`
	Admin1Id    int64  `json:"admin1_id" bson:"admin1_id"`
	Admin1Name  string `json:"admin1_name" bson:"admin1_name"`
	Admin1Code  string `json:"admin1_code" bson:"admin1_code"`
	Latitude    string `json:"latitude" bson:"latitude"`
	Longitude   string `json:"longitude" bson:"longitude"`
	WikiDataId  string `json:"wiki_data_id" bson:"wiki_data_id"`
}

type CityFilter struct {
	Name       *string `json:"name,omitempty"`
	CountryID  *int    `json:"countryId,omitempty"`
	Admin1Code *string `json:"admin1Code,omitempty"`
}
