$(document).ready(function () {

    // memorizzo l'ultimo id selezionato
    var lastId, 
        // recupero l'elemento menu attraverso l'id
        menu = $('#menu'), 
        // altezza del menu
        altezza = menu.outerHeight() + 15, 
        // elementi ancora del menu
        elementi = menu.find('a'), 
        // ricerco i valori del collegamento per ciascuna ancora
        insegui = elementi.map(function () {
            var item = $($(this).attr('href'));
            if (item.length) { return item; }
        });

    // quando clicco su un elemento del menu
    elementi.click(function (e) {
        // ricerco il nuovo valore del collegamento
        var href = $(this).attr('href'), 
            // e il suo offset
            offsetTop = href === '#' ? 0 : $(href).offset().top - altezza + 1; 
        // animo il movimento sino all'offset dell'elemento
        $('html, body').stop().animate({ scrollTop: offsetTop }, 300);
        // evito il comportamento predefinito 
        e.preventDefault();
    });

    // funzione scorrimento
    $(window).scroll(function () {
        // primo elemento
        var fromTop = $(this).scrollTop() + altezza;
        // scorro
        var cur = insegui.map(function () {
            if ($(this).offset().top < fromTop)
                return this;
        });
        // ultimo elemento
        cur = cur[cur.length - 1];
        // id attuale
        var id = cur && cur.length ? cur[0].id : '';
        // se l'id è cambiato
        if (lastId !== id) {
            // aggiorno la variabile 
            lastId = id;
            // sposto la classe che indica la posizione dalla precedente alla attuale nel menu
            elementi.parent().removeClass('active').end().filter('[href=#' + id + ']').parent().addClass('active');
        }
    });
});

$(function () {
    
    // ricerco l'elemento pull
    var pull = $('#pull');
    // ricerco l'elemento menu
    menu = $('#menu');
    // ricerco l'altezza
    altezza = menu.height();

    // quando clicco su pull
    $(pull).on('click', function (e) {
        // evito l'azione predefinita
        e.preventDefault();
        // ma ativo lo slide
        menu.slideToggle();
    });
});

$(window).resize(function () {
    // ricerco la profondità della finestra attuale
    var w = $(window).width();
    // quando è maggiore di 320px e il pseudoselettore di menu mi indica che è nascosto
    if (w > 320 && menu.is(':hidden')) {
        // rimuovo l'attributo di stile per nasconderlo
        menu.removeAttr('style');
    }
});