package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"forum/database"
	"forum/models"
)

type Response struct {
	Exist        bool   `json:"exist"`
	HomePage     string `json:"homePage"`
	SignUpSignIn string `json:"signUpsignIn"`
}

var Home = `<html lang="en">

<head>
    <meta charset="UTF-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
    <meta name="viewport" content="width= , initial-scale=1.0" />
    <title>Real Time Forum</title>
    <link rel="preconnect" href="https://fonts.googleapis.com" />
    <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin />
    <link href="https://fonts.googleapis.com/css2?family=Poppins:wght@300;400;500;600&display=swap" rel="stylesheet" />
    <link rel="stylesheet" href="https://unicons.iconscout.com/release/v2.1.6/css/unicons.css" />
    <link rel="stylesheet" href="./CSS/styles.css" />
</head>

<body>
    <nav>
        <div class="container">
            <h2 class="logo">Real Time Forum</h2>
        </div>
    </nav>

    <main>
        <div class="container">
            <div class="left">
                <a class="profile">
                    <div class="profile-pic">
                        <img src="./images/profile-8.jpg">
                    </div>
                    <div class="handle">
                        <h4>Xander</h4>
                    </div>
                </a>
                <div class="sidebar">
                    <a class="menu-item active">
                        <span><i class="uil uil-home"></i></span>
                        <h3>Home</h3>
                    </a>
                    <a class="menu-item" id="notifications">
                        <span><i class="uil uil-bell"><small class="notification-count">9+</small></i></span>
                        <h3>Notifications</h3>
                        <div class="notifications-popup">
                            <div>
                                <div class="profile-pic">
                                    <img src="./images/profile-10.jpg">
                                </div>
                                <div class="notification-body">
                                    <b>Abigail Willey</b> accepted your friend request
                                    <small class="text-muted">2 DAYS AGO</small>
                                </div>
                            </div>
                            <div>
                                <div class="profile-pic">
                                    <img src="./images/profile-11.jpg">
                                </div>
                                <div class="notification-body">
                                    <b>Varun Nair</b> commented on your post
                                    <small class="text-muted">1 HOUR AGO</small>
                                </div>
                            </div>
                            <div>
                                <div class="profile-pic">
                                    <img src="./images/profile-12.jpg">
                                </div>
                                <div class="notification-body">
                                    <b>Marry Opmong</b> and 210 other liked your post
                                    <small class="text-muted">4 MINUTES AGO</small>
                                </div>
                            </div>
                            <div>
                                <div class="profile-pic">
                                    <img src="./images/profile-13.jpg">
                                </div>
                                <div class="notification-body">
                                    <b>Wilson Fisk</b> started following you
                                    <small class="text-muted"> 11 HOURS AGO</small>
                                </div>
                            </div>
                        </div>
                    </a>
                    <a class="menu-item" id="messages-notifications">
                        <i class="uil uil-envelope"><small class="notification-count">6</small></i></span>
                        <h3>Messages</h3>
                    </a>
                </div>
            </div>



            <div class="middle">
                <form class="create-post">

                    <input type="submit" value="Create Post" class="btn btn-primary">
                </form>

                <div class="feeds">
                    <div class="feed">
                        <div class="head">

                        </div>
                        <div class="user">
                            <div class="profile-pic">
                                <img src="images/profile-14.jpg" alt="">
                            </div>
                            <div class="info">
                                <h3>Lana Rose</h3>
                                <small>Dubai, 15 MINUTES AGO</small>
                            </div>
                            <SPAN class="edit"><i class="uil uil-ellipsis-h"></i></SPAN>
                        </div>

                        <div class="photo">
                            <img src="images/feed-1.jpg" alt="">
                        </div>

                        <div class="action-button">
                            <div class="interaction-button">
                                <span><i class="uil uil-thumbs-up"></i></span>
                                <span><i class="uil uil-comment"></i></span>
                                <span><i class="uil uil-share"></i></span>
                            </div>
                            <div class="bookmark">
                                <span><i class="uil uil-bookmark"></i></span>
                            </div>
                        </div>

                        <div class="liked-by">
                            <span><img src="images/profile-15.jpg"></span>
                            <span><img src="images/profile-16.jpg"></span>
                            <span><img src="images/profile-17.jpg"></span>
                            ,<p>Liked by <b>Enrest Achiever</b>snd <b>220 others</b></p>
                        </div>

                        <div class="caption">
                            <p><b>Lana Rose</b>Lorem ipsum dolor storiesquiquam eius.
                                <span class="hash-tag">#lifestyle</span>
                            </p>
                        </div>
                        <div class="comments text-muted">View all 130 comments</div>
                    </div>

                    <div class="feed">
                        <div class="head">

                        </div>
                        <div class="user">
                            <div class="profile-pic">
                                <img src="images/profile-15.jpg" alt="">
                            </div>
                            <div class="info">
                                <h3>Chris Brown</h3>
                                <small>New York, 1 HOUR AGO</small>
                            </div>
                            <SPAN class="edit"><i class="uil uil-ellipsis-h"></i></SPAN>
                        </div>

                        <div class="photo">
                            <img src="images/feed-2.jpg" alt="">
                        </div>

                        <div class="action-button">
                            <div class="interaction-button">
                                <span><i class="uil uil-thumbs-up"></i></span>
                                <span><i class="uil uil-comment"></i></span>
                                <span><i class="uil uil-share"></i></span>
                            </div>
                            <div class="bookmark">
                                <span><i class="uil uil-bookmark"></i></span>
                            </div>
                        </div>

                        <div class="liked-by">
                            <span><img src="images/profile-2.jpg"></span>
                            <span><img src="images/profile-4.jpg"></span>
                            <span><img src="images/profile-6.jpg"></span>
                            ,<p>Liked by <b>Enrest Achiever</b>snd <b>188 others</b></p>
                        </div>

                        <div class="caption">
                            <p><b>Chirs Brown</b>Lorem ipsum dolor storiesquiquam eius.
                                <span class="hash-tag">#lifestyle</span>
                            </p>
                        </div>
                        <div class="comments text-muted">View all 40 comments</div>
                    </div>

                    <div class="feed">
                        <div class="head">

                        </div>
                        <div class="user">
                            <div class="profile-pic">
                                <img src="images/profile-16.jpg" alt="">
                            </div>
                            <div class="info">
                                <h3>John Samron</h3>
                                <small>Amsterdam, 7 HOURS AGO</small>
                            </div>
                            <SPAN class="edit"><i class="uil uil-ellipsis-h"></i></SPAN>
                        </div>

                        <div class="photo">
                            <img src="images/feed-3.jpg" alt="">
                        </div>

                        <div class="action-button">
                            <div class="interaction-button">
                                <span><i class="uil uil-thumbs-up"></i></span>
                                <span><i class="uil uil-comment"></i></span>
                                <span><i class="uil uil-share"></i></span>
                            </div>
                            <div class="bookmark">
                                <span><i class="uil uil-bookmark"></i></span>
                            </div>
                        </div>

                        <div class="liked-by">
                            <span><img src="images/profile-3.jpg"></span>
                            <span><img src="images/profile-5.jpg"></span>
                            <span><img src="images/profile-7.jpg"></span>
                            ,<p>Liked by <b>Enrest Achiever</b>snd <b>130 others</b></p>
                        </div>

                        <div class="caption">
                            <p><b>John Samron</b>Lorem ipsum dolor storiesquiquam eius.
                                <span class="hash-tag">#lifestyle</span>
                            </p>
                        </div>
                        <div class="comments text-muted">View all 15 comments</div>
                    </div>

                    <div class="feed">
                        <div class="head">

                        </div>
                        <div class="user">
                            <div class="profile-pic">
                                <img src="images/profile-17.jpg" alt="">
                            </div>
                            <div class="info">
                                <h3>Kareena Joshua</h3>
                                <small>USA, 3 HOURS AGO</small>
                            </div>
                            <SPAN class="edit"><i class="uil uil-ellipsis-h"></i></SPAN>
                        </div>

                        <div class="photo">
                            <img src="images/feed-4.jpg" alt="">
                        </div>

                        <div class="action-button">
                            <div class="interaction-button">
                                <span><i class="uil uil-thumbs-up"></i></span>
                                <span><i class="uil uil-comment"></i></span>
                                <span><i class="uil uil-share"></i></span>
                            </div>
                            <div class="bookmark">
                                <span><i class="uil uil-bookmark"></i></span>
                            </div>
                        </div>

                        <div class="liked-by">
                            <span><img src="images/profile-8.jpg"></span>
                            <span><img src="images/profile-10.jpg"></span>
                            <span><img src="images/profile-12.jpg"></span>
                            ,<p>Liked by <b>Enrest Achiever</b>snd <b>280 others</b></p>
                        </div>

                        <div class="caption">
                            <p><b>Kareena Joshua</b>Lorem ipsum dolor storiesquiquam eius.
                                <span class="hash-tag">#lifestyle</span>
                            </p>
                        </div>
                        <div class="comments text-muted">View all 110 comments</div>
                    </div>

                    <div class="feed">
                        <div class="head">

                        </div>
                        <div class="user">
                            <div class="profile-pic">
                                <img src="images/profile-18.jpg" alt="">
                            </div>
                            <div class="info">
                                <h3>Dan Smith</h3>
                                <small>Paris, 1 DAY AGO</small>
                            </div>
                            <SPAN class="edit"><i class="uil uil-ellipsis-h"></i></SPAN>
                        </div>

                        <div class="photo">
                            <img src="images/feed-5.jpg" alt="">
                        </div>

                        <div class="action-button">
                            <div class="interaction-button">
                                <span><i class="uil uil-thumbs-up"></i></span>
                                <span><i class="uil uil-comment"></i></span>
                                <span><i class="uil uil-share"></i></span>
                            </div>
                            <div class="bookmark">
                                <span><i class="uil uil-bookmark"></i></span>
                            </div>
                        </div>

                        <div class="liked-by">
                            <span><img src="images/profile-9.jpg"></span>
                            <span><img src="images/profile-11.jpg"></span>
                            <span><img src="images/profile-13.jpg"></span>
                            ,<p>Liked by <b>Enrest Achiever</b>snd <b>420 others</b></p>
                        </div>

                        <div class="caption">
                            <p><b>Dan Smith</b>Lorem ipsum dolor storiesquiquam eius.
                                <span class="hash-tag">#lifestyle</span>
                            </p>
                        </div>
                        <div class="comments text-muted">View all 120 comments</div>
                    </div>

                    <div class="feed">
                        <div class="head">

                        </div>
                        <div class="user">
                            <div class="profile-pic">
                                <img src="images/profile-19.jpg" alt="">
                            </div>
                            <div class="info">
                                <h3>Karim Benzema</h3>
                                <small>Mumbai, 30 MINUTES AGO</small>
                            </div>
                            <SPAN class="edit"><i class="uil uil-ellipsis-h"></i></SPAN>
                        </div>

                        <div class="photo">
                            <img src="images/feed-6.jpg" alt="">
                        </div>

                        <div class="action-button">
                            <div class="interaction-button">
                                <span><i class="uil uil-thumbs-up"></i></span>
                                <span><i class="uil uil-comment"></i></span>
                                <span><i class="uil uil-share"></i></span>
                            </div>
                            <div class="bookmark">
                                <span><i class="uil uil-bookmark"></i></span>
                            </div>
                        </div>

                        <div class="liked-by">
                            <span><img src="images/profile-15.jpg"></span>
                            <span><img src="images/profile-14.jpg"></span>
                            <span><img src="images/profile-17.jpg"></span>
                            ,<p>Liked by <b>Enrest Achiever</b>snd <b>150 others</b></p>
                        </div>

                        <div class="caption">
                            <p><b>Karim Benzema</b>Lorem ipsum dolor storiesquiquam eius.
                                <span class="hash-tag">#lifestyle</span>
                            </p>
                        </div>
                        <div class="comments text-muted">View all 30 comments</div>
                    </div>
                    <div class="feed">
                        <div class="head">

                        </div>
                        <div class="user">
                            <div class="profile-pic">
                                <img src="images/profile-20.jpg" alt="">
                            </div>
                            <div class="info">
                                <h3>Srishti Tirkey</h3>
                                <small>Bangalore, 11 HOURS AGO</small>
                            </div>
                            <SPAN class="edit"><i class="uil uil-ellipsis-h"></i></SPAN>
                        </div>

                        <div class="photo">
                            <img src="images/feed-7.jpg" alt="">
                        </div>

                        <div class="action-button">
                            <div class="interaction-button">
                                <span><i class="uil uil-thumbs-up"></i></span>
                                <span><i class="uil uil-comment"></i></span>
                                <span><i class="uil uil-share"></i></span>
                            </div>
                            <div class="bookmark">
                                <span><i class="uil uil-bookmark"></i></span>
                            </div>
                        </div>

                        <div class="liked-by">
                            <span><img src="images/profile-15.jpg"></span>
                            <span><img src="images/profile-13.jpg"></span>
                            <span><img src="images/profile-10.jpg"></span>
                            ,<p>Liked by <b>Enrest Achiever</b>snd <b>530 others</b></p>
                        </div>

                        <div class="caption">
                            <p><b>Srishti Tirkey</b>Lorem ipsum dolor storiesquiquam eius.
                                <span class="hash-tag">#lifestyle</span>
                            </p>
                        </div>
                        <div class="comments text-muted">View all 190 comments</div>
                    </div>
                </div>
            </div>

            <div class="right">
                <div class="messages">
                    <div class="heading">
                        <h4>Messages</h4>
                    </div>

                    <div class="message">
                        <div class="profile-pic">
                            <img src="images/profile-17.jpg">
                            <!-- <div class="active"></div> -->
                        </div>
                        <div class="message-body">
                            <h5>Kareena Joshua</h5>
                            <p class="text-muted">Just woke up bruh..</p>
                        </div>
                    </div>
                    <div class="message">
                        <div class="profile-pic">
                            <img src="images/profile-18.jpg">
                            <!-- <div class="active"></div> -->
                        </div>
                        <div class="message-body">
                            <h5>Dan Smith</h5>
                            <p class="text-bold">2 New Messages</p>
                        </div>
                    </div>
                    <div class="message">
                        <div class="profile-pic">
                            <img src="images/profile-15.jpg">
                        </div>
                        <div class="message-body">
                            <h5>Chris Brown</h5>
                            <p class="text-muted">Lol u right</p>
                        </div>
                    </div>
                    <div class="message">
                        <div class="profile-pic">
                            <img src="images/profile-14.jpg">
                        </div>
                        <div class="message-body">
                            <h5>Lana Rose</h5>
                            <p class="text-bold">Birthday tomorrow!!</p>
                        </div>
                    </div>
                    <div class="message">
                        <div class="profile-pic">
                            <img src="images/profile-11.jpg">
                        </div>
                        <div class="message-body">
                            <h5>Varun Nair</h5>
                            <p class="text-muted">Ssup?</p>
                        </div>
                    </div>
                    <div class="message">
                        <div class="profile-pic">
                            <img src="images/profile-1.jpg">
                            <!-- <div class="active"></div> -->
                        </div>
                        <div class="message-body">
                            <h5>Jahnvi Doifode</h5>
                            <p class="text-bold">3 New Messages</p>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </main>

    <script src="./JS/index.js"></script>
</body>

</html>`
var SignUpIn = `<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.4.2/css/all.min.css">
	<link rel="stylesheet" href="./CSS/styles.css">
	<title>Real Time Forum</title>
</head>
<body>
	<!-- <div class="loader"></div> -->
	<div class="container" id="container">
		<div class="form-container sign-up">
			<form id="registerForm">
				<div class="title">Create Account</div>
				<span>or use your email for registration</span>
				<input id="nickname" type="text" name="pseudo" placeholder="Nickname">
				<input id="firstname" type="text" name="firstName" placeholder="First Name">
				<input id="lastname" type="text" name="lastName" placeholder="Last Name">
				<input id="age" type="number" name="age" placeholder="Age">
				<div class="gender">
					<p>Gender</p>
					<select id="gender" name="gender">
						<option value="Male">Male</option>
						<option value="Female">Female</option>
					</select>
				</div>
				<input id="email" type="email" name="email" placeholder="Email">
				<input id="password" type="password" name="password" placeholder="Password">
				<button class="signup" disabled>Sign Up</button>
			</form>
			<div id="successMessage" class="successMessage">
				Registered ✅!
			</div>
		</div>
		<div class="form-container sign-in">
			<form id="loginForm">
				<h1>Sign In</h1>
				<span>or use your email password</span>
				<input id="loginMail" type="email" placeholder="Email">
				<input id="loginPassword" type="password" placeholder="Password">
				<button>Sign In</button>
			</form>
		</div>
		<div class="toggle-container">
			<div class="toggle">
				<div class="toggle-panel toggle-left">
					<h1>Welcome Back!</h1>
					<p>Enter your personal details to use all of site features</p>
					<button class="hidden" id="login">Sign In</button>
				</div>
				<div class="toggle-panel toggle-right">
					<h1>Hello, Friend!</h1>
					<p>Register with your personal details to use all of site features</p>
					<button class="hidden" id="register">Sign Up</button>
				</div>
			</div>
		</div>
	</div>
	<script src="./JS/app.js"></script>
</body>
</html>`

var ErrorPage = `<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <link rel="stylesheet" href=./CSS/styleerror.css>
  <title>{{.Code}}</title>
</head>
<body>
  <div class="contain">
   <p class="codeError">
    {{.Code}}
   </p>
   <p class="Description">
    {{ .Message }}
   </p>
    <button><a href="/">Go Home</a></button>
  </div>
</body>
</html>`

// Fonctions CRUD pour Users
func GetUser(w http.ResponseWriter, r *http.Request, id string) {

	user, err := database.GetUserByID(id)
	if err != nil {
		fmt.Println("Error getting user: ", err)
		http.Error(w, "Utilisateur non trouvé", http.StatusNotFound)
		return
	}

	jsonResponse2(w, http.StatusOK, user)
	fmt.Println("status: ", http.StatusOK)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := database.GetAllUsers()
	if err != nil {
		fmt.Println("Error getting users: ", err)
		http.Error(w, "Erreur lors de la récupération des utilisateurs", http.StatusInternalServerError)
		return
	}

	jsonResponse2(w, http.StatusOK, users)
	fmt.Println("status: ", http.StatusOK)
	fmt.Println("*** Liste des utilisateurs ENVOYER AVEC SUCCÈS ***")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.Users

	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	_, err := database.CreateUser(&newUser)
	if err != nil {
		fmt.Println("Error creating user: ", err)
		http.Error(w, "Erreur lors de la création de l'utilisateur", http.StatusInternalServerError)
		return
	}

	jsonResponse2(w, http.StatusCreated, newUser)
	fmt.Println("status: ", http.StatusCreated)
}

func UpdateUser(w http.ResponseWriter, r *http.Request, id string) {

	user, err := database.GetUserByID(id)
	if err != nil {
		http.Error(w, "Utilisateur non trouvé", http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = database.UpdateUser(user)
	if err != nil {
		fmt.Println("Error updating user: ", err)
		http.Error(w, "Erreur lors de la mise à jour de l'utilisateur", http.StatusInternalServerError)
		return
	}

	jsonResponse2(w, http.StatusOK, user)
	fmt.Println("status: ", http.StatusOK)
}

func DeleteUser(w http.ResponseWriter, r *http.Request, id string) {

	err := database.DeleteUser(id)
	if err != nil {
		http.Error(w, "Utilisateur non trouvé", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	fmt.Println("status: ", http.StatusNoContent)
}

// Fonction utilitaire pour extraire l'ID à partir du chemin de l'URL
func extractIDFromPath(path string) int {
	parts := strings.Split(path, "/")
	id, _ := strconv.Atoi(parts[len(parts)-1])
	return id
}

func CheckActiveSession(r *http.Request) (*models.Users, bool) {
	var user *models.Users

	if cookie, err := r.Cookie("sessionId"); err != nil {
		return nil, false
	} else {
		user, err = database.GetUserByID(cookie.Value)

		if err != nil {
			return nil, false
		}
	}
	return user, true
}

func handleActiveSession(w http.ResponseWriter, r *http.Request) {
	var res Response
	var user *models.Users
	var exist bool

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	user, exist = CheckActiveSession(r)

	if user == nil {
		exist = false
	}

	res = Response{
		Exist:        exist,
		HomePage:     Home,
		SignUpSignIn: SignUpIn,
	}

	jsonRes, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonRes)
}

// if cookie, err := r.Cookie("sessionId"); err != nil {
// 	exist = false
// 	res = Response{
// 		Exist:        exist,
// 		HomePage:     ``,
// 		SignUpSignIn: SignUpIn,
// 	}
// 	jsonRes, err := json.Marshal(res)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(jsonRes)
// 	return

// } else {
// 	val := cookie.Value
// 	user, err = database.GetUserByID(val)
// 	if user != nil && err == nil {
// 		exist = true
// 		res.Exist = exist
// 		res.HomePage = Home
// 		res.SignUpSignIn = SignUpIn
// 	}

// }
