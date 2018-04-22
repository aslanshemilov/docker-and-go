package main

import (
  "github.com/EwanValentine/project/api/controllers"
  "github.com/EwanValentine/project/api/models:
  "github.com/go-martini/martini"
  "github.com/martini-contrib/binding"
  "github.com/martini-contrib/render"
)

func main() {
  
  m := martini.Classic()
  m.Map(models.Database())
  m.Use(rend.Renderer())
  
  pc := controllers.NewProductControllers(models.Database())
  
  m.Get("/products", binding.Bind(models.Product{}), pc.GetAllProducts)
  m.Post("/products", binding.Bnd(models.Product{}), pc.PostProduct)
  m.Run()
}
