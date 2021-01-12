package main

var Users = map[string]User{}
var Sessions = map[string]string{}

type User struct{
  username string
  FName string
  lName string
  passwordHash []byte
  Notes notesList
}


type note struct{
  Id int
  Content string
  Date string
}

type notesList []note


// notes list methods
func (n notesList) DeleteNote(id int) notesList{
  for i, note := range n{
    if note.Id == id{
      n = append(n[:i], n[i+1:]...)
      return n
    }
  }
  return n
}

func (n notesList) EditNote(id int, content string) notesList{
  for i, note := range n{
    if note.Id == id{
      n[i].Content = content
      return n
    }
  }
  return n
}
// ////////////
