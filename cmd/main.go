package main

import (
	"context"
	mysql "demo-service/client"
	"demo-service/datastores"
	"demo-service/repository"
	"demo-service/route"
	"log"
	"os"
	"time"
)

func main() {
	defer mysql.Disconnect()

	{
		loc, err := time.LoadLocation("Asia/Ho_Chi_Minh")
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		time.Local = loc
	}

	client := mysql.GetClient
	repo := repository.New(client)
	datastores.Migrate(client(context.TODO()))

	{
		h := route.NewHTTPHandler(repo)
		h.Logger.Fatal(h.Start(":3033"))
	}
}
