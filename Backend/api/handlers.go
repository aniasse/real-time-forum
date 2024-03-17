package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"time"

	"forum/database"
	"forum/models"

	"golang.org/x/crypto/bcrypt"
)

type Response struct {
	Exist        bool   `json:"exist"`
	HomeHead     string `json:"homeHead"`
	HomePage     string `json:"homePage"`
	SignHead     string `json:"signHead"`
	SignUpSignIn string `json:"signUpIn"`
	NickName     string `json:"nickname"`
}

type CookieData struct {
	CookieValue string `json:"cookieValue"`
}

type PostWithUser struct {
	ID       int    `json:"id"`
	UserID   string `json:"userId"`
	Category string `json:"category"`
	Content  string `json:"content"`
	Date     string `json:"date"`
	Nickname string `json:"nickname"`
}

var Homehead = `<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<link rel="stylesheet" href="https://unicons.iconscout.com/release/v2.1.6/css/unicons.css" />
<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.5.1/css/all.min.css" />
<link href='https://unpkg.com/boxicons@2.1.1/css/boxicons.min.css' rel='stylesheet'>
<script src='https://kit.fontawesome.com/a076d05399.js' crossorigin='anonymous'></script>
<link rel="stylesheet" href="/static/CSS/home.css">
<title>Real Time Forum</title>
</head>`

var Home = `<body>
    <header>
        <h2 class="logo">Real Time Forum</h2>
        <nav>
            <div class="nav-content">
                <div class="toggle-btn">
                    <i class='bx bx-plus'></i>
                </div>
                <span class="first">
                    <a href="#"><i class='bx bxs-home'></i></a>
                </span>
                <span class="second">
                    <a href="#"><i class='bx bxs-message-alt'></i></a>
                </span>
                <span class="thirth">
                    <a href="#"><i class='bx bxs-bell'></i></a>
                </span>
            </div>
        </nav>
        <img src="./static/images/notification.gif" alt="Notification" class="notifgif" style="display:none;">
        <div class="navigation">
            <a class="button" href="#">
                <img src="./static/images/user.png" alt="logout">
                <div class="logout">LOGOUT</div>
            </a>
        </div>
    </header>
    <div id="Message" class="Message"></div>
    <div class="content">
        <div class="notifs">
            <div class="title">
                <span>
                    <h5>Notifications</h5>
                </span>
            </div>
            <div class="notif">
                <div class="profil-pic">
                    <img src="./static/images/user.png" alt="">
                </div>
                <div class="notif-body">
                    <b>Abigail Willey</b> Send you a message
                    <small class="text-muted">2 DAYS AGO</small>
                </div>
            </div>
            <div class="notif">
                <div class="profil-pic">
                    <img src="./static/images/user.png" alt="">
                </div>
                <div class="notif-body">
                    <b>Abigail Willey</b> Send you a message
                    <small class="text-muted">2 DAYS AGO</small>
                </div>
            </div>
            <div class="notif">
                <div class="profil-pic">
                    <img src="./static/images/user.png" alt="">
                </div>
                <div class="notif-body">
                    <b>Abigail</b> Send you a message
                    <small class="text-muted">2 DAYS AGO</small>
                </div>
            </div>
            <div class="notif">
                <div class="profil-pic">
                    <img src="./static/images/user.png" alt="">
                </div>
                <div class="notif-body">
                    <b>Abigail Willey</b> Send you a message
                    <small class="text-muted">2 DAYS AGO</small>
                </div>
            </div>
            <div class="notif">
                <div class="profil-pic">
                    <img src="./static/images/user.png" alt="">
                </div>
                <div class="notif-body">
                    <b>Abigail Willey</b> Send you a message
                    <small class="text-muted">2 DAYS AGO</small>
                </div>
            </div>
            <div class="notif">
                <div class="profil-pic">
                    <img src="./static/images/user.png" alt="">
                </div>
                <div class="notif-body">
                    <b>Abigail Willey</b> Send you a message
                    <small class="text-muted">2 DAYS AGO</small>
                </div>
            </div>
            <div class="notif">
                <div class="profil-pic">
                    <img src="./static/images/user.png" alt="">
                </div>
                <div class="notif-body">
                    <b>Abigail Willey</b> Send you a message
                    <small class="text-muted">2 DAYS AGO</small>
                </div>
            </div>
        </div>
        <div class="posts">
            <div class="create">
                <div class="user">
                    <div class="profil-pic">
                        <img src="./static/images/user.png" alt="">
                    </div>
                    <h6>Xander</h6>
                </div>
                <button class="button-create">Create</button>
            </div>
            <div class="createpost">
                <div class="close">
                    <img src="./static/images/close.png" alt="close" aria-details="close">
                </div>
                <div class="categories">
                    <div class="box">
                        <select id="categorie">
                            <option>News</option>
                            <option>Tech</option>
                            <option>Computing</option>
                            <option>Sport</option>
                            <option>Gaming</option>
                        </select>
                    </div>
                </div>
                <div class="topost">
                    <textarea aria-label="#" name="post" id="post" placeholder="What's happening"></textarea>
                </div>
                <button class="sub" id="postButton">Post</button>
            </div>
            <div class="post">
                <div class="poster">
                    <div class="profil-pic">
                        <img src="./static/images/user.png" alt="">
                    </div>
                    <div class="info">
                        <h4>Xander</h4>
                        <small class="text-muted">1 DAYS AGO</small>
                    </div>
                </div>
                <div class="cat"><span>Holiday</span></div>
                <div class="ctn">
                    <p class="text">It's holiday everybody is happy</p>
                </div>
                <div class="comment">
                    <img src="./static/images/comment.png" alt="">
                    <div class="comments">
                        <div class="all-com">
                            <div class="usr">
                                <div class="profil-pic">
                                    <img src="./static/images/user.png" alt="">
                                </div>
                                <h6>Username</h6>
                            </div>
                            <p>Great</p>
                            <div class="usr">
                                <div class="profil-pic">
                                    <img src="./static/images/user.png" alt="">
                                </div>
                                <h6>Username</h6>
                            </div>
                            <p>Great</p>
                            <div class="usr">
                                <div class="profil-pic">
                                    <img src="./static/images/user.png" alt="">
                                </div>
                                <h6>Username</h6>
                            </div>
                            <p>Great</p>
                        </div>
                        <div class="tocom">
                            <textarea aria-label="#" name="com" id="com" placeholder="comment..."></textarea>
                            <button class="btn-com">Comment</button>
                        </div>
                </div>
                </div>
            </div>
            <div class="post">
                <div class="poster">
                    <div class="profil-pic">
                        <img src="./static/images/user.png" alt="">
                    </div>
                    <div class="info">
                        <h4>Xander</h4>
                        <small class="text-muted">1 DAYS AGO</small>
                    </div>
                </div>
                <div class="cat"><span>Holiday</span></div>
                <div class="ctn">
                    <p class="text">It's holiday everybody is happy</p>
                </div>
                <div class="comment">
                    <img src="./static/images/comment.png" alt="">
                </div>
                <div class="comments">
                    <div class="all-com">
                        <div class="usr">
                            <div class="profil-pic">
                                <img src="./static/images/user.png" alt="">
                            </div>
                            <h6>Username</h6>
                        </div>
                        <p>Great</p>
                        <div class="usr">
                            <div class="profil-pic">
                                <img src="./static/images/user.png" alt="">
                            </div>
                            <h6>Username</h6>
                        </div>
                        <p>Great</p>
                        <div class="usr">
                            <div class="profil-pic">
                                <img src="./static/images/user.png" alt="">
                            </div>
                            <h6>Username</h6>
                        </div>
                        <p>Great</p>
                    </div>
                    <div class="tocom">
                        <textarea aria-label="#" name="com" id="com" placeholder="comment..."></textarea>
                        <button class="btn-com">Comment</button>
                    </div>
                </div>
            </div>
            <div class="post">
                <div class="poster">
                    <div class="profil-pic">
                        <img src="./static/images/user.png" alt="">
                    </div>
                    <div class="info">
                        <h4>Xander</h4>
                        <small class="text-muted">1 DAYS AGO</small>
                    </div>
                </div>
                <div class="cat"><span>Holiday</span></div>
                <div class="ctn">
                    <p class="text">It's holiday everybody is happy</p>
                </div>
                <div class="comment">
                    <img src="./static/images/comment.png" alt="">
                </div>
                <div class="comments">
                    <div class="all-com">
                        <div class="usr">
                            <div class="profil-pic">
                                <img src="./static/images/user.png" alt="">
                            </div>
                            <h6>Username</h6>
                        </div>
                        <p>Great</p>
                        <div class="usr">
                            <div class="profil-pic">
                                <img src="./static/images/user.png" alt="">
                            </div>
                            <h6>Username</h6>
                        </div>
                        <p>Great</p>
                        <div class="usr">
                            <div class="profil-pic">
                                <img src="./static/images/user.png" alt="">
                            </div>
                            <h6>Username</h6>
                        </div>
                        <p>Great</p>
                    </div>
                    <div class="tocom">
                        <textarea aria-label="#" name="com" id="com" placeholder="comment..."></textarea>
                        <button class="btn-com">Comment</button>
                    </div>
                </div>
            </div>
            <div class="post">
                <div class="poster">
                    <div class="profil-pic">
                        <img src="./static/images/user.png" alt="">
                    </div>
                    <div class="info">
                        <h4>Xander</h4>
                        <small class="text-muted">1 DAYS AGO</small>
                    </div>
                </div>
                <div class="cat"><span>Holiday</span></div>
                <div class="ctn">
                    <p class="text">It's holiday everybody is happy</p>                
                </div>
                <div class="comment">
                    <img src="./static/images/comment.png" alt="">
                </div>
                <div class="comments">
                    <div class="all-com">
                        <div class="usr">
                            <div class="profil-pic">
                                <img src="./static/images/user.png" alt="">
                            </div>
                            <h6>Username</h6>
                        </div>
                        <p>Great</p>
                        <div class="usr">
                            <div class="profil-pic">
                                <img src="./static/images/user.png" alt="">
                            </div>
                            <h6>Username</h6>
                        </div>
                        <p>Great</p>
                        <div class="usr">
                            <div class="profil-pic">
                                <img src="./static/images/user.png" alt="">
                            </div>
                            <h6>Username</h6>
                        </div>
                        <p>Great</p>
                    </div>
                    <div class="tocom">
                        <textarea aria-label="#" name="com" id="com" placeholder="comment..."></textarea>
                        <button class="btn-com">Comment</button>
                    </div>
                </div>
            </div>
            <div class="post">
                <div class="poster">
                    <div class="profil-pic">
                        <img src="./static/images/user.png" alt="">
                    </div>
                    <div class="info">
                        <h4>Xander</h4>
                        <small class="text-muted">1 DAYS AGO</small>
                    </div>
                </div>
                <div class="cat"><span>Holiday</span></div>
                <div class="ctn">
                    <p class="text">It's holiday everybody is happy</p>
                </div>
                <div class="comment">
                    <img src="./static/images/comment.png" alt="">
                </div>
                <div class="comments">
                    <div class="all-com">
                        <div class="usr">
                            <div class="profil-pic">
                                <img src="./static/images/user.png" alt="">
                            </div>
                            <h6>Username</h6>
                        </div>
                        <p>Great</p>
                        <div class="usr">
                            <div class="profil-pic">
                                <img src="./static/images/user.png" alt="">
                            </div>
                            <h6>Username</h6>
                        </div>
                        <p>Great</p>
                        <div class="usr">
                            <div class="profil-pic">
                                <img src="./static/images/user.png" alt="">
                            </div>
                            <h6>Username</h6>
                        </div>
                        <p>Great</p>
                    </div>
                    <div class="tocom">
                        <textarea aria-label="#" name="com" id="com" placeholder="comment..."></textarea>
                        <button class="btn-com">Comment</button>
                    </div>
                </div>
            </div>
            <div class="post">
                <div class="poster">
                    <div class="profil-pic">
                        <img src="./static/images/user.png" alt="">
                    </div>
                    <div class="info">
                        <h4>Xander</h4>
                        <small class="text-muted">1 DAYS AGO</small>
                    </div>
                </div>
                <div class="cat"><span>Holiday</span></div>
                <div class="ctn">
                    <p class="text">It's holiday everybody is happy</p>
                </div>
                <div class="comment">
                    <img src="./static/images/comment.png" alt="">
                </div>
                <div class="comments">
                    <div class="all-com">
                        <div class="usr">
                            <div class="profil-pic">
                                <img src="./static/images/user.png" alt="">
                            </div>
                            <h6>Username</h6>
                        </div>
                        <p>Great</p>
                        <div class="usr">
                            <div class="profil-pic">
                                <img src="./static/images/user.png" alt="">
                            </div>
                            <h6>Username</h6>
                        </div>
                        <p>Great</p>
                        <div class="usr">
                            <div class="profil-pic">
                                <img src="./static/images/user.png" alt="">
                            </div>
                            <h6>Username</h6>
                        </div>
                        <p>Great</p>
                    </div>
                    <div class="tocom">
                        <textarea aria-label="#" name="com" id="com" placeholder="comment..."></textarea>
                        <button class="btn-com">Comment</button>
                    </div>
                </div>
            </div>
        </div>
        <div class="messages">
            <div class="title">
                <span>
                    <h5>Messages</h5>
                </span>
            </div>
            <div class="message">
                <div class="usr">
                    <div class="inf">
                        <div class="profil-pic"><img src="./static/images/user.png" alt=""></div>
                        <div>Shadow</div>
                    </div>
                    <div class="stat"></div>
                </div>
                <div class="discus">
                    <div class="from-usr">Hello <p class="message-date">17/02/2024</p></div>
                    <div class="from-usr">How are u ?<p class="message-date">17/02/2024</p></div>
                    <div class="from-exp">Hi <p class="message-date">17/02/2024</p></div>
                    <div class="from-exp">I'm fine and u ? <p class="message-date">17/02/2024</p></div>
                    <div class="from-usr">Yeah <p class="message-date">17/02/2024</p></div>
                    <div class="to-send">
                        <textarea name="sms" id="sms" placeholder="Type a message..."></textarea>
                        <div class="send"><img src="./static/images/send.png" alt="send"></div>
                    </div>
                </div>
            </div>
            <div class="message">
                <div class="usr">
                    <div class="inf">
                        <div class="profil-pic"><img src="./static/images/user.png" alt=""></div>
                        <div>Shadow</div>
                    </div>
                    <div class="stat"></div>
                </div>
                <div class="discus">
                    <div class="from-usr">Hello <p class="message-date">17/02/2024</p></div>
                    <div class="from-usr">How are u ?<p class="message-date">17/02/2024</p></div>
                    <div class="from-exp">Hi <p class="message-date">17/02/2024</p></div>
                    <div class="from-exp">I'm fine and u ? <p class="message-date">17/02/2024</p></div>
                    <div class="from-usr">Yeah <p class="message-date">17/02/2024</p></div>
                    <div class="to-send">
                        <textarea name="sms" id="sms" placeholder="Type a message..."></textarea>
                        <div class="send"><img src="./static/images/send.png" alt="send"></div>
                    </div>
                </div>
            </div>
            <div class="message">
                <div class="usr">
                    <div class="inf">
                        <div class="profil-pic"><img src="./static/images/user.png" alt=""></div>
                        <div>Shadow</div>
                    </div>
                    <div class="stat"></div>
                </div>
                <div class="discus">
                    <div class="from-usr">Hello <p class="message-date">17/02/2024</p></div>
                    <div class="from-usr">How are u ?<p class="message-date">17/02/2024</p></div>
                    <div class="from-exp">Hi <p class="message-date">17/02/2024</p></div>
                    <div class="from-exp">I'm fine and u ? <p class="message-date">17/02/2024</p></div>
                    <div class="from-usr">Yeah <p class="message-date">17/02/2024</p></div>
                    <div class="to-send">
                        <textarea name="sms" id="sms" placeholder="Type a message..."></textarea>
                        <div class="send"><img src="./static/images/send.png" alt="send"></div>
                    </div>
                </div>
            </div>
            <div class="message">
                <div class="usr">
                    <div class="inf">
                        <div class="profil-pic"><img src="./static/images/user.png" alt=""></div>
                        <div>Shadow</div>
                    </div>
                    <div class="stat"></div>
                </div>
                <div class="discus">
                    <div class="from-usr">Hello <p class="message-date">17/02/2024</p></div>
                    <div class="from-usr">How are u ?<p class="message-date">17/02/2024</p></div>
                    <div class="from-exp">Hi <p class="message-date">17/02/2024</p></div>
                    <div class="from-exp">I'm fine and u ? <p class="message-date">17/02/2024</p></div>
                    <div class="from-usr">Yeah <p class="message-date">17/02/2024</p></div>
                    <div class="to-send">
                        <textarea name="sms" id="sms" placeholder="Type a message..."></textarea>
                        <div class="send"><img src="./static/images/send.png" alt="send"></div>
                    </div>
                </div>
            </div>
            <div class="message">
                <div class="usr">
                    <div class="inf">
                        <div class="profil-pic"><img src="./static/images/user.png" alt=""></div>
                        <div>Shadow</div>
                    </div>
                    <div class="stat"></div>
                </div>
                <div class="discus">
                    <div class="from-usr">Hello <p class="message-date">17/02/2024</p></div>
                    <div class="from-usr">How are u ?<p class="message-date">17/02/2024</p></div>
                    <div class="from-exp">Hi <p class="message-date">17/02/2024</p></div>
                    <div class="from-exp">I'm fine and u ? <p class="message-date">17/02/2024</p></div>
                    <div class="from-usr">Yeah <p class="message-date">17/02/2024</p></div>
                    <div class="to-send">
                        <textarea name="sms" id="sms" placeholder="Type a message..."></textarea>
                        <div class="send"><img src="./static/images/send.png" alt="send"></div>
                    </div>
                </div>
            </div>
            <div class="message">
                <div class="usr">
                    <div class="inf">
                        <div class="profil-pic"><img src="./static/images/user.png" alt=""></div>
                        <div>Shadow</div>
                    </div>
                    <div class="stat"></div>
                </div>
                <div class="discus">
                    <div class="from-usr">Hello <p class="message-date">17/02/2024</p></div>
                    <div class="from-usr">How are u ?<p class="message-date">17/02/2024</p></div>
                    <div class="from-exp">Hi <p class="message-date">17/02/2024</p></div>
                    <div class="from-exp">I'm fine and u ? <p class="message-date">17/02/2024</p></div>
                    <div class="from-usr">Yeah <p class="message-date">17/02/2024</p></div>
                    <div class="to-send">
                        <textarea name="sms" id="sms" placeholder="Type a message..."></textarea>
                        <div class="send"><img src="./static/images/send.png" alt="send"></div>
                    </div>
                </div>
            </div>
            <div class="message">
                <div class="usr">
                    <div class="inf">
                        <div class="profil-pic"><img src="./static/images/user.png" alt=""></div>
                        <div>Shadow</div>
                    </div>
                    <div class="stat"></div>
                </div>
                <div class="discus">
                    <div class="from-usr">Hello <p class="message-date">17/02/2024</p></div>
                    <div class="from-usr">How are u ?<p class="message-date">17/02/2024</p></div>
                    <div class="from-exp">Hi <p class="message-date">17/02/2024</p></div>
                    <div class="from-exp">I'm fine and u ? <p class="message-date">17/02/2024</p></div>
                    <div class="from-usr">Yeah <p class="message-date">17/02/2024</p></div>
                    <div class="to-send">
                        <textarea name="sms" id="sms" placeholder="Type a message..."></textarea>
                        <div class="send"><img src="./static/images/send.png" alt="send"></div>
                    </div>
                </div>
            </div>
        </div>
    </div>
    </body>
`
var Signhead = `
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.4.2/css/all.min.css">
	<link rel="stylesheet" href="/static/CSS/styles.css">
	<title>Real Time Forum</title>
</head>`

var SignUpIn = `
<body>
    <div class="container" id="container">
        <div id="Message" class="Message"></div>
        <div class="form-container sign-up">
            <form id="registerForm">
                <div class="title">Create Account</div>
                <span>or use your email for registration</span>
                <input id="nickname" type="text" name="pseudo" placeholder="Nickname (4 to 8 alpha-num chars)">
                <input id="firstname" type="text" name="firstName" placeholder="First Name">
                <input id="lastname" type="text" name="lastName" placeholder="Last Name">
                <input id="age" type="number" name="age" placeholder="Age (min age 14 max 60)">
                <div class="gender">
                    <p>Gender</p>
                    <select id="gender" name="gender">
                        <option value="Male">Male</option>
                        <option value="Female">Female</option>
                    </select>
                </div>
                <input id="email" type="text" name="email" placeholder="Email (ex: a1@gmail.com)">
                <input id="password" type="password" name="password" placeholder="Password (min 4 chars, no spaces)">
                <button class="signup" disabled>Sign Up</button>
            </form>
        </div>
        <div class="form-container sign-in">
            <form id="loginForm">
                <h1>Sign In</h1>
                <span>or use your email password</span>
                <input id="loginMail" type="text" placeholder="Email or Nickname">
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
</body>
`

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
    <button><a>Go Home</a></button>
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

func handleFirstPage(w http.ResponseWriter, r *http.Request) {
	err := renderTemplateWithLayout(w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
}

func CheckActiveSession(r *http.Request) (*models.Users, bool) {
	var user *models.Users
	var data CookieData

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return nil, false
	} else {
		cookie := data.CookieValue
		user, err = database.GetUserByID(cookie)

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
		jsonResponse(w, http.StatusMethodNotAllowed, "Invalid request method")
		fmt.Println("Les données d'identification sont invalides: ", http.StatusBadRequest)
		return
	} else {
		user, exist = CheckActiveSession(r)
		nickname := ""
		if user == nil {
			exist = false
		} else {
			nickname = user.Nickname
		}

		res = Response{
			Exist:        exist,
			HomeHead:     Homehead,
			HomePage:     Home,
			SignHead:     Signhead,
			SignUpSignIn: SignUpIn,
			NickName:     nickname,
		}

		jsonResponse2(w, http.StatusOK, res)
	}
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	var user models.Users
	var login models.Register

	// Lire les données JSON de la requête
	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		jsonResponse(w, http.StatusBadRequest, "Bad Request")
		fmt.Println("Les données d'identification sont invalides: ", http.StatusBadRequest)
		return
	}

	// Recherche l'email ou le nickname de l'utilisateur dans la base de données
	err := database.DB.QueryRow("SELECT * FROM users WHERE Email = ? OR Nickname = ?", login.Credential, login.Credential).
		Scan(&user.ID, &user.Nickname, &user.Firstname, &user.Lastname, &user.Email, &user.Gender, &user.Age, &user.Password, &user.SessionExpiry)

	if err != nil {
		if err == sql.ErrNoRows {
			jsonResponse(w, http.StatusUnauthorized, "Bad Credentials ❌")
			return
		}
		fmt.Println(err) // Journalisation de l'erreur pour le débogage
		jsonResponse(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	// Vérification du mot de passe
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
	fmt.Println(login.Password)
	if err != nil {
		jsonResponse(w, http.StatusUnauthorized, "Bad Credentials ❌")
		fmt.Println("Mot de passe incorrect")
		return
	}

	// Calcul de l'heure d'expiration de la session (7 jours plus tard)
	sessionExpiry := time.Now().Add(7 * 24 * time.Hour)

	// Vérification si l'utilisateur a déjà une session
	var existingSessionID int
	err = database.DB.QueryRow("SELECT Id FROM sessions WHERE UserId = ?", user.ID).Scan(&existingSessionID)
	if err == nil {
		// Si l'utilisateur a déjà une session, mettre à jour la session existante
		_, err = database.DB.Exec("UPDATE sessions SET SessionExpiry = ? WHERE Id = ?", sessionExpiry, existingSessionID)
		if err != nil {
			fmt.Println(err)
			jsonResponse(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
	} else if err == sql.ErrNoRows {
		// Si l'utilisateur n'a pas de session, insérer une nouvelle session
		_, err = database.DB.Exec("INSERT INTO sessions (UserId, SessionExpiry) VALUES (?, ?)", user.ID, sessionExpiry)
		if err != nil {
			fmt.Println(err)
			jsonResponse(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
	} else {
		fmt.Println(err)
		jsonResponse(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	response := LoginSuccessResponse{
		Status:   201,
		Message:  "Success ✅",
		UserID:   user.ID,
		HomePage: Home,
		HomeHead: Homehead,
	}

	jsonResponse2(w, http.StatusOK, response)
}

func handleLogout(w http.ResponseWriter, r *http.Request) {
	var req struct {
		UserID string `json:"UserId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		jsonResponse(w, http.StatusBadRequest, "Bad Request")
		fmt.Println("Erreur lors du décodage de la requête:", err)
		return
	}

	_, err := database.DB.Exec("DELETE FROM sessions WHERE UserId = ?", req.UserID)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, "Erreur lors de la suppression de la session")
		fmt.Println("Erreur lors de la suppression de la session:", err)
		return
	}

	// Répond avec un succes
	jsonResponse(w, http.StatusOK, "Disconnected ✅")
}

func renderTemplateWithLayout(w http.ResponseWriter) error {
	w.WriteHeader(200)
	page, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println(err)
		return err
	}

	return page.Execute(w, "")
}

type CheckUserRequest struct {
	Nickname string `json:"nickname"`
}

type CheckUserResponse struct {
	Exists bool `json:"exists"`
}

func handleGetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		jsonResponse(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	var req CheckUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		jsonResponse(w, http.StatusBadRequest, "Bad Request")
		fmt.Println("Erreur lors du decodage: ", http.StatusBadRequest)
		return
	}

	exists, err := userExists(req.Nickname)
	if err != nil {
		fmt.Println(err) // Journalisation de l'erreur pour le débogage
		jsonResponse(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	res := CheckUserResponse{Exists: exists}
	jsonResponse2(w, 200, res)
}

func userExists(nickname string) (bool, error) {

	var count int
	err := database.DB.QueryRow("SELECT COUNT(*) FROM users WHERE Nickname = ?", nickname).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
