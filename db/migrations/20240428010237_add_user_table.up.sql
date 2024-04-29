CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email VARCHAR NOT NULL UNIQUE,
    name VARCHAR(50) NOT NULL,
    password VARCHAR(255) NOT NULL,
    deleted_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP 
);

CREATE INDEX IF NOT EXISTS user_id ON users (id);


DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'cat_race_enum') THEN
        CREATE TYPE cat_race_enum AS ENUM ('Persian', 'Maine Coon', 'Siamese', 'Ragdoll', 'Bengal', 'Sphynx', 'British Shorthair', 'Abyssinian', 'Scottish Fold', 'Birman');
    END IF;
END $$;


DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'cat_sex_enum') THEN
        CREATE TYPE cat_sex_enum AS ENUM ('male', 'female');
    END IF;
END $$;


CREATE TABLE IF NOT EXISTS cats (
    id SERIAL PRIMARY KEY,  
    user_id INT NOT NULL REFERENCES users(id),
    name VARCHAR(30) NOT NULL,
    race cat_race_enum NOT NULL,
    sex cat_sex_enum NOT NULL,
    age_in_month INT NOT NULL,
    description VARCHAR(200) NOT NULL,
    has_matched BOOLEAN NOT NULL DEFAULT false,
    image_urls TEXT[] NOT NULL,
    deleted_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP 
);

CREATE INDEX IF NOT EXISTS cat_id ON cats (id);
CREATE INDEX IF NOT EXISTS cat_user_id ON cats (user_id);
CREATE INDEX IF NOT EXISTS cat_race ON cats (race);
CREATE INDEX IF NOT EXISTS cat_sex ON cats (sex);
CREATE INDEX IF NOT EXISTS cat_matched ON cats (has_matched);
CREATE INDEX IF NOT EXISTS cat_age ON cats (age_in_month);
CREATE INDEX IF NOT EXISTS cat_name ON cats (name);


CREATE TABLE IF NOT EXISTS matches (
    id SERIAL PRIMARY KEY,
    issued_by INT NOT NULL REFERENCES users(id),
    match_cat_id INT NOT NULL REFERENCES cats(id),
    user_cat_id INT NOT NULL REFERENCES cats(id),
    is_approved BOOLEAN NOT NULL DEFAULT false,
    message VARCHAR(120) NOT NULL,
    deleted_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP 
);

CREATE INDEX IF NOT EXISTS match_id ON matches (id);