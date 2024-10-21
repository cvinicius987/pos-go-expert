package main

import "net/http"

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", Home)
	mux.Handle("/blog", blog{title: "My Blog"})

	http.ListenAndServe(":3000", mux)
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home"))
}

type blog struct {
	title string
}

func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Blog " + b.title))
}
