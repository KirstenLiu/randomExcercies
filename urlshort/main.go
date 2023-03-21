package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gophercises/urlshort/test"
)

func main() {
    //flag usage: -yml in cmd
    ymlFile := flag.String("yml", "none", "input yml file path")
    jsonFile := flag.String("json", "none", "input json file path")
    redisPath := flag.String("redis", "none", "redis client adress")
    
    flag.Parse()

    mux := defaultMux()

    // Build the MapHandler using the mux as the fallback
    pathsToUrls := map[string]string{
        "/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
        "/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
    }
    //test.Test()
    mapHandler := test.MapHandler(pathsToUrls, mux)

    // Build the YAMLHandler using the mapHandler as the
    // fallback
    /*yaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`*/
    //json := []byte(`[{"path":"/urlshort", "url":"https://github.com/gophercises/urlshort"}, {"path":"/urlshort-final", "url": "https://github.com/gophercises/urlshort/tree/solution"}]`)
    
    fmt.Println("before reading")
    if *ymlFile != "none" {
        yaml, err := ioutil.ReadFile(*ymlFile)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println(yaml)
        yamlHandler, err := test.YAMLHandler([]byte(yaml), mapHandler)
        if err != nil {
            panic(err)
        }
        fmt.Println("Starting the server on :8080")
        http.ListenAndServe(":8080", yamlHandler)
    }else if *jsonFile != "none" {
        json, err := ioutil.ReadFile(*jsonFile)
        jsonHandler, err := test.JSONHandler(json, mapHandler)
        if err != nil{
            log.Fatal(err)
        }
        fmt.Println("Starting the server on :8080")
        http.ListenAndServe(":8080", jsonHandler)
    }else if *redisPath != "none"{
        dbHandler := test.DbHandler(mapHandler)
        
        fmt.Println("Starting the server on :8080")
        http.ListenAndServe(":8080", dbHandler)
    }
}

func defaultMux() *http.ServeMux {
    mux := http.NewServeMux()
    mux.HandleFunc("/", hello)
    return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, world!")
}
