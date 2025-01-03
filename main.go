// main.go
package main

import (
	"log"
	"net/http"
	"fmt"
	"html/template"
	"path/filepath"
)

func main() {
	const PORT = ":3030"
	fs := http.FileServer(http.Dir("./static"))
	// http.Handle("/", fs)
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", serveTemplate)

	fmt.Printf("Listening on %s...\n", PORT)
	fmt.Printf("Server starting on http://localhost%s/templated_page.html\n",PORT)
	// fmt.Printf("Server starting on http://localhost%s\n", PORT)

	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Fatal(err)
	}
	
	
}


func serveTemplate(w http.ResponseWriter, r *http.Request) {
	lp := filepath.Join("templates", "layout.html")
	fp := filepath.Join("templates", filepath.Clean(r.URL.Path))

	tmpl, _ := template.ParseFiles(lp, fp)
	tmpl.ExecuteTemplate(w, "layout", nil)
}

// // main.go
// package main

// import (
// 	"log"
// 	"net/http"
// 	"fmt"
// 	// "html/template"
// 	// "path/filepath"
// )

// func main() {
// 	const PORT = ":3030"
// 	fs := http.FileServer(http.Dir("./static/images"))
// 	http.Handle("/", fs)
// 	// http.Handle("/static/", http.StripPrefix("/static/", fs))
// 	// http.HandleFunc("/", serveTemplate)

// 	fmt.Printf("Listening on %s...\n", PORT)
// 	fmt.Printf("Server starting on http://localhost%s/templated_page.html\n",PORT)
// 	// fmt.Printf("Server starting on http://localhost%s\n", PORT)

// 	err := http.ListenAndServe(PORT, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
	
	
// }
