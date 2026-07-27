// Package spamcheck is a tiny client for Postmark's SpamCheck API
// (https://spamcheck.postmarkapp.com/doc/). It scores a raw RFC-822 email
// against SpamAssassin rules and returns the score and, optionally, the full
// matched-rule breakdown plus report.
package spamcheck

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/psyb0t/ctxerrors"
)

const (
	// DefaultBaseURL is the public Postmark SpamCheck API base URL.
	DefaultBaseURL = "https://spamcheck.postmarkapp.com"

	filterPath      = "/filter"
	contentTypeJSON = "application/json"
	defaultTimeout  = 30 * time.Second
)

// Report selects how much detail the API returns.
type Report string

const (
	// ReportShort asks for the spam score only.
	ReportShort Report = "short"
	// ReportLong asks for the score plus the matched-rule breakdown and report.
	ReportLong Report = "long"
)

// Rule is a single SpamAssassin rule that fired, with its score contribution.
type Rule struct {
	Score       string `json:"score"`
	Description string `json:"description"`
}

// Result is the SpamCheck API response. Rules and Report are only populated
// when the check was requested with ReportLong.
type Result struct {
	Success bool   `json:"success"`
	Score   string `json:"score"`
	Rules   []Rule `json:"rules"`
	Report  string `json:"report"`
	Message string `json:"message"` // set when Success is false
}

// ScoreValue parses Score into a float64. The API returns the SpamAssassin
// score as a string; higher means more spammy.
func (r *Result) ScoreValue() (float64, error) {
	score, err := strconv.ParseFloat(r.Score, 64)
	if err != nil {
		return 0, ctxerrors.Wrapf(err, "parse score %q", r.Score)
	}

	return score, nil
}

// request is the SpamCheck API request body.
type request struct {
	Email   string `json:"email"`
	Options string `json:"options"`
}

// HTTPDoer is the slice of *http.Client this package depends on. Swap in a fake
// to test without hitting the network.
type HTTPDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client talks to the SpamCheck API.
type Client struct {
	httpDoer HTTPDoer
	baseURL  string
}

// Option configures a Client.
type Option func(*Client)

// WithHTTPDoer sets a custom HTTP client — a fake in tests, or an http.Client
// with a specific timeout or transport.
func WithHTTPDoer(doer HTTPDoer) Option {
	return func(c *Client) {
		c.httpDoer = doer
	}
}

// WithBaseURL overrides the API base URL. Useful for tests or a proxy.
func WithBaseURL(baseURL string) Option {
	return func(c *Client) {
		c.baseURL = baseURL
	}
}

// New builds a Client. With no options it uses an http.Client with a 30s
// timeout against the public Postmark endpoint.
func New(opts ...Option) *Client {
	client := &Client{
		httpDoer: &http.Client{Timeout: defaultTimeout}, //nolint:exhaustruct // only Timeout matters
		baseURL:  DefaultBaseURL,
	}

	for _, opt := range opts {
		opt(client)
	}

	return client
}

// Check submits rawEmail (a full RFC-822 message, headers included) to the
// SpamCheck API. It returns an error if the email is empty, the request fails,
// the API responds with a non-200 status, the response cannot be decoded, or
// the API reports success=false.
func (c *Client) Check(
	ctx context.Context,
	rawEmail string,
	report Report,
) (*Result, error) {
	if rawEmail == "" {
		return nil, ctxerrors.Wrap(ErrEmptyEmail, "check")
	}

	body, err := json.Marshal(request{Email: rawEmail, Options: string(report)})
	if err != nil {
		return nil, ctxerrors.Wrap(err, "marshal request")
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		c.baseURL+filterPath,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, ctxerrors.Wrap(err, "build request")
	}

	req.Header.Set("Content-Type", contentTypeJSON)

	resp, err := c.httpDoer.Do(req)
	if err != nil {
		return nil, ctxerrors.Wrap(err, "do request")
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, ctxerrors.Wrapf(ErrUnexpectedStatus, "status %d", resp.StatusCode)
	}

	var result Result
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, ctxerrors.Wrap(ErrDecodeResponse, err.Error())
	}

	if !result.Success {
		return nil, ctxerrors.Wrapf(ErrAPIFailure, "message: %q", result.Message)
	}

	return &result, nil
}

// Check runs a one-off spam check with a default client against the public
// Postmark endpoint. For repeated calls, build a Client with New and reuse it.
func Check(
	ctx context.Context,
	rawEmail string,
	report Report,
) (*Result, error) {
	return New().Check(ctx, rawEmail, report)
}
