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
    document.head.appendChild(script);
}

const container = document.getElementById('container');
const registerBtn = document.getElementById('register');
const loginBtn = document.getElementById('login');
const loginForm = document.getElementById('loginForm');
const registerForm = document.getElementById('registerForm');


registerBtn.addEventListener('click', () => {
    container.classList.add("active");
});
loginBtn.addEventListener('click', () => {
    container.classList.remove("active");
});

// Définir les expressions régulières pour chaque champ
const regexMap = {
    pseudo: /^[a-zA-Z0-9]{4,}$/,
    firstName: /^[a-zA-Z]{2,}$/,
    lastName: /^[a-zA-Z]{2,}$/,
    age: /^(1[4-9]|[2-5][0-9]|60)$/,
    email: /^[\w-]+(\.[\w-]+)*@([\w-]+\.)+[a-zA-Z]{2,7}$/,
    password: /^(?=.*\d)(?=.*[a-z])(?=.*[A-Z]).{8,}$/,
};

// Récupérer tous les champs du formulaire
const button = document.querySelector('.signup');
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
button.addEventListener('mouseover', function () {
    if (checkValidity()) {
        button.style.backgroundColor = '#512da8';
        button.style.cursor = 'pointer';
        button.disabled = false
    } else {
        button.style.backgroundColor = 'red';
        button.style.cursor = 'not-allowed';
        button.disabled = true
        // event.preventDefault();
    }
});

var Head = `<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<link rel="stylesheet" href="https://unicons.iconscout.com/release/v2.1.6/css/unicons.css" />
<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.5.1/css/all.min.css" />
<link href='https://unpkg.com/boxicons@2.1.1/css/boxicons.min.css' rel='stylesheet'>
<script src='https://kit.fontawesome.com/a076d05399.js' crossorigin='anonymous'></script>
<link rel="stylesheet" href="/static/CSS/home.css">
<title>Real Time Forum</title>
</head>`

function deleteCookie(name) {
    document.cookie = name + "=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";
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

//Login
loginForm.addEventListener('submit', async (event) => {
    event.preventDefault();
    const email = document.getElementById('loginMail').value;
    const password = document.getElementById('loginPassword').value;
    console.log(email, password);
    printLoader(false);
    try {
        const response = await fetch('http://localhost:8080/api/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ Email: email, Password: password }),
        });

        const data = await response.json();
        if (data.Status === 201) {
            setTimeout(() => {
                addingHtmlClass();
                addingHead();
                addingBody(data);
            }, 1500);
            createCookie("sessionID", data.ID, 7)
            return
        }
        messageToPrint(data);

    } catch (error) {
        console.error('Login error:', error);
    }
});

//Ajout de la class homepage a html
const addingHtmlClass = () => {
    // Sélection de la balise <html>
    const htmlElement = document.querySelector('html');
    
    if (!htmlElement.classList.contains('homepage')) htmlElement.classList.add('homepage');
}
//Ajout du Head 
const addingHead = () => {
    const newHead = document.createElement('head');
    const oldHead = document.querySelector('head');
    newHead.innerHTML = Head;
    document.documentElement.replaceChild(newHead, oldHead);
}

//Ajout du body
const addingBody = (data) => {
    document.body.innerHTML = '';
    document.body.insertAdjacentHTML('afterbegin', data.homePage);
    loadScript('/static/JS/home.js');
}

//Print loader
const printLoader = (validForm) => {
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

registerForm.addEventListener('submit', async (event) => {
    event.preventDefault();

    printLoader(true);

    const email = document.getElementById('email').value;
    const password = document.getElementById('password').value;
    const nickname = document.getElementById('nickname').value;
    const firstname = document.getElementById('firstname').value;
    const lastname = document.getElementById('lastname').value;
    const age = document.getElementById('age').value;
    const newAge = parseInt(age, 10);
    const gender = document.getElementById('gender').value;

    console.log(email, password);

    try {
        const response = await fetch('http://localhost:8080/api/register', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ Email: email, Password: password, Nickname: nickname, Firstname: firstname, Lastname: lastname, Age: newAge, Gender: gender }),
        });

        const data = await response.json();
        console.log('Register success:', data);
        messageToPrint(data);

    } catch (error) {
        console.error('Register error:', error);
    }
});

// });

const messageToPrint = (data) => {
    setTimeout(function () {
        let toPrint = document.getElementById('Message');
        toPrint.innerText = data.message
        toPrint.style.visibility = 'visible';
        setTimeout(() => {
            toPrint.style.visibility = 'hidden';
            if (data.Status !== 201) return
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

    console.log(genderVal);
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
