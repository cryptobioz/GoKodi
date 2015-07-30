package main

import (
  "os"
  "os/exec"
  "net/url"
)

func watch_movie(host string, movie string){
  file, err := get_movie(host, movie)
  if err != ""{
    display_error("BAD_MOVIE_NAME", movie)
  }else{
    display_info("Your movie is running ...")
    file = url.QueryEscape(file)
    var link string = "http://"+host+"/vfs/"+file
    cmd := exec.Command("vlc", link)
    cmd.Run()
  }
}

func watch_serie(host string, serie string, episode string){
  /*id, err := get_serie(host, serie)
  if err != ""{
    display_error("BAD_SERIE_NAME", serie)
  }else{
    file, err := get_episode(host, serie, episode)
    if err != ""{
        display_error("BAD_EPISODE_NAME", episode)
    }else{
      display_info("Your episode is running ...")
      file = url.QueryEscape(file)
      var link string = "http://"+host+"/vfs/"+file
      cmd := exec.Command("vlc "+link)
      cmd.Run()
    }
  }*/
}
func main() {

  if(len(os.Args) < 4){
    display_help()
  }else{
    var host string = os.Args[1]
    switch os.Args[2] {
    case "list":
      switch os.Args[3]{
      case "movies":
        display_list_movies(host)
      case "series":
        display_list_series(host)
      case "episodes":
        display_list_episodes()
      default:
        display_error("BAD_MEDIA", os.Args[3])
      }
    case "watch":
      if len(os.Args) == 5{
        //watch_serie(host, os.Args[3], os.Args[4])
      }else{
        watch_movie(host, os.Args[3])
      }

    default:
      display_error("BAD_OPTIONS", os.Args[2])
    }
  }
}
