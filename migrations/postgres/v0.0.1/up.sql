CREATE TABLE IF NOT EXISTS hash (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    password text,
    chg_0 text,
    chg_1 text,
    chg_2 text,
    chg_3 text,
    chg_4 text,
    chg_5 text,
    chg_6 text,
    chg_7 text
);

CREATE TABLE IF NOT EXISTS challenge (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    chg_0 int,
    chg_1 int,
    chg_2 int,
    chg_3 int,
    chg_4 int,
    chg_5 int,
    chg_6 int,
    chg_7 int
);

INSERT INTO challenge VALUES (DEFAULT, 10, 20, 30, 40, 50, 60, 70, 80);
