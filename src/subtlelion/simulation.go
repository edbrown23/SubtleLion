package subtlelion

import (
    // "fmt"
)

type simulation struct {
    EndStatus chan int
    Control chan int
    Running bool
    World world
}


func NewSimulation() *simulation {
    esChan := make(chan int)
    controlChan := make(chan int, 1)
    w := NewWorld();
    s := simulation{esChan, controlChan, false, *w}
    return &s
}


func (s *simulation) RunSimulation() {
    s.Running = true;
    for s.Running {
        if 1 == <- s.Control {
            s.Running = false
        }
        

    }
    s.EndStatus <- 0
}

