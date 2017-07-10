
// Change periodically the images in the slideshow
$(function() {
  // If there's no slideshow, it's not a page where images should rotate
  if ($('ul.slideshow').length > 0) {
    window.setInterval(function slide() {
      var actual = $('ul.slideshow').children('li.show').first(),
          next = actual.next();
          next = (next.length) ? next : actual.prevAll().last() ;
      next.addClass('show');
      actual.removeClass('show');
    }, 5000);
  }
});
