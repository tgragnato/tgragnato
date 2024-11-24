let currentSlide = 0;
const slides = document.querySelectorAll('.slideshow-item');
function showSlide(index) {
  slides.forEach((slide, i) => {
    slide.classList.toggle('is-active', i === index);
  });
}
function prevSlide() {
  currentSlide = (currentSlide > 0) ? currentSlide - 1 : slides.length - 1;
  showSlide(currentSlide);
}
function nextSlide() {
  currentSlide = (currentSlide < slides.length - 1) ? currentSlide + 1 : 0;
  showSlide(currentSlide);
}
document.addEventListener('DOMContentLoaded', () => {
  showSlide(currentSlide);
});