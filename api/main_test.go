package api

import (
	"os"
	"testing"
	"time"

	db "github.com/DMV-Nicolas/DevoraBank/db/sqlc"
	"github.com/DMV-Nicolas/DevoraBank/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store db.Store, tokenSymmetricKey string) *Server {
	config := util.Config{
		TokenSymmetricKey:   tokenSymmetricKey,
		AccessTokenDuration: time.Minute,
	}
	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
