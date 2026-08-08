# Changelog

All notable changes per release. Versions follow [semver](https://semver.org).

## v1.0.3 — 2026-08-08

Repository infrastructure only. No library code changed.

- Added the imported-by badge: a count of the public packages importing this
  module, linking to `importers.md` on the `badges` branch — the importing
  repositories, grouped, package counts descending, and flagged when the owner
  differs from this repo's.
- It measures **blast radius, not adoption**. The number tells you how much
  breaks when an exported name moves; the external mark tells you whether any of
  that is someone else's problem, which is what decides how strictly the module
  has to be versioned. It reads `0` today — a confirmed zero, which is itself
  useful: nothing downstream breaks yet.
- Refreshed weekly rather than daily, because pkg.go.dev's crawl lags
  publication by days and each run drags the full test suite along (the badges
  job needs the coverage artifact). The whole pipeline runs rather than a
  badges-only job: the badge publisher republishes only what a run produced, so
  a badge-only refresh would delete the coverage, version and license badges.
- The cron slot is derived from a hash of the repository name rather than
  chosen — GitHub cron has no randomness, and its scheduler sheds queued runs
  hardest at the round times a human would pick.

## v1.0.2 — 2026-08-01

Infrastructure only. No library code changed — every commit since v1.0.1
touches `.github/workflows/`.

- The pipeline was split: building and publishing stay in `pipeline.yml`, and
  everything that leaves the host now lives beside it in `mirror-and-archive.yml`.
- The repo is mirrored to Codeberg as well as GitLab.
- It is archived to the Wayback Machine, Software Heritage and archive.org.
- Issues opened on either mirror are copied back to GitHub every six hours, and
  closed here when the original closes.
- Pull requests are switched off on the mirrors: they are force-pushed from
  GitHub, so anything merged there would be destroyed by the next sync. Issues
  and forking stay enabled.

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
