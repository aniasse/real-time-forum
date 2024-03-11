import { handleRedirection } from './tools.js'; 

document.querySelector('.button').addEventListener('click', async () => {
     handleRedirection()
     window.history.replaceState(null, null, "/");
     location.reload();
})