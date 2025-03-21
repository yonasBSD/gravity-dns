package extconfig

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtconfig(t *testing.T) {
	assert.NotNil(t, Get())
	assert.NotNil(t, Get().Dirs())
	assert.NotNil(t, Get().EtcdClient())
	assert.True(t, strings.HasSuffix(Get().Listen(1234), "1234"))
}

func TestTransport(t *testing.T) {
	assert.NotNil(t, Transport())
}

func TestGetIP(t *testing.T) {
	ip, err := Get().GetIP()
	assert.NotNil(t, ip)
	assert.Nil(t, err)
}

func TestVersion(t *testing.T) {
	assert.Equal(t, Version+"+test", FullVersion())
}
