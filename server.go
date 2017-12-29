package bayamo

import (
  "errors"
  "io"
  "net"
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

  // addClient adds a client into the clients map
  addClient(c io.ReadWriteCloser)

  // Write start a download task and write the response into a file
  Write(ctx context.Context, job *_proto.DownloadJob) (res *_proto.DownloadJobResult, err error)
}

type clientConnection struct {
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
  var conn net.Conn
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
  go s.grpcServer.Serve(ln)
  for {
    conn, err = ln.Accept()
    if err != nil {
      conn.Close()
      return errors.New("Error serving grpc")
    }
    println("New connection")
    s.addClient(conn)
  }

  return err
}

// addClient adds a client into the clients map
func (s *server) addClient(c io.ReadWriteCloser) {}

// Write start a download task and write the response into a file
func (s *server) Write(ctx context.Context, job *_proto.DownloadJob) (res *_proto.DownloadJobResult, err error){
  var outputPath string = job.OutputDir

  println("Starting a download task:", job.Url)

  res = &_proto.DownloadJobResult{Ok: true,
    OutputPath: outputPath,
  }

  return res, err
}
