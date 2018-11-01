## v0.3.3 (Released 2018-10-08)

ADDITIONS

- Expose `TypeCode` on Addenda records

IMPROVEMENTS

- Run as unprivileged user in Docker image
- `cmd/readACH`: add -json to print `File` as JSON
- `cmd/writeACH`: validate file before writing
- `cmd/writeACH`: add -json to output `File` as JSON

BUG FIXES

- Fix reading batches out from JSON endpoint
- Fix plaintext ACH file rendering endpoint

## v0.3.2 (Released 2018-10-05)

ADDITIONS

- Handle pre-flight requests and proxy back CORS headers

## v0.3.1 (Released 2018-10-05)

ADDITIONS

- Add `X-Total-Count` on GET responses
- Proxy back CORS headers if sent on the request

BUG FIXES

- Drop requirement for /-suffix on GET /files
- Don't trample content-type in encodeResponse

## v0.3.0 (Released 2018-09-26)

FEATURES

- Added HTTP Server
- SEC Code CIE (Customer-Initiated Entry) ([#209](https://github.com/moov-io/ach/issues/209))
- Support IAT ([#211](https://github.com/moov-io/ach/issues/211))
- IAT Returns ([#233](https://github.com/moov-io/ach/issues/233))
- Support CTX ([#212](https://github.com/moov-io/ach/issues/212))

IMPROVEMENTS

- Added admin HTTP service.
- Added `GET /ping` route (unauthed)

BUG FIXES

- Fixes to parsing and validation.
