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
    addinghead(errorHead)
    addingbody(errorBody)
    loadScript('/static/JS/error.js')
}

function loadScript(scriptUrl) {
    const scripts = document.querySelectorAll('script');

    // Parcourir tous les scripts
    scripts.forEach(script => {
        // Vérifier si le script a un src contenant "/static/JS/sign.js"
        if (script.src && script.src.includes("/static/JS/sign.js")) {
            // Supprimer le script s'il correspond
            script.remove();
        }
    });

    const script = document.createElement('script');
    script.src = scriptUrl;
    document.body.appendChild(script);
}


document.getElementById('register').addEventListener('click', () => {
    document.getElementById('container').classList.add("active");
});
document.getElementById('login').addEventListener('click', () => {
    document.getElementById('container').classList.remove("active");
});

// Définir les expressions régulières pour chaque champ
const regexMap = {
    pseudo: /^[a-zA-Z0-9]{4,8}$/,
    firstName: /^[a-zA-Z]{2,}$/,
    lastName: /^[a-zA-Z]{2,}$/,
    age: /^(1[4-9]|[2-5][0-9]|60)$/,
    email: /^[\w-]+(\.[\w-]+)*@([\w-]+\.)+[a-zA-Z]{2,7}$/,
    password: /^[!-~]{4,}$/,
};

// Récupérer tous les champs du formulaire
const btn = document.querySelector('.signup');
const inputs = document.querySelectorAll('.form-container.sign-up input');

// Pour chaque champ, ajouter un écouteur d'événements
inputs.forEach((input) => {
    input.addEventListener('input', function () {
        const key = this.getAttribute('name');
        const regex = regexMap[key];
        if (regex.test(this.value)) {
            this.style.borderColor = 'green';
        } else {
            this.style.borderColor = 'red';
        }
    });
});

// Fonction pour vérifier si tous les champs sont valides
function checkValidity() {
    const greenInputs = Array.from(inputs).filter((input) => input.style.borderColor === 'green');
    return greenInputs.length === 6;
}

// Ajouter un écouteur d'événements au bouton
btn.addEventListener('mouseover', function () {
    if (checkValidity()) {
        btn.style.backgroundColor = '#512da8';
        btn.style.cursor = 'pointer';
        btn.disabled = false
    } else {
        btn.style.backgroundColor = 'red';
        btn.style.cursor = 'not-allowed';
        btn.disabled = true
        // event.preventDefault();
    }
});

// function deleteCookie(name) {
//     document.cookie = name + "=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";
// }

function deleteCookie(name) {
    createCookie(name, "", -1); // expiration dans le passé
}

function createCookie(name, value, exp) {
    var expiration = "";
    if (exp) {
        var date = new Date();
        date.setTime(date.getTime() + (exp * 24 * 60 * 60 * 1000));
        expiration = "; expires=" + date.toUTCString();
    }
    document.cookie = name + "=" + value + expiration + "; path=/";
}

// //Ajout de la class homepage a html
const addingHtmlClass = () => {
    // Sélection de la balise <html>
    const htmlElement = document.querySelector('html');

    if (!htmlElement.classList.contains('homepage')) htmlElement.classList.add('homepage');
}
//Ajout du Head 
const addinghead = (data) => {
    const newHead = document.createElement('head');
    const oldHead = document.querySelector('head');
    newHead.innerHTML = data.homeHead;
    document.documentElement.replaceChild(newHead, oldHead);
}

//Ajout du body
const addingbody = (data) => {
    document.body.innerHTML = '';
    document.body.insertAdjacentHTML('afterbegin', data.homePage);
    loadScript('/static/JS/home.js');
}

//Login
document.getElementById('loginForm').addEventListener('submit', async (event) => {
    event.preventDefault();
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
        console.log(data);
        if (data.Status === 201) {
            setTimeout(() => {
                addingHtmlClass();
                addinghead(data);
                addingbody(data);
                createCookie("sessionID", data.userID, 7)
            }, 1500)
        } else {
            messageToPrint(data);
        }

    } catch (error) {
        handleError(error)
    }
});



//Print loader
let printLoader = (validForm) => {
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

//Register
document.getElementById('registerForm').addEventListener('submit', async (event) => {
    event.preventDefault();

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
        handleError(error)
    }
});


const messageToPrint = (data) => {
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

function validateForm() {
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
