//go:build real

// These tests hit the LIVE Postmark SpamCheck API to detect upstream response
// shape changes. They are excluded from the default build and CI — run locally:
//
//	make test-real
//
// A failure here means Postmark changed the API (or it's down), not that this
// package is broken.
package spamcheck_test

import (
	"context"
	"testing"

	spamcheck "github.com/psyb0t/go-spamcheck"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const spammyEmail = `From: sender@example.com
To: me@example.com
Subject: FREE MONEY

Congratulations you WON $1000000 click http://spam.example VIAGRA cheap pills`

func TestReal_CheckShort(t *testing.T) {
	t.Parallel()

	result, err := spamcheck.Check(
		context.Background(),
		spammyEmail,
		spamcheck.ReportShort,
	)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.True(t, result.Success)

	score, err := result.ScoreValue()
	require.NoError(t, err)
	assert.Positive(t, score)
}

func TestReal_CheckLong(t *testing.T) {
	t.Parallel()

	result, err := spamcheck.Check(
		context.Background(),
		spammyEmail,
		spamcheck.ReportLong,
	)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.True(t, result.Success)
	assert.NotEmpty(t, result.Report, "long mode should return a report")
	require.NotEmpty(t, result.Rules, "long mode should return matched rules")
	assert.NotEmpty(t, result.Rules[0].Description)
}
