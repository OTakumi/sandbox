//go:build tools

package tools

import (
	_ "github.com/99designs/gqlgen"
	_ "github.com/volatiletech/sqlboiler/v4"
	_ "github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql"
	_ "github.com/golang-migrate/migrate/v4"
)
