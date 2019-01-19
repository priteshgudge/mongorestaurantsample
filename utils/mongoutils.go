package mongoutils

import (
	"bitbucket.org/agrostar/realtime-dashboard/utils/loggerutils"
	"gopkg.in/mgo.v2"
)

type MongoAuthObject struct {
	DBname string
}

//MongoSession stores mongo session
var MongoSession *mgo.Session

//RegisterMongoSession creates a new mongo session
func RegisterMongoSession(mongoURI string) (*mgo.Session, error) {
	logger := loggerutils.GetLogger()
	var err error
	MongoSession, err = mgo.Dial(mongoURI)
	if err != nil {
		logger.Errorf("Error in Connecting Mongo")
		panic(err)
	}
	return MongoSession, nil
}
