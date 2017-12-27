package bayamo

import (
  "io/ioutil"
  "net/http"
  "os"
)

type Downloader interface {
  // Download data from url into a file
  Get(url string) (f *os.File, err error)
}

//
//  fileDownloader
//
type fileDownloader struct {}

// Download data from url into a file
func (dwn *fileDownloader) Get(url string) (f *os.File, err error){
  var res *http.Response
  var body []byte

  res, err = http.Get(url)
  if err != nil {
    return f, err
  }
  defer res.Body.Close()

  body, err = ioutil.ReadAll(res.Body)
  if err != nil {
    return f, err
  }

  f, err = os.Create("test_file.txt")
  if err != nil {
    return f, err
  }

  _, err = f.Write(body)

  return f, err
}

//
//  Downloader
//
func NewDownloader() Downloader {
  return new(fileDownloader)
}
