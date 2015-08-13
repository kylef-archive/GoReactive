package GoReactive


import (
  "testing"
  "github.com/stvp/assert"
)


func TestSubjectSendingCompletion(t *testing.T) {
  subject := Subject {}
  completed := false

  subject.Subscribe(func(value interface{}) {},
                    func() { completed = true },
                    func(err error) {})

  subject.SendCompletion()

  assert.True(t, completed)
}


func TestSubjectSendingError(t *testing.T) {
  subject := Subject {}
  failed := false

  subject.Subscribe(func(value interface{}) {},
                    func() {},
                    func(err error) { failed = true })

  subject.SendError(nil)

  assert.True(t, failed)
}


func TestSubjectSendingNext(t *testing.T) {
  subject := Subject {}
  next := false

  subject.Subscribe(func(value interface{}) { next = true },
                    func() {},
                    func(err error) {})

  subject.SendNext(nil)

  assert.True(t, next)
}


func TestSubjectNewObservable(t *testing.T) {
  observable := NewObservable(func(subject *Subject) {
    subject.SendCompletion()
  })

  completed := false

  observable.Subscribe(func(value interface{}) {},
                    func() { completed = true },
                    func(err error) {})
  assert.True(t, completed)
}

