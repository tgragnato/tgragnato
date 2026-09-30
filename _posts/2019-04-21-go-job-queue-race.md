---
title: A Go Job Queue, a Misnamed Stack, and a Race
title_marker: Go concurrency
description: A small producer-consumer example that exposes locking and ticker-lifecycle mistakes
layout: default
lang: en
---

This Go experiment has one goroutine producing jobs and another consuming them. Despite the name `stack`, `pop` removes `jobs[0]`, so the data structure behaves as a FIFO queue.

It is more useful today as a concurrency review than as a reusable queue. `pop` checks the slice before taking the mutex, so concurrent access can race with `push`; meanwhile, calling `time.Tick` inside each loop iteration creates new tickers instead of reusing one. The shutdown handshake is also harder to reason about than a context or a closed channel. These are compact examples of why synchronization must cover the whole state transition, not just the mutation.

```go
type job string
type stack struct {
  sync.Mutex
  jobs []job
}

func (s *stack) push(in job) {
  s.Lock()
  if s.jobs == nil || len(s.jobs) == 0 {
    s.jobs = []job{in}
  } else {
    s.jobs = append(s.jobs, in)
  }
  s.Unlock()
}

func (s *stack) pop() (job, error) {
  if s.jobs == nil || len(s.jobs) == 0 {
    return "", fmt.Errorf("pila vuota")
  }
  out := s.jobs[0]
  s.Lock()
  if len(s.jobs) == 1 {
    s.jobs = []job{}
  } else {
    s.jobs = s.jobs[1:]
  }
  s.Unlock()
  return out, nil
}

func main() {
  done := make(chan bool)

  var items stack

  go func(items *stack) {
    for {
      select {
      case <-time.Tick(time.Second / 3):
        items.push(job(fmt.Sprintf("job %v", time.Now())))
      }
    }
  }(&items)

  go func(items *stack, done chan bool) {
    for {
      select {
      case <-time.Tick(time.Second / 5):
        item, err := items.pop()
        if err != nil {
          log.Println(err)
        } else {
          log.Println(item)
        }
      case <-done:
        log.Println("chiudo")
        time.Sleep(2 * time.Second)
        done <- true
        return
      }
    }
  }(&items, done)

  log.Println("sono il main")
  time.Sleep(30 * time.Second)
  done <- true
  pippo := <-done
  log.Printf("fine del main %v", pippo)
}
```
