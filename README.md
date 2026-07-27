# go-spamcheck

[![Go Reference](https://pkg.go.dev/badge/github.com/psyb0t/go-spamcheck.svg)](https://pkg.go.dev/github.com/psyb0t/go-spamcheck)
[![CI](https://github.com/psyb0t/go-spamcheck/actions/workflows/pipeline.yml/badge.svg?branch=master)](https://github.com/psyb0t/go-spamcheck/actions/workflows/pipeline.yml)
[![coverage](https://raw.githubusercontent.com/psyb0t/go-spamcheck/badges/coverage.svg)](https://github.com/psyb0t/go-spamcheck/actions/workflows/pipeline.yml)
[![version](https://raw.githubusercontent.com/psyb0t/go-spamcheck/badges/version.svg)](https://github.com/psyb0t/go-spamcheck/tags)
[![license](https://raw.githubusercontent.com/psyb0t/go-spamcheck/badges/license.svg)](LICENSE)

A tiny Go client for [Postmark's SpamCheck API](https://spamcheck.postmarkapp.com/doc/). Feed it a raw email, get back a SpamAssassin score — no account, no API key, no bullshit.

## What the fuck does it do?

You hand it a full RFC-822 email (headers and all), it POSTs it to Postmark's public SpamCheck endpoint, and hands you back:

- a **score** (SpamAssassin points — higher = spammier), and
- optionally the **rule breakdown** + human-readable **report** of every rule that fired.

That's it. One dependency ([`ctxerrors`](https://github.com/psyb0t/ctxerrors) for error context), stdlib `net/http` for the request. No account or token required — the endpoint is public.

## Installation

```bash
go get github.com/psyb0t/go-spamcheck
```

## Usage

### One-off check

```go
package main

import (
	"context"
	"fmt"
	"log"

	spamcheck "github.com/psyb0t/go-spamcheck"
)

func main() {
	rawEmail := `From: sender@example.com
To: me@example.com
Subject: FREE MONEY!!!

Congratulations you WON $1000000, click http://spam.example to claim.`

	result, err := spamcheck.Check(context.Background(), rawEmail, spamcheck.ReportShort)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Score:", result.Score) // e.g. "4.8"
}
```

### Reusable client + full report

For repeated calls, build a `Client` once and reuse it. Ask for `ReportLong`
to get the matched rules and the full report:

```go
client := spamcheck.New()

result, err := client.Check(ctx, rawEmail, spamcheck.ReportLong)
if err != nil {
	log.Fatal(err)
}

score, err := result.ScoreValue() // parse "4.8" -> 4.8
if err != nil {
	log.Fatal(err)
}

if score >= 5.0 {
	fmt.Println("probably spam")
}

for _, rule := range result.Rules {
	fmt.Printf("%s\t%s\n", rule.Score, rule.Description)
}
fmt.Println(result.Report)
```

### Options

`New` takes functional options:

- `WithBaseURL(url)` — point at a different base URL (a proxy, or a test server).
- `WithHTTPDoer(doer)` — swap the HTTP client (custom timeout/transport, or a fake in tests). Any type with `Do(*http.Request) (*http.Response, error)` works.

```go
client := spamcheck.New(
	spamcheck.WithHTTPDoer(&http.Client{Timeout: 5 * time.Second}),
)
```

## Errors

`Check` returns typed sentinels you can match with `errors.Is`:

| Error | Meaning |
|-------|---------|
| `ErrEmptyEmail` | the raw email you passed was empty |
| `ErrUnexpectedStatus` | the API responded with a non-200 status |
| `ErrDecodeResponse` | the response body couldn't be decoded as JSON |
| `ErrAPIFailure` | the API responded with `success: false` (its message is wrapped in) |

Every error carries file/line/function context via `ctxerrors`.

## Testing

```bash
make test           # unit tests (mocked HTTP, no network)
make test-coverage  # unit tests + coverage gate
make test-real      # hit the LIVE Postmark API — local only, detects upstream API changes
```

The default suite mocks the HTTP layer, so it's fast and offline. The `-real`
suite (build tag `real`) hits the actual Postmark endpoint and asserts the
response shape still matches — run it locally when you want to catch upstream
API drift. It is excluded from CI.

## License

MIT. See [LICENSE](LICENSE).
