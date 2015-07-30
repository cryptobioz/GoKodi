package main

import (
  "net/http"
  "io/ioutil"
  "encoding/json"
  "github.com/bitly/go-simplejson"
  "log"
)


type ResponseListMovies struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Limits struct {
			End   int `json:"end"`
			Start int `json:"start"`
			Total int `json:"total"`
		} `json:"limits"`
		Movies []struct {
			File    string `json:"file"`
			Label   string `json:"label"`
			Movieid int    `json:"movieid"`
		} `json:"movies"`
	} `json:"result"`
}


type ResponseListSeries struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Limits struct {
			End   int `json:"end"`
			Start int `json:"start"`
			Total int `json:"total"`
		} `json:"limits"`
		Tvshows []struct {
			Label         string `json:"label"`
			Originaltitle string `json:"originaltitle"`
			Title         string `json:"title"`
			Tvshowid      int    `json:"tvshowid"`
		} `json:"tvshows"`
	} `json:"result"`
}



func get_list_movies(host string) ([]string){
  data := []byte(`{"jsonrpc": "2.0", "method": "VideoLibrary.GetMovies", "params": { "filter": {"field": "playcount", "operator": "is", "value": "0"}, "limits": { "start" : 0, "end": 75 }, "properties" : ["file"], "sort": { "order": "ascending", "method": "label", "ignorearticle": true } }, "id": "libMovies"}`)
  js, err := simplejson.NewJson(data)
  if err != nil{
    log.Fatal(err)
  }
  b, _ := js.Encode()
  res, err := http.Get("http://"+host+"/jsonrpc?request="+string(b))
  if err != nil{
    log.Fatal(err)
  }
  defer res.Body.Close()
  body, _ := ioutil.ReadAll(res.Body)
  response := &ResponseListMovies{}
  json.Unmarshal([]byte(body), &response)

  var movies []string
  for _, movie := range response.Result.Movies{
    movies = append(movies, movie.Label)
  }
  return movies
}

func get_movie(host string, media string) (file string, flag string){
  data := []byte(`{"jsonrpc": "2.0", "method": "VideoLibrary.GetMovies", "params": { "filter": {"field": "playcount", "operator": "is", "value": "0"}, "limits": { "start" : 0, "end": 75 }, "properties" : ["file"], "sort": { "order": "ascending", "method": "label", "ignorearticle": true } }, "id": "libMovies"}`)
  js, err := simplejson.NewJson(data)
  if err != nil{}
  b, _ := js.Encode()
  res, err := http.Get("http://"+host+"/jsonrpc?request="+string(b))
  if err != nil{}
  defer res.Body.Close()
  body, _ := ioutil.ReadAll(res.Body)
  response := &ResponseListMovies{}
  json.Unmarshal([]byte(body), &response)

  for _, movie := range response.Result.Movies{
    if(movie.Label == media){
      return movie.File, ""
    }
  }
  return "", "This movie isn't in Media Center."
}

func get_list_series(host string) ([]string){
  data := []byte(`{"jsonrpc": "2.0", "method": "VideoLibrary.GetTVShows", "params": { "filter": {"field": "playcount", "operator": "is", "value": "0"}, "limits": { "start" : 0, "end": 1000 }, "properties": ["title"], "sort": { "order": "ascending", "method": "label" } }, "id": "libTvShows"}`)
  js, err := simplejson.NewJson(data)
  if err != nil{
    log.Fatal(err)
  }
  b, _ := js.Encode()
  res, err := http.Get("http://"+host+"/jsonrpc?request="+string(b))
  if err != nil{
    log.Fatal(err)
  }
  defer res.Body.Close()
  body, _ := ioutil.ReadAll(res.Body)
  response := &ResponseListSeries{}
  json.Unmarshal([]byte(body), &response)
  var series []string
  for _, serie := range response.Result.Tvshows{
    series = append(series, serie.Label)
  }
  return series
}

func get_serie(host string, media string) (id int, flag string){
  data := []byte(`{"jsonrpc": "2.0", "method": "VideoLibrary.GetTVShows", "params": { "filter": {"field": "playcount", "operator": "is", "value": "0"}, "limits": { "start" : 0, "end": 1000 }, "properties": ["title"], "sort": { "order": "ascending", "method": "label" } }, "id": "libTvShows"}`)
  js, err := simplejson.NewJson(data)
  if err != nil{}
  b, _ := js.Encode()
  res, err := http.Get("http://"+host+"/jsonrpc?request="+string(b))
  if err != nil{}
  defer res.Body.Close()
  body, _ := ioutil.ReadAll(res.Body)
  response := &ResponseListSeries{}
  json.Unmarshal([]byte(body), &response)

  for _, serie := range response.Result.Tvshows{
    if(serie.Label == media){
      return serie.Tvshowid, ""
    }
  }
  return 0, "Error"
}

func get_list_episodes(host string) ([]string){
  data := []byte(`{"jsonrpc": "2.0", "method": "VideoLibrary.GetTVShows", "params": {"filter": {"field": "playcount", "operator": "is", "value": "0"}, "limits": { "start" : 0, "end": 1000 }, "properties": ["title"], "sort": { "order": "ascending", "method": "label" } }, "id": "libTvShows"}`)
  js, err := simplejson.NewJson(data)
  if err != nil{
    log.Fatal(err)
  }
  b, _ := js.Encode()
  res, err := http.Get("http://"+host+"/jsonrpc?request="+string(b))
  if err != nil{
    log.Fatal(err)
  }
  defer res.Body.Close()
  body, _ := ioutil.ReadAll(res.Body)
  response := &ResponseListSeries{}
  json.Unmarshal([]byte(body), &response)
  var series []string
  for _, serie := range response.Result.Tvshows{
    series = append(series, serie.Label)
  }
  return series
}
