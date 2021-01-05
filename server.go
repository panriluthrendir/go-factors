package main

import (
    "fmt"
    "html/template"
    "net/http"
    "os"
    "strconv"
)

var PORT = os.Getenv("PORT")
var templates = template.Must(template.ParseFiles("form.html"))
var primes = primesUnder(10e6)

func renderTemplate(w http.ResponseWriter, tmpl string, factors string) {
    err := templates.ExecuteTemplate(w, tmpl+".html", factors)
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
            factors, _ := factorize(num, primes)
            renderTemplate(w, "form", formatFactors(num, factors))
        }
    }
}

func formatFactors(n int, factors map[int]int) string {
    result := fmt.Sprint("%d = ", n)
    
    for p, deg := range factors {
        if deg > 1 {
            result += fmt.Sprintf(" * %d^%d", p, deg)
        } else {
            result += fmt.Sprintf(" * %d", p)
        }                              
    }
    return result                             
}

func main() {
    http.HandleFunc("/", handle)
    http.ListenAndServe(":"+PORT, nil)
}
