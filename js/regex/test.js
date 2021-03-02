const fs = require('fs');

const regex = /(<AdaptationSet mimeType=\"audio\/mp4\".*>((?!Ada).|\n)*<Role.*value=\"main\".*>((?!Ada).|\n)*<\/AdaptationSet>)/;

const regexAlternate = /(<AdaptationSet mimeType=\"audio\/mp4\".*>((?!Ada).|\n)*<Role.*value=\"dub\".*>((?!Ada).|\n)*<\/AdaptationSet>)/;

fs.readFile('/Users/sapooman/Downloads/sample_manifest.mpd', (err, data) => {
    const manifest = data.toString();

    console.log(regexAlternate.test(manifest));


    if (manifest.match(regex)) {
        console.log(manifest.match(regex)[0]);
    } else {
        console.log('no match found..');
    }
    
});

