package main

import(
  "net/http"
  "github.com/julienschmidt/httprouter"
  "html/template"
  "time"
  "strconv"
  "log"
)

var tpl *template.Template

func init(){
    tpl = template.Must(template.ParseGlob("templates/*"))
}

// Index route "/"
func index(w http.ResponseWriter, req *http.Request, _ httprouter.Params){
  user, err := getUser(req)
  if err != nil{
    http.Redirect(w, req, "/login", http.StatusSeeOther)
    return
  }

  data := struct{
                  IsLoggedIn bool
                  FName string
                  Notes notesList
                }{true, user.FName, user.Notes}

  err = tpl.ExecuteTemplate(w, "index.gohtml", data)
  HandleError(w, err)
}


// Add note route "/add"
func add(w http.ResponseWriter, req *http.Request, _ httprouter.Params){
  user, err := getUser(req)
  if err != nil{
    http.Redirect(w, req, "/login", http.StatusSeeOther)
    return
  }

  input := req.FormValue("note")

  if !ValidateInput(input){
      http.Redirect(w, req, "/", http.StatusSeeOther)
  }

  newNote := note{len(user.Notes) + 1, input, time.Now().Format("02-Jan-2006 15:04:05")}

  user.Notes = append(user.Notes, newNote)

  Users[user.username] = user

  http.Redirect(w, req, "/", http.StatusSeeOther)
}



// Edit note route "/edit"
func edit(w http.ResponseWriter, req *http.Request, _ httprouter.Params){
  user, err := getUser(req)
  if err != nil{
    http.Redirect(w, req, "/login", http.StatusSeeOther)
    return
  }

  id, _ := strconv.Atoi(req.FormValue("id"))
  input := req.FormValue("note")

  user.Notes = user.Notes.EditNote(id, input)

  Users[user.username] = user

  http.Redirect(w, req, "/", http.StatusSeeOther)
}


// Delete note route "/delete"
func delete(w http.ResponseWriter, req *http.Request, _ httprouter.Params){
  user, err := getUser(req)
  if err != nil{
    http.Redirect(w, req, "/login", http.StatusSeeOther)
    return
  }

  id , _ :=  strconv.Atoi(req.FormValue("Id"))
  user.Notes = user.Notes.DeleteNote(id)

  Users[user.username] = user

  http.Redirect(w, req, "/", http.StatusSeeOther)
}


// Error handler
func HandleError(w http.ResponseWriter, err error){
  if err != nil{
    http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
  }
}
