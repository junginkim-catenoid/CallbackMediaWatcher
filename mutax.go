package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

type WatcherRequest2 struct {
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
	mux.HandleFunc("/register", PostRegisterHandler2).Methods("POST")
	mux.HandleFunc("/complete", PostCompleteHandler1).Methods("POST")

	return mux
}

func PostRegisterHandler2(w http.ResponseWriter, r *http.Request) {
	var watcher WatcherRequest

	b, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatal(err)
	}

	err = r.ParseForm()
	if err != nil {
		panic(err)
	}

	fmt.Println("request.Form::")
	for key, value := range r.Form {
		fmt.Printf("Key:%s, Value:%s\n", key, value)
	}
	fmt.Println("\nrequest.PostForm::")
	for key, value := range r.PostForm {
		fmt.Printf("Key:%s, Value:%s\n", key, value)
	}

	fmt.Printf("\nName field in Form:%s\n", r.Form["api_key"])
	fmt.Printf("\nName field in PostForm:%s\n", r.PostForm["api_key"])
	fmt.Printf("\nHobbies field in FormValue:%s\n", r.FormValue("api_key"))

	w.WriteHeader(http.StatusOK)

	fmt.Println(watcher)
	json.NewEncoder(w).Encode(map[string]string{"msg": "ok"})

	fmt.Println(string(b))
}

func PostCompleteHandler1(w http.ResponseWriter, r *http.Request) {
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

func mutax() {

	println("Hello CallbackMediaWatcher app")

	http.ListenAndServe(":3002", MakeWebHandler())
}
