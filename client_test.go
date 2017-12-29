package bayamo

import (
  "testing"
)

func TestClientCreateNew(t *testing.T){
  var c *client = NewClient()
  if c == nil {
    t.Error("Error creating a new client")
  }

}
