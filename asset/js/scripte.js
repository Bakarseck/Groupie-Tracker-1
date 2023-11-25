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


document.addEventListener("DOMContentLoaded", function() {
    // Obtenez l'élément audio
    var monAudio = document.getElementById("monAudio");

    // Obtenez les boutons
    var btnPlay = document.getElementById("btnPlay");
    var btnStop = document.getElementById("btnStop");
    var btnNext = document.getElementById("btnNext");

    // Liste des fichiers audio
    var listeFichiersAudio = [
        "asset/audio/ash.mp3",
        "asset/audio/david.mp3",
        "asset/audio/Dido.mp3",
        "asset/audio/Zokush .mp3",
    ];

    // Index de la piste audio actuelle
    var indexPisteActuelle = 0;

    // Ajoutez un gestionnaire d'événements pour le clic sur le bouton Play
    btnPlay.addEventListener("click", function() {
        monAudio.src = listeFichiersAudio[indexPisteActuelle];
        monAudio.play();
        document.getElementById("image").style.animation= "rotate 2s linear infinite";
    });

    // Ajoutez un gestionnaire d'événements pour le clic sur le bouton Stop
    btnStop.addEventListener("click", function() {
        monAudio.pause();
        monAudio.currentTime = 0;
        document.getElementById("image").style.animation= "";
    });

    // Ajoutez un gestionnaire d'événements pour le clic sur le bouton Next
    btnNext.addEventListener("click", function() {
        // Incrémentez l'index de la piste actuelle
        indexPisteActuelle = (indexPisteActuelle + 1) % listeFichiersAudio.length;

        // Changez la source audio et démarrez la lecture
        monAudio.src = listeFichiersAudio[indexPisteActuelle];
        monAudio.play();
        document.getElementById("image").style.animation= "rotate 2s linear infinite";
    });
});




// Appeler la fonction d'initialisation de la carte lorsque la page est chargée
function convertDMSToDD(degrees, minutes, seconds, direction) {
    let dd = parseFloat(degrees) + parseFloat(minutes) / 60 + parseFloat(seconds) / (60 * 60);
    return direction === 'S' || direction === 'W' ? dd * -1 : dd;
}

function Mapping() {
    let mesDonnees;
    fetch("api")
        .then(response => {
            if (!response.ok) {
                throw new Error("Network response was not ok");
            }
            return response.json();
        })
        .then(data => {
            mesDonnees = data;
            console.log("Données récupérées:", mesDonnees);

            // Create and initialize Leaflet map object
            var map = L.map('map').setView([0, 0], 2); // Adjust the initial view as needed

            // Load map tiles
            L.tileLayer('https://{s}.basemaps.cartocdn.com/light_all/{z}/{x}/{y}.png', {
                attribution: 'Data <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors, '
                + 'Map tiles &copy; <a href="https://carto.com/attribution">CARTO</a>'
            }).addTo(map);

            // Custom icon for the markers
            var customIcon = L.icon({
                iconUrl: 'asset/img/point.png',
                iconSize: [32, 32],
                iconAnchor: [16, 32],
                popupAnchor: [0, -32]
            });

            // Iterate through coordinates and add markers with custom icon
            mesDonnees.forEach(coordinate => {
                // Extract degrees, minutes, seconds, and direction from the string
                const latParts = coordinate.Lat.match(/[0-9.]+|[A-Z]+/g);
                const lngParts = coordinate.Lng.match(/[0-9.]+|[A-Z]+/g);

                // Convert degrees, minutes, seconds to decimal degrees
                const lat = convertDMSToDD(latParts[0], latParts[1], latParts[2], latParts[3]);
                const lng = convertDMSToDD(lngParts[0], lngParts[1], lngParts[2], lngParts[3]);

                L.marker([lat, lng], { icon: customIcon }).addTo(map);
            });

            // Geosearch options
            var options = {
                key: 'YOUR-GEOSEARCH-KEY',
                position: 'topright',
            };

            // Add geosearch to the map
            var geosearchControl = L.Control.openCageGeosearch(options).addTo(map);
        })
        .catch(error => {
            console.error("Erreur lors de la récupération des données :", error);
        });
}

document.querySelector("button").addEventListener("click", e => {
    el = document.getElementById("map")
    el.style.width = "750px";   
    el.style.height = "500px";
    Mapping();
})



