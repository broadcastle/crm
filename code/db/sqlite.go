package db

import (
	"path"

	"github.com/jinzhu/gorm"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	_ "github.com/jinzhu/gorm/dialects/sqlite" // Needed in order to use sqlite.
)

// InitSQLite starts the sqlite database.
func InitSQLite() {

	var err error

	file := viper.GetString("db.path")

	if file == "" {

		home, err := homedir.Dir()
		if err != nil {
			logrus.Fatal(err)
		}

		file = path.Join(home, ".crm.db")

		logrus.Debugf("unable to get a database path from the configuration file\nusing default location of %s", file)

	}

	DB, err = gorm.Open("sqlite3", file)
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Debugf("using %s as database", file)
}
