# go-semaphore

Implements basic semaphore and time limited semaphore on go language.

# Examples

## Basic Semaphore

```go
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
```

## Time limieted semaphore

```go
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
```

# Test

```
$ go test
```

# Author

**Jun Kimura**

* <http://github.com/bluele>
* <junkxdev@gmail.com>