package manage

import (
	"broadcastle.co/code/crm/code/db"
	"github.com/sirupsen/logrus"
)

// Init starts the database.
func Init() {
	logrus.Debug("starting the usage of the database")
	db.Init()
}

// Close ends the manage package.
func Close() {
	logrus.Debug("ending the usage of the database")
	db.Close()
}
