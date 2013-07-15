package webserver

import (
    "testing"
)

func TestRegisterCallback(t *testing.T) {
    x := Newurls()
    x.RegisterCallback("test", testFunc)
}

func testFunc() {
    
}