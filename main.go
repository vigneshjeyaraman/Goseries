package main
import (
  "net/http"
  "fmt"
  "encoding/json"
)

type Response struct  {
  UserId int `json:"userId"`
  Id int `json: "id"`
  Title string `json:"title"`
  Completed bool `json:"completed"`
}



func main(){
  resp, _ := http.Get("https://jsonplaceholder.typicode.com/todos")
  todos := make([]Response, 0)
  json.NewDecoder(resp.Body).Decode(&todos)
	fmt.Println(todos)
}
