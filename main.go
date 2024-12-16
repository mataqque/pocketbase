package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
        app := pocketbase.New()

    // serves static files from the provided public dir (if exists)
    app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
        e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./pb_public"), false))
        
        e.Router.GET("/hola", func(c echo.Context) error {
            //get list from database
            fmt.Println("prove")
            return c.JSON(http.StatusOK, map[string]string{"message": "Hola mundo"})
        }, apis.ActivityLogger(app))
        
        return nil
    })

    if err := app.Start(); err != nil {
        log.Fatal(err)
    }
    
}