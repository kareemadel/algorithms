const left = (i) => 2*i+1;
const right = (i) => 2*i+2;
const parent = (i) => Math.trunc((i - 1) / 2);

module.exports = {
    left,
    right,
    parent
};
