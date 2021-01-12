package main

import(
  "net/url"
  "strings"
  "errors"
)

// Checks single input field not blank
func ValidateInput(i string) bool {
  if len(strings.ReplaceAll(i, " ", "")) == 0{
    return false
  }
  return true
}


func ValidateRegisterForm(form url.Values) error {
  errorMsg := ""
  // check fields not blank
  for k, v := range form{
    if !ValidateInput(v[0]){
      errorMsg = errorMsg + "Please enter a value for " + k + "\n"
    }
  }

  // check username available
  _, ok := Users[form["username"][0]]
  if ok{
    errorMsg = errorMsg + "Username unavailable\n"
  }

  // check passwords match
  if form["password"][0] != form["passwordAgain"][0]{
    errorMsg = errorMsg + "Passwords do not match\n"
  }

  if errorMsg == ""{
    return nil
  } else {
      return errors.New(errorMsg)
  }
}


func ValidateLoginForm(form url.Values) error {
  errorMsg := ""
  // check fields not blank
  for k, v := range form{
    if !ValidateInput(v[0]){
      errorMsg = errorMsg + "Please enter a value for " + k + "\n"
    }
  }

  if errorMsg == ""{
    return nil
  } else {
      return errors.New(errorMsg)
  }

}
