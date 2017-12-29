package bayamo

import (
  "bufio"
  "errors"
  "os"
  "os/signal"
  "syscall"
  "time"

  "google.golang.org/grpc"
  context "golang.org/x/net/context"
  _proto "github.com/shavit/bayamo/proto"
)

type client struct {
  downloadClient _proto.DownloaderServiceClient
}

func NewClient() (c *client){
  return new(client)
}

func (c *client) Dial() (err error){
  var ch chan os.Signal = make(chan os.Signal)
  var conn *grpc.ClientConn
  var timeout time.Duration = time.Duration(10 * time.Millisecond)
  var host string = os.Getenv("SERVER_HOSTNAME")

  conn, err = grpc.Dial(host, grpc.WithTimeout(timeout), grpc.WithInsecure())
  if err != nil {
    return errors.New("Error dialing to remote server " + host)
  }
  defer conn.Close()
  go func() {
    signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
    <-ch
    println("\n\nDisconnecting from the server\n")
    conn.Close()
    os.Exit(0)
  }()

  c.downloadClient = _proto.NewDownloaderServiceClient(conn)

  // Choose a job
  c.typeToSend()

  return err
}

func (c *client) typeToSend() {
  var err error
  var res *_proto.DownloadJobResult
  var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
  var job *_proto.DownloadJob

  print("> ")
  for scanner.Scan(){
    job = &_proto.DownloadJob{
      Url: scanner.Text(),
      OutputDir: "out",
    }

    // Start the job
    res, err = c.work(job)
    if err != nil {
      println("Error sending a message:", err.Error())
    }

    if res == nil {
      os.Exit(0)
    }

    if res.Ok == false {
      println("Error completing the download task for", job.Url)
    } else {
      println("Saved to", res.OutputPath)
    }

    print("> ")
  }


}

func (c *client) work(job *_proto.DownloadJob) (res *_proto.DownloadJobResult, err error){
  var ctx context.Context
  var cancel context.CancelFunc

  ctx, cancel = context.WithTimeout(context.Background(), time.Duration(200 * time.Millisecond))
  res, err = c.downloadClient.Write(ctx, job)
  cancel()
  if err != nil {
    return res, err
  }

  return res, err
}
