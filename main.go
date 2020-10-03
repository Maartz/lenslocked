package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"lenslocked.com/views"
)

var (
	homeView    *views.View
	contactView *views.View
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	if err := homeView.Template.ExecuteTemplate(w, homeView.Layout, nil); err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := contactView.Template.ExecuteTemplate(w, contactView.Layout, nil); err != nil {
		panic(err)
	}
}
func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	_, _ = fmt.Fprint(w, "<ul><li><strong>Q: Is this course hard?</strong> A: Not so, in fact, it's pretty easy for now.</li><li><strong>Q: Is this course mqde for me?</strong> A: You've got a 30 days cashback guarantee.</li><li><strong>Q: What about beginners?</strong> A: This course is tailored to fit for both beginners and more advanced developers.</li><li><strong>Q: Why is it not like, 10 bucks or so like Udemy?</strong> A: Because I really worked to produce this one.</li></ul>")
}

func four_oh_four(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>Ooops, something's missing here... We should fire someone 🤷‍♂️</h1>")
}
func main() {
	homeView = views.NewView("bootstrap", "views/home.gohtml")
	contactView = views.NewView("bootstrap", "views/contact.gohtml")
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	r.NotFoundHandler = http.HandlerFunc(four_oh_four)
	http.ListenAndServe(":3000", r)
}
