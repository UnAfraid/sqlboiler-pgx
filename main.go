package main

import (
	"github.com/UnAfraid/sqlboiler-pgx/driver"
	"github.com/volatiletech/sqlboiler/v4/drivers"
)

func main() {
	drivers.DriverMain(&driver.PostgresDriver{})
}
