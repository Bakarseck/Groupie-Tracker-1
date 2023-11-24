$(document).ready(function () {
    var arrow = $('<div class="arrow">&#9660;</div>');

    $('body').append(arrow);

    $(window).scroll(function () {
        if ($(this).scrollTop() <= 25) {
            arrow.addClass('active').fadeIn();
        } else {
            arrow.removeClass('active').fadeOut();
        }
    });

    arrow.click(function () {
        $('html, body').animate({ scrollTop: 0 }, 'slow');
    });
});




// Appeler la fonction d'initialisation de la carte lorsque la page est chargée

function Mapping() {
    let mesDonnees;  // Déclarer la variable à l'extérieur de la chaîne de promesses 
    fetch("api")
        .then(response => response.json())  // Convertir la réponse en JSON
        .then(data => {
            mesDonnees = data;  // Assigner les données à la variable
            console.log("Données récupérées:", mesDonnees);
        })
        .catch(error => console.error("Erreur lors de la récupération des données :", error));
    document.getElementById("map").innerHTML = `<iframe width="425" height="350" src="https://www.openstreetmap.org/export/embed.html?bbox=-22.510986328125%2C9.459898921269597%2C-11.195068359375002%2C18.44834670293207&amp;layer=mapnik" style="border: 1px solid black"></iframe><br/><small><a href="https://www.openstreetmap.org/#map=7/13.998/-16.853">View Larger Map</a></small>`;

  
}

document.querySelector("button").addEventListener("click", e => {
    Mapping()
})


