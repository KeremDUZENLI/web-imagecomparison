const fs = require('fs');
const path = require('path');

const imagesDir = path.join(__dirname, '../images');
const outputFile = path.join(__dirname, '../images.json');

fs.readdir(imagesDir, (err, files) => {
  if (err) throw err;

  const imageFiles = files
    .filter(file => {
      const ext = path.extname(file).toLowerCase();
      return ['.jpg', '.jpeg', '.png', '.gif', '.webp'].includes(ext);
    })
    .sort((a, b) => {
      const aNum = parseInt(a.match(/\d+/));
      const bNum = parseInt(b.match(/\d+/));
      return aNum - bNum;
    });

  fs.writeFile(outputFile, JSON.stringify(imageFiles, null, 2), err => {
    if (err) throw err;
    console.log('images.json has been saved.');
  });
});
