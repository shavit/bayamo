package bayamo

import (
  "errors"
  "time"

  "google.golang.org/grpc"
  context "golang.org/x/net/context"
  _proto "github.com/shavit/bayamo/proto"
)

type Client interface {
  Dial() (err error)
}

type client struct {
  downloadClient _proto.DownloaderServiceClient
}

func NewClient() (c *client){
  return new(client)
}

func (c *client) Dial() (err error){
  var conn *grpc.ClientConn
  var timeout time.Duration = time.Duration(10 * time.Second)
  var res *_proto.DownloadJobResult
  var job *_proto.DownloadJob

  conn, err = grpc.Dial(":2400", grpc.WithTimeout(timeout), grpc.WithInsecure())
  if err != nil {
    return errors.New("Error dialing to remote server")
  }
  defer conn.Close()
  c.downloadClient = _proto.NewDownloaderServiceClient(conn)

  res, err = c.downloadClient.Write(context.Background(), job)
  if err != nil {
    return err
  }

  println(res)

  return err
}
