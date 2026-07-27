package spamcheck

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type fakeDoer struct {
	resp *http.Response
	err  error
}

func (f *fakeDoer) Do(*http.Request) (*http.Response, error) {
	return f.resp, f.err
}

func TestClient_Check(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name          string
		status        int
		body          string
		report        Report
		wantErr       error
		wantScore     string
		wantRuleCount int
	}{
		{
			name:      "success short returns score only",
			status:    http.StatusOK,
			body:      `{"success":true,"score":"2.7"}`,
			report:    ReportShort,
			wantScore: "2.7",
		},
		{
			name:   "success long returns rules and report",
			status: http.StatusOK,
			body: `{"success":true,"score":"5.1",` +
				`"rules":[{"score":"1.1","description":"HTML_MESSAGE"},` +
				`{"score":"4.0","description":"URI_HEX"}],"report":"pts rule"}`,
			report:        ReportLong,
			wantScore:     "5.1",
			wantRuleCount: 2,
		},
		{
			name:    "api failure surfaces ErrAPIFailure",
			status:  http.StatusOK,
			body:    `{"success":false,"message":"Missing required parameter 'email'"}`,
			report:  ReportShort,
			wantErr: ErrAPIFailure,
		},
		{
			name:    "non-200 surfaces ErrUnexpectedStatus",
			status:  http.StatusInternalServerError,
			body:    `boom`,
			report:  ReportShort,
			wantErr: ErrUnexpectedStatus,
		},
		{
			name:    "malformed json surfaces ErrDecodeResponse",
			status:  http.StatusOK,
			body:    `{not json`,
			report:  ReportShort,
			wantErr: ErrDecodeResponse,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			server := httptest.NewServer(http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					assert.Equal(t, http.MethodPost, r.Method)
					assert.Equal(t, filterPath, r.URL.Path)
					assert.Equal(t, contentTypeJSON, r.Header.Get("Content-Type"))

					var req request

					require.NoError(t, json.NewDecoder(r.Body).Decode(&req))
					assert.NotEmpty(t, req.Email)
					assert.Equal(t, string(tc.report), req.Options)

					w.WriteHeader(tc.status)
					_, _ = io.WriteString(w, tc.body)
				},
			))
			t.Cleanup(server.Close)

			client := New(WithBaseURL(server.URL))

			result, err := client.Check(
				context.Background(),
				"From: a@b.c\n\nhello",
				tc.report,
			)

			if tc.wantErr != nil {
				require.ErrorIs(t, err, tc.wantErr)
				assert.Nil(t, result)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, result)
			assert.True(t, result.Success)
			assert.Equal(t, tc.wantScore, result.Score)
			assert.Len(t, result.Rules, tc.wantRuleCount)
		})
	}
}

func TestClient_CheckEmptyEmail(t *testing.T) {
	t.Parallel()

	result, err := New().Check(context.Background(), "", ReportShort)
	require.ErrorIs(t, err, ErrEmptyEmail)
	assert.Nil(t, result)
}

func TestClient_CheckTransportError(t *testing.T) {
	t.Parallel()

	wantErr := errors.New("connection refused")
	client := New(WithHTTPDoer(&fakeDoer{resp: nil, err: wantErr}))

	result, err := client.Check(context.Background(), "raw email", ReportShort)
	require.ErrorIs(t, err, wantErr)
	assert.Nil(t, result)
}

func TestClient_CheckBuildRequestError(t *testing.T) {
	t.Parallel()

	// A control character in the base URL makes NewRequestWithContext fail.
	client := New(WithBaseURL("http://\x7f"))

	result, err := client.Check(context.Background(), "raw email", ReportShort)
	require.Error(t, err)
	assert.Nil(t, result)
}

func TestResult_ScoreValue(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name    string
		score   string
		want    float64
		wantErr bool
	}{
		{name: "positive", score: "2.7", want: 2.7, wantErr: false},
		{name: "zero", score: "0", want: 0, wantErr: false},
		{name: "negative", score: "-1.5", want: -1.5, wantErr: false},
		{name: "not a number", score: "abc", want: 0, wantErr: true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			result := &Result{
				Success: true,
				Score:   tc.score,
				Rules:   nil,
				Report:  "",
				Message: "",
			}

			got, err := result.ScoreValue()
			if tc.wantErr {
				require.Error(t, err)

				return
			}

			require.NoError(t, err)
			assert.InDelta(t, tc.want, got, 0.0001)
		})
	}
}

func TestCheckPackageLevel(t *testing.T) {
	t.Parallel()

	result, err := Check(context.Background(), "", ReportShort)
	require.ErrorIs(t, err, ErrEmptyEmail)
	assert.Nil(t, result)
}
