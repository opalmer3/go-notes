package main

import(
  "net/http"
  "github.com/julienschmidt/httprouter"
  "log"
)

func main()  {
  router := httprouter.New()
  router.ServeFiles("/static/*filepath", http.Dir("static"))

  // Authentication
  router.GET("/register", register)
  router.POST("/register", handleRegister)
  router.GET("/login", login)
  router.POST("/login", handleLogin)
  router.GET("/logout", logout)
  
  router.GET("/", index)
  router.POST("/add", add)
  router.POST("/edit", edit)
  router.POST("/delete", delete)

  log.Fatal(http.ListenAndServe(":8080", router))
}
