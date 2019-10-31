package paratils

import (
	"database/sql"
	"net/url"

	"github.com/ParaServices/paralog"
	"github.com/magicalbanana/dbal"
	"go.uber.org/zap"
)

// SetupPostgresConnection conencts to a postgres database while pinging the
// database to ensure that the connection is ative
func SetupPostgresConnection(dbCreds *url.URL, dbPingTime int, lgr paralog.Logger) (dbal.DBAL, error) {
	db, openErr := sql.Open("postgres", dbCreds.String())
	if openErr != nil {
		return nil, openErr
	}
	lgrFunc := func(msg string) {
		lgr.Info(msg)
	}
	if lgr != nil {
		lgr.Info("Pinging dtabase", zap.String("host", dbCreds.Host))
	}
	pingErr := dbal.PingDatabase(db, dbPingTime, lgrFunc)
	if pingErr != nil {
		return nil, pingErr
	}
	return dbal.New(db), nil
}
