package main

import "fmt"
import "github.com/JustinBeckwith/go-yelp/yelp"

func main() {

  client := yelp.New(options)
  result, err := client.DoSimpleSearch("coffee", "seattle")


}
