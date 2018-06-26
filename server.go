package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/knakk/rdf"
	"github.com/knakk/sparql"
)

const bookQueries = `
# Comment are ignored

# tag: book-query
PREFIX rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#>
PREFIX dbo: <http://dbpedia.org/ontology/>
PREFIX rdfs: <http://www.w3.org/2000/01/rdf-schema#>
PREFIX dc: <http://purl.org/dc/elements/1.1/>

select distinct ?name ?author ?thumbnail ?country ?comment ?genre ?title where{
?book rdf:type dbo:Book;
      rdfs:label ?name;
      rdfs:comment ?comment;
      dbp:author ?author;
      dbo:thumbnail ?thumbnail;
      dbp:country ?country;
      dbp:genre ?genre;
      dbp:title ?title

FILTER((str(?title)="{{.Bookname}}") && (lang(?comment)="en")).
}limit 1
`

//FILTER( (lang(?name)="zh") && (str(?title)="{{.Bookname}}") && (lang(?comment)="zh")).

const bookRecommend = `
# Comment are ignored

# tag: book-recommend
prefix rdf:<http://www.w3.org/1999/02/22-rdf-syntax-ns#>
prefix rdfs: <http://www.w3.org/2000/01/rdf-schema#>
prefix dbo: <http://dbpedia.org/ontology/>
prefix dbp: <http://dbpedia.org/property/>
prefix foaf: <http://xmlns.com/foaf/0.1/>
select distinct ?name where {
?au        rdfs:label "{{.Author}}"@en .
?book		dbp:author ?au .
            rdfs:label ?name
}limit 5
`

const filmQueries = `
# Comment are ignored

# tag: film-query
prefix rdf:<http://www.w3.org/1999/02/22-rdf-syntax-ns#>
prefix rdfs: <http://www.w3.org/2000/01/rdf-schema#>
prefix dbo: <http://dbpedia.org/ontology/>
prefix dbp: <http://dbpedia.org/property/>
prefix foaf: <http://xmlns.com/foaf/0.1/>
select distinct ?name ?description ?director ?comment ?homepage where {
?film       rdf:type dbo:Film;
            rdfs:label ?name;
            dbo:abstract ?description;
            dbo:director ?drc;
            rdfs:comment ?comment;
			foaf:homepage ?homepage.
?drc        rdfs:label ?director
FILTER((str(?name)="{{.Filmname}}")&&(lang(?description)="en")&&(lang(?comment)="en"))
}limit 1
`

//Kingdom of Heaven
// const filmQueries = `
// # Comment are ignored

// # tag: film-query

// `

const filmRecommend = `
# Comment are ignored

# tag: film-recommend
prefix rdf:<http://www.w3.org/1999/02/22-rdf-syntax-ns#>
prefix rdfs: <http://www.w3.org/2000/01/rdf-schema#>
prefix dbo: <http://dbpedia.org/ontology/>
prefix dbp: <http://dbpedia.org/property/>
prefix foaf: <http://xmlns.com/foaf/0.1/>
select distinct ?name where {
?drc        rdfs:label "{{.Director}}"@en.
?film		dbo:director ?drc.
?film       rdfs:label ?name
}limit 5
`

// type data struct {
// 	Name string `json:"name"`
// }

// type netResponse struct {
// 	Response map[string]string `json:"response"`
// }

func QueryBook(s string, itemName string, bankName string) (map[string][]rdf.Term, error) {
	f := bytes.NewBufferString(s)
	bank := sparql.LoadBank(f)

	q, err := bank.Prepare(bankName, struct{ Bookname string }{itemName})
	if err != nil {
		//fmt.Println(err, 0)
		return nil, err
	}
	//fmt.Println(q)
	repo, err := sparql.NewRepo("http://dbpedia.org/sparql")
	if err != nil {
		//fmt.Println(err, 1)
		return nil, err
	}
	res, err := repo.Query(q)
	if err != nil {
		fmt.Println(err, 2)
		return nil, err
	}
	rst := res.Bindings()
	//fmt.Println(rst)
	return rst, nil
}

func RecommendBook(s string, itemName string, bankName string) (map[string][]rdf.Term, error) {
	f := bytes.NewBufferString(s)
	bank := sparql.LoadBank(f)
	q, err := bank.Prepare(bankName, struct{ Author string }{itemName})
	if err != nil {
		//fmt.Println(err, 0)
		return nil, err
	}
	//fmt.Println(q)
	repo, err := sparql.NewRepo("http://dbpedia.org/sparql")
	if err != nil {
		//fmt.Println(err, 1)
		return nil, err
	}
	res, err := repo.Query(q)
	if err != nil {
		fmt.Println(err, 2)
		return nil, err
	}
	rst := res.Bindings()
	//fmt.Println(rst)
	return rst, nil
}

func QueryFilm(s string, itemName string, bankName string) (map[string][]rdf.Term, error) {
	f := bytes.NewBufferString(s)
	bank := sparql.LoadBank(f)
	q, err := bank.Prepare(bankName, struct{ Filmname string }{itemName})
	if err != nil {
		//fmt.Println(err, 0)
		return nil, err
	}
	fmt.Println(q)
	repo, err := sparql.NewRepo("http://dbpedia.org/sparql")
	if err != nil {
		//fmt.Println(err, 1)
		return nil, err
	}
	res, err := repo.Query(q)
	if err != nil {
		fmt.Println(err, 2)
		return nil, err
	}
	rst := res.Bindings()
	fmt.Println(rst)
	return rst, nil
}

func RecommendFilm(s string, itemName string, bankName string) (map[string][]rdf.Term, error) {
	f := bytes.NewBufferString(s)
	bank := sparql.LoadBank(f)
	q, err := bank.Prepare(bankName, struct{ Director string }{itemName})
	if err != nil {
		//fmt.Println(err, 0)
		return nil, err
	}
	//fmt.Println(q)
	repo, err := sparql.NewRepo("http://dbpedia.org/sparql")
	if err != nil {
		//fmt.Println(err, 1)
		return nil, err
	}
	res, err := repo.Query(q)
	if err != nil {
		fmt.Println(err, 2)
		return nil, err
	}
	rst := res.Bindings()
	fmt.Println(rst)
	return rst, nil
}

// func HttpResponse(w http.ResponseWriter, v interface{}) {
// 	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 	w.WriteHeader(http.StatusOK)
// 	content, err := json.Marshal(v)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	w.Write(content)
// }

func HelloServer(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world!\n")
}

func Query(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		fmt.Println("start1")
		// resultResponse := netResponse{}
		// resultResponse.Response = make(map[string]string)
		itemType := r.FormValue("itemtype")
		itemName := r.FormValue("itemname")
		if itemType == "图书" {
			rst1, err := QueryBook(bookQueries, itemName, "book-query")
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(rst1["author"][0].String())
			recommend1, err := RecommendBook(bookRecommend, rst1["author"][0].String(), "book-recommend")
			if err != nil {
				fmt.Println(err)
			}
			returnBook(w, rst1, recommend1)
		} else if itemType == "电影" {
			fmt.Println("start!")
			rst2, err := QueryFilm(filmQueries, itemName, "film-query")
			if err != nil {
				fmt.Println(err, 1)
			}
			fmt.Println(rst2["director"][0].String())
			recommend2, err := RecommendFilm(filmRecommend, rst2["director"][0].String(), "film-recommend")
			if err != nil {
				fmt.Println(err, 2)
			}
			returnFilm(w, rst2, recommend2)
		} else {

		}
	default:
	}
}

// func TestRecv(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case "POST":
// 		nameValue := r.FormValue("name")
// 		fmt.Println(nameValue)
// 		mydata := data{}
// 		mydata.Name = nameValue
// 		sendBuf, err := json.Marshal(mydata)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		HttpResponse(w, sendBuf)
// 		w.Write(sendBuf)
// 	default:
// 		return
// 	}
// }

func returnBook(w http.ResponseWriter, result map[string][]rdf.Term, recommend map[string][]rdf.Term) {
	io.WriteString(w, "<html>")
	io.WriteString(w, "<body>")
	io.WriteString(w, "<h1>Query Result<h1>")
	io.WriteString(w, "<table class=\"sparql\" border=\"1\">")
	io.WriteString(w, "<tr>")
	io.WriteString(w, "<th>name</th>")
	io.WriteString(w, "<th>author</th>")
	io.WriteString(w, "<th>thumbnail</th>")
	io.WriteString(w, "<th>country</th>")
	io.WriteString(w, "<th>comment</th>")
	io.WriteString(w, "<th>genre</th>")
	io.WriteString(w, "<th>title</th>")
	io.WriteString(w, "</tr>")
	for i := 0; i < len(result["name"]); i++ {
		io.WriteString(w, "<tr>")
		io.WriteString(w, fmt.Sprintf("<td><pre>%s</pre></td>", result["name"][i].String()))
		io.WriteString(w, fmt.Sprintf("<td><pre>%s</pre></td>", result["author"][i].String()))
		io.WriteString(w, fmt.Sprintf("<td><a href=\"%s\">%s</a></td>", result["thumbnail"][i].String(), result["thumbnail"][i].String()))
		io.WriteString(w, fmt.Sprintf("<td><pre>%s</pre></td>", result["country"][i].String()))
		io.WriteString(w, fmt.Sprintf("<td><pre>%s</pre></td>", result["comment"][i].String()))
		io.WriteString(w, fmt.Sprintf("<td><pre>%s</pre></td>", result["genre"][i].String()))
		io.WriteString(w, fmt.Sprintf("<td><pre>%s</pre></td>", result["title"][i].String()))
		io.WriteString(w, "</tr>")
	}
	io.WriteString(w, "</table>")
	io.WriteString(w, "<h1>Recommend Result<h1>")
	io.WriteString(w, "<table class=\"sparql\" border=\"1\">")
	io.WriteString(w, "<tr>")
	io.WriteString(w, "<th>book</th>")
	//io.WriteString(w, "<th>title</th>")
	io.WriteString(w, "</tr>")
	for i := 0; i < len(result["name"]); i++ {
		io.WriteString(w, "<tr>")
		//io.WriteString(w, fmt.Sprintf("<td><a href=\"%s\">%s</a></td>", result["book"][i].String(), result["book"][i].String()))
		io.WriteString(w, fmt.Sprintf("<td><pre>%s</pre></td>", result["name"][i].String()))
		io.WriteString(w, "</tr>")
	}
	io.WriteString(w, "</table>")
	io.WriteString(w, "</body>")
	io.WriteString(w, "</html>")
}

func returnFilm(w http.ResponseWriter, result map[string][]rdf.Term, recommend map[string][]rdf.Term) {
	//?name ?description ?director ?comment ?homepage
	io.WriteString(w, "<html>")
	io.WriteString(w, "<body>")
	io.WriteString(w, "<h1>Query Result<h1>")
	io.WriteString(w, "<table class=\"sparql\" border=\"1\">")
	io.WriteString(w, "<tr>")
	io.WriteString(w, "<th>name</th>")
	io.WriteString(w, "<th>description</th>")
	io.WriteString(w, "<th>director</th>")
	io.WriteString(w, "<th>comment</th>")
	io.WriteString(w, "<th>homepage</th>")
	//io.WriteString(w, "<th>Year</th>")
	//io.WriteString(w, "<th>film</th>")
	io.WriteString(w, "</tr>")
	for i := 0; i < len(result["name"]); i++ {
		io.WriteString(w, "<tr>")
		io.WriteString(w, fmt.Sprintf("<td><pre>%s</pre></td>", result["name"][i].String()))
		io.WriteString(w, fmt.Sprintf("<td><pre>%s</pre></td>", result["description"][i].String()))
		io.WriteString(w, fmt.Sprintf("<td><pre>%s</pre></td>", result["director"][i].String()))
		io.WriteString(w, fmt.Sprintf("<td><pre>%s</pre></td>", result["comment"][i].String()))
		io.WriteString(w, fmt.Sprintf("<td><pre>%s</pre></td>", result["homepage"][i].String()))
		//io.WriteString(w, fmt.Sprintf("<td><pre>%s</pre></td>", result["Year"][i].String()))
		//io.WriteString(w, fmt.Sprintf("<td><pre>%s</pre></td>", result["Year"][i].String()))
		//io.WriteString(w, fmt.Sprintf("<td><a href=\"%s\">%s</a></td>", result["thumbnail"][i].String(), result["thumbnail"][i].String()))
		io.WriteString(w, "</tr>")
	}
	io.WriteString(w, "</table>")
	io.WriteString(w, "<h1>Recommend Result<h1>")
	io.WriteString(w, "<table class=\"sparql\" border=\"1\">")
	io.WriteString(w, "<tr>")
	//io.WriteString(w, "<th>book</th>")
	io.WriteString(w, "<th>name</th>")
	io.WriteString(w, "</tr>")
	for i := 0; i < len(result["name"]); i++ {
		io.WriteString(w, "<tr>")
		//io.WriteString(w, fmt.Sprintf("<td><a href=\"%s\">%s</a></td>", result["book"][i].String(), result["book"][i].String()))
		io.WriteString(w, fmt.Sprintf("<td><pre>%s</pre></td>", result["name"][i].String()))
		io.WriteString(w, "</tr>")
	}
	io.WriteString(w, "</table>")
	io.WriteString(w, "</body>")
	io.WriteString(w, "</html>")
}

func main() {
	http.HandleFunc("/hello", HelloServer)
	http.HandleFunc("/query", Query)
	log.Fatal(http.ListenAndServe(":12345", nil))
}
