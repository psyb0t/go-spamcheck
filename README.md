# go-spamcheck

https://godoc.org/github.com/psyb0t/go-spamcheck/spamcheck

A simple go wrapper for [Postmark’s Spamcheck API](http://spamcheck.postmarkapp.com/doc)

## usage

```golang
package main

import (
	"github.com/psyb0t/go-spamcheck/spamcheck"
	"fmt"
	"encoding/json"
)

func main() {
	rawEmail := `ALL THE RAW EMAIL CONTENT WITH HEADERS AND SHIT`
	result := spamcheck.Check(rawEmail, true)

	fmt.Println("Score:", result.Score)

	rules, _ := json.MarshalIndent(result.Rules, "", "  ")
	fmt.Println("Rules:\n", string(rules))

	fmt.Println("Report:\n", result.Report)

	return
}
```

## test run

```
git clone https://github.com/psyb0t/go-spamcheck.git
cd go-spamcheck
go build
./go-spamcheck
Score: 3.9
Rules:
 [
  {
    "score": "0.0",
    "description": "ADMINISTRATOR NOTICE: The query to URIBL was blocked. See http://wiki.apache.org/spamassassin/DnsBlocklists#dnsbl-block for more information. [URIs: xxxpersonals.com]"
  },
  {
    "score": "0.8",
    "description": "BODY: Bayes spam probability is 40 to 60% [score: 0.5000]"
  },
  {
    "score": "0.0",
    "description": "BODY: Test for Invalidly Named or Formatted Colors in HTML"
  },
  {
    "score": "0.8",
    "description": "BODY: HTML and text parts are different"
  },
  {
    "score": "0.7",
    "description": "BODY: Message only has text/html MIME parts"
  },
  {
    "score": "0.0",
    "description": "BODY: HTML included in message"
  },
  {
    "score": "0.1",
    "description": "Message has a DKIM or DK signature, not necessarily valid"
  },
  {
    "score": "0.0",
    "description": "Multipart message only has text/html MIME parts"
  },
  {
    "score": "0.1",
    "description": "DKIM or DK signature exists, but is not valid"
  },
  {
    "score": "1.4",
    "description": "Missing Date: header"
  }
]
Report:
  pts rule                    description
---- ----------------------  --------------------------------------------------
 0.0 URIBL_BLOCKED           ADMINISTRATOR NOTICE: The query to URIBL was
                             blocked. See
                             http://wiki.apache.org/spamassassin/DnsBlocklists…
                             #dnsbl-block for more information. [URIs:
                             xxxpersonals.com]
 0.8 BAYES_50                BODY: Bayes spam probability is 40 to 60% [score:
                             0.5000]
 0.0 T_KAM_HTML_FONT_INVALID BODY: Test for Invalidly Named or Formatted Colors
                             in HTML
 0.8 MPART_ALT_DIFF          BODY: HTML and text parts are different
 0.7 MIME_HTML_ONLY          BODY: Message only has text/html MIME parts
 0.0 HTML_MESSAGE            BODY: HTML included in message
 0.1 DKIM_SIGNED             Message has a DKIM or DK signature, not
                             necessarily valid
 0.0 MIME_HTML_ONLY_MULTI    Multipart message only has text/html MIME parts
 0.1 DKIM_INVALID            DKIM or DK signature exists, but is not valid
 1.4 MISSING_DATE            Missing Date: header
```
