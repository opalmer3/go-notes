package main

import(
  "net/http"
  "github.com/julienschmidt/httprouter"
  "golang.org/x/crypto/bcrypt"
)


func register(w http.ResponseWriter, req *http.Request,  _ httprouter.Params){
  if isLoggedIn(req){
    http.Redirect(w, req, "/", http.StatusSeeOther)
    return
  }

  err := tpl.ExecuteTemplate(w, "register.gohtml", nil)
  HandleError(w, err)
}


func handleRegister(w http.ResponseWriter, req *http.Request,  _ httprouter.Params){
  if isLoggedIn(req){
    http.Redirect(w, req, "/", http.StatusSeeOther)
    return
  }

  err := req.ParseForm()
  HandleError(w, err)

  err = ValidateRegisterForm(req.Form)

  if err != nil{
    err = tpl.ExecuteTemplate(w, "register.gohtml", err)
    return
  }

  username, fName, lName := req.FormValue("username"), req.FormValue("fName"), req.FormValue("lName")
  passwordHash, _ := bcrypt.GenerateFromPassword([]byte(req.FormValue("password")), bcrypt.DefaultCost)
  var notes notesList

  Users[username] = User{username, fName, lName, passwordHash, notes}

  createSession(w, username)


  http.Redirect(w, req, "/", http.StatusSeeOther)

}


func login(w http.ResponseWriter, req *http.Request,  _ httprouter.Params){
  if isLoggedIn(req){
    http.Redirect(w, req, "/", http.StatusSeeOther)
    return
  }

  err := tpl.ExecuteTemplate(w, "login.gohtml", nil)
  HandleError(w, err)
}


func handleLogin(w http.ResponseWriter, req *http.Request,  _ httprouter.Params){
  if isLoggedIn(req){
    http.Redirect(w, req, "/", http.StatusSeeOther)
    return
  }

  err := ValidateLoginForm(req.Form)

  if err != nil{
    err = tpl.ExecuteTemplate(w, "login.gohtml", err)
    return
  }

  un := req.FormValue("username")
  p := req.FormValue("password")

  user, ok := Users[un]
  if !ok {
    err = tpl.ExecuteTemplate(w, "login.gohtml", "Invalid credentials")
    return
  }

  err = bcrypt.CompareHashAndPassword(user.passwordHash, []byte(p))
  if err != nil{
    err = tpl.ExecuteTemplate(w, "login.gohtml", "Invalid credentials")
    return
  }

  createSession(w, un)

  http.Redirect(w, req, "/", http.StatusSeeOther)
}


func logout(w http.ResponseWriter, req *http.Request,  _ httprouter.Params){
  if !isLoggedIn(req){
    http.Redirect(w, req, "/", http.StatusSeeOther)
    return
  }

  c, _ := req.Cookie("session")

  destroySession(w, c)

  http.Redirect(w, req, "/login", http.StatusSeeOther)
}
