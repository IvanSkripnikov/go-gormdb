package gormdb

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	openConnectionOnDB = promauto.NewCounter(prometheus.CounterOpts{
		Name: "success_open_new_connection_on_db_service_total",
		Help: "The total number of open connections to the service database",
	})
	failedOpenConnectionOnDB = promauto.NewCounter(prometheus.CounterOpts{
		Name: "failed_connection_opening_on_db_service_total",
		Help: "The total number of unsuccessful connections to the service database",
	})
	appliedMigrations = promauto.NewCounter(prometheus.CounterOpts{
		Name: "applied_migrations_total",
		Help: "The total number of migrations applied",
	})
	failedApplyMigrations = promauto.NewCounter(prometheus.CounterOpts{
		Name: "failed_apply_migrations_total",
		Help: "The total number of migrations failed to apply",
	})
)
