package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"practice-AITU/internal"
	"practice-AITU/pkg"
)

func main() {
	http.HandleFunc("/", mainHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/whois", whoisHandler)
	fmt.Println("Running on http://localhost:8080/")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		fmt.Println("There is an error with parsing the HTML file:", err)
		return
	}

	err = html.Execute(w, nil)
	if err != nil {
		fmt.Println("There is an error with executing the HTML file:", err)
		return
	}
}

func whoisHandler(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("./templates/whois.html")
	if err != nil {
		fmt.Println("There is an error with parsing the HTML file:", err)
		return
	}
	domain := r.FormValue("domainName")
	whoisURL := fmt.Sprintf("https://www.whois.com/whois/%s", domain)

	resp, err := http.Get(whoisURL)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to retrieve WHOIS information", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read WHOIS response", http.StatusInternalServerError)
		return
	}

	whois := parseWhois(string(body))
	geoLocation := internal.GetGeoLocation(domain)

	whois.Country = geoLocation.Country
	whois.CountryCode = geoLocation.CountryCode
	whois.Latitude = geoLocation.Latitude
	whois.Longitude = geoLocation.Longitude
	if len(whois.Domain) == 0 {
		fmt.Fprintln(w, "There is error with domen")
		http.Redirect(w, r, "/", 200)
		return
	}
	err = html.Execute(w, whois)
	if err != nil {
		fmt.Println("There is an error with executing the HTML file:", err)
		return
	}
}

func parseWhois(html string) pkg.WhoisResponse {
	var whois pkg.WhoisResponse

	whois.Domain = pkg.ExtractValue(html, `<h1>(.*?)<\/h1>`)
	whois.Registrar = pkg.ExtractValue(html, `Registrar:<\/div><div class="df-value">(.*?)<\/div>`)
	whois.RegisteredOn = pkg.ExtractValue(html, `Registered On:<\/div><div class="df-value">(.*?)<\/div>`)
	whois.ExpiresOn = pkg.ExtractValue(html, `Expires On:<\/div><div class="df-value">(.*?)<\/div>`)
	whois.Status = pkg.ExtractMultipleValues(html, `Status:<\/div><div class="df-value">(.*?)<\/div>`)
	whois.NameServers = pkg.ExtractMultipleValues(html, `Name Servers:<\/div><div class="df-value">(.*?)<\/div>`)
	whois.Organization = pkg.ExtractValue(html, `Registrant Organization:<\/div><div class="df-value">(.*?)<\/div>`)
	whois.Country = pkg.ExtractValue(html, `Registrant Country:<\/div><div class="df-value">(.*?)<\/div>`)

	return whois
}
