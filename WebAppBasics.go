package main

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>This is Go go go!!</h1>")
	fmt.Fprintf(w, "<p>We can insert %s and %s", "html tags", "<strong>variables</strong>")
}

func hello_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<h1>This is Go go go!!</h1>
	<p>Inside the Hello!!!</p>
	<p>And is a Multiline Print</p>`)
}

func main() {
	http.HandleFunc("/", index_handler)       // HandleFunc(WebBrowserPath,HandlerFunction)
	http.HandleFunc("/hello/", hello_handler) // HandlerFunction: function that handles the representation in webpage
	http.ListenAndServe(":8000", nil)         // Server connection Declaration; ListenAndServe(PortNumber,ServerConfigDetails{Kept as nil})

}
