const fs = require('fs');

fs.readFile('demofile', (err, data) => {
    console.log(data.toString());
});