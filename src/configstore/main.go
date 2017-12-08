package main

import (
	"log"
	"os"

	"configstore/db"
	"configstore/views"

	"github.com/gin-gonic/gin"
)

// configuration should really be in a file
const (
	DBConnect = "host=localhost user=postgres dbname=db sslmode=disable password=pass"
	Listen    = "127.0.0.1:8078"
)

const (
	usageString = "usage: configstore [rollback_last|migrate|run]"
)

func main() {

	dbconn, err := db.NewPostgresConfigStore(DBConnect)
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) <= 1 {
		log.Fatal(usageString)
	}

	switch os.Args[1] {

	case "rollback_last":
		err = dbconn.RollbackLast()
		if err != nil {
			log.Fatal(err)
		}

	case "migrate":
		err = dbconn.Migrate()
		if err != nil {
			log.Fatal(err)
		}

	case "run":

		view := views.ConfigView{Store: dbconn}
		router := gin.Default()

		router.POST("/get_config", GinAdapter(views.GetRequest{}, view.Get))

		log.Fatal(router.Run(Listen))

	default:

		log.Fatal(usageString)

	}
}
