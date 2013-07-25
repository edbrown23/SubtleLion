package subtlelion

import (
    "testing"
)

func TestSimulation(t *testing.T) {
    s := NewSimulation()
    go s.RunSimulation()
    s.Control <- 1
    ret := <- s.EndStatus
    if ret != 0 {
        t.Fail()
    }
}
