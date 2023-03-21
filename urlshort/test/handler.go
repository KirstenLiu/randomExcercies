package test

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis"
	"gopkg.in/yaml.v2"
)

/*func Test(){
    fmt.Println("testtest")
}*/
// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request){
        if redirectUrl, ok := pathsToUrls[r.URL.Path]; ok {
            http.Redirect(w, r, redirectUrl, http.StatusFound)
        } 
        fallback.ServeHTTP(w, r)
    }
}

func DbHandler(fallback http.Handler) http.HandlerFunc {
    redisClient := redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
        Password: "",
        DB: 0,
    })

    err := redisClient.Set("/urlshort", "https://github.com/gophercises/urlshort", 0).Err()
    if err != nil{
        fmt.Println(err)
    }

    err2 := redisClient.Set("/urlshort-final", "https://github.com/gophercises/urlshort/tree/solution", 0).Err()
    if err2 != nil{
        fmt.Println(err)
    }

    return func(w http.ResponseWriter, r *http.Request) {
        redirectUrl, err := redisClient.Get(r.URL.Path).Result()
        if err != nil {
            fmt.Println(err)
        }

        http.Redirect(w, r, redirectUrl, http.StatusFound)
        fallback.ServeHTTP(w, r)
    }
}

/*func DbHandler(fallback http.Handler) http.HandlerFunc{
    //database, err := NewDataBase(RedisAddr)
    client := redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
        Password: "",
        DB: 0,
    })
    
    err := client.Set("/urlshort", "https://github.com/gophercises/urlshort", 0).Err()
    if err != nil {
        log.Fatal(err)
    }
    err = client.Set("/urlshort-final", "https://github.com/gophercises/urlshort/tree/solution", 0).Err()
    if err != nil {
        log.Fatal(err)
    }

    return func (w http.ResponseWriter, r *http.Request){
        path := r.URL.Path
        fmt.Println(path)
        redirect, err := client.Get(path).Result()
        //fmt.Println(redirect)
        //fmt.Println(err)
        if err == nil {
            http.Redirect(w, r, redirect, http.StatusFound)
        }
        fallback.ServeHTTP(w, r)
        log.Fatal(err)
    }
}*/

type Redirect struct {
    Path string
    Url string
}

func parseYaml(yml []byte)([]Redirect, error){
    var r []Redirect
    err := yaml.Unmarshal(yml, &r)
    if err != nil {
        log.Fatalf("error: %v", err)
    }
    return r, err
}

func buildMap(r []Redirect) map[string]string {
    pathMap := make(map[string]string)
    
    for _, rs := range(r){
        pathMap[rs.Path] = rs.Url
    } 
    return pathMap
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
    parsedYaml, err := parseYaml(yml)
    if err != nil {
        return nil, err
    }
    pathMap := buildMap(parsedYaml)
    return MapHandler(pathMap, fallback), nil
}

//json has array!!!![{"key":"value"}, {"key2":"value2"}]
func parseJson(j []byte) ([]Redirect, error){
    var r []Redirect

    err := json.Unmarshal(j, &r)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(r)
    return r, err
}

func JSONHandler(json []byte, fallback http.Handler) (http.HandlerFunc, error){
    parsedJson, err := parseJson(json)
    if err != nil{
        return nil, err
    }
    pathMap := buildMap(parsedJson)
    return MapHandler(pathMap, fallback), nil
}
