package main

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
	"github.com/stripe/stripe-go/v84"

	"school-mentors/bookings"
	_ "school-mentors/migrations"
)

func main() {
	godotenv.Load("../.env")

    stripe.Key = os.Getenv("STRIPE_API_KEY")

	app := pocketbase.New()

    app.OnServe().BindFunc(func(se *core.ServeEvent) error {
        // Serve service worker
        se.Router.GET("/sw.js", func(e *core.RequestEvent) error {
            fsys := os.DirFS("./pb_public")
            e.Response.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
            e.Response.Header().Set("Pragma", "no-cache")
            e.Response.Header().Set("Expires", "0")
            return e.FileFS(fsys, "sw.js")
        })

        // SPA + Prerender fallback: serve static files, check for .html, or fallback to index.html
        se.Router.GET("/{path...}", func(e *core.RequestEvent) error {
            path := e.Request.PathValue("path")
            fsys := os.DirFS("./pb_public")

            // 1. Try to serve the static file as is (assets, direct matches)
            err := e.FileFS(fsys, path)
            if err == nil {
                return nil
            }

            // 2. Try to serve path.html (for prerendered routes)
            if path != "" && !strings.Contains(path, ".") {
                htmlErr := e.FileFS(fsys, strings.TrimSuffix(path, "/")+".html")
                if htmlErr == nil {
                    return nil
                }
            }

            // 3. Fallback to index.html for SPA routing, 
            // except for API and internal PocketBase routes.
            if !strings.HasPrefix(path, "api/") && !strings.HasPrefix(path, "_/") {
                return e.FileFS(fsys, "index.html")
            }

            return err
        })

        // Register custom API routes
        bookings.API(se, app)

        return se.Next()
    })

    isGoRun := strings.HasPrefix(os.Args[0], os.TempDir())
    migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
        Automigrate: isGoRun,
    })

	// Bookings
	bookings.Hooks(app)
	bookings.Crons(app)

    if err := app.Start(); err != nil {
        log.Fatal(err)
    }
}

