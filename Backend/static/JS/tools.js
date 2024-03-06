// export function checkSession() {
//     // Faire une requête API pour vérifier la session
//     fetch('http://localhost:8080/api/checksession', {
//         method: 'GET',
//         headers: {
//             'Content-Type': 'application/json',
//         },
//         credentials: 'include',  // Ajoutez cette ligne pour inclure les cookies dans la requête
//     })
//     .then(response => {
//         if (response.ok) {
//             // La session est valide, récupérez les données JSON de la réponse
//             return response.json();
//         } else {
//             // La session n'est pas valide, l'utilisateur doit se connecter
//             console.log('Session is not valid, proceed with login');
//             return Promise.reject('Session not validd');
//         }
//     })
//     .then(data => {
//         // Traitement des données JSON de la réponse
//         console.log('Session is valid');
//         console.log('User ID:', data.userID);
//         // Rediriger vers la page d'accueil ou effectuer d'autres actions nécessaires
//         window.location.href = '/homePage';
//     })
//     .catch(error => {
//         // Gestion des erreurs
//         console.error('Error checking session:', error);
//     });
// }

// //Recuperer le cookie
// function getCookieValue(cookieName) {
//     var name = cookieName + "=";
//     var decodedCookie = decodeURIComponent(document.cookie);
//     var cookieArray = decodedCookie.split(';');
//     for (var i = 0; i < cookieArray.length; i++) {
//         var cookie = cookieArray[i];
//         while (cookie.charAt(0) == ' ') {
//             cookie = cookie.substring(1);
//         }
//         if (cookie.indexOf(name) == 0) {
//             var cookieValue = cookie.substring(name.length, cookie.length);
//             // Vérifier si le cookie est expiré
//             var cookieExpires = cookieValue.split(';')[1];
//             if (!cookieExpires || new Date(cookieExpires.trim()) > new Date()) {
//                 return cookieValue.split(';')[0]; // Retourne la valeur du cookie
//             } else {
//                 // Cookie expiré, donc supprimer le cookie
//                 document.cookie = cookieName + "=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";
//                 return ""; // Retourne une chaîne vide si le cookie est expiré
//             }
//         }
//     }
//     return ""; // Retourne une chaîne vide si le cookie n'est pas trouvé
// }

// export async function checkSession() {
//     try {
//         console.log(result);
//         const cookieValue = getCookieValue("sessionID");

//         const response = await fetch('/api/activeSession', {
//             method: 'POST',
//             headers: {
//                 'Content-Type': 'application/json',
//             },
//             body: JSON.stringify({ cookieValue: cookieValue }),
//         });

//         if (!response.ok) {
//             throw new Error('Erreur lors de la requête');
//         }

//         const data = await response.json();
        
//         if (data.Exist) {
//             return { success: true, content: data.homePage, nickname: data.nickname };
//         } else {
//             return { success: false, content: data.signUpsignIn };
//         }
//     } catch (error) {
//         console.error('Erreur lors de la récupération de la session:', error);
//         return { error: error.message };
//     }
// }