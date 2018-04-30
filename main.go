package main

import (
	"flag"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"Gateway/auth"
	"os"
)

type Prox struct {
	target *url.URL
	proxy  *httputil.ReverseProxy
}

var (
	UserUrl, _ = url.Parse("http://user:3000")
	StudentUrl, _ = url.Parse("http://localhost:3001")
)



var TargetServers = map[string]*httputil.ReverseProxy{"random" : httputil.NewSingleHostReverseProxy(UserUrl),
	"user" : httputil.NewSingleHostReverseProxy(StudentUrl)}

func HandleCors(r *http.Request, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Cache-Control, Postman-Token")
	w.Write([]byte(""))
}
func AddCorsHeader(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

func handle(w http.ResponseWriter, r *http.Request) {
	if (r.Method == http.MethodOptions) {
		HandleCors(r, w)
		return
	} else {
		AddCorsHeader(w)
	}
	targetServer := MsFinder(r.URL.Path)
	ctx := auth.HandleAuth(r, w)
	if (ctx != nil) {
		q := r.URL.Query()
		role := ctx["role"].(string)
		q.Add("role", role)
		r.URL.RawQuery = q.Encode()
		targetServer.ServeHTTP(w, r)
	}

}



//func (p *Prox) parseWhiteList(r *http.Request) bool {
//	for _, regexp := range p.routePatterns {
//		fmt.Println(r.URL.Path)
//		if regexp.MatchString(r.URL.Path) {
//			return true
//		}
//	}
//	fmt.Println("Not accepted routes %x", r.URL.Path)
//	return false
//}


//user/get

func MsFinder(pattern string) *httputil.ReverseProxy {
	res := strings.Split(pattern, "/")
	return TargetServers[res[1]]
}

func main() {
	var env = os.Getenv("GO_ENV");

	if env == "local" {
		UserUrl, _ = url.Parse("http://localhost:3000")
		StudentUrl, _ = url.Parse("http://localhost:3001")
	}

	const (
		defaultPort = ":3006"
		defaultPortUsage = "default server port, ':80', ':8080'..."
	)
	// flags
	port := flag.String("port", defaultPort, defaultPortUsage)
	auth.InitiateJwt()
	// server
	http.HandleFunc("/", handle)
	http.ListenAndServe(*port, nil)

}

