package mime

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHeaderUserAgent(t *testing.T) {
	assert.Equal(t, "User-Agent", HeadUserAgent)
}
func TestHeaderContentType(t *testing.T) {
	assert.Equal(t, "Content-Type", HeaderContentType)
}
func TestContentTypeJson(t *testing.T) {
	assert.Equal(t, "application/json", ContentTypeJson)
}
func TestContentTypeXml(t *testing.T) {
	assert.Equal(t, "application/xml", ContentTypeXml)
}