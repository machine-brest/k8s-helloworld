package srv

import (
	"github.com/go-pg/pg/v10"
	"github.com/labstack/echo/v4"
	"github.com/machine-brest/k8s-helloworld/srv/data"
	"net/http"
)

//Login Authorize user credentials
func Login(db *pg.DB) echo.HandlerFunc {

	if db == nil {
		panic("data where not instantiated")
	}

	return func(c echo.Context) error {

		var user data.User
		res, err := db.QueryOne(&user, "SELECT * FROM users WHERE name = ?", "usr1")
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound)
		}

		_ = res

		return nil
	}
}
