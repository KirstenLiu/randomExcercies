package main

import (
	"bufio"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gophercises/cyoa"
)

type handlerWrapper struct {
	filename string
}

func CustomPathFn(r *http.Request) string{
	going := strings.TrimSpace(r.URL.Path)

	if r.URL.Path == "/story/" || r.URL.Path == "/story"{
		going = "intro"
	}else{
		//going = going[len("story")+2:]
		going = going[1:]
	}
	//fmt.Println(going)
	//going = going[1:]
	fmt.Println("customPathFn:")
	fmt.Println(going)

	return going
}

func main(){
	fmt.Println("hi for main")
	//bonus: cmd game
	jsonPath := flag.String("input", "/Users/kirstenliu/go/src/gophercises/cyoa/gopher.json", "give json file path:")
	//portFirstTry := 8080
	portByExample := 3030
	port := flag.Int("port", portByExample, "give a port: ")
	stories := cyoa.ReadJson(*jsonPath)
	//storyWalker(story)
	
	//try copy internet example
	//https://lets-go.alexedwards.net/sample/02.09-serving-static-files.html

	//fmt.Println(stories)
	out := fmt.Sprintf("localhost:%d", *port)
	fmt.Println("http server second try starting:", out)
	/*tmplStr := `
	<!DOCTYPE html>
<html>
    <head>
        <title>Choose Your Own Advanture</title>
    </head>
    <body>
        <section class="page">
            <h1>
                {{.Title}}
            </h1>
            {{range .Story}}
                <p>{{.}}</p>
            {{end}}

            {{if .Options}}
                <ul>
                    {{range .Options}}
                        <li>
                            <a href="{{.Arc}}">{{.Text}} </a>
                        </li>
                    {{end}}
                
                </ul>
            {{else}}
                <h3>~The End~</h3>
            {{end}}
        </section>
    </body>
</html>
	`
	t := template.Must(template.New("").Parse(tmplStr))
	handler2 := cyoa.NewHandler(stories, cyoa.OptTmpl(t), cyoa.OptFcPath(CustomPathFn))*/
	handler2 := cyoa.NewHandler(stories, cyoa.OptFcPath(CustomPathFn))
	
	log.Fatal(http.ListenAndServe(fmt.Sprintf("localhost:%d", *port), handler2)) 
	
	/*fmt.Println("http server first try starting:")
	handler := handlerWrapper{filename: *jsonPath}
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.homePageHandler)
	//mux.HandleFunc("/new-york", newYorkHandler)
	http.ListenAndServe(portFirstTry, mux)*/

	
}

/*func findChapter(handle string, stories cyoa.Story) cyoa.Chapter{
	for _, s := range(stories){
		if s.Handle == handle {
			return s.Content
		}
	}

	fmt.Println("can't find the defined chapter: ", handle)
	err := cyoa.Chapter{
		Title: "NOFOUND",
	}
	return err
}*/



func (hw handlerWrapper) homePageHandler(w http.ResponseWriter, r *http.Request){
	stories := cyoa.ReadJson(hw.filename)

	going := r.URL.Path

	if going != "/" {
		going = strings.Trim(going, "/")
	}

	//story := findChapter(going, )

	data := cyoa.Chapter{
		Title: stories[going].Title,
		Story: stories[going].Story,
		Options: stories[going].Options,
	}

	fmt.Println(going)

	tmpl, err := template.ParseFiles("./intro.html")
	if err != nil {
		PrintErr := fmt.Errorf("Problem when creating template of home.html: %v", err)
		fmt.Println(PrintErr.Error())
		return
	}
	
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
	
	//going = r.URL.Path
	//fmt.Println(going)
}

func storyWalker(story cyoa.Stories){
	going := "intro"
	arcNum := 1

	for {
		fmt.Println(story[going].Title)
		for _, txt := range(story[going].Story){
			fmt.Println(txt)
		}

		if len(story[going].Options) == 0 {
			fmt.Println("End of story. Greetings.")
			break
		}

		for _, ops := range(story[going].Options) {
			fmt.Println(ops.Text)
			//fmt.Println("choose this route by hitting following archive:", ops.Arc)
			fmt.Println("choose this route by hitting following number:", arcNum)
			arcNum++
		}

		fmt.Println("Select your choice from above archives:")

		reader := bufio.NewReader(os.Stdin)
		arcChoice, err := reader.ReadString('\n')
		if err != nil {
			PrintErr := fmt.Errorf("can't read input choice from cmd: %v", err)
			fmt.Println(PrintErr.Error())
		}
		chooseNum, _ := strconv.Atoi(strings.TrimSuffix(arcChoice, "\n"))
		
		going = story[going].Options[chooseNum-1].Arc
		fmt.Println("Going to this route: ", story[going].Title)
		arcNum = 1
	}
}


//func newYorkHandler(w http.ResponseWriter, r *http.Request){
//	fmt.Println("Helllloooooo NY")
	/*type data struct {
		Key string
		New string
	}

	var d data
	d.Key = r.URL.Query().Get("key")
	d.New = r.URL.Query().Get("new")

	jData, _ := json.Marshal(d)*/

	/*jData, _ := json.Marshal(r.URL.Query())

	w.Write([]byte(jData))*/	
//}