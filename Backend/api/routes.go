package api

import (
	"net/http"
)

func Router() {

	// Endpoints
	http.HandleFunc("/", handleFirstPage)
	http.HandleFunc("/api/activeSession", handleActiveSession)
	http.HandleFunc("/api/login", handleLogin)
	http.HandleFunc("/api/register", handleRegister)
	http.HandleFunc("/api/posts", handleGetPosts)
	http.HandleFunc("/api/createPost", handleCreatingPost)
	http.HandleFunc("/api/comment", handleComment)
	http.HandleFunc("/api/getComments", handleGetComments)
	http.HandleFunc("/api/getUsers", handleGetUsers)
	http.HandleFunc("/api/getDiscussions", handleGettingDiscus)
	http.HandleFunc("/api/getNotifs", handleNotification)
	http.HandleFunc("/api/handleError", handleError)
	http.HandleFunc("/api/checkUser", handleGetUser)
	http.HandleFunc("/api/logout", handleLogout)
	http.HandleFunc("/ws", handleConnections)

}
