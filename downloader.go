package bayamo

import (
  "io/ioutil"
  "net/http"
  "os"
  "strings"
)

type Downloader interface {
  // Download data from url into a file
  Get(url string) (f *os.File, err error)
}

type configuration struct {}

//
//  fileDownloader
//
type fileDownloader struct {
  outputDir string
  Filename string
}

// Download data from url into a file
func (dwn *fileDownloader) Get(url string) (f *os.File, err error){
  var res *http.Response
  var body []byte

  var urlParts []string = strings.Split(url, "/")
  dwn.Filename = urlParts[len(urlParts)-1]

  res, err = http.Get(url)
  if err != nil {
    return f, err
  }
  defer res.Body.Close()

  body, err = ioutil.ReadAll(res.Body)
  if err != nil {
    return f, err
  }

  f, err = os.Create(dwn.outputDir + "/" + dwn.Filename)
  if err != nil {
    return f, err
  }

  _, err = f.Write(body)

  return f, err
}

//
//  Downloader
//
func NewDownloader(outputDir string) Downloader {
  os.MkdirAll(outputDir, os.ModePerm)

  return &fileDownloader{
    outputDir: outputDir,
  }
}
