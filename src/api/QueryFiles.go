package api

import (
	"net/http"
	"encoding/json"
	"fmt"
	"os"
	"steam"
)

func StreamsHandler(w http.ResponseWriter, r *http.Request) {
	files, err := steam.QueryFiles("sfp")
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	output, err := json.Marshal(files)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshalling json: %q\n", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}
