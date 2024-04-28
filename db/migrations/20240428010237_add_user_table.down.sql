DROP TABLE IF EXISTS matches;
DROP TABLE IF EXISTS cats;
DROP TABLE IF EXISTS users;

-- Drop enum types if they exist
DO $$
BEGIN
    IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'cat_race_enum') THEN
        DROP TYPE cat_race_enum;
    END IF;
END $$;

DO $$
BEGIN
    IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'cat_sex_enum') THEN
        DROP TYPE cat_sex_enum;
    END IF;
END $$;
