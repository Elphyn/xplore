package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"rxplore/internals/daemon"
	"rxplore/internals/tui"
)

func main() {
	info := daemon.DirInfoReq{
		Path:         "/home/vlad/Documents/",
		FilterHidden: true,
	}
	jsonBytes, err := json.Marshal(info)
	if err != nil {
		panic(err)
	}

	resp, err := http.Post("http://127.0.0.1:8080/api/files", "application/json", bytes.NewBuffer(jsonBytes))
	if err != nil {
		panic(err)
	}

	var files daemon.DirInfoPacked
	json.NewDecoder(resp.Body).Decode(&files)

	tui.StartTUI(files.Files)
}
