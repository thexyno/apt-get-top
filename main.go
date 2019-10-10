package main

import (
	"log"
	"fmt"
        "strings"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `#!/bin/sh
# This is a joke
# haha
echo "Presented by apt-get.top, your TOP apt installer"
apt-get install -y%s
                `, strings.ReplaceAll(r.URL.Path, "/", " "))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
