package scanners_test

import (
	"github.com/roadsigns/httpct/pkg/http"
	"github.com/roadsigns/httpct/pkg/scanners"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestXFrameOptionsExistsAndIsValidWithDeny(t *testing.T) {
	var headers http.Headers
	headers.AddHeader(http.Header{
		Title:   "X-Frame-Options",
		Content: "DENY",
	})

	securityHeader := scanners.XFrameOptions(headers)
	assert.True(t, securityHeader.Exists)
	assert.True(t, securityHeader.Valid)
}

func TestXFrameOptionsExistsAndIsValidWithSameOrigin(t *testing.T) {
	var headers http.Headers
	headers.AddHeader(http.Header{
		Title:   "X-Frame-Options",
		Content: "SAMEORIGIN",
	})

	securityHeader := scanners.XFrameOptions(headers)
	assert.True(t, securityHeader.Exists)
	assert.True(t, securityHeader.Valid)
}

func TestXFrameOptionsExistsAndIsNotValid(t *testing.T) {
	var headers http.Headers
	headers.AddHeader(http.Header{
		Title:   "X-Frame-Options",
		Content: "",
	})

	securityHeader := scanners.XFrameOptions(headers)
	assert.True(t, securityHeader.Exists)
	assert.False(t, securityHeader.Valid)
}

func TestXFrameOptionsDoesNotExist(t *testing.T) {
	var headers http.Headers
	headers.AddHeader(http.Header{
		Title:   "X-Fake-Header",
		Content: "Not Valid",
	})

	securityHeader := scanners.XFrameOptions(headers)
	assert.False(t, securityHeader.Exists)
	assert.False(t, securityHeader.Valid)
}
