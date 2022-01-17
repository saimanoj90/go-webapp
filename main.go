package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/contact", contactHandler)
	http.HandleFunc("/faq", faqHandler)
	fmt.Println("About to start the server...")
	http.ListenAndServe(":9000", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to my awesome site!</h1>"))
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte("Contact: <a href=\"mailto:chsaimanoj90@gmail.com\">chsaimanoj90@gmail.com</a>"))
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte("<ul><li>Do u have trail version?<br>Yes</li><li>do u operate in europe?<br>No</li><li>who are your customers?<br>No Customers Yet</li></ul>"))
}
