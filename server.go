package main

import (
    "fmt"
    "net/http"
    "os"
    "strconv"
    "encoding/json"
)

func main() {
    PORT := os.Getenv("PORT")
    PRIMES := primesUnder(10e5)

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        keys, present := r.URL.Query()["num"]
        if present {
            num, err := strconv.Atoi(keys[0])
            if err != nil {
                fmt.Fprint(w, err)
            } else {
                factors, _ := factorize(num, PRIMES)
                fmt.Fprint(w, string(json.Marshal(factors)))
            }
        }
    })

    http.ListenAndServe(":"+PORT, nil)
}
