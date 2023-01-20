class Ball {
  /**
   * @param {HTMLCanvasElement} canvas
   * @param {CanvasRenderingContext2D} ctx
   */
  constructor(canvas, ctx) {
    /** @type {HTMLCanvasElement} */
    this.canvas = canvas;

    /** @type {CanvasRenderingContext2D} */
    this.ctx = ctx;

    /** @type {number} */
    this.radius = 10;

    /** @type {number} */
    this.vx = Math.random();

    /** @type {number} */
    this.vy = Math.random();

    /** @type {number} */
    this.x = this.getRandomPosition(canvas.width);

    /** @type {number} */
    this.y = this.getRandomPosition(canvas.height);
  }

  /**
   * @param {number} size
   */
  getRandomPosition(size) {
    let pos = Math.random() * size;
    pos = Math.max(pos, this.radius);
    pos = Math.min(pos, size - this.radius);
    return pos;
  }

  draw() {
    this.ctx.beginPath();
    this.ctx.arc(this.x, this.y, this.radius, 0, 2 * Math.PI);
    this.ctx.closePath();
    this.ctx.fill();
  }

  /**
   * @param {number} dt
   */
  move(dt) {
    // Check if ball hits wall
    if (
      this.x + this.vx * dt < this.radius ||
      this.x + this.vx * dt > this.canvas.width - this.radius
    ) {
      this.vx *= -1;
    } else if (
      this.y + this.vx * dt < this.radius ||
      this.y + this.vx * dt > this.canvas.height - this.radius
    ) {
      this.vy *= -1;
    }

    // Update position
    this.x += this.vx * dt;
    this.y += this.vy * dt;
  }
}
