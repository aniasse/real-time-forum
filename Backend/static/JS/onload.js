// import { checkSession } from "./tools";

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

        const data = await response.json();
        
        if (data.Exist) {
            return { success: true, head: data.homeHead, content: data.homePage, nickname: data.nickname };
        } else {
            return { success: false, head: data.signHead, content: data.signUpIn };
        }
    } catch (error) {
        console.error('Erreur lors de la récupération de la session:', error);
        return { error: error.message };
    }
}

//Au chargement
document.addEventListener('DOMContentLoaded', () => {

    checkSession()
        .then(result => {
            console.log("checkSession result:", result);
            if (result.success) {
                addingHead(result.head)
                addingBody(result.content)
                loadScript('/static/JS/home.js');
            } else {
                addingHead(result.head)
                addingBody(result.content)

                loadScript('/static/JS/sign.js');
            }
        })
        .catch(error => {
            console.error('Erreur lors de la récupération de la session:', error);
            // Gérer l'erreur ici
        });
});

//Ajout du Head 
const addingHead = (head) => {
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


