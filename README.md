# GoReactive

[![Build Status](https://img.shields.io/circleci/project/kylef/GoReactive/master.svg)](https://circleci.com/gh/kylef/GoReactive)

GoReactive is a Go library for [Function Reactive Programming](https://en.wikipedia.org/wiki/Functional_reactive_programming), a library for representing and consuming asyncronous data streams with Observables.

## Usage

### Observables

At it's simplest form, an `Observable` is an interface which allows you to subscribe to events, the next value, the completion or the failure.

```go
observable.Subscribe(
  func(value interface{}) { /* a new value has been sent */ },
  func() { /* the stream has completed */ },
  func(err error) { /* the stream has errored */ }
)
```

#### Transforming an Observable

Using `Map`, you can transform an Observables next values. For example, we can use the following to create a new `Observable` from an `Observable` of integers by multiplying their value.

```go
transformedObservable := Map(observable, func(value interface{}) interface{} {
  return value.(int) * 2
})
```

#### Filtering an Observable

Using `Filter` and `Exclude` you can filter or exclude next values from an `Observable`. We can use the following to create a new `Observable` from an `Observable` of integers, filtering for all even numbers.

```go
filteredObservable := Filter(observable, func(value interface{}) bool {
  return (value.(int) % 2) == 0
})
```

### Subject

A `Subject` is an interface for creating Observables, a `Subject` contains functions to emit new values into your Observable.

```go
subject.SendNext("Kyle")
subject.SendNext("Katie")
subject.SendCompletion()
```

```go
subject.SendError(err)
```

#### Creating a new subject using `NewObservable`

`NewObservable` is a function for creating an Observable, which when subscribed, it will call a callback allowing you to send events using a `Subject`.

```go
observable := NewObservable(func(subject *Subject) Disposable {
  subject.SendNext("Hello, my subscriber.")
  return nil
})
```

## License

GoReactive is licensed under the BSD license. See [LICENSE](LICENSE) for more info.

