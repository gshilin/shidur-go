package main

import (
  "github.com/gshilin/shidur-go/config"
  "github.com/gshilin/shidur-go/routes"
  "os"
)

func main() {

  app := config.NewApp(".")

  routes.Setup(app)
  app.Negroni.Run(":" + os.Getenv("PORT"))
}
