package scanners_test

import (
	"github.com/roadsigns/httpct/pkg/http"
	"github.com/roadsigns/httpct/pkg/scanners"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestXXssProtectionExistsAndIsValid(t *testing.T) {
	var headers http.Headers
	headers.AddHeader(http.Header{
		Title:   "X-Xss-Protection",
		Content: "1; mode=block",
	})

	securityHeader := scanners.XXssProtection(headers)
	assert.True(t, securityHeader.Exists)
	assert.True(t, securityHeader.Valid)
}

func TestXXssProtectionExistsAndIsValidWithZero(t *testing.T) {
	var headers http.Headers
	headers.AddHeader(http.Header{
		Title:   "X-Xss-Protection",
		Content: "0",
	})

	securityHeader := scanners.XXssProtection(headers)
	assert.True(t, securityHeader.Exists)
	assert.True(t, securityHeader.Valid)
}

func TestXXssProtectionExistsAndIsValidWithOne(t *testing.T) {
	var headers http.Headers
	headers.AddHeader(http.Header{
		Title:   "X-Xss-Protection",
		Content: "1",
	})

	securityHeader := scanners.XXssProtection(headers)
	assert.True(t, securityHeader.Exists)
	assert.True(t, securityHeader.Valid)
}

func TestXXssProtectionExistsAndIsNotValid(t *testing.T) {
	var headers http.Headers
	headers.AddHeader(http.Header{
		Title:   "X-Xss-Protection",
		Content: "Not Valid",
	})

	securityHeader := scanners.XXssProtection(headers)
	assert.True(t, securityHeader.Exists)
	assert.False(t, securityHeader.Valid)
}

func TestXXssProtectionDoesNotExist(t *testing.T) {
	var headers http.Headers
	headers.AddHeader(http.Header{
		Title:   "X-Fake-Header",
		Content: "Not Valid",
	})

	securityHeader := scanners.XXssProtection(headers)
	assert.False(t, securityHeader.Exists)
	assert.False(t, securityHeader.Valid)
}
