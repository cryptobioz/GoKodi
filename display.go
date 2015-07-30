package main

import (
  "os"
  "fmt"
)

func display_help(){
  fmt.Println("Usage: media-center HOST OPTIONS MEDIA")
  fmt.Println("Options :")
  fmt.Println("     list    => Display list of media.")
  fmt.Println("     watch   => Watch media")
  fmt.Println("Media: movies | series")
  os.Exit(0)
}

func display_error(flag string, data string){
  switch flag{
  case "BAD_OPTIONS":
    fmt.Println("Error : You have to choose an option in this list {list, watch}")
  case "BAD_MEDIA":
    fmt.Println("Error : You have to choose an option in this list {movies, series}")
  case "BAD_MOVIE_NAME":
    fmt.Println("Error : This movie isn't in Media Center.\nPlease run : mediacen list movies")
  case "BAD_SERIE_NAME":
    fmt.Println("Error : This tvshow isn't in Media Center.\nPlease run : mediacen list series")
  case "BAD_EPISODE_NAME":
    fmt.Println("Error : This episode isn't in Media Center.\nPlease run : mediacen lists series")
  }
}

func display_list_movies(host string){
  var movies []string = get_list_movies(host)
  for _, movie := range movies{
    fmt.Println(movie)
  }
}
func display_list_series(host string){
  var series []string = get_list_series(host)
  for _, serie := range series{
    fmt.Println(serie)
  }
}
func display_list_episodes(){

}

func display_info(message string){
  fmt.Println("[*] "+message)
}
