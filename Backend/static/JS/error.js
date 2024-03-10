//Ajout du Head 
const addingHead = (head) => {
    const newHead = document.createElement('head');
    const oldHead = document.querySelector('head');
    newHead.innerHTML = head;
    document.documentElement.replaceChild(newHead, oldHead);
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

const button = document.querySelector('.button')
console.log(button);

button.addEventListener('click', () => {
    window.history.replaceState(null, null, "/");
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
    addingHead(OnloadHead)
    loadScript('/static/JS/onload.js')
})