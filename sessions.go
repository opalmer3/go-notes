package main

import(
  "net/http"
  "github.com/satori/go.uuid"
  "errors"
)

const(
	SessionLengthDays = 30
)

func createSession(w http.ResponseWriter, username string){

  sId,_ := uuid.NewV4()

  http.SetCookie(w, &http.Cookie{
    Name: "session",
    Value: sId.String(),
    MaxAge: SessionLengthDays * 86400,
  })

  Sessions[sId.String()] = username

}

func destroySession(w http.ResponseWriter, c *http.Cookie){

  // delete(Sessions, c.Value) Delete cookie from table

  http.SetCookie(w, &http.Cookie{
    Name: "session",
    Value: "",
    MaxAge: -1,
  })

}

func getUser(req *http.Request) (User, error){
  c, err := req.Cookie("session")
  if err != nil{
    return User{}, errors.New("User not logged in")
  }

  _, ok := Sessions[c.Value]
  if !ok{
    return User{}, errors.New("User not logged in")
  }

  un, _ := Sessions[c.Value]
  user, _ := Users[un]

  return user, nil
}

func isLoggedIn(req *http.Request) bool {
  c, err := req.Cookie("session")
  if err != nil{
    return false
  }

  _, ok := Sessions[c.Value]
  if !ok{
    return false
  }

  return true
}
