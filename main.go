package main
import (
  "encoding/json"
  "fmt"
  "strings"
  
)
type book_read struct {
  Name string `json:"name"`
  Author string `json:"author"`
}
type decode struct{
  Name string `json:"name"`
  PhoneNumber []string `json:"phone_number"`
  FavoriteBand []string `json:"favorite_band"`
  BooksRead []book_read `json:"books_read"`
  Address struct {
    City string `json:"city"`
    Pincode string `json:"pincode"`
  }
  
}
func main(){
  json_ := `
  {
    "name": "Vignesh",
    "phone_number": [
      "xxxxxxxx29",
      "xxxxxxxx34"
    ],
    "favorite_band": [
      "gnr",
      "metallica",
      "pink floyd"
    ],
    "books_read": [
      {
        "name": "only love is read",
        "author": "Brian weiss"
      },
      {
        "name": "many lives many master",
        "author": "Brian weiss"
      }
    ],
    "address": {
      "city": "noida",
      "pincode": "201304"
    }
  }`
  
  var d decode
  json.NewDecoder(strings.NewReader(json_)).Decode(&d)
  fmt.Println(d)
}
