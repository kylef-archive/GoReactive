package GoReactive


// Disposable

type newDisposable struct {
  dispose func()
}

func (disposable newDisposable) Dispose() {
  disposable.dispose()
}

func NewDisposable(closure func()) Disposable {
  return newDisposable{ dispose: closure }
}


// Subject

type Subject struct {
  onNext func(interface{})
  onCompletion func()
  onError func(error)
  dispose func()

  stopped bool
  disposed bool
}

func (subject *Subject) IsStopped() bool {
  return subject.stopped
}


func (subject *Subject) IsDisposed() bool {
  return subject.disposed
}


func (subject *Subject) SendNext(value interface{}) {
  if !subject.IsStopped() && subject.onNext != nil {
    subject.onNext(value)
  }
}


func (subject *Subject) SendCompletion() {
  if !subject.IsStopped() {
    subject.stopped = true

    if subject.onCompletion != nil {
      subject.onCompletion()
    }
  }
}


func (subject *Subject) SendError(err error) {
  if !subject.IsStopped() {
    subject.stopped = true

    if subject.onError != nil {
      subject.onError(err)
    }
  }
}


func (subject *Subject) Dispose() {
  if !subject.IsDisposed() {
    subject.disposed = true

    if subject.dispose != nil {
      subject.dispose()
    }
  }
}


func (subject *Subject) Subscribe(next func(interface{}), completion func(), failure func(error)) Disposable {
  subject.onNext = next
  subject.onCompletion = completion
  subject.onError = failure
  return subject
}


//

type newSubject struct {
  subject *Subject
  subscribed func(*Subject) Disposable
}


func (ns *newSubject) Subscribe(next func(interface{}), completion func(), failure func(error)) Disposable {
  ns.subject.Subscribe(next, completion, failure)
  disposable := ns.subscribed(ns.subject)

  if disposable == nil {
    disposable = NewDisposable(nil)
  }

  ns.subject.dispose = func() { disposable.Dispose() }

  return ns.subject
}


func NewObservable(closure func(*Subject) Disposable) Observable {
  return &newSubject{ subject: &Subject{}, subscribed: closure }
}


/// Returns an Observable from a slice
func NewObservableSlice(values []interface{}) Observable {
  return NewObservable(func(subject *Subject) Disposable {
    for _, value := range values {
      subject.SendNext(value)
    }
    subject.SendCompletion()
    return nil
  })
}


