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
  assert.True(t, subject.IsStopped())
}


func TestSubjectSendingError(t *testing.T) {
  subject := Subject {}
  failed := false

  subject.Subscribe(func(value interface{}) {},
                    func() {},
                    func(err error) { failed = true })

  subject.SendError(nil)

  assert.True(t, failed)
  assert.True(t, subject.IsStopped())
}


func TestSubjectSendingNext(t *testing.T) {
  subject := Subject {}
  next := false

  subject.Subscribe(func(value interface{}) { next = true },
                    func() {},
                    func(err error) {})

  subject.SendNext(nil)

  assert.True(t, next)
  assert.False(t, subject.IsStopped())
}


func TestSubjectNewObservable(t *testing.T) {
  observable := NewObservable(func(subject *Subject) Disposable {
    subject.SendCompletion()
    return nil
  })

  completed := false

  observable.Subscribe(func(value interface{}) {},
                    func() { completed = true },
                    func(err error) {})
  assert.True(t, completed)
}

func TestSubjectNewObservableDisposable(t *testing.T) {
  disposed := false
  observable := NewObservable(func(subject *Subject) Disposable {
    return NewDisposable(func() {
      disposed = true
      assert.True(t, subject.IsDisposed())
    })
  })

  disposable := observable.Subscribe(nil, nil, nil)
  disposable.Dispose()

  assert.True(t, disposed)
}

