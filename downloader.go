package bayamo

import (
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
  return f, err
}

//
//  Downloader
//
func NewDownloader() Downloader {
  return new(fileDownloader)
}
