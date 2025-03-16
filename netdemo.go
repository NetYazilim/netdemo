package netdemo

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
)

func HIndex(w http.ResponseWriter, r *http.Request) {
	var d struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		fmt.Fprint(w, `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Net Demo</title>
</head>
<body>
    <h5>Hello, World!</h5>
    <ul>
        <li><a href="/env">envaironment variables</a></li>
        <li><a href="/header">request header info</a></li>
        <li><a href="/user">user info</a></li>
    </ul>
</body>
</html>
`)
		return
	}
	if d.Name == "" {
		fmt.Fprint(w, "Hello, World!")
		return
	}
	fmt.Fprintf(w, "Hello, %s!", html.EscapeString(d.Name))
}
