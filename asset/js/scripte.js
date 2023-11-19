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
