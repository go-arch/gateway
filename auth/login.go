package auth

import (
	"net/http"
	"net/rpc"
	"log"
	"fmt"
	"encoding/json"
	"Gateway/utils"
)

type Args struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Password string `json:"password"`
	Role string `json:"role"`
}

type AuthResponse struct {
	StatusCode          int `json:"statusCode"`
	Data        interface{} `json:"data"`
}

func LoginHandler(r *http.Request, w http.ResponseWriter) {

	client, err := rpc.DialHTTP("tcp", "localhost" + ":3001")
	if err != nil {
		log.Fatal("dialing:", err)
	}// Synchronous call
	var args Args
	decoder := json.NewDecoder(r.Body);
	err = decoder.Decode(&args);

	Fatal(err);
	var reply User
	err = client.Call("Usr.GetUser", args.Email, &reply)


	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Println(reply.Password,reply.Email)
	if (reply.Email != "" && reply.Password == args.Password){
		token,err := GenerateJwtToken(reply)
		Fatal(err);
		utils.HandleJsonResponse(w,AuthResponse{200,token});

	} else {
		w.WriteHeader(http.StatusUnauthorized) //TODO code repeated
		fmt.Fprint(w, "Email or Password is Wrong")
	}
}


//func main() {
//	client, err := rpc.DialHTTP("tcp", "localhost" + ":1234")
//	if err != nil {
//		log.Fatal("dialing:", err)
//	}// Synchronous call
//	args := "kamo.rahul@gmail.com"
//	var reply User
//	err = client.Call("Usr.GetUser", args, &reply)
//	if err != nil {
//		log.Fatal("arith error:", err)
//	}
//	fmt.Println(reply)
//}

