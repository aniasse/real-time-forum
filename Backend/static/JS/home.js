//Check session
async function checkSession() {
    try {
        const cookieValue = getCookieValue("sessionID");
        const response = await fetch('/api/activeSession', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ cookieValue: cookieValue }),
        });

        if (!response.ok) {
            throw new Error('Erreur lors de la requête');
        }

        return await response.json();
    } catch (error) {
        console.error('Erreur lors de la récupération de la session:', error);
        return { error: error.message };
    }
}

//Get Posts and display
async function fetchAndDisplayPosts() {
    try {
        const response = await fetch("/api/posts");
        const posts = await response.json();

        const postsContainer = document.querySelector('.posts');

        posts.reverse().forEach(post => {
            const formattedDate = formatDate(post.date);
            const postElement = `
                <div class="poster">
                    <div class="profil-pic">
                        <img src="./static/images/user.png" alt="">
                    </div>
                    <div class="info">
                        <h4>${post.nickname}</h4>
                        <small class="text-muted">${formattedDate}</small>
                    </div>
                </div>
                <div class="cat"><span>${post.category}</span></div>
                <div class="ctn">
                    <p class="text">${post.content}</p>
                </div>
                <div class="comment" data-post-id="${post.id}">
                    <img src="./static/images/comment.png" alt="">
                    <div class="comments" data-post-id="${post.id}"> </div>
                </div>
            `;
            const postDiv = document.createElement('div');
            postDiv.classList.add('post');
            postDiv.innerHTML = postElement;
            postDiv.setAttribute('data-post-id', post.id);
            postsContainer.appendChild(postDiv);
        });
    } catch (error) {
        console.error('Erreur lors de la récupération des posts:', error);
    }
}

//On home.js load
async function init() {
    const sessionResult = await checkSession();

    if (sessionResult.exist) {
        const create = `
            <div class="create">
                <div class="user">
                    <div class="profil-pic">
                        <img src="./static/images/user.png" alt="">
                    </div>
                    <h6>${sessionResult.nickname}</h6>
                </div>
                <div class="button-create">Create</div>
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
        `;
        // document.body.innerHTML = sessionResult.content; // Insérer la page d'accueil
        document.querySelector('.posts').innerHTML = '';
        document.querySelector('.posts').insertAdjacentHTML('beforeend', create);
        await fetchAndDisplayPosts(); // Récupérer et afficher les posts après l'affichage de la structure
    } else {
        document.body.innerHTML = sessionResult.content; // Insérer la page de connexion
        loadScript('/static/JS/sign.js');
    }

    viewPostBox()
    createAPost()
    viewComments()
    getUsers()
}

init();

//view post box
const viewPostBox = () => {
    const ShowCreatePost = document.querySelector('.button-create');
    const createPost = document.querySelector('.createpost');
    const closePostBox = document.querySelector('.close')
    // const seeComs = document.querySelectorAll('.comment')
    // const coms = document.querySelectorAll('.comments')

    const showCreate = () => {
        if (createPost.style.display === 'none' || createPost.style.display === '') {
            createPost.style.display = 'flex';
        } else {
            createPost.style.display = 'none';
        }
    }

    ShowCreatePost.addEventListener('click', () => {
        showCreate()
    });
    closePostBox.addEventListener('mousedown', () => {
        if (document.querySelector('.createpost').style.display !== 'none') document.querySelector('.createpost').style.display = 'none'
    })
}

//View comments
const viewComments = async () => {
    const commentImages = document.querySelectorAll('.comment img');

    commentImages.forEach(image => {
        image.addEventListener('click', async () => {
            const sessionResult = await checkSession();
            if (!sessionResult.exist) {
                // Redirection vers la page de connexion
                return;
            }

            const postId = image.parentNode.getAttribute('data-post-id');
            const commentsDiv = image.parentNode.querySelector('.comments');

            try {
                const response = await fetch(`/api/getComments`, {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({ PostId: postId }),
                });

                if (!response.ok) {
                    throw new Error('Erreur réseau : ' + response.status);
                }

                const comments = await response.json();

                commentsDiv.innerHTML = ''; // Nettoyer la division des commentaires

                if (comments.length === 0) {
                    commentsDiv.innerHTML = '<p>No comments yet.</p>';
                    const commentBox = document.createElement('div');
                    commentBox.className = 'tocom';
                    commentBox.innerHTML = `
                        <textarea aria-label="#" name="com" id="com" placeholder="comment..."></textarea>
                        <button class="btn-com">Comment</button>
                    `;
                    commentsDiv.appendChild(commentBox);
                } else {
                    const allComs = document.createElement('div');
                    allComs.className = 'all-com';

                    comments.reverse().forEach(comment => {
                        const commentDiv = document.createElement('div');
                        commentDiv.classList.add('usr');

                        const profilPic = document.createElement('div');
                        profilPic.classList.add('profil-pic');
                        const img = document.createElement('img');
                        img.src = './static/images/user.png';
                        profilPic.appendChild(img);

                        const username = document.createElement('h6');
                        username.textContent = comment.username;

                        commentDiv.appendChild(profilPic);
                        commentDiv.appendChild(username);

                        const content = document.createElement('p');
                        content.textContent = comment.content;

                        allComs.appendChild(commentDiv);
                        allComs.appendChild(content);
                    });

                    const commentBox = document.createElement('div');
                    commentBox.className = 'tocom';
                    commentBox.innerHTML = `
                        <textarea aria-label="#" name="com" id="com" placeholder="comment..."></textarea>
                        <button class="btn-com">Comment</button>
                    `;

                    allComs.appendChild(commentBox);
                    commentsDiv.appendChild(allComs);
                }
                // Affichage/masquage des commentaires
                commentsDiv.style.display = (commentsDiv.style.display === 'none' || commentsDiv.style.display === '') ? 'flex' : 'none';

                //Poster un commentaire
                postComment()
            } catch (error) {
                console.error("Erreur lors du chargement des commentaires:", error);
            }
        });
    });
}


//Post a comment
const postComment = () => {
    console.log('a');
    const commentButtons = document.querySelectorAll('.btn-com');

    console.log('b');

    commentButtons.forEach(button => {
        button.addEventListener('click', async () => {

            printCharging()

            const sessionResult = await checkSession();
            const userId = getCookieValue("sessionID");
            if (!sessionResult.exist || userId === "") {
                // Redirection vers la page de connexion
                return;
            }

            const commentText = button.parentNode.querySelector('textarea').value;
            const postId = button.parentNode.parentNode.parentNode.getAttribute('data-post-id');

            try {
                const response = await fetch(`/api/comment`, {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({ UserId: userId, PostId: postId, Content: commentText }),
                });

                if (!response.ok) {
                    //{status: 400, statusText: 'Bad Request'}
                    throw new Error('Erreur réseau : ' + response.status);
                }

                const data = await response.json();
                messageToShow(data)
                button.parentNode.querySelector('textarea').value = '';
            } catch (error) {
                console.error("Erreur lors de l'envoi du commentaire:", error);
                //Redirection Bad Request
            }
        });
    });
}



//Messages
// const messages = document.querySelectorAll('.messages .usr')
// const discus = document.querySelectorAll('.discus')

// messages.forEach((mes, i) => {
//     mes.addEventListener('click', () => {
//         if (discus[i].style.display === 'none' || discus[i].style.display === '') discus[i].style.display = 'flex';
//         else discus[i].style.display = 'none';
//     })
// })


//Responsivity
const nav = document.querySelector("nav"),
    toggleBtn = nav.querySelector(".toggle-btn");
toggleBtn.addEventListener("click", () => {
    nav.classList.toggle("open");
});

// js code to make draggable nav
function onDrag({ movementY }) { //movementY gets mouse vertical value
    const navStyle = window.getComputedStyle(nav), //getting all css style of nav
        navTop = parseInt(navStyle.top), // getting nav top value & convert it into string
        navHeight = parseInt(navStyle.height), // getting nav height value & convert it into string
        windHeight = window.innerHeight; // getting window height
    nav.style.top = navTop > 0 ? `${navTop + movementY}px` : "1px";
    if (navTop > windHeight - navHeight) {
        nav.style.top = `${windHeight - navHeight}px`;
    }
}

//this function will call when user click mouse's button and  move mouse on nav
nav.addEventListener("mousedown", () => {
    nav.addEventListener("mousemove", onDrag);
});
//these function will call when user relase mouse button and leave mouse from nav
nav.addEventListener("mouseup", () => {
    nav.removeEventListener("mousemove", onDrag);
});
nav.addEventListener("mouseleave", () => {
    nav.removeEventListener("mousemove", onDrag);
});

const homeIcon = document.querySelector('.first');
const notifIcon = document.querySelector('.thirth');
const smsIcon = document.querySelector('.second');

const notifs = document.querySelector('.notifs');
const posts = document.querySelector('.posts');
const smss = document.querySelector('.messages');

const makeResponsive = (selectedSection) => {
    // Masquer tous les éléments par défaut
    notifs.style.display = 'none';
    posts.style.display = 'none';
    smss.style.display = 'none';

    // Afficher la section sélectionnée
    selectedSection.style.display = 'flex';
    selectedSection.style.minWidth = '200px';
}

notifIcon.addEventListener('click', () => {
    makeResponsive(notifs);
});

homeIcon.addEventListener('click', () => {
    makeResponsive(posts);
});

smsIcon.addEventListener('click', () => {
    makeResponsive(smss);
});
// Fonction pour gérer la responsivité
const handleResponsivity = () => {
    const screenWidth = window.innerWidth;

    if (screenWidth > 920) {
        notifs.style.display = 'flex';
        notifs.style.height = '80vh'
        posts.style.display = 'flex';
        smss.style.display = 'flex';
        smss.style.height = '80vh'
    } else {
        notifs.style.display = 'none';
        notifs.style.height = '80vh'
        posts.style.display = 'flex';
        smss.style.display = 'none';
        smss.style.height = '80vh';
    }
};
// Appel initial pour gérer la responsivité au chargement de la page
handleResponsivity();
// Écouteur d'événements pour la taille de l'écran
window.addEventListener('resize', handleResponsivity);



//Access to cookie
function getCookieValue(cookieName) {
    var name = cookieName + "=";
    var decodedCookie = decodeURIComponent(document.cookie);
    var cookieArray = decodedCookie.split(';');
    for (var i = 0; i < cookieArray.length; i++) {
        var cookie = cookieArray[i];
        while (cookie.charAt(0) == ' ') {
            cookie = cookie.substring(1);
        }
        if (cookie.indexOf(name) == 0) {
            var cookieValue = cookie.substring(name.length, cookie.length);
            // Vérifier si le cookie est expiré
            var cookieExpires = cookieValue.split(';')[1];
            if (!cookieExpires || new Date(cookieExpires.trim()) > new Date()) {
                return cookieValue.split(';')[0]; // Retourne la valeur du cookie
            } else {
                // Cookie expiré, donc supprimer le cookie
                document.cookie = cookieName + "=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";
                return ""; // Retourne une chaîne vide si le cookie est expiré
            }
        }
    }
    return ""; // Retourne une chaîne vide si le cookie n'est pas trouvé
}

// Fonction pour formater la date
function formatDate(dateString) {
    const options = {
        day: "2-digit",
        month: "long",
        year: "numeric",
        hour: "numeric",
        minute: "numeric"
    };
    const date = new Date(dateString);
    return date.toLocaleDateString("fr-FR", options);
}

//Create Post
const createAPost = () => {
    const postButton = document.getElementById("postButton");

    postButton.addEventListener('click', async () => {
        printCharging()
        const categorie = document.getElementById("categorie").value;
        const postContent = document.getElementById("post").value;
        const userId = getCookieValue("sessionID")
        if (userId === "") {
            //redirect homePage
        }
        // event.preventDefault();
        try {
            const response = await fetch('/api/createPost', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ UserId: userId, Category: categorie, PostContent: postContent }),
            });
            if (response.ok) {
                const data = await response.json();
                if (data.status === 201) {
                    messageToShow(data)
                    setTimeout(() => {
                        document.getElementById("post").value = ''
                        document.querySelector('.createpost').style.display = 'none';
                    }, 1500)
                    return
                }
                messageToShow(data)
            } else {
                console.log('Erreur de requete', response.status);
                //Gestion d'erreur
            }
        } catch (error) {
            console.log('Error', error);
        }
    })
}

//Message to print 
const messageToShow = (data) => {
    setTimeout(function () {
        let toPrint = document.getElementById('Message');
        toPrint.innerText = data.message
        toPrint.style.visibility = 'visible';
        setTimeout(() => {
            toPrint.style.visibility = 'hidden';
            if (data.status !== 201) return
            container.classList.remove("active");
            inputs.forEach(input => {
                input.value = '';
            })
        }, 1000)
    }, 500);
}


//Print loader
const printCharging = () => {
    let div = document.createElement('div');
    let loader = document.createElement('div')
    div.classList = 'fullScreenDiv';
    loader.classList = 'loader'
    div.appendChild(loader)
    document.body.appendChild(div);
    setTimeout(() => {
        if (document.body.contains(div)) document.body.removeChild(div);
    }, 1000);
}


//Get Users for Messages

async function getUsers() {
    const sessionResult = await checkSession();
    const userId = getCookieValue("sessionID");
    if (!sessionResult.exist || userId === "") {
        // Redirection vers la page de connexion
        return;
    }

    try {
        const response = await fetch(`/api/getUsers`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ UserId: userId }),
        });

        if (!response.ok) {
            throw new Error('Erreur réseau : ' + response.status);
        }

        const data = await response.json();

        const messages = document.querySelector('.messages');
        messages.innerHTML = '';

        const title = `
            <div class="title">
                <span>
                    <h5>Messages</h5>
                </span>
            </div>
        `;

        messages.insertAdjacentHTML('beforeend', title);

        data.forEach(user => {
            const userDiv = `
                <div class="message">
                    <div class="usr">
                        <div class="inf">
                            <div class="profil-pic"><img src="./static/images/user.png" alt=""></div>
                            <div>${user.nickname}</div>
                        </div>
                        <div class="stat"></div>
                    </div>
                    <div class="discus"></div>
                </div>
            `;
            messages.insertAdjacentHTML('beforeend', userDiv);
        });

    } catch (error) {
        console.error("Erreur lors de la récupération des utilisateurs:", error);
        // Gestion de l'erreur
    }
}

