const fs   = require('fs').promises;
const path = require('path');

async function main() {
  const projectRoot = path.resolve(__dirname, '..');
  const imagesDir   = path.join(projectRoot, 'images');
  const outputFile  = path.join(imagesDir, 'images.json');

  const files = await fs.readdir(imagesDir);
  const imageFiles = files
    .filter(file => {
      const ext = path.extname(file).toLowerCase();
      return ['.jpg', '.jpeg', '.png', '.gif', '.webp'].includes(ext);
    })
    .sort((a, b) => {
      const aNum = parseInt(a.match(/\d+/), 10) || 0;
      const bNum = parseInt(b.match(/\d+/), 10) || 0;
      return aNum - bNum;
    });

  await fs.writeFile(outputFile, JSON.stringify(imageFiles, null, 2));
  console.log(`\t\u2705 Wrote ${imageFiles.length} entries to "${outputFile}"`);
}

main().catch(err => {
  console.error('\t\u274c Error generating images.json:', err);
  process.exit(1);
});
