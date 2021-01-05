package main

import (
    "net/http"
    "os"
    "strconv"
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
            factors, _ := factorize(num, PRIMES)
            renderTemplate(w, "form", factors)
        }
    }
}

func main() {
    PORT := os.Getenv("PORT")
    PRIMES := primesUnder(10e6)

    http.HandleFunc("/", handle)
    http.ListenAndServe(":"+PORT, nil)
}
