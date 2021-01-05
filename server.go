package main

import (
    "fmt"
    "html/template"
    "net/http"
    "os"
    "strconv"
)

func renderTemplate(w http.ResponseWriter, tmpl string, factors string) {
    err := template.ExecuteTemplate(w, tmpl+"html", factors)
    if err != nil {
        fmt.Fprint(w, err)
    } 
}

func handle(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        renderTemplate(w, "form", "")
    } else {
        num, err := strconv.Atoi(r.FormValue("number"))
        if err != nil {
            fmt.Fprint(w, err)
        } else {
            primes := primesUnder(10e6)
            factors, _ := factorize(num, primes)
            renderTemplate(w, "form", fmt.Sprint(factors))
        }
    }
}

func main() {
    PORT := os.Getenv("PORT")
    
    var templates = template.Must(template.ParseFiles("form.html"))

    http.HandleFunc("/", handle)
    http.ListenAndServe(":"+PORT, nil)
}
