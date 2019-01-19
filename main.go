package main

import (
	"fmt"
	"os"

	dbrepo "github.com/priteshgudge/mongosample/dbrepository"
	mongoutils "github.com/priteshgudge/mongosample/utils"
)

func main() {
	//pass mongohost through the environment
	mongoSession, _ := mongoutils.RegisterMongoSession(os.Getenv("MONGO_HOST"))

	dbname := "restaurant"
	repoaccess := dbrepo.NewMongoRepository(mongoSession, dbname)
	fmt.Println(repoaccess)
	//Run sample commands
}
