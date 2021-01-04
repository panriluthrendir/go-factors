package module github.com/panriluthrendir/go-test-heroku

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	PORT := os.Getenv("PORT")
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World!")
	})

	http.ListenAndServe(":" + PORT, nil)
}
