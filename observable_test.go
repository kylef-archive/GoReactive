package GoReactive


import (
  "testing"
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


func TestSkipSkipsValues(t *testing.T) {
  observable := oneToFiveObservable()
  skippedObservable := Skip(observable, 2)

  values := []int{}
  skippedObservable.Subscribe(
      func(value interface{}) { values = append(values, value.(int)) },
      func() {},
      func(err error) {})

  assert.Equal(t, values, []int{3, 4, 5})
}


func TestMap(t *testing.T) {
  observable := oneToFiveObservable()

  mappedObservable := Map(observable, func(value interface{}) interface{} {
    return value.(int) * 2
  })

  values := []int{}
  mappedObservable.Subscribe(
      func(value interface{}) { values = append(values, value.(int)) },
      func() {},
      func(err error) {})

  assert.Equal(t, values, []int{2, 4, 6, 8, 10})
}


func TestFilter(t *testing.T) {
  observable := oneToFiveObservable()

  filteredObservable := Filter(observable, func(value interface{}) bool {
    return (value.(int) % 2) == 0
  })

  values := []int{}
  filteredObservable.Subscribe(
      func(value interface{}) { values = append(values, value.(int)) },
      func() {},
      func(err error) {})

  assert.Equal(t, values, []int{2, 4})
}


func TestExclude(t *testing.T) {
  observable := oneToFiveObservable()

  filteredObservable := Exclude(observable, func(value interface{}) bool {
    return (value.(int) % 2) == 0
  })

  values := []int{}
  filteredObservable.Subscribe(
      func(value interface{}) { values = append(values, value.(int)) },
      func() {},
      func(err error) {})

  assert.Equal(t, values, []int{1, 3, 5})
}

