package subtlelion

import (
    "testing"
)

func TestNewWorld(t *testing.T) {
    w := NewWorld()
    if w == nil {
        t.Fail()
    }
}
