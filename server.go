package bayamo

import (
  "errors"
  "io"
  "net"
  "os"
  "sync"

  "google.golang.org/grpc"
  "google.golang.org/grpc/credentials"
  "golang.org/x/net/context"
  _proto "github.com/shavit/bayamo/proto"
)

type Server interface{
  // AddCredentials adds credentials to the server
  AddCredentials(crtFile, keyFile string) (err error)

  // Start starts the server
  Start() (err error)

  // Write start a download task and write the response into a file
  Write(ctx context.Context, job *_proto.DownloadJob) (res *_proto.DownloadJobResult, err error)
}

type clientConnection struct {
  id int
  conn io.ReadWriteCloser
}

type server struct {
  grpcServer *grpc.Server
  clients map[chan<- string]*clientConnection
  Creds credentials.TransportCredentials
  *sync.RWMutex
}

func NewServer() (s *server){
  return &server{
    clients: make(map[chan<- string]*clientConnection),
    RWMutex: new(sync.RWMutex),
  }
}

// AddCredentials adds credentials to the server
func (s *server) AddCredentials(crtFile, keyFile string) (err error){
  s.Creds, err = credentials.NewServerTLSFromFile(crtFile, keyFile)
  if err != nil {
    return err
  }

  return err
}

func (s *server) Start() (err error){
  var ln net.Listener
  var host string = "0.0.0.0:2400"
  // var opts grpc.ServerOption = grpc.Creds(s.Creds)
  // s.grpcServer = grpc.NewServer(opts)
  s.grpcServer = grpc.NewServer()

  ln, err = net.Listen("tcp", host)
  if err != nil {
    return errors.New("Error listening to " + host)
  }

  _proto.RegisterDownloaderServiceServer(s.grpcServer, s)

  println("Listening on", host)
  s.grpcServer.Serve(ln)

  return err
}

// Write start a download task and write the response into a file
func (s *server) Write(ctx context.Context, job *_proto.DownloadJob) (res *_proto.DownloadJobResult, err error){
  var outputPath string = job.OutputDir + "/downloads"
  var f *os.File
  var info os.FileInfo
  var dwn Downloader = NewDownloader(outputPath)

  res = &_proto.DownloadJobResult{Ok: false,
    OutputPath: outputPath,
  }

  println("Starting a download task:", job.Url)
  f, err = dwn.Get(job.Url)
  if err != nil {
    println("Error downloading from", job.Url, ".", err.Error())
    return res, err
  } else {
    res.Ok = true
  }

  info, err = f.Stat()
  if err != nil {
    return res, err
  }
  println("Filename:", info.Name(), "| File size:", info.Size())
  f.Close()

  return res, err
}
