package chserver

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/ssh"

	"github.com/renatomb/open-rport/db/sqlite"
	"github.com/renatomb/open-rport/server/api/session"
	"github.com/renatomb/open-rport/server/bearer"
	chshare "github.com/renatomb/open-rport/share/logger"
)

var testLog = chshare.NewLogger("chserver-test", chshare.LogOutput{File: os.Stdout}, chshare.LogLevelDebug)
var hour = time.Hour

var DataSourceOptions = sqlite.DataSourceOptions{WALEnabled: false}

type mockConnection struct {
	ssh.Conn
	closed bool
}

func (m *mockConnection) Close() error {
	m.closed = true
	return nil
}

func newEmptyAPISessionCache(t *testing.T) *session.Cache {
	storage, err := session.NewSqliteProvider(":memory:", DataSourceOptions)
	require.NoError(t, err)
	c, err := session.NewCache(context.Background(), bearer.DefaultTokenLifetime, cleanupAPISessionsInterval, storage, nil)
	require.NoError(t, err)
	return c
}
