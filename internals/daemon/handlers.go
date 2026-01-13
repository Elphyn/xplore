package daemon

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// POST request for api/files gives this json request
// in future there would be parameters for serverside soring probably
type DirInfoReq struct {
	Path string `json:"path"`
}

// Types for response on api/files
type FileInfo struct {
	Name  string `json:"name"`
	IsDir bool   `json:"isDir"`
}

type DirInfoResponse struct {
	Path  string     `json:"path"`
	Files []FileInfo `json:"files"`
}

func HandleHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello from the server")
}

func HandleDirInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var parsed DirInfoReq
	err := json.NewDecoder(r.Body).Decode(&parsed)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if parsed.Path == "" {
		http.Error(w, "Invalid PATH", http.StatusBadRequest)
		return
	}

	files, err := os.ReadDir(parsed.Path)
	if err != nil {
		panic(err)
	}

	reqResponse := DirInfoResponse{
		Path: parsed.Path,
	}
	for _, file := range files {
		reqResponse.Files = append(reqResponse.Files, FileInfo{
			IsDir: file.IsDir(),
			Name:  file.Name(),
		})
	}

	if err := json.NewEncoder(w).Encode(reqResponse); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
