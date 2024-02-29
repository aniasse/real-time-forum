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

document.addEventListener('DOMContentLoaded', () => {

    fetch('/api/activeSession', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
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

