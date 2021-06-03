const o = {
    A: '1',
    B: '2',
    C: '3',
    D: '4'
};

for(let i in o) {
    console.log(o[i]);
}

const moment = require('moment');

const date = '2021-05-21';

let start = moment(`${date}T00:00:00+09:00`);

let end = moment(start).add(1, 'days');

console.log(start.toDate());
console.log(end.toDate());

const arr = ['d', 'c', 'b']

console.log(arr.sort());

console.log(moment(date).isValid());

console.log(end.add(-7, 'days').format('YYYY-MM-DD'));