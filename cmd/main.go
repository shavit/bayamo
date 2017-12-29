package main

import (
  "os"

  "github.com/shavit/bayamo"
)

func printUsage(){
  println("Options: server, repl")
}

func main(){
  if len(os.Args) <= 1 {
    printUsage()
    os.Exit(1)
  }

  switch os.Args[1] {
  case "server":
    var err error
    var server bayamo.Server = bayamo.NewServer()
    err = server.Start()
    if err != nil {
      panic(err)
    }
    break
  case "repl":
    var err error
    client := bayamo.NewClient()
    err = client.Dial()
    if err != nil {
      panic(err)
    }
    break
  default:
    printUsage()
    os.Exit(1)
  }
}
