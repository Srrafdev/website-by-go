package box

import (
	b32 "encoding/base32"
	"fmt"
	"net/http"
	"strconv"
)

func DownloadStringHandler(w http.ResponseWriter, r *http.Request) {
	ResCookie, errCooke := r.Cookie("lastRes")
	if errCooke != nil {
		fmt.Println("ERROR-COOKIE-DOWNLOAD", errCooke)
	}

	content, errB32 := b32.StdEncoding.DecodeString(ResCookie.Value)
	if errB32 != nil {
		fmt.Println("ERROR-B32-DOWNLOAD", errB32)
	}

	// Set headers
	w.Header().Set("Content-Disposition", "attachment; filename=txt-asci-art.txt")
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Length", strconv.Itoa(len(content)))

	// Write the content to the response
	w.Write(content)

	_, err := w.Write(content)
	if err != nil {
		http.Error(w, "Error downloading content", http.StatusInternalServerError)
		return
	}
}
