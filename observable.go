package GoReactive


type Observer interface {
  Next(value interface{})
  Error(err error)
  Completion()
}


type Observable interface {
  Subscribe(next func(interface{}), completion func(), failure func(error))
}


// Map

type mappedObservable struct {
  observable Observable
  transform func(interface{}) interface{}
}


func (observable mappedObservable) Subscribe(next func(interface{}), completion func(), failure func(error)) {
  observable.observable.Subscribe(func(value interface{}) {
                                    next(observable.transform(value))
                                  },
                                  completion, failure)
}


func Map(observable Observable, transform func(interface{}) interface{}) Observable {
  return mappedObservable{ observable: observable, transform: transform }
}


// Filter

type filterObservable struct {
  observable Observable
  filter func(interface{}) bool
}


func (observable filterObservable) Subscribe(next func(interface{}), completion func(), failure func(error)) {
  observable.observable.Subscribe(func(value interface{}) {
                                    if observable.filter(value) {
                                      next(value)
                                    }
                                  },
                                  completion, failure)
}


func Filter(observable Observable, filter func(interface{}) bool) Observable {
  return filterObservable{ observable: observable, filter: filter }
}


func Exclude(observable Observable, exclude func(interface{}) bool) Observable {
  return Filter(observable, func(value interface{}) bool { return !exclude(value) })
}

