package main

import (
    "net/http"
    "os"
    "strconv"
    "template"
)

func renderTemplate(w http.ResponseWriter, tmpl string, factors map[int]int) {
    t, _ := template.ParseFiles(tmpl + ".html")
    t.Execute(w, factors)
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
