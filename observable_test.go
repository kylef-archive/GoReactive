package GoReactive


import (
  "testing"
  "time"
  "github.com/stvp/assert"
)


/// Returns an observable which emits 1, 2, 3, 4, and 5 then completion
func oneToFiveObservable() Observable {
  return NewObservable(func(subject *Subject) {
    subject.SendNext(1)
    subject.SendNext(2)
    subject.SendNext(3)
    subject.SendNext(4)
    subject.SendNext(5)
    subject.SendCompletion()
  })
}

/// Returns an observable which emits 1, 1, 2, 3, 3, 4, 5 and 5 then completion
func oneishToFiveObservable() Observable {
  return NewObservable(func(subject *Subject) {
    subject.SendNext(1)
    subject.SendNext(1)
    subject.SendNext(2)
    subject.SendNext(3)
    subject.SendNext(3)
    subject.SendNext(4)
    subject.SendNext(5)
    subject.SendNext(5)
    subject.SendCompletion()
  })
}


/// Subscribes to an observable returning once it completes or fails
func wait(t *testing.T, observable Observable) []interface{} {
  values := []interface{}{}
  completed := false
  failed := false

  observable.Subscribe(
      func(value interface{}) { values = append(values, value) },
      func() { completed = true },
      func(err error) { failed = true })

  for !completed && !failed {
    time.Sleep(500 * time.Millisecond)
  }

  if failed {
    t.FailNow()
  }

  return values
}


func TestSkipSkipsValues(t *testing.T) {
  observable := oneToFiveObservable()
  values := wait(t, Skip(observable, 2))

  assert.Equal(t, values, []interface{}{3, 4, 5})
}


func TestDistinctUntilChanged(t *testing.T) {
  observable := oneishToFiveObservable()
  values := wait(t, DistinctUntilChanged(observable))

  assert.Equal(t, values, []interface{}{1, 2, 3, 4, 5})
}


func TestMap(t *testing.T) {
  observable := oneToFiveObservable()
  mappedObservable := Map(observable, func(value interface{}) interface{} {
    return value.(int) * 2
  })
  values := wait(t, mappedObservable)

  assert.Equal(t, values, []interface{}{2, 4, 6, 8, 10})
}


func TestFilter(t *testing.T) {
  observable := oneToFiveObservable()
  filteredObservable := Filter(observable, func(value interface{}) bool {
    return (value.(int) % 2) == 0
  })
  values := wait(t, filteredObservable)

  assert.Equal(t, values, []interface{}{2, 4})
}


func TestExclude(t *testing.T) {
  observable := oneToFiveObservable()
  filteredObservable := Exclude(observable, func(value interface{}) bool {
    return (value.(int) % 2) == 0
  })
  values := wait(t, filteredObservable)

  assert.Equal(t, values, []interface{}{1, 3, 5})
}

