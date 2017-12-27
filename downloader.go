package bayamo

type Downloader interface {}

type fileDownloader struct {}

func NewDownloader() Downloader {
  return new(fileDownloader)
}
