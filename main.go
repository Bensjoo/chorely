package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintln(w, `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>HTMX Hello World</title>
    <script src="https://unpkg.com/htmx.org"></script>
</head>
<body>
    <h1 hx-get="/hello" hx-trigger="load" hx-swap="outerHTML">Loading...</h1>
</body>
</html>
        `)
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "<h1>Chorely!</h1>")
	})

	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}
