package main

import (
    "fmt"
    "html/template"
    "net/http"
    "os"
    "strconv"
)

func renderTemplate(w http.ResponseWriter, tmpl string, factors map[int]int) {
    t, err := template.ParseFiles(tmpl + ".html")
    t.Execute(w, factors)
    if err != nil {
        fmt.Fprint(w, err)
    }
}

func handle(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        renderTemplate(w, "form", nil)
    } else {
        num, err := strconv.Atoi(r.FormValue("number"))
        if err != nil {
            renderTemplate(w, "form", nil)
        } else {
            primes := primesUnder(10e6)
            factors, _ := factorize(num, primes)
            renderTemplate(w, "form", factors)
        }
    }
}

func main() {
    PORT := os.Getenv("PORT")

    http.HandleFunc("/", handle)
    http.ListenAndServe(":"+PORT, nil)
}
