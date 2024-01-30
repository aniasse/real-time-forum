const container = document.getElementById('container');
const registerBtn = document.getElementById('register');
const loginBtn = document.getElementById('login');

registerBtn.addEventListener('click', () => {
    container.classList.add("active");
});

loginBtn.addEventListener('click', () => {
    container.classList.remove("active");
});

// Définir les expressions régulières pour chaque champ
const regexMap = {
    pseudo: /^[a-zA-Z0-9]{4,}$/, // Remplacer par votre expression régulière
    firstName: /^[a-zA-Z]{2,}$/, // Remplacer par votre expression régulière
    lastName: /^[a-zA-Z]{2,}$/, // Remplacer par votre expression régulière
    age: /^(1[4-9]|[2-5][0-9]|60)$/, // Remplacer par votre expression régulière
    email: /^[\w-]+(\.[\w-]+)*@([\w-]+\.)+[a-zA-Z]{2,7}$/, // Remplacer par votre expression régulière
    password: /^(?=.*\d)(?=.*[a-z])(?=.*[A-Z]).{8,}$/, // Remplacer par votre expression régulière
};

// Récupérer tous les champs du formulaire
const button =  document.querySelector('.signup');

const inputs = document.querySelectorAll('input');

console.log(button);


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
document.querySelector('.signup').addEventListener('mouseover', function () {
    // event.preventDefault();
    if (checkValidity()) {
        button.style.backgroundColor = '#512da8';
        button.style.cursor = 'pointer';
    }else{
        button.style.backgroundColor = 'red';
        button.style.cursor = 'not-allowed';
    }
});



// console.log(login);


// function handleLoginSuccess() {
//     // Supposons que la connexion réussit ici
//     var signin = document.querySelector('.home')
//     var login = document.querySelector('.loginsignup')
//     login.style.display = 'none';
//     signin.style.display = 'block';
//     event.preventDefault();
// }

// document.addEventListener("DOMContentLoaded", function() {
//     var signin = document.querySelector('.home');
//     signin.style.display = 'none';
//     console.log(signin);
// });



// document.querySelector('.signin').addEventListener('click', handleLoginSuccess)