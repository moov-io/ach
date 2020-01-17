## v0.11.0 (Released 2020-01-16)

ADDITIONS

- admin: add a handler to print the version on 'GET /version'

IMPROVEMENTS

- http/bind: rename ofac as watchman

BUILD

- Update module prometheus/client_golang to v1.3.0
- Update Copyright headers for 2020
- chore(deps): update module hashicorp/golang-lru to v0.5.4

## v0.10.0 (Released 2019-08-13)

BREAKING CHANGES

We've renamed `http.GetRequestID` and `http.GetUserID` from `http.Get*Id` to match Go's preference for `ID` suffixes.

ADDITIONS

- idempotent: add [`Header(*http.Request) string`](https://godoc.org/github.com/moov-io/base/idempotent#Header) and `HeaderKey`
- http/bind: add Wire HTTP service/port binding
- http/bind: add customers port
- http/bind: rename gl to accounts
- time: expose ISO 8601 format

BUG FIXES

- http: respond with '429 PreconditionFailed' if X-Idempotency-Key has been seen before

IMPROVEMENTS

- idempotent: bump up max header length
- admin: bind on a random port and return it in BindAddr on `:0`
- build: enable windows in TravisCI

## v0.9.0 (Released 2019-03-04)

ADDITIONS

- admin: Added `AddLivenessCheck` and `AddReadinessCheck` for HTTP health checks

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
