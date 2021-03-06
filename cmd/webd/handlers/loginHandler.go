package handlers

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	handlermodels "github.com/SmitSheth/Mini-twitter/cmd/webd/handlers/models"
	authpb "github.com/SmitSheth/Mini-twitter/internal/auth/authentication"
	userpb "github.com/SmitSheth/Mini-twitter/internal/user/userpb"
	"google.golang.org/grpc/status"
)

// Login is the handler for /login
func Login(w http.ResponseWriter, r *http.Request) {
	var user handlermodels.LoginRequest
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		APIResponse(w, r, http.StatusInternalServerError, "Error in reading request", make(map[string]string)) // send data to client side
		return
	}
	err = json.Unmarshal([]byte(body), &user)
	if err != nil {
		APIResponse(w, r, http.StatusUnauthorized, "Error in unmarshalling request", make(map[string]string)) // send data to client side
		return
	}
	//code to check if user exists
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := AuthClient.CheckAuthentication(ctx, &authpb.UserCredential{Username: user.Email, Password: user.Password})
	if err != nil {
		APIResponse(w, r, http.StatusInternalServerError, "Database not responding", make(map[string]string))
		return
	}

	if res.Authenticated {
		user, err := UserServiceClient.GetUserIdByUsername(r.Context(), &userpb.UserName{Email: user.Email})
		if err != nil {
			e, _ := status.FromError(err)
			if e.Message() == "wrong credentials" {
				APIResponse(w, r, http.StatusConflict, "Invalid credentials", make(map[string]string))
			} else {
				APIResponse(w, r, http.StatusInternalServerError, "Database not responding", make(map[string]string))
				return
			}
		}

		authToken, err := AuthClient.GetAuthToken(ctx, &authpb.UserId{UserId: user.UserId})
		if err != nil {
			APIResponse(w, r, http.StatusInternalServerError, "Database not responding", make(map[string]string))
		}

		cookie := http.Cookie{Name: "sessionId", Value: url.QueryEscape(authToken.Token), Path: "/", HttpOnly: true}
		http.SetCookie(w, &cookie)

		APIResponse(w, r, http.StatusOK, "Login successful", make(map[string]string)) // send data to client side
	} else {
		APIResponse(w, r, http.StatusUnauthorized, "Login unsuccessful", make(map[string]string)) // send data to client side
	}

}
