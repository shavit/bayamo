package bayamo

import (
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

  _, err = dwn.Get("")
  if err != nil {
    t.Error("Error downloading a file: " + err.Error())
  }
}
