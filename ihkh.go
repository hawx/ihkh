package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"hawx.me/code/ihkh/flickr"
	"hawx.me/code/ihkh/handler"
	"hawx.me/code/route"
	"hawx.me/code/serve"
)

func main() {
	const usage = `Usage: ihkh [OPTIONS]

  A minimalist Flickr viewer.

 OPTIONS
   --user-id ID      # Your Flickr userid, like XXXXXXXX@XXX
   --api-key KEY     # Your Flickr API key
   --page-size NUM   # Number of photos per page (default: 10)
   --web PATH        # Path to 'web' directory (default: 'web')

   --port PORT       # Port to serve on (default: 8080)
   --socket PATH     # Socket to serve at, instead
`

	var (
		port     = flag.String("port", "8080", "")
		socket   = flag.String("socket", "", "")
		userID   = flag.String("user-id", "", "")
		apiKey   = flag.String("api-key", "", "")
		pageSize = flag.Int("page-size", 10, "")
		webPath  = flag.String("web", "web", "")
	)

	flag.Usage = func() { fmt.Fprint(os.Stderr, usage) }
	flag.Parse()

	client := flickr.New(*apiKey)

	user, err := client.UserInfo(*userID)
	if err != nil {
		log.Println(err)
		return
	}

	templates, err := template.ParseGlob(*webPath + "/template/*.gotmpl")
	if err != nil {
		log.Println(err)
		return
	}

	userInfo := handler.UserInfo{
		ID:         *userID,
		PhotosURL:  user.PhotosURL,
		ProfileURL: user.ProfileURL,
		UserName:   user.Username,
		RealName:   user.Realname,
	}

	index := handler.Index(client, userInfo, *pageSize, templates)
	route.HandleFunc("/:page", index.Show())

	sets := handler.Sets(client, userInfo, *pageSize, templates)
	route.HandleFunc("/sets", sets.List())
	route.HandleFunc("/sets/:param", sets.Show())
	route.HandleFunc("/sets/:param/:page", sets.Show())

	tags := handler.Tags(client, userInfo, *pageSize, templates)
	route.HandleFunc("/tags", tags.List())
	route.HandleFunc("/tags/:param", tags.Show())
	route.HandleFunc("/tags/:param/:page", tags.Show())

	route.Handle("/public/*path", http.StripPrefix("/public", http.FileServer(http.Dir(*webPath+"/static"))))

	serve.Serve(*port, *socket, route.Default)
}
