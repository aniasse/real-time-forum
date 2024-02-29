// getting HTML elements
const nav = document.querySelector("nav"),
    toggleBtn = nav.querySelector(".toggle-btn");
toggleBtn.addEventListener("click", () => {
    nav.classList.toggle("open");
});

// js code to make draggable nav
function onDrag({ movementY }) { //movementY gets mouse vertical value
    const navStyle = window.getComputedStyle(nav), //getting all css style of nav
        navTop = parseInt(navStyle.top), // getting nav top value & convert it into string
        navHeight = parseInt(navStyle.height), // getting nav height value & convert it into string
        windHeight = window.innerHeight; // getting window height
    nav.style.top = navTop > 0 ? `${navTop + movementY}px` : "1px";
    if (navTop > windHeight - navHeight) {
        nav.style.top = `${windHeight - navHeight}px`;
    }
}

//this function will call when user click mouse's button and  move mouse on nav
nav.addEventListener("mousedown", () => {
    nav.addEventListener("mousemove", onDrag);
});
//these function will call when user relase mouse button and leave mouse from nav
nav.addEventListener("mouseup", () => {
    nav.removeEventListener("mousemove", onDrag);
});
nav.addEventListener("mouseleave", () => {
    nav.removeEventListener("mousemove", onDrag);
});

//Create Posts
const create = document.querySelector('.button-create');
const createPost = document.querySelector('.createpost');
const closePostBox = document.querySelector('.close')
const seeComs = document.querySelectorAll('.comment')
const coms = document.querySelectorAll('.comments')

create.addEventListener("click", () => {
    if (createPost.style.display === 'none' || createPost.style.display === '') createPost.style.display = 'flex'
    else createPost.style.display = 'none';
})

closePostBox.addEventListener('click', () => {
    if (createPost.style.display !== 'none') createPost.style.display = 'none'
})

seeComs.forEach((seeCom, i) => {
    seeCom.addEventListener('click', () => {
        if (coms[i].style.display === 'none' || coms[i].style.display === '') coms[i].style.display = 'flex';
        else coms[i].style.display = 'none';
    });
});

//Messages
const messages = document.querySelectorAll('.messages .usr')
const discus = document.querySelectorAll('.discus')

messages.forEach((mes, i) => {
    mes.addEventListener('click', () => {
        if (discus[i].style.display === 'none' || discus[i].style.display === '') discus[i].style.display = 'flex';
        else discus[i].style.display = 'none';
    })
})


//Responsivity
const homeIcon = document.querySelector('.first');
const notifIcon = document.querySelector('.thirth');
const smsIcon = document.querySelector('.second');

const notifs = document.querySelector('.notifs');
const posts = document.querySelector('.posts');
const smss = document.querySelector('.messages');

const makeResponsive = (selectedSection) => {
    // Masquer tous les éléments par défaut
    notifs.style.display = 'none';
    posts.style.display = 'none';
    smss.style.display = 'none';

    // Afficher la section sélectionnée
    selectedSection.style.display = 'flex';
    selectedSection.style.width = '100%';
}

notifIcon.addEventListener('click', () => {
    makeResponsive(notifs);
});

homeIcon.addEventListener('click', () => {
    makeResponsive(posts);
});

smsIcon.addEventListener('click', () => {
    makeResponsive(smss);
});


// Fonction pour gérer la responsivité
const handleResponsivity = () => {
    const screenWidth = window.innerWidth;

    if (screenWidth > 920) {
        notifs.style.display = 'flex';
        notifs.style.height = '50vh'
        posts.style.display = 'flex';
        smss.style.display = 'flex';
        smss.style.height = '55vh'
    } else {
        notifs.style.display = 'none';
        notifs.style.height = '80vh'
        posts.style.display = 'flex';
        smss.style.display = 'none';
        smss.style.height = '80vh'
    }
};

// Appel initial pour gérer la responsivité au chargement de la page
handleResponsivity();

// Écouteur d'événements pour la taille de l'écran
window.addEventListener('resize', handleResponsivity);