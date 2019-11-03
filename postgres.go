package paratils

import (
	"database/sql"
	"net/url"
	"time"

	"github.com/ParaServices/paralog"
	"github.com/magicalbanana/dbal"
	"go.uber.org/zap"
)

type PostgresConfig struct {
	URL             url.URL
	PingTime        int
	ConnMaxLifetime time.Duration
	MaxIdleConns    int
	MaxOpenConns    int
}

// SetupPostgresConnection conencts to a postgres database while pinging the
// database to ensure that the connection is ative
func SetupPostgresConnection(cfg *PostgresConfig, lgr paralog.Logger) (dbal.DBAL, error) {
	db, openErr := sql.Open("postgres", cfg.URL.String())
	if openErr != nil {
		return nil, openErr
	}
	lgrFunc := func(msg string) {
		lgr.Info(msg)
	}
	if lgr != nil {
		lgr.Info("Pinging dtabase", zap.String("host", cfg.URL.Host))
	}
	pingErr := dbal.PingDatabase(db, cfg.PingTime, lgrFunc)
	if pingErr != nil {
		return nil, pingErr
	}
	return dbal.New(db), nil
}
