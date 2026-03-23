package container_test

import (
	"testing"

	"HATCH_APP/test/container"

	"github.com/stretchr/testify/require"
)

func Test_SetupPostgres(t *testing.T) {
	t.Parallel()

	db, teardown := container.SetupPostgres(t)
	defer teardown()

	require.NoError(t, db.Ping())
}
