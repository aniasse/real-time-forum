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
    document.head.appendChild(script);
}

//Recuperer le cookie
function getCookieValue(cookieName) {
    var name = cookieName + "=";
    var decodedCookie = decodeURIComponent(document.cookie);
    var cookieArray = decodedCookie.split(';');
    for(var i = 0; i < cookieArray.length; i++) {
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

//Au chargement
document.addEventListener('DOMContentLoaded', () => {

    const cookieValue = getCookieValue("sessionID")

    fetch('/api/activeSession', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ cookieValue: cookieValue }),
    })
        .then(response => response.json())
        .then(data => {
            if (data.Exist) {
                document.body.innerHTML = '';
                document.body.insertAdjacentHTML('afterbegin', data.homePage)
                loadScript('/static/JS/home.js');
            } else {
                document.body.innerHTML = '';
                document.body.insertAdjacentHTML('afterbegin', data.signUpsignIn)
                loadScript('/static/JS/sign.js');
            }
        })
        .catch((error) => {
            console.log("errorr");
            console.log('Error:', error);
        });
});



