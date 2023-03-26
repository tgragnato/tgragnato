function f() {
    var some = [];
    while(some.length < 1e6) {
        some.push(some.length);
    }
    function unused() { some; } // causes massive memory leak
    return function() {};
}
  
var a = [];
var interval = setInterval(function() {
    var len = a.push(f());
    document.getElementById('count').innerHTML = len + ' / 500';
    if(len >= 500) {
        clearInterval(interval);
    }
}, 10);