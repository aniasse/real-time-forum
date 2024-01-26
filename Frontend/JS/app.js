document.addEventListener('DOMContentLoaded', () => {
    const loginForm = document.getElementById('loginForm');
    const registerForm = document.getElementById('registerForm');

    loginForm.addEventListener('submit', (event) => {
        event.preventDefault();
        const email = document.getElementById('email').value;
        const password = document.getElementById('password').value;

        // Faire la requête API pour le login
        fetch('http://localhost:8080/api/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ email, password }),
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
        const newEmail = document.getElementById('Email').value;
        const newPassword = document.getElementById('Password').value;
        console.log(newEmail, newPassword);
        // Faire la requête API pour l'inscription
        fetch('http://localhost:8080/api/register', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ Email: newEmail, Password: newPassword }),
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
