package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/muhlemmer/pagination"
)

type Page struct {
	Articles   []result
	Pagination *pagination.Pagination
}

const (
	pageSize      = 4 //Amount of entries per page
	paginationMax = 9
	paginationPos = 3
)

var templates = template.Must(template.ParseFiles("template/layout.html", "template/pagination.html"))

func invalidArgument(w http.ResponseWriter) {
	http.Error(w, "Invalid argument(s)", http.StatusBadRequest)
}

func notFound(w http.ResponseWriter, r *http.Request) {
	e := "Page not found"
	log.Fatal(e, "path", r.URL.Path, "host", r.RemoteAddr)
	http.Error(w, e, http.StatusNotFound)
}

func serverError(w http.ResponseWriter) {
	http.Error(w, "Internal server error", http.StatusInternalServerError)
}

//Extracts the page number from the request arguments
func pageNumber(r *http.Request) (p int, err error) {
	if r.FormValue("page") == "" {
		p = 1
		return
	}
	p, err = strconv.Atoi(r.FormValue("page"))
	if err != nil {
		log.Println("Invalid argument. Page: ", r.FormValue("page"), "err", err)
		return
	}
	return
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		notFound(w, r)
		return
	}

	p, err := pageNumber(r)
	if err != nil {
		invalidArgument(w)
		return
	}

	results, err := lipsum.Query(query{p, pageSize})
	if err != nil {
		log.Println("Query error: ", err)
		serverError(w)
		return
	}

	a := pagination.Args{
		Max:     paginationMax,
		Pos:     paginationPos,
		Page:    p,
		Records: len(results),
		Total:   lipsum.Count(),
		Size:    pageSize,
	}

	pag, err := pagination.New(a)
	if err != nil {
		if err.Error() == pagination.ErrPageNo {
			log.Println(err.Error())
			invalidArgument(w)
			return
		}
		log.Println(err.Error())
		serverError(w)
		return
	}

	view := Page{
		Articles:   results,
		Pagination: pag,
	}

	if err = templates.ExecuteTemplate(w, "layout", view); err != nil {
		log.Println("Template error: ", err.Error())
		serverError(w)
		return
	}
}

func main() {
	http.HandleFunc("/", rootHandler)

	log.Println("Listening...")
	log.Panic(http.ListenAndServe(":8080", nil).Error())
}
