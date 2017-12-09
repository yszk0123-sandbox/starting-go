package main

import (
	"io"
	"net/http"
)

func helloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, `
<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>hello</title>
</head>
<body>
<h1>Hello, world!</h1>
</body>
</html>
`)
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":3000", nil)
}
