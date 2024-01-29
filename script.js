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
    pseudo: /^[a-zA-Z0-9].{2,}$/, // Remplacer par votre expression régulière
    firstName: /^[a-zA-Z]{2,}$/, // Remplacer par votre expression régulière
    lastName: /^[a-zA-Z]{2,}$/, // Remplacer par votre expression régulière
    age: /^(1[4-9]|[2-5][0-9]|60)$/, // Remplacer par votre expression régulière
    email: /^[\w-]+(\.[\w-]+)*@([\w-]+\.)+[a-zA-Z]{2,7}$/, // Remplacer par votre expression régulière
    password: /^(?=.*\d)(?=.*[a-z])(?=.*[A-Z]).{8,}$/, // Remplacer par votre expression régulière
};

// Récupérer tous les champs du formulaire
const inputs = document.querySelectorAll('input');

let signup = document.querySelector('.signup')

console.log(signup);

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
    return Array.from(inputs).every((input) => input.style.borderColor === 'green');
}

// Ajouter un écouteur d'événements au bouton
document.querySelector('.signup').addEventListener('click', function (event) {
    if (!checkValidity()) {
        event.preventDefault();
    }
});
