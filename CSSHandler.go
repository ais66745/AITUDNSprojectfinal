package internal

import "net/http"

func CssHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/css")
	http.ServeFile(writer, request, "./static/style.css")
}
