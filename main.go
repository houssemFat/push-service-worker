package main


import (
  "net/http"
  "html/template"

)

// the cache directory
const cacheDir = "./cache"

var templates = template.Must(template.ParseFiles(
        "./templates/index.html"))

func handler(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("./templates/index.html")
    t.Execute(w, nil)
    sendMail()
}

func main() {
    h := http.NewServeMux()
    h.Handle("/push/images/", http.StripPrefix("/push/images/", http.FileServer(http.Dir("./push/images/"))))
    h.Handle("/push/", http.StripPrefix("/push/", http.FileServer(http.Dir("./push/"))))
    h.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./public/"))))
    h.HandleFunc("/", handler)
    http.ListenAndServe(":8080", h)
}
