package bayamo

import (
  "net/http"
  "net/http/httptest"
  "os"
  "testing"
)

func TestCreateDownloader(t *testing.T){
  var dwn Downloader = NewDownloader("out")

  if dwn == nil {
    t.Error("Found nil while creating a downloader")
  }
}

func TestFileDownloaderGenerateName(t *testing.T){
  var dwn Downloader = NewDownloader("tmp_dir")
  var filename string = "example.txt"
  var path string = dwn.generateName(filename)

  if path == "" {
    t.Error(path)
  }

  os.RemoveAll("tmp_dir")
}

func TestFileDownloaderGet(t *testing.T){
  var dwn Downloader = NewDownloader("out")
  var err error
  var serv *httptest.Server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request){
    http.ServeFile(w, req, "Dockerfile")
  }))
  defer serv.Close()

  _, err = dwn.Get(serv.URL + "/Dockerfile?query=string")
  if err != nil {
    t.Error("Error downloading a file: " + err.Error())
  }

  os.RemoveAll("out")
}

func TestFileDownloaderGet404(t *testing.T){
  var dwn Downloader = NewDownloader("out")
  var err error
  var serv *httptest.Server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request){
    w.WriteHeader(http.StatusNotFound)
    http.StatusText(http.StatusNotFound)
  }))
  defer serv.Close()

  _, err = dwn.Get(serv.URL + "/not-found")
  if err == nil {
    t.Error("Should not proceed if not found")
  }

  os.RemoveAll("out")
}

func TestFileDownloaderExtractFilenameFromURL(t *testing.T){
  var dwn Downloader = NewDownloader("out")
  var serv *httptest.Server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request){
    http.ServeFile(w, req, "Dockerfile")
  }))
  defer serv.Close()

  dwn.Get(serv.URL + "/Dockerfile")
  if dwn.(*fileDownloader).Filename != "Dockerfile" {
    t.Error("Wrong filename, got " + dwn.(*fileDownloader).Filename + " while expecting Dockerfile")
  }

  os.RemoveAll("out")
}
