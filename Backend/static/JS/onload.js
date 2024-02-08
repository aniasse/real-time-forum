document.addEventListener('DOMContentLoaded', () => {

    fetch('/api/activeSession', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
    })
    .then(response => response.json())
    .then(data => {
        console.log(data);
        if (data.Exist) {
            document.body.innerHTML = '';
            document.body.insertAdjacentHTML('afterbegin', data.homePage)
            loadScript('/static/JS/home.js');
        }else{
            document.body.innerHTML = '';
            document.body.insertAdjacentHTML('afterbegin', data.signUpsignIn)
            loadScript('/static/JS/sign.js');
        }
    })
    .catch((error) => {
        console.log("errorr");
        console.log('Error:', error);
    });
    function loadScript(scriptUrl) {
        const script = document.createElement('script');
        script.src = scriptUrl;
        document.head.appendChild(script);
    }
});