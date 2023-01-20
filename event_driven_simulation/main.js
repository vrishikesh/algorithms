/** @type {HTMLCanvasElement} */
const canvas = document.getElementById("canvas");

// canvas.width = window.innerWidth;
// canvas.height = window.innerHeight;

/** @type {CanvasRenderingContext2D} */
const ctx = canvas.getContext("2d");

let then = null;
const fps = 1000 / 60;
const speed = 4;

/** @type {Array<Ball>} */
const balls = [];
for (let i = 0; i < 10; i++) {
  balls.push(new Ball(canvas, ctx));
}

function loop(now) {
  if (then === null) {
    then = now;
  }

  const delta = now - then;

  if (delta < fps) {
    window.requestAnimationFrame(loop);
    return;
  }

  then = now;

  ctx.clearRect(0, 0, canvas.width, canvas.height);
  for (let i = 0; i < balls.length; i++) {
    balls[i].move(speed);
    balls[i].draw();
  }

  window.requestAnimationFrame(loop);
}

window.requestAnimationFrame(loop);
