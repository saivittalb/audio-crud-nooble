CREATE TABLE shorts (
    audioid SERIAL PRIMARY KEY,
    title TEXT,
    description TEXT,
    category TEXT,
    audiofile BYTEA,
    creatorname TEXT,
    creatoremail TEXT,
    date DATE
);