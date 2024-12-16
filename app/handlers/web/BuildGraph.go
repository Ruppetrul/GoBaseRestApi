package web

import (
	"firstRest/actions"
	"fmt"
	"net/http"
)

func BuildGraph(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")

	svg, err := actions.Build("BTCUSD")
	if err != nil {
		fmt.Println("Error building records:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	_, err2 := w.Write([]byte("<HTML><BODY>" + svg + "</BODY></HTML>"))

	if err2 != nil {
		fmt.Println("Error marshaling records:", err2)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
