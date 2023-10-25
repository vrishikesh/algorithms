/**
 * @param {number[]} height
 * @return {number}
 */
var maxArea = function (height) {
  let max = 0;
  let p1 = 0;
  let p2 = height.length-1;
  while (p1 < p2) {
    const h = Math.min(height[p1], height[p2]);
    const w = p2 - p1;
    const area = h * w;
    max = Math.max(max, area);
    if (height[p1] <= height[p2]) {
      p1++;
    } else {
      p2--;
    }
  }
  return max
};

console.log(maxArea([1, 8, 6, 2, 5, 4, 8, 3, 7]));
