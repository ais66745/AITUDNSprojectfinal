package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"practice-AITU/pkg"
)

func GetGeoLocation(domain string) pkg.GeoLocation {
	url := fmt.Sprintf("http://ip-api.com/json/%s", domain)
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Failed to retrieve geolocation data:", err)
		return pkg.GeoLocation{}
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Failed to read geolocation response:", err)
		return pkg.GeoLocation{}
	}

	var geoLocation pkg.GeoLocation
	err = json.Unmarshal(body, &geoLocation)
	if err != nil {
		fmt.Println("Failed to parse geolocation data:", err)
		return pkg.GeoLocation{}
	}

	return geoLocation
}
