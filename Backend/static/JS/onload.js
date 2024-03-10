function goodURL() {
    const currentURL = window.location.href;

    return (currentURL === "http://localhost:8080/") || (currentURL === "http://localhost:8080/#")
}

//Recuperer cookie
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

        const data = await response.json();

        if (data.exist) {
            console.log(data);
            return { success: true, head: data.homeHead, content: data.homePage, nickname: data.nickname };
        } else {
            return { success: false, head: data.signHead, content: data.signUpIn };
        }
    } catch (error) {
        handleError(error)
    }
}

// //Ajout de la class homepage a html
const addingHtmlClas = () => {
    // Sélection de la balise <html>
    const htmlElement = document.querySelector('html');

    if (!htmlElement.classList.contains('homepage')) htmlElement.classList.add('homepage');
}

//Ajout du Head 
const addingheader = (head) => {
    const newHead = document.createElement('head');
    const oldHead = document.querySelector('head');
    newHead.innerHTML = head;
    document.documentElement.replaceChild(newHead, oldHead);
}

//Ajout du body
const addingBody = (page) => {
    document.body.innerHTML = '';
    document.body.insertAdjacentHTML('afterbegin', page);
}

//Charger le fichier js de la page
function loadScript(scriptUrl) {

    const scripts = document.querySelectorAll('script');

    // Parcourir tous les scripts
    scripts.forEach(script => {
        // Vérifier si le script a un src contenant "/static/JS/sign.js"
        if (script.src.includes('/static/JS/sign.js') || script.src.includes('/static/JS/home.js')) {            // Supprimer le script s'il correspond
            script.remove();
        }
    });

    const script = document.createElement('script');
    script.src = scriptUrl;
    document.body.appendChild(script);
}


async function gettingErrorPage(status, message) {

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

//
async function handleError(error) {
    console.error('Error:', error);

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
    //Redirection
    addingheader(errorHead)
    addingBody(errorBody)
    loadScript('/static/JS/error.js')
}


//Au chargement
async function afterPageLoad() {
    if (!goodURL()) {
        await handleError({ response: { status: 404 } });  // Simuler une réponse 404
        return;
    }

    try {
        const result = await checkSession();
        if (result.success) {
            addingHtmlClas();
            addingheader(result.head)
            addingBody(result.content)
            loadScript('/static/JS/home.js');
        } else {
            addingheader(result.head)
            addingBody(result.content)
            loadScript('/static/JS/sign.js');
        }
    } catch (error) {
        handleError(error)
    }
}

afterPageLoad()




