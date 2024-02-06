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
    pseudo: /^[a-zA-Z0-9]{4,}$/,
    firstName: /^[a-zA-Z]{2,}$/,
    lastName: /^[a-zA-Z]{2,}$/,
    age: /^(1[4-9]|[2-5][0-9]|60)$/,
    email: /^[\w-]+(\.[\w-]+)*@([\w-]+\.)+[a-zA-Z]{2,7}$/,
    password: /^(?=.*\d)(?=.*[a-z])(?=.*[A-Z]).{8,}$/,
};

// Récupérer tous les champs du formulaire
const button =  document.querySelector('.signup');
const inputs = document.querySelectorAll('input');

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
    if (checkValidity()) {
        button.style.backgroundColor = '#512da8';
        button.style.cursor = 'pointer';
    }else{
        button.style.backgroundColor = 'red';
        button.style.cursor = 'not-allowed';
        // event.preventDefault();
    }
});

document.addEventListener('DOMContentLoaded', () => {
    const loginForm = document.getElementById('loginForm');
    const registerForm = document.getElementById('registerForm');

    loginForm.addEventListener('submit', (event) => {
        event.preventDefault();
        const email = document.getElementById('loginMail').value;
        const password = document.getElementById('loginPassword').value;
        console.log(email, password);
        //requête API pour le login
        fetch('http://localhost:8080/api/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ Email: email, Password: password }),
        })
            .then(response => response.json())
            .then(data => {
                console.log('Login success:', data);
            })
            .catch(error => {
                console.error('Login error:', error);
            });
    });

    registerForm.addEventListener('submit', (event) => {
        event.preventDefault();

        let div = document.createElement('div');
        div.classList = 'fullScreenDiv';
        let loader = document.createElement('div')
        loader.classList = 'loader'
        div.appendChild(loader)
        document.body.appendChild(div);

        setTimeout(()=> {
            document.body.removeChild(div);
            if (!validateForm()) return
        }, 2000);
        

        const email = document.getElementById('email').value;
        const password = document.getElementById('password').value;
        const nickname = document.getElementById('nickname').value;
        const firstname = document.getElementById('firstname').value;
        const lastname = document.getElementById('lastname').value;
        const age = document.getElementById('age').value;
        const newAge = parseInt(age, 10);
        const gender = document.getElementById('gender').value;

        console.log(email, password);
        //requête API pour l'inscription
        fetch('http://localhost:8080/api/register', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ Email: email, Password: password, Nickname: nickname, Firstname: firstname, Lastname: lastname, Age: newAge, Gender: gender }),
        })
            .then(response => response.json())
            .then(data => {
                console.log('Register success:', data);
            })
            .catch(error => {
                console.error('Register error:', error);
            });
    });
});

const printRegistered = () => {
    setTimeout(function() {
        document.getElementById('successMessage').style.visibility = 'visible';
        setTimeout(()=> {
            document.getElementById('successMessage').style.visibility = 'hidden';
        },1500)
    }, 1000); 
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
