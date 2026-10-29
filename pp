// Debounce: run after the user stops doing something
function debounce(fn, delay = 300) {
  let timer;
  return function (...args) {
    clearTimeout(timer);
    timer = setTimeout(() => fn.apply(this, args), delay);
  };
}

// Throttle: run at most once every X ms
function throttle(fn, limit = 300) {
  let lastCall = 0;
  return function (...args) {
    const now = Date.now();
    if (now - lastCall >= limit) {
      lastCall = now;
      fn.apply(this, args);
    }
  };
}

// Usage
const onSearch = debounce((q) => console.log("Searching:", q), 500);
const onScroll = throttle(() => console.log("Scrolled"), 200);
