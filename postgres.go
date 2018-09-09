package paratils

import (
	"database/sql"

	"github.com/ParaServices/paralog"
	"github.com/magicalbanana/dbal"
)

// SetupPostgresConnection conencts to a postgres database while pinging the
// database to ensure that the connection is ative
func SetupPostgresConnection(dbCreds string, dbPingTime int, lgr paralog.Logger) (dbal.DBAL, error) {
	db, openErr := sql.Open("postgres", dbCreds)
	if openErr != nil {
		return nil, openErr
	}
	lgrFunc := func(msg string) {
		lgr.Info(msg)
	}
	pingErr := dbal.PingDatabase(db, dbPingTime, lgrFunc)
	if pingErr != nil {
		return nil, pingErr
	}
	return dbal.New(db), nil
}
