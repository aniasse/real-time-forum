export const addingHtmlClass = () => {
    const htmlElement = document.querySelector('html');

    if (!htmlElement.classList.contains('homepage')) htmlElement.classList.add('homepage');
}

//Ajout du Head 
export const addHead = (head) => {
    const newHead = document.createElement('head');
    const oldHead = document.querySelector('head');
    newHead.innerHTML = head;
    document.documentElement.replaceChild(newHead, oldHead);
}

//Ajout du body
export const addBody = (page) => {
    document.body.innerHTML = '';
    document.body.insertAdjacentHTML('afterbegin', page);
}

export function loadScript(scriptUrl) {

    const scripts = document.querySelectorAll('script');

    // Parcourir tous les scripts
    if (scripts.length !== 0) {
        scripts.forEach(script => {
            // Vérifier si le script a un src contenant "/static/JS/sign.js"
            if (script.src.includes('/static/JS/sign.js') || script.src.includes('/static/JS/home.js') || script.src.includes('/static/JS/error.js')) {
                script.remove();
            }
        });
    }

    const script1 = document.createElement('script');
    const script = document.createElement('script');
    script1.src = '/static/JS/tools.js';
    script.src = scriptUrl;
    script1.type = 'module'
    script.type = "module"
    document.body.appendChild(script1);
    document.body.appendChild(script);
}

export async function gettingErrorPage(status, message) {

    try {
        const response = await fetch('/api/handleError', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ Status: status, Message: message })
        })

        const data = await response.json();

        return { status: data.status, message: data.message }

    } catch (error) {
        console.log("error", error);
        handleError(error)
    }
}

export async function handleError(error) {
    console.log('Error:', error);

    let data; // Variable pour stocker les données d'erreur

    // Gérer différents types d'erreurs
    if (error.response) {
        switch (error.response.status) {
            case 404:
                data = await gettingErrorPage("404", "Page Not Found");
                break;
            case 405:
                data = await gettingErrorPage("405", "Method Not Allowed");
                break;
            case 400:
                data = await gettingErrorPage("400", "Bad Request");
                break;
            case 401:
                data = await gettingErrorPage("401", "Session Expired");
                break;
            default:
                // Gérer les autres codes d'état HTTP
                data = await gettingErrorPage("500", "Something went wrong. Please try again later.");
        }
    } else {
        // Autre erreur non-HTTP (erreur JavaScript, etc.)
        data = await gettingErrorPage("500", "Something went wrong. Please try again later.");
    }


    var errorHead = `
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Error Page</title>
        <link rel="stylesheet" href="/static/CSS/error.css">
    </head>
    `

    var errorBody = `
    <body>
        <div class="container">
            <div class="status">${data.status}</div>
            <div class="message">${data.message}</div>
            <div class="button">GO HOME</div>
        </div>
    </body>
    `

    // Redirection
    addHead(errorHead)
    addBody(errorBody)
    loadScript('/static/JS/error.js')
}

export function getCookieValue(cookieName) {
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
            var cookieExpires = cookie.split(';')[1];
            if (!cookieExpires || new Date(cookieExpires.trim()) > new Date()) {
                // Retourne la valeur du cookie et la date d'expiration
                return /*{ value:*/ cookieValue.split(';')[0]//, expires: cookieExpires.trim() };
            } else {
                // Cookie expiré, donc supprimer le cookie
                document.cookie = cookieName + "=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";
                return ""
            }
        }
    }
    return "" // Retourne une chaîne vide si le cookie n'est pas trouvé
}


export function goodURL() {
    const currentURL = window.location.pathname;

    return (currentURL === "/") || (currentURL === "/#")
}

export async function checkSession() {

    try {
        const cookieValue = getCookieValue("sessionID");

        const response = await fetch('/api/activeSession', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ cookieValue: cookieValue }),
        });

        const data = await response.json();

        if (data.exist) {

            return { success: true, head: data.homeHead, content: data.homePage, nickname: data.nickname, userId: cookieValue };
        } else {
            return { success: false, head: data.signHead, content: data.signUpIn };
        }
    } catch (error) {
        console.log('error', error);
        handleError(error)
    }
}

export function deleteCookie(name) {
    createCookie(name, "", -1); // expiration dans le passé
}

export function createCookie(name, value, exp) {
    var expiration = "";
    if (exp) {
        var date = new Date();
        date.setTime(date.getTime() + (exp * 24 * 60 * 60 * 1000));
        expiration = "; expires=" + date.toUTCString();
    }
    document.cookie = name + "=" + value + expiration + "; path=/";
}

export async function handleLogin() {
    console.log('handleLogin');
    const email = document.getElementById('loginMail').value;
    const password = document.getElementById('loginPassword').value;
    printLoader(false);
    try {
        const response = await fetch('/api/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ Credential: email, Password: password }),
        });

        const data = await response.json();

        if (data.status === 201) {
            setTimeout(() => {
                addingHtmlClass();
                addHead(data.homeHead);
                addBody(data.homePage);
                createCookie("sessionID", data.userID, 7)
                loadScript('/static/JS/home.js')
            }, 1500)
        } else {
            messageToPrint(data);
        }

    } catch (error) {
        console.log('error', error);
        handleError(error)
    }
}

export function validateForm() {
    // Récupérer les éléments du formulaire par leur ID
    var pseudoElement = document.getElementById('nickname');
    var firstNameElement = document.getElementById('firstname');
    var lastNameElement = document.getElementById('lastname');
    var ageElement = document.getElementById('age');
    var emailElement = document.getElementById('email');
    var passwordElement = document.getElementById('password');
    var gender = document.getElementById('gender')

    // Récupérer les valeurs des éléments du formulaire
    var pseudoValue = pseudoElement.value;
    var firstNameValue = firstNameElement.value;
    var lastNameValue = lastNameElement.value;
    var ageValue = ageElement.value;
    var emailValue = emailElement.value;
    var passwordValue = passwordElement.value;
    var genderVal = gender.value;

    // Vérifier chaque champ avec son expression régulière correspondante
    for (var key in regexMap) {
        var regex = regexMap[key];
        var value = eval(key + 'Value');
        if (!regex.test(value)) {
            alert('Please enter valid ' + key);
            return false;
        }
    }

    // Si tous les champs sont valides, retourner true
    return (genderVal === "Male" || genderVal === "Female");
}

export let printLoader = (validForm) => {
    let div = document.createElement('div');
    let loader = document.createElement('div')
    div.classList = 'fullScreenDiv';
    loader.classList = 'loader'
    div.appendChild(loader)
    document.body.appendChild(div);
    setTimeout(() => {
        if (document.body.contains(div)) document.body.removeChild(div);
        if (validForm) validateForm()
    }, 1000);
}

export const messageToPrint = (data) => {
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

export async function handleRegister() {
    console.log('handleRegister')
    printLoader(true);

    const email = document.getElementById('email').value;
    const password = document.getElementById('password').value;
    const nickname = document.getElementById('nickname').value;
    const firstname = document.getElementById('firstname').value;
    const lastname = document.getElementById('lastname').value;
    const newAge = document.getElementById('age').value;
    // const newAge = parseInt(age, 10);
    const gender = document.getElementById('gender').value;


    try {
        const response = await fetch('api/register', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ Email: email, Password: password, Nickname: nickname, Firstname: firstname, Lastname: lastname, Age: newAge, Gender: gender }),
        });

        const data = await response.json();
        messageToPrint(data);

    } catch (error) {
        console.log('error', error);
        handleError(error)
    }
}

/* Home */

export const messageToShow = (data) => {
    setTimeout(function () {
        let toPrint = document.getElementById('Message');
        toPrint.innerText = data.message
        toPrint.style.visibility = 'visible';
        setTimeout(() => {
            toPrint.style.visibility = 'hidden';
        }, 1000)
    }, 500);
}

export function formatDate(dateString) {
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

export const printCharging = () => {
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

export const removeHtmlClass = () => {
    // Sélection de la balise <html>
    const htmlElement = document.querySelector('html');

    // Vérification si la classe "homepage" est présente
    if (htmlElement.classList.contains('homepage')) {
        // Suppression de la classe "homepage"
        htmlElement.classList.remove('homepage');
    }
};


export async function handleLogout() {
    console.log('handleLogout');
    printCharging()

    const sessionResult = await checkSession()

    if (!sessionResult.success) {
        await handleError({ response: { status: 404 } });
        return;
    }

    const userId = sessionResult.userId

    try {
        const response = await fetch(`/api/logout`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ UserId: userId })
        })

        if (!response.ok) {
            await handleError({ response: { status: 500 } });  // Simuler une réponse 404
            return;
        }

        deleteCookie("sessionID")

        var OnloadHead = `
            <head>
                <meta charset="UTF-8">
                <meta name="viewport" content="width=device-width, initial-scale=1.0">
                <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.4.2/css/all.min.css">
                <link rel="stylesheet" href="/static/CSS/styles.css">
                <title>Real Time Forum</title>
            </head>
            `
        removeHtmlClass()
        document.body.innerHTML = ''
        addHead(OnloadHead)
        loadScript('/static/JS/onload.js')
        window.history.replaceState(null, null, "/");
        location.reload();
    } catch (error) {
        console.log('error', error);
        handleError(error)
    }
}

export async function fetchAndDisplayPosts() {
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
        console.log('Erreur lors de la récupération des posts:', error);
        handleError(error)
    }
}

export async function init() {

    printCharging();

    const sessionResult = await checkSession();

    if (sessionResult.success) {
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

    viewPostBox();
    createAPost();
    viewComments();
    getUsers();
    inableWebsocket()
}


export const viewPostBox = () => {
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
export const viewComments = async () => {
    const commentImages = document.querySelectorAll('.comment img');

    commentImages.forEach(image => {
        image.addEventListener('click', async () => {
            const sessionResult = await checkSession();
            if (!sessionResult.success) {
                await handleError({ response: { status: 404 } });
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
                    await handleError({ response: { status: 500 } });  // Simuler une réponse 404
                    return;
                }

                const comments = await response.json();

                commentsDiv.innerHTML = '';

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

                        const username = document.createElement('h5');
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
                console.log("Erreur lors du chargement des commentaires:", error);
                handleError(error)
            }
        });
    });
}


//Post a comment
export const postComment = () => {
    const commentButtons = document.querySelectorAll('.btn-com');

    commentButtons.forEach(button => {
        button.addEventListener('click', async () => {

            printCharging()

            const sessionResult = await checkSession();
            if (!sessionResult.success) {
                // Redirection vers la page de connexion
                await handleError({ response: { status: 404 } });
                return;
            }
            const userId = sessionResult.userId

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
                    await handleError({ response: { status: 500 } });  // Simuler une réponse 404
                    return;
                }

                const data = await response.json();
                messageToShow(data)
                button.parentNode.querySelector('textarea').value = '';
            } catch (error) {
                console.log("Erreur lors de l'envoi du commentaire:", error);
                handleError(error)
                //Redirection Bad Request
            }
        });
    });
}


export const createAPost = () => {
    const postButton = document.getElementById("postButton");

    postButton.addEventListener('click', async () => {
        printCharging()
        const categorie = document.getElementById("categorie").value;
        const postContent = document.getElementById("post").value;
        const userId = getCookieValue("sessionID")
        if (userId === "") {
            await handleError({ response: { status: 404 } });
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
                    setTimeout(async () => {
                        document.getElementById("post").value = ''
                        document.querySelector('.createpost').style.display = 'none';
                    }, 1500)
                    return
                }
                messageToShow(data)
            } else {
                //Gestion d'erreur
            }
        } catch (error) {
            handleError(error)
        }
    })
}


export async function getUsers() {
    const sessionResult = await checkSession();
    if (!sessionResult.success) {
        await handleError({ response: { status: 404 } });
        return;
    }

    const userId = sessionResult.userId

    try {
        const response = await fetch(`/api/getUsers`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ UserId: userId }),
        });

        if (!response.ok) {
            await handleError({ response: { status: 500 } });  // Simuler une réponse 404
            return;
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
            const statusColor = user.status === 'green' ? 'green' : 'red';
            const userDiv = `
                <div class="message">
                    <div class="usr">
                        <img src="./static/images/user.png" alt="">
                        <p>${user.nickname}</p>
                        <div class="stat" style="background: ${statusColor};"></div>
                    </div>
                </div>
            `;
            messages.insertAdjacentHTML('beforeend', userDiv);
        });
        getDiscutions();

    } catch (error) {
        console.log("Erreur lors de la récupération des utilisateurs:", error);
        handleError(error)
        // Gestion de l'erreur
    }
}

export function formatMDate(dateString) {
    const [date, time] = dateString.split(' ');
    const [day, month, year] = date.split('/');
    return `${day}/${month}/${year} à ${time}`;
}

// Initialisation websocket
export const inableWebsocket = () => {
    const homepage = document.querySelector('.homepage')

    if (!homepage) {
        console.warn("La page d'accueil n'existe pas. WebSocket est désactivé.");
        return; 
    }
    
    const socket = new WebSocket("ws://localhost:8080/ws");

    socket.addEventListener("message", async function (event) {
        const message = JSON.parse(event.data);
        console.log("Message reçu:", message);
        await getUsers()
        await handleReceivedMessage(message);
    });
    
     async function isPrintable(message) {
        const sessionResult = await checkSession()
        if (!sessionResult.hasOwnProperty('nickname')) return false
        const receiver = document.querySelector('.sms .usr p').textContent
        return message.sender === receiver && message.receiver === sessionResult.nickname
    }
    
     async function handleReceivedMessage(message) {
        // Logique pour traiter le message reçu
        if (await isPrintable(message)) {
            // Exemple : Afficher le message dans le chat box
            printMessage(message.content, "from-exp", message.timestamp);
        }
    }
}

// Fonction pour envoyer un message au serveur via WebSocket
export function sendMessage(sender, receiver, content, time) {
    let message = {
        Sender: sender,
        Receiver: receiver,
        Content: content,
        Timestamp: time,
    };
    console.log("le message", message);
    socket.send(JSON.stringify(message));
}


// Fonction pour envoyer un message du chat box via WebSocket
export function sendChatMessage(sender, receiver) {
   const messageInput = document.getElementById('sms');
   const messageContent = messageInput.value.trim();
   if (messageContent === "") {
       return;
   }
   let timestamp = formatMDate(new Date().toLocaleString())
   sendMessage(sender, receiver, messageContent, timestamp);
   messageInput.value = ""; // Effacer le contenu après l'envoi
   printMessage(messageContent, "from-usr", timestamp)
}

export function printMessage(messageContent, from, timestamp) {
    const discus = document.querySelector('.discus')
    let smsDiv = document.createElement('div')
    smsDiv.className = from
    smsDiv.textContent = messageContent
    let date = document.createElement('p')
    date.className = 'sms-date'
    date.textContent = timestamp
    smsDiv.appendChild(date)
    discus.appendChild(smsDiv)
}

export function createMessageDiv(message) {
    let msgDiv = document.createElement('div');
    msgDiv.className = message.from === 'user' ? 'from-usr' : 'from-exp';
    msgDiv.textContent = message.text;

    let messageDate = document.createElement('p');
    messageDate.className = 'sms-date';
    messageDate.textContent = message.date;

    msgDiv.appendChild(messageDate);

    return msgDiv;
}

export async function getDiscutions() {
    const discussions = document.querySelectorAll('.messages .message');

    discussions.forEach(discus => {
        discus.addEventListener('click', async () => {
            const sessionResult = await checkSession();

            if (!sessionResult.success) {
                await handleError({ response: { status: 404 } });
                return;
            }

            const senderNickname = sessionResult.nickname

            const nickname = discus.querySelector('.usr p').textContent;

            try {
                const response = await fetch(`/api/getDiscussions`, {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({ SenderNickname: senderNickname, ReceiverNickname: nickname })
                });

                if (!response.ok) {
                    await handleError({ response: { status: 500 } });  // Simuler une réponse 404
                    return;
                }

                const messages = await response.json();
                let fullDiv = document.querySelector('.fullScreenDiv');
                if (!fullDiv) {
                    fullDiv = document.createElement('div');
                    fullDiv.className = 'fullScreenDiv';
                }
                fullDiv.innerHTML = '';

                let sms = document.createElement('div')
                sms.className = 'sms'
                let smsInner = `
                <div class="usr">
                    <img src="./static/images/goback.png" class="goback" alt="">
                    <img src="./static/images/user.png" alt="">
                    <p>${nickname}</p>
                </div>
                `
                sms.insertAdjacentHTML('beforeend', smsInner);

                let writeSms = `
                <div class="to-send">
                    <textarea name="sms" id="sms" placeholder=" Type a message..."></textarea>
                    <img id="sendButton" src="./static/images/send.png" alt="Send">
                </div>
                `
                let discus = document.createElement('div');
                discus.className = 'discus';

                if (messages) {
                    messages.forEach(message => {

                        let msgDiv = createMessageDiv(message)

                        discus.appendChild(msgDiv);
                    });
                }

                sms.append(discus)
                sms.innerHTML += writeSms
                fullDiv.appendChild(sms)

                // Ajouter le fullDiv à un conteneur existant dans votre page
                document.body.appendChild(fullDiv);

                document.getElementById('sendButton').addEventListener("click", async function () {
                    const sessionResult = await checkSession()
                    if (!sessionResult.hasOwnProperty('userId')) {
                        return await handleError({ response: { status: 500 } })
                    }

                    const receiverName = document.querySelector('.sms .usr p').textContent
                    console.log('receveiver', receiverName);

                    try {
                        const response = await fetch('/api/checkUser', {
                            method: 'POST',
                            headers: {
                                'Content-Type': 'application/json'
                            },
                            body: JSON.stringify({ nickname: receiverName })
                        });

                        const data = await response.json();

                        if (data.exists) {
                            sendChatMessage(sessionResult.nickname, receiverName);
                        } else {
                            return await handleError({ response: { status: 400 } })
                        }
                    } catch (error) {
                        console.log('Erreur lors de la vérification de l\'utilisateur :', error);
                        handleError(error)
                    }
                });

                // Ajouter l'événement pour vider fullDiv en cliquant sur l'image .goback
                const gobackImg = fullDiv.querySelector('.goback');
                if (gobackImg) {
                    gobackImg.addEventListener('click', () => {
                        document.body.removeChild(fullDiv);
                    });
                }

            } catch (error) {
                console.log("Error: getting Discuss", error);
                handleError(error)
            }
        });
    });
}

export function changeUrl(newUrl) {
    window.history.pushState({}, '', newUrl);
    // Mettez à jour le contenu de la page en fonction du newUrl
}

export function handleRedirection() {
    changeUrl('/');

    var OnloadHead = `
    <head>
	    <meta charset="UTF-8">
	    <meta name="viewport" content="width=device-width, initial-scale=1.0">
	    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.4.2/css/all.min.css">
	    <link rel="stylesheet" href="/static/CSS/styles.css">
	    <title>Real Time Forum</title>
    </head>
    `
    document.body.innerHTML = ''
    addHead(OnloadHead)
    loadScript('/static/JS/onload.js')
}


export async function afterPageLoad() {
    // let redirected
    if (!goodURL()) {
        // redirected = await handleError({ response: { status: 404 } });  // Simuler une réponse 404
        return await handleError({ response: { status: 404 } });
    }

    try {
        const result = await checkSession();
        if (result.success) {
            addingHtmlClass();
            addHead(result.head)
            addBody(result.content)
            loadScript('/static/JS/home.js');
        } else {
            addHead(result.head)
            addBody(result.content)
            loadScript('/static/JS/sign.js');
        }
    } catch (error) {
        console.log('error', error);
        handleError(error)
    }
}