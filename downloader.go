package bayamo

import (
  "encoding/hex"
  "errors"
  "fmt"
  "io/ioutil"
  "net/http"
  "os"
  "math/rand"
  "net/url"
  "strings"
  "time"
)

type Downloader interface {
  // generateName generates a unique filename
  generateName(name string) (path string)

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

// generateName generates a unique filename
func (dwn *fileDownloader) generateName(name string) (path string){
  var rndData []byte = make([]byte, 12)
  var fileName []byte = []byte(name)[0:4]
  var fileParts []string = strings.Split(name, ".")
  var ext string = fileParts[len(fileParts)-1]
  y, m, d := time.Now().Date()
  var dirPath string = fmt.Sprintf("%s/%d/%d/%d", dwn.outputDir, y, m, d)

  // Create the sub directories
  os.MkdirAll(dirPath, os.ModePerm)

  // Generate a file name
  rand.Read(rndData)
  path = fmt.Sprintf("%s/%s-%s.%s", dirPath, hex.EncodeToString(fileName),
    hex.EncodeToString(rndData), ext)

  return path
}

// Download data from url into a file
func (dwn *fileDownloader) Get(rawUrl string) (f *os.File, err error){
  var res *http.Response
  var body []byte
  var filePath string
  var _url *url.URL
  var urlParts []string

  _url, err = url.Parse(rawUrl)
  if err != nil {
    return f, err
  }

  // Sanitize the file name and path
  urlParts = strings.Split(_url.EscapedPath(), "/")
  dwn.Filename = urlParts[len(urlParts)-1]
  filePath = dwn.generateName(dwn.Filename)

  // Start the download
  res, err = http.Get(rawUrl)
  if err != nil {
    return f, err
  }
  if res.StatusCode != 200 {
    return f, errors.New(fmt.Sprintf("Resource unavailable. Expected 200, got %s", res.Status))
  }
  defer res.Body.Close()

  body, err = ioutil.ReadAll(res.Body)
  if err != nil {
    return f, err
  }

  // If O_CREAT and O_EXCL are set, open() shall fail if the file exists
  f, err = os.OpenFile(filePath, os.O_CREATE|os.O_RDWR|os.O_EXCL, 0666)
  if err != nil {
    println("Error openning a file")
    return f, err
  }

  // Write the data to disk
  _, err = f.Write(body)
  if err == nil {
    println("Success writing a file")
    // Report
  } else {
    panic(err)
  }

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
