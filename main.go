package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type WatcherRequest struct {
	APIKey       string `json:"api_key"`
	APIReference string `json:"api_reference"`
	WatcherFiles []struct {
		Kind               int     `json:"kind"`
		ModuleKey          string  `json:"module_key"`
		IsAudioFile        int     `json:"is_audio_file"`
		ChecksumType       int     `json:"checksum_type"`
		UseEncryption      int     `json:"use_encryption"`
		PhysicalPath       string  `json:"physical_path"`
		UploadPath         string  `json:"upload_path"`
		Filesize           int     `json:"filesize"`
		FilesizeStr        string  `json:"filesize_str"`
		Lastmodified       int     `json:"lastmodified"`
		LastmodifiedStr    string  `json:"lastmodified_str"`
		ContentProviderKey string  `json:"content_provider_key"`
		Key                string  `json:"key"`
		ContentPath        string  `json:"content_path"`
		SnapshotPath       string  `json:"snapshot_path"`
		Title              string  `json:"title"`
		Md5                string  `json:"md5"`
		Format             string  `json:"format"`
		Duration           float64 `json:"duration"`
		PassthroughAhead   int     `json:"passthrough_ahead"`
		ChannelKeys        string  `json:"channel_keys"`
	} `json:"watcher_files"`
}

func MakeWebHandler() http.Handler {
	mux := mux.NewRouter()
	mux.HandleFunc("/register", PostRegisterHandler).Methods("POST")
	mux.HandleFunc("/complete", PostCompleteHandler).Methods("POST")

	return mux
}

func PostRegisterHandler(w http.ResponseWriter, r *http.Request) {
	var watcher WatcherRequest
	err := json.NewDecoder(r.Body).Decode(&watcher)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	//lastId++
	//student.Id = lastId
	//students[lastId] = student
	//w.WriteHeader(http.StatusCreated)
}

func PostCompleteHandler(w http.ResponseWriter, r *http.Request) {
	//var student Student
	//err := json.NewDecoder(r.Body).Decode(&student)
	//if err != nil {
	//	w.WriteHeader(http.StatusBadRequest)
	//	return
	//}
	//lastId++
	//student.Id = lastId
	//students[lastId] = student
	//w.WriteHeader(http.StatusCreated)
}

func main() {

	println("Hello CallbackMediaWatcher app")

	http.ListenAndServe(":3002", MakeWebHandler())
}
