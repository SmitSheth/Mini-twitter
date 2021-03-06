package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	handlermodels "github.com/SmitSheth/Mini-twitter/cmd/webd/handlers/models"
	authpb "github.com/SmitSheth/Mini-twitter/internal/auth/authentication"
	userpb "github.com/SmitSheth/Mini-twitter/internal/user/userpb"
)

// UserHandler is the handler for /users. It is for getting a list of all users, or a specific user.
func UserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		users, err := UserServiceClient.GetAllUsers(r.Context(), &userpb.Void{})
		tempArr := make([]*userpb.AccountInformation, len(users.UserList))
		for i, userObj := range users.UserList {
			tempArr[i] = userObj.AccountInformation
		}
		respMessage := handlermodels.GetUsersResponse{tempArr}
		body, err := json.Marshal(respMessage)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Header().Set("content-type", "application/json")
		w.Write(body)
	case "POST":
		reqMessage := handlermodels.GetUserRequest{}
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			APIResponse(w, r, http.StatusBadRequest, "Could not read request", make(map[string]string))
			return
		}
		err = json.Unmarshal(b, &reqMessage)
		if err != nil {
			APIResponse(w, r, http.StatusBadRequest, "Could not read request", make(map[string]string))
			return
		}
		//u := application.GetUser(reqMessage.UserID)
		userId := r.Context().Value("user").(*authpb.UserId)
		user, err := UserServiceClient.GetUser(r.Context(), &userpb.UserId{UserId: userId.UserId})
		body, err := json.Marshal(user.AccountInformation)
		w.Header().Set("content-type", "application/json")
		w.Write(body)
	default:
		http.Error(w, "Only GET allowed", http.StatusMethodNotAllowed)
	}
}

// UserFollowingHandler is the handler for /user/following. It is used for getting a list of users
// the current user is following
func UserFollowingHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		user := r.Context().Value("user").(*authpb.UserId)
		users, err := UserServiceClient.GetFollowing(r.Context(), &userpb.UserId{UserId: user.UserId})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//users, err := application.GetFollowing(user.UserId)
		tempArr := make([]*userpb.AccountInformation, len(users.UserList))
		for i, userObj := range users.UserList {
			tempArr[i] = userObj.AccountInformation
		}
		respMessage := handlermodels.GetUsersResponse{tempArr}
		body, err := json.Marshal(respMessage)
		if err != nil {
			APIResponse(w, r, http.StatusInternalServerError, "Cannot get user following list", make(map[string]string))
		}
		w.Header().Set("content-type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
		w.Write(body)
	default:
		http.Error(w, "Only GET allowed", http.StatusMethodNotAllowed)
	}
}

// UserNotFollowingHandler is the handler for /user/notfollowing. It is used for getting a list of users
// the current user is not following
func UserNotFollowingHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		user := r.Context().Value("user").(*authpb.UserId)
		users, err := UserServiceClient.GetNotFollowing(r.Context(), &userpb.UserId{UserId: user.UserId})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tempArr := make([]*userpb.AccountInformation, len(users.UserList))
		for i, userObj := range users.UserList {
			tempArr[i] = userObj.AccountInformation
		}
		respMessage := handlermodels.GetUsersResponse{tempArr}
		body, err := json.Marshal(respMessage)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Header().Set("content-type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
		w.Write(body)
	default:
		http.Error(w, "Only GET allowed", http.StatusMethodNotAllowed)
	}
}
