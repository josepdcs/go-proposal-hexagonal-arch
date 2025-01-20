package testutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApp(t *testing.T) {
	a := App()
	assert.NotNil(t, a)
}

func TestAcquireFiberCtx(t *testing.T) {
	app := App()
	defer Shutdown(app)

	ctx := AcquireFiberCtx(app)
	assert.NotNil(t, ctx)
}
