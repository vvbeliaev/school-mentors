package main

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"

	_ "school-mentors/migrations"
)

func main() {
	godotenv.Load("../.env")

	app := pocketbase.New()

    app.OnServe().BindFunc(func(se *core.ServeEvent) error {
        se.Router.GET("/{path...}", apis.Static(os.DirFS("./pb_public"), false))

        return se.Next()
    })

    isGoRun := strings.HasPrefix(os.Args[0], os.TempDir())
    migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
        Automigrate: isGoRun,
    })

    if err := app.Start(); err != nil {
        log.Fatal(err)
    }
}