package gokhttp_sslpinning

import (
	"context"
	gokhttp "github.com/BRUHItsABunny/gOkHttp"
	gokhttp_requests "github.com/BRUHItsABunny/gOkHttp/requests"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewSSLPinningOption(t *testing.T) {
	pinner := NewSSLPinningOption()
	err := pinner.AddPin("github.com", true, "sha256\\Gs+dT9kUC17nDYZXH52mKzGnlUU/Q5mS0UruTQW3H0U=")
	require.NoError(t, err, "pinner.AddPin: errored unexpectedly.")

	hClient, err := gokhttp.NewHTTPClient(pinner)
	require.NoError(t, err, "NewHTTPClient: errored unexpectedly.")

	req, err := gokhttp_requests.MakeGETRequest(context.Background(), "https://github.com")
	require.NoError(t, err, "requests.MakeGETRequest: errored unexpectedly.")

	_, err = hClient.Do(req)
	require.NoError(t, err, "hClient.Do: errored unexpectedly.")
}
