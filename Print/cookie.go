package box

import (
	"fmt"
	"net/http"
)

func SetCookieHandler(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:   "my_cookie",
		Value:  "cookie_value",
	}
	http.SetCookie(w, cookie)
	fmt.Fprintf(w, "Cookie has been set")
}

func GetCookieHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("my_cookie")
	if err != nil {
		if err == http.ErrNoCookie {
			fmt.Fprintf(w, "Cookie not found")
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Cookie value: %s", cookie.Value)
}
