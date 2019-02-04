## moov-io/base/admin

Package admin implements an `http.Server` which can be used for operations and monitoring tools. It's designed to be shipped (and ran) inside an existing Go service.

Here's an example of adding `admin.Server` to serve Prometheus metrics:

```Go
import (
    "fmt"
    "os"

    "github.com/moov-io/base/admin"

    "github.com/go-kit/kit/log"
)

var logger log.Logger

// in main.go or cmd/server/main.go

adminServer := admin.NewServer(*adminAddr)
go func() {
	logger.Log("admin", fmt.Sprintf("listening on %s", adminServer.BindAddr()))
	if err := adminServer.Listen(); err != nil {
		err = fmt.Errorf("problem starting admin http: %v", err)
		logger.Log("admin", err)
		// errs <- err // send err to shutdown channel
	}
}()
defer adminServer.Shutdown()
```
