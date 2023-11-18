package ulid

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestJson(t *testing.T) {

	type foo struct {
		Val     ULID
		Pointer *ULID
	}

	f := foo{}

	err := json.Unmarshal([]byte(`{"Val": "01FFQY1GMHXBC29AYD6SMVMQGX", "Pointer": "01FFQY1FQDMCNEWBQMGHHV7P2J"}`), &f)
	require.NoError(t, err)

	err = json.Unmarshal([]byte(`{}`), &f)
	require.NoError(t, err)

}
