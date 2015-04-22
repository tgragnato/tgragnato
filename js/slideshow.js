$(document).ready(function() {    
  
  // 4 secondi per ogni immagine
  cambia(4000);

});

function cambia(speed) {

  // aggiungo un elemento alla lista per la descriione
  $('ul.slideshow').append('<li id="slideshow-caption" class="caption"><div class="slideshow-caption-container"><p></p></div></li>');

  // imposto l'opacità a 0 per ogni immagine
  $('ul.slideshow li').css({opacity: 0.0});
  
  // visualizzo la prima immagine
  $('ul.slideshow li:first').css({opacity: 1.0}).addClass('show');
  
  // visualizzo la didascalia della prima immagine
  $('#slideshow-caption p').html($('ul.slideshow li.show').find('img').attr('alt'));
    
  // visualizzo la didascalia
  $('#slideshow-caption').css({opacity: 0.6, bottom:0});
  
  // chiamo gallery 
  var timer = setInterval('gallery()',speed);
  
  // interrompo lo scorrimento quando il puntatore è su una immagine
  $('ul.slideshow').hover(
    function () {
      clearInterval(timer); 
    },  
    function () {
      timer = setInterval('gallery()',speed);     
    }
  );  
}

function gallery() {

  // fallback per quando l'utente non importa show sulla prima immagine
  var current = ($('ul.slideshow li.show')?  $('ul.slideshow li.show') : $('#ul.slideshow li:first'));

  // cerco di evitare i problemi sulla velocità di scorrimento
  if(current.queue('fx').length == 0) {

    // passo alla prossima immagine
    var next = ((current.next().length) ? ((current.next().attr('id') == 'slideshow-caption')? $('ul.slideshow li:first') :current.next()) : $('ul.slideshow li:first'));
      
    // ricerco la prossima didascalia
    var desc = next.find('img').attr('alt');  
  
    // imposto l'effetto di fade-in effect per la prossima immagine
    next.css({opacity: 0.0}).addClass('show').animate({opacity: 1.0}, 1000);
    
    // nascondo la didascalia, poi la cambio e la visualizzo
    $('#slideshow-caption').slideToggle(300, function () { 
      $('#slideshow-caption p').html(desc); 
      $('#slideshow-caption').slideToggle(500); 
    });   
  
    // nascondo l'immagine corrente
    current.animate({opacity: 0.0}, 1000).removeClass('show');

  }
}