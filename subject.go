package GoReactive


type Subject struct {
  onNext func(interface{})
  onCompletion func()
  onError func(error)
}


func (subject *Subject) SendNext(value interface{}) {
  if (subject.onNext != nil) {
    subject.onNext(value)
  }
}


func (subject *Subject) SendCompletion() {
  if (subject.onCompletion != nil) {
    subject.onCompletion()
  }
}


func (subject *Subject) SendError(err error) {
  if (subject.onError != nil) {
    subject.onError(err)
  }
}


func (subject *Subject) Subscribe(next func(interface{}), completion func(), failure func(error)) {
  subject.onNext = next
  subject.onCompletion = completion
  subject.onError = failure
}


//

type newSubject struct {
  subject *Subject
  subscribed func(*Subject)
}


func (ns *newSubject) Subscribe(next func(interface{}), completion func(), failure func(error)) {
  ns.subject.Subscribe(next, completion, failure)
  ns.subscribed(ns.subject)
}


func NewObservable(closure func(*Subject)) Observable {
  return &newSubject{ subject: &Subject{}, subscribed: closure }
}

