## v0.8.0 (Released 2019-02-01)

ADDITIONS

- Added `Has` and `Match` functions to support type-based error handling

## v0.7.0 (Released 2019-01-31)

ADDITIONS

- Add `ID() string` to return a random identifier.

## v0.6.0 (Released 2019-01-25)

ADDITIONS

- admin: [`Server.AddHandler`](https://godoc.org/github.com/moov-io/base/admin#Server.AddHandler) for extendable commands
- http/bind: Add [Fed](https://github.com/moov-io/fed) service

## v0.5.1 (Released 2019-01-17)

BUG FIXES

- http: fix panic in ResponseWriter.WriteHeader

## v0.5.0 (Released 2019-01-17)

BUG FIXES

- http: don't panic if nil idempotent.Recorder is passed to ResponseWriter

ADDITIONS

- http/bind: Add [OFAC](https://github.com/moov-io/ofac) and [GL](https://github.com/moov-io/gl) services
- k8s: Add [`Inside()`](https://godoc.org/github.com/moov-io/base/k8s#Inside) for cluster awareness.
- docker: Add [`Enabled()`](https://godoc.org/github.com/moov-io/base/docker#Enabled) for compatability checks.

## v0.4.0 (Released 2019-01-11)

BREAKING CHANGES

- time: default times to UTC rather than Eastern.

## v0.3.1 (Released 2019-01-09)

- error: Add `ParseError` and `ErrorList` types.
- time: Prevent negative times in `NewTime(t time.Time)`

## v0.3.0 (Released 2019-01-07)

ADDITIONS

- Add ParseError and ErrorList. (See: [moov-io/base #23](https://github.com/moov-io/base/issues/23))

## v0.2.1 (Released 2019-01-03)

BUG FIXES

- http: Add OPTIONS to Access-Control-Allow-Methods

## v0.2.0 (Released 2018-12-18)

ADDITIONS

- Add `base.Time` as an embedded `time.Time` with banktime methods. (AddBankingDay, IsWeekend)

## v0.1.0 (Released 2018-12-17)

- Initial release
