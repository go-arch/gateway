package auth

import (
	"net/http"
)

var (
	PublicApi = map[string]interface{}{"auth" : "/auth"}
	WhiteListApis = []string{"/login" , "/user/set", "anotherPublicApi"}
)

func HandleAuth(r *http.Request,w http.ResponseWriter)  map[string]interface {}{
	if(r.URL.Path == "/login"){
	 	LoginHandler(r,w)
		return nil
	} else if(WhitelistFinder(r.URL.Path) == true ) {
		return nil
	} else {
		return ValidateMiddleware(w,r)
	}
	return nil
}




func WhitelistFinder(url string) bool{
	for _, element := range WhiteListApis {
		if(element == url){
			return true;
			break;
		}
	}
	return false
}



