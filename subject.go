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
}


func (subject *Subject) SendNext(value interface{}) {
  if subject.onNext != nil {
    subject.onNext(value)
  }
}


func (subject *Subject) SendCompletion() {
  if subject.onCompletion != nil {
    subject.onCompletion()
  }
}


func (subject *Subject) SendError(err error) {
  if subject.onError != nil {
    subject.onError(err)
  }
}


func (subject *Subject) Dispose() {
  if subject.dispose != nil {
    subject.dispose()
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

  return disposable
}


func NewObservable(closure func(*Subject) Disposable) Observable {
  return &newSubject{ subject: &Subject{}, subscribed: closure }
}


