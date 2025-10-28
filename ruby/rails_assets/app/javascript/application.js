// Configure your import map in config/importmap.rb. Read more: https://github.com/rails/importmap-rails
import "@hotwired/turbo-rails"
import "controllers"

window.addEventListener('load', function(event) {
  var div = document.createElement('div');
  div.innerText = "And hello from Javascript!";

  var body = document.querySelector('body');
  body.appendChild(div);
});
