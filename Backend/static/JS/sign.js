import { handleLogin, handleRegister } from './tools.js'; 

const regist = document.getElementById('register')
if (regist) {
    regist.addEventListener('click', () => {
        document.getElementById('container').classList.add("active");
    });
}

const logi = document.getElementById('login')
if (logi){
    logi.addEventListener('click', () => {
        document.getElementById('container').classList.remove("active");
    });
}

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
// const inputs = document.querySelectorAll('.form-container.sign-up input');

// Récupérer la div contenant le formulaire
const formContainer = document.querySelector('.form-container.sign-up');

// Récupérer le formulaire
const registerForm = formContainer.querySelector('#registerForm');

// Récupérer tous les champs de saisie du formulaire
const inputs = registerForm.querySelectorAll('input');

if (inputs) {
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
}


// Fonction pour vérifier si tous les champs sont valides
function checkValidity() {
    const greenInputs = Array.from(inputs).filter((input) => input.style.borderColor === 'green');
    return greenInputs.length === 6;
}

// Ajouter un écouteur d'événements au bouton
const btn = document.querySelector('.signup');
if (btn) {
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
}

//Login
const logBtn = document.getElementById('loginForm')
if (logBtn) {
    logBtn.addEventListener('submit', async (event) => {
        event.preventDefault();
        await handleLogin()
    });
}

//Register
const regBtn = document.getElementById('registerForm')
if (regBtn) {
    regBtn.addEventListener('submit', async (event) => {
        event.preventDefault();
        await handleRegister()
    });
}
