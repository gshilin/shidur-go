// CGO_ENABLED=0 go build -o shidur-slides .; strip shidur-slides && upx -9 shidur-slides && cp shidur-slides /media/sf_D_DRIVE/projects/

package main

import (
	"os"

	"shidur-go/config"
	"shidur-go/routes"
)

func main() {

	app := config.NewApp(".")

	routes.Setup(app)
	app.Negroni.Run(":" + os.Getenv("PORT"))
}
