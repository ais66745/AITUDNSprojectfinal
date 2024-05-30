package pkg

type GeoLocation struct {
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Latitude    float64 `json:"lat"`
	Longitude   float64 `json:"lon"`
}

type WhoisResponse struct {
	Domain       string   `json:"domain"`
	Registrar    string   `json:"registrar"`
	RegisteredOn string   `json:"registered_on"`
	ExpiresOn    string   `json:"expires_on"`
	Status       []string `json:"status"`
	NameServers  []string `json:"name_servers"`
	Organization string   `json:"organization"`
	Country      string   `json:"country"`
	CountryCode  string   `json:"country_code"`
	Latitude     float64  `json:"latitude"`
	Longitude    float64  `json:"longitude"`
}
