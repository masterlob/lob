package infrastructure

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestApplication_StartupAndStop(t *testing.T) {
	app := newApplication(AppOptions{IsProduction: false})
	app.Start()

	client := http.Client{}
	require.Eventually(t,
		func() bool {
			req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("http://localhost%s/", app.Addr), nil)
			resp, err := client.Do(req)
			return err == nil && resp.StatusCode == http.StatusNotFound
		},
		2*time.Second,
		5*time.Millisecond,
	)

	require.NoError(t, app.Stop())
}
