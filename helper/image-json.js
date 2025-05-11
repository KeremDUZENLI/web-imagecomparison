const fs   = require('fs').promises;
const path = require('path');

const imagesDir  = path.join(__dirname, '../frontend/images');
const outputFile = path.join(__dirname, '../frontend/images.json');

async function createImageJson() {
  const files      = await fs.readdir(imagesDir);
  const imageFiles = files.filter(file => {const ext = path.extname(file).toLowerCase();
    return ['.jpg', '.jpeg', '.png', '.gif', '.webp'].includes(ext);
  })    
  .sort((a, b) => {
    const aNumber = parseInt(a.match(/\d+/));
    const bNumber = parseInt(b.match(/\d+/));
    return aNumber - bNumber;
  });

  await fs.writeFile(outputFile, JSON.stringify(imageFiles, null, 2));
}

createImageJson();
