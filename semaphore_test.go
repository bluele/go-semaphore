package semaphore_test

import (
	"github.com/bluele/go-semaphore"
	"testing"
	"time"
)

func TestSemaphore(t *testing.T) {
	permit := 3
	sem := semaphore.NewSemaphore(permit)
	sem.Aquire(1)
	if sem.Available() != permit-1 {
		t.Errorf("sem.Available() should be %v", permit-1)
	}
	sem.Release()
	if sem.Available() != permit {
		t.Errorf("sem.Available() should be %v", permit)
	}
}

func TestTimeLimitedSemaphore(t *testing.T) {
	permit := 3
	sem := semaphore.NewTimeLimitedSemaphore(permit, time.Second)
	sem.Aquire(1)
	if sem.Available() != permit-1 {
		t.Errorf("sem.Available() should be %v", permit-1)
	}
	sem.Release()
	if sem.Available() != permit-1 {
		t.Errorf("sem.Available() should be %v", permit-1)
	}
	time.Sleep(2 * time.Second)
	if sem.Available() != permit {
		t.Errorf("sem.Available() should be %v", permit)
	}
}
