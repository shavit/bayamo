package bayamo

import (
  "testing"
)

func TestClientDial(t *testing.T){
  var err error
  var c Client = NewClient()
  if c == nil {
    t.Error("Error creating a new client")
  }

  // err = c.Dial()
  if err != nil {
    t.Error(err)
  }
}
