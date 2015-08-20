package main

import (
  "github.com/gshilin/shidur-slides/config"
  "github.com/gshilin/shidur-slides/routes"
  "os"
)

func main() {
  app := config.NewApp(".")
  routes.Setup(app)
  app.Negroni.Run(":" + os.Getenv("PORT"))
}
