package bayamo

import (
  "net/http"
  "net/http/httptest"
  "os"
  "testing"
)

func TestCreateDownloader(t *testing.T){
  var dwn Downloader = NewDownloader()

  if dwn == nil {
    t.Error("Found nil while creating a downloader")
  }
}

func TestFileDownloaderGet(t *testing.T){
  var dwn Downloader = NewDownloader()
  var err error
  var serv *httptest.Server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request){
    http.ServeFile(w, req, "Dockerfile")
  }))
  defer serv.Close()

  _, err = dwn.Get(serv.URL + "/Dockerfile")
  if err != nil {
    t.Error("Error downloading a file: " + err.Error())
  }

  os.Remove("test_file.txt")
}
