package main

import(
  "github.com/julienschmidt/httprouter"
  "gopkg.in/mgo.v2"
  "net/http"

  "github.com/anish2604/instagram-backend-api/controllers"
)

func main(){

  r := httprouter.New()
  uc := controllers.NewUserController(getSession())
  r.GET("/users/:id", uc.GetUser)
  r.POST("/users", uc.CreateUser)
  http.ListenAndServe("localhost:3000", r)
}

func getSession() *mgo.Session{

  s, err := mgo.Dial("mongodb://localhost:27017")
  if err != nil {
      panic(err)
  }
  return s
}
