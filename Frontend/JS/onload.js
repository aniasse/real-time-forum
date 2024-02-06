document.addEventListener('DOMContentLoaded', () => {

    fetch('http://localhost:8080/api/activeSession', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
    })
    .then(response => response.json())
    .then(data => {
        if (data.Exist) {
            document.body.innerHTML = '';
            document.body.insertAdjacentHTML('afterbegin', data.HomePage)
        }else{
            document.body.innerHTML = '';
            document.body.insertAdjacentHTML('afterbegin', data.SignUpSignIn)
        }
    })
    .catch((error) => {
        console.log('Error:', error);
    });
});