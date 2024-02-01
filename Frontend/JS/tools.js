export function checkSession() {
    // Faire une requête API pour vérifier la session
    fetch('http://localhost:8080/api/checksession', {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
        },
        credentials: 'include',  // Ajoutez cette ligne pour inclure les cookies dans la requête
    })
    .then(response => {
        if (response.ok) {
            // La session est valide, récupérez les données JSON de la réponse
            return response.json();
        } else {
            // La session n'est pas valide, l'utilisateur doit se connecter
            console.log('Session is not valid, proceed with login');
            return Promise.reject('Session not valid');
        }
    })
    .then(data => {
        // Traitement des données JSON de la réponse
        console.log('Session is valid');
        console.log('User ID:', data.userID);
        // Rediriger vers la page d'accueil ou effectuer d'autres actions nécessaires
        window.location.href = '/home';
    })
    .catch(error => {
        // Gestion des erreurs
        console.error('Error checking session:', error);
    });
}
