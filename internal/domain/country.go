package domain

type Country struct {
	ID             int    `bson:"id" json:"id"`
	Name           string `bson:"name" json:"name"`
	ISO3           string `bson:"iso3" json:"iso3"`
	ISO2           string `bson:"iso2" json:"iso2"`
	NumericCode    string `bson:"numeric_code" json:"numeric_code"`
	PhoneCode      string `bson:"phone_code" json:"phone_code"`
	Capital        string `bson:"capital" json:"capital"`
	Currency       string `bson:"currency" json:"currency"`
	CurrencyName   string `bson:"currency_name" json:"currency_name"`
	CurrencySymbol string `bson:"currency_symbol" json:"currency_symbol"`
	TLD            string `bson:"tld" json:"tld"`
	Native         string `bson:"native" json:"native"`
	Region         string `bson:"region" json:"region"`
	Subregion      string `bson:"subregion" json:"subregion"`
	Timezones      []*Timezone
	Translations   *Translations `bson:"translations" json:"translations"`
	Latitude       string        `bson:"latitude" json:"latitude"`
	Longitude      string        `bson:"longitude" json:"longitude"`
	Emoji          string        `bson:"emoji" json:"emoji"`
	EmojiU         string        `bson:"emojiU" json:"emojiU"`
}

type Translations struct {
	Kr   *string `bson:"kr" json:"kr,omitempty"`
	PtBr *string `bson:"pt-BR" json:"ptBR,omitempty"`
	Pt   *string `bson:"pt" json:"pt,omitempty"`
	Nl   *string `bson:"nl" json:"nl,omitempty"`
	Hr   *string `bson:"hr" json:"hr,omitempty"`
	Fa   *string `bson:"fa" json:"fa,omitempty"`
	De   *string `bson:"de" json:"de,omitempty"`
	Es   *string `bson:"es" json:"es,omitempty"`
	Fr   *string `bson:"fr" json:"fr,omitempty"`
	Ja   *string `bson:"ja" json:"ja,omitempty"`
	It   *string `bson:"it" json:"it,omitempty"`
	Cn   *string `bson:"cn" json:"cn,omitempty"`
	Tr   *string `bson:"tr" json:"tr,omitempty"`
}

type Timezone struct {
	ZoneName      *string `json:"zoneName,omitempty"`
	GmtOffset     *int    `json:"gmtOffset,omitempty"`
	GmtOffsetName *string `json:"gmtOffsetName,omitempty"`
	Abbreviation  *string `json:"abbreviation,omitempty"`
	TzName        *string `json:"tzName,omitempty"`
}
