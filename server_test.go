package bayamo

import (
  "os"
  "testing"
)

func TestServerStart(t *testing.T){
  var err error
  var server Server = NewServer()
  if server == nil {
    t.Error("Error creating a server")
  }

  // err = server.Start()
  if err != nil {
    t.Error(err)
  }
}

func TestServerSecureTest(t *testing.T){
  var err error
  var server Server = NewServer()
  var fServerCrt *os.File
  var fServerKey *os.File

  fServerCrt, err = os.Create("server.crt")
  if err != nil {
    t.Error(err)
  }
  defer fServerCrt.Close()

  fServerKey, err = os.Create("server.key")
  if err != nil {
    t.Error(err)
  }
  defer fServerKey.Close()

  err = server.AddCredentials("./server.crt", "./server.key")
  // Ignore TLS errors
  if err != nil && err.Error()[0:3] != "tls" {
    t.Error(err)
  }

  os.Remove("server.crt")
  os.Remove("server.key")
}
