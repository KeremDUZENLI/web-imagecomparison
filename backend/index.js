import express  from 'express';
import cors     from 'cors';
import pkg      from 'pg';
import dotenv   from 'dotenv';

dotenv.config();

const { Pool } = pkg;

const pool = new Pool({
  host:     process.env.DB_HOST,
  user:     process.env.DB_USER,
  password: process.env.DB_PASSWORD,
  database: process.env.DB_NAME,
  port:     process.env.DB_PORT,
});

pool.connect()
  .then(() => console.log('✅ Connected to PostgreSQL'))
  .catch(err => console.error('❌ Connection error:', err.stack));

const app = express();
app.use(cors());
app.use(express.json());

const initDB = async () => {
  const createTableQuery = `
    CREATE TABLE IF NOT EXISTS votes (
      id SERIAL PRIMARY KEY,
      userName TEXT,
      imageA TEXT,
      imageB TEXT,
      imageWinner TEXT,
      imageLoser TEXT,
      eloWinnerPrevious INTEGER,
      eloWinnerNew INTEGER,
      eloLoserPrevious INTEGER,
      eloLoserNew INTEGER,
      created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );
  `;
  try {
    await pool.query(createTableQuery);
    console.log('✅ Votes table is ready');
  } catch (err) {
    console.error('❌ Failed to initialize database:', err);
  }
};

app.post('/api/votes', async (req, res) => {
  const vote = req.body;

  const query = `
    INSERT INTO votes (
      userName, imageA, imageB, imageWinner, imageLoser, 
      eloWinnerPrevious, eloWinnerNew, eloLoserPrevious, eloLoserNew
    ) 
    VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) 
    RETURNING *`;

  const values = [
    vote.userName, vote.imageA, vote.imageB, vote.imageWinner, vote.imageLoser, 
    vote.eloWinnerPrevious, vote.eloWinnerNew, vote.eloLoserPrevious, vote.eloLoserNew
  ];

  try {
    const { rows } = await pool.query(query, values);
    res.json(rows[0]);
  } catch (err) {
    console.error('❌ Database insert error:', err);
    res.status(500).send('Database error');
  }
});

const PORT = 3000;
app.listen(PORT, async () => {
  await initDB();
  console.log(`✅ API is running at http://localhost:${PORT}`);
});
