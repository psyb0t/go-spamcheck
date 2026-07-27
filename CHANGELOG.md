# Changelog

All notable changes per release. Versions follow [semver](https://semver.org).

## v1.0.1 — 2026-07-27

Lint tooling: `modernize` → built-in `go fix`.

- `make lint` / `make lint-fix` now use Go 1.26's built-in `go fix` (`-diff`
  check in `lint`, apply in `lint-fix`) instead of the `modernize` analyzer, and
  the `modernize` tool directive is dropped from `go.mod`. No library code changed.

## v1.0.0 — 2026-07-27

First tagged release — a full modernization of the original pre-modules package.

- **Go modules + Go 1.26.** The package had no `go.mod` at all (GOPATH-era). It
  now builds as a proper module with vendored dependencies.
- **New API — no more panics.** `Check(ctx, rawEmail, report)` returns
  `(*Result, error)`; the old code called `panic(err)` on every failure. A
  reusable `Client` (built with `New` + `WithBaseURL` / `WithHTTPDoer` options)
  replaces the package-level-only function, and a package-level `Check` remains
  for one-off calls. `Report` is a typed enum (`ReportShort` / `ReportLong`).
- **Context-aware.** Requests are built with `http.NewRequestWithContext`, so
  callers control cancellation and timeouts.
- **Typed error sentinels** (`ErrEmptyEmail`, `ErrUnexpectedStatus`,
  `ErrDecodeResponse`, `ErrAPIFailure`), matchable with `errors.Is` and wrapped
  with `github.com/psyb0t/ctxerrors` for file/line/function context.
- **Bug fix:** the old code ran `defer resp.Body.Close()` before checking the
  request error, which would panic on a nil response when the request failed.
- **`Result.ScoreValue()`** helper parses the string score into a `float64`.
- **Tests:** table-driven `testify` unit tests that mock the HTTP layer via
  `httptest` (fast, offline). A separate `-real` build-tagged suite
  (`make test-real`) hits the live Postmark API to detect upstream response
  shape changes; it is excluded from CI.
- **Tooling:** Makefile, `golangci-lint`, GitHub Actions CI, and self-hosted
  coverage / version / license badges.
- **Breaking:** the import path is now `github.com/psyb0t/go-spamcheck` (root
  package `spamcheck`), replacing the old `github.com/psyb0t/go-spamcheck/spamcheck`.
