package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

type Sitemap struct {
	//Locations []Location `xml:"sitemap"`
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Locations []string `xml:"url>loc"`
	LastMod   []string `xml:"url>lastmod"`
	ChngFreq  []string `xml:"url>changefreq"`
}

type NewsVal struct {
	Locations string
	ChngFreq  string
}

type NewsDisp struct { // FrontEnd Display Fields
	Title    string
	NewsShow map[string]NewsVal
}

/*type Location struct {
	Loc string `xml:"loc"` // convert Loc tags to string
}

func (l Location) String() string { // Converitng the slice Locations elements into strings
	return fmt.Sprintf(l.Loc)
}*/

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>News App</h1>")
}

func news_handler(w http.ResponseWriter, r *http.Request) {

	var s Sitemap
	var n News
	newsMap := make(map[string]NewsVal)
	resp, _ := http.Get("https://timesofindia.indiatimes.com/travel/staticsitemap/htarticles/sitemap-index.xml") // Gets the xml in bytes
	bytes, _ := ioutil.ReadAll(resp.Body)                                                                        // Reading the response in bytes
	xml.Unmarshal(bytes, &s)                                                                                     // Parsing the XML according to the tags mentioned in the structure s
	//bodystr := string(bytes)                                               // Converting the bytes to string
	//fmt.Println(bodystr)
	//resp.Body.Close()
	//fmt.Println(s.Locations)
	for /*index*/ _, location := range s.Locations { // range returns index,data in the index
		resp, _ := http.Get(location) // Gets the xml in bytes
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &n)
		for idx, _ := range n.LastMod {
			newsMap[n.LastMod[idx]] = NewsVal{n.Locations[idx], n.ChngFreq[idx]}
		}
	}

	p := NewsDisp{Title: "News Aggregator", NewsShow: newsMap}
	t, _ := template.ParseFiles("FrontEnd.html") // The FrontEnd template is defined in the given html file
	fmt.Println(t.Execute(w, p))                 // Usable for showing Error messages on the terminal
}

func main() {

	http.HandleFunc("/", index_handler)
	http.HandleFunc("/NewsView", news_handler)
	http.ListenAndServe(":8000", nil)

}
