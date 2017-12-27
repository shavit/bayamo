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
