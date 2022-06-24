package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type WatcherRequest struct {
	APIKey       string `json:"api_key"`
	APIReference string `json:"api_reference"`
}

type ReqWatcherFiles []struct {
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
}

type Dto struct {
	Error       int         `json:"error"`
	Message     string      `json:"message"`
	WatcherFile ContentDto  `json:"watcher_file"`
	Result      ContentsDto `json:"result"`
}

type ContentDto struct {
	Error   int           `json:"error"`
	Message string        `json:"message"`
	Result  []WatcherFile `json:"result"`
}

type ContentsDto struct {
	ErrorCode    int          `json:"error_code"`
	ErrorDetail  string       `json:"error_detail"`
	WatcherFiles []ContentDto `json:"watcher_files"`
}

type WatcherFile struct {
	ErrorCode                  int    `json:"error_code"`
	ErrorDetail                string `json:"error_detail"`
	ContentProviderKey         string `json:"content_provider_key"`
	Key                        string `json:"key"`
	MediaContentId             string `json:"media_content_id"`
	ContentPath                string `json:"content_path"`
	UploadPath                 string `json:"upload_path"`
	IsAudioFile                int    `json:"is_audio_file"`
	ChecksumType               int    `json:"checksum_type"`
	SnapshotPath               string `json:"snapshot_path"`
	PhysicalPath               string `json:"physical_path"`
	DeleteWatcherFileUploadUrl string `json:"delete_watcher_file_upload_url"`
}

func PostRegisterHandler(w http.ResponseWriter, r *http.Request) {

	var watcher WatcherRequest
	var reqWatcher ReqWatcherFiles
	var dto Dto

	//headerContentType := r.Header.Get("Content-Type")
	//if strings.Contains("application/x-www-form-urlencoded;", headerContentType) {
	//	w.WriteHeader(http.StatusUnsupportedMediaType)
	//	return
	//}

	err := r.ParseForm()
	if err != nil {
		log.Fatal(err.Error())
	}
	watcher.APIKey = r.FormValue("api_key")
	watcher.APIReference = r.FormValue("api_reference")

	err = json.Unmarshal([]byte(r.FormValue("watcher_files")), &reqWatcher)
	if err != nil {
		log.Fatal("ee")
	}

	var data []WatcherFile
	err = json.Unmarshal([]byte(r.FormValue("watcher_files")), &data)
	if err != nil {
		log.Fatal("dd")
	}

	contentDtos := ContentDto{Error: 0, Message: "ok", Result: data}

	fmt.Println(contentDtos)
	//strResult, err := json.Marshal(result)
	//if err != nil {
	//	log.Fatal("d")
	//}

	json.NewEncoder(w).Encode(dto)
	//json.NewEncoder(w).Encode(watcherFiels)

	return
}

func PostCompleteHandler(w http.ResponseWriter, r *http.Request) {
	headerContentTtype := r.Header.Get("Content-Type")
	if strings.Contains("application/x-www-form-urlencoded;", headerContentTtype) {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}
	r.ParseForm()
	fmt.Println("\nrequest.PostForm::")
	for key, value := range r.PostForm {
		fmt.Printf("Key:%s, Value:%s\n", key, value)
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "ok",
	})

	w.WriteHeader(200)
	return
}

func PostDeleteHandler(w http.ResponseWriter, r *http.Request) {
	headerContentTtype := r.Header.Get("Content-Type")
	if strings.Contains("application/x-www-form-urlencoded;", headerContentTtype) {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}
	r.ParseForm()

	json.NewEncoder(w).Encode(map[string]string{
		"error":   "",
		"message": "ok",
	})

	w.WriteHeader(200)
	return
}

func main() {
	http.Handle("/register", http.HandlerFunc(PostRegisterHandler))
	http.Handle("/complete", http.HandlerFunc(PostCompleteHandler))
	http.Handle("/delete", http.HandlerFunc(PostDeleteHandler))

	http.ListenAndServe(":3002", nil)
}
