package main

import (
	b32 "encoding/base32"
	"fmt"
	"net/http"
	"text/template"

	box "box/Print"
)

var (
	Tp1 *template.Template
	er  error
)

type dataValu struct {
	Text string
	Font string
	Res  string
}

func main() {
	Tp1, er = template.ParseFiles("templates/pageDefoult.html")

	http.Handle("/styles.css", http.FileServer(http.Dir("./templates/css/")))
	http.HandleFunc("/", DefoultHundlr)
	http.HandleFunc("/ascii-art", AsciiHundler)

	http.HandleFunc("/set-cookie", box.SetCookieHandler)
	http.HandleFunc("/get-cookie", box.GetCookieHandler)

	http.HandleFunc("/download", box.DownloadStringHandler)

	fmt.Println("Server started on port 8080...")
	fmt.Println(http.ListenAndServe(":8080", nil))
}

func DefoultHundlr(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		http.ServeFile(w, r, "./templates/errors/404.html")
		return
	}

	var data dataValu

	if r.Method == http.MethodGet {
		data.Text = ""
		if er != nil {
			w.WriteHeader(http.StatusInternalServerError)
			http.ServeFile(w, r, "./templates/errors/500.html")
			return
		}

		Tp1.Execute(w, data)

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		http.ServeFile(w, r, "./templates/errors/405.html")
		return
	}
}

func AsciiHundler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art" {
		w.WriteHeader(http.StatusNotFound)
		http.ServeFile(w, r, "./templates/errors/404.html")
		return
	}

	var data dataValu
	if r.Method == http.MethodPost {
		r.ParseForm()
		fontValue := r.Form.Get("font")
		textValue := "\n" + r.FormValue("inputText")

		if box.CheckArgs(fontValue) == "" || textValue == "" || len(textValue) > 2000 {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "./templates/errors/400.html")
			return
		}

		Str, err := box.ChangeToAsciiArt(textValue, fontValue)

		ResCookie := http.Cookie{
			Name:   "lastRes",
			Value:  b32.StdEncoding.EncodeToString([]byte(Str)),
			MaxAge: 3600 * 24, // Cookie expires 24h
		}

		http.SetCookie(w, &ResCookie)

		data.Text = textValue
		data.Font = fontValue
		data.Res = Str

		if er != nil || err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			http.ServeFile(w, r, "./templates/errors/500.html")
			return
		}
		Tp1.Execute(w, data)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		http.ServeFile(w, r, "./templates/errors/405.html")
		return
	}
}
