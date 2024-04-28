PRAGMA foreign_keys = ON;

CREATE TABLE IF NOT EXISTS deities (
    deity_id TEXT NOT NULL PRIMARY KEY,
    name TEXT COLLATE NOCASE,
    description TEXT COLLATE NOCASE,
    gender TEXT
);

CREATE TABLE IF NOT EXISTS sports (
    sport_id TEXT NOT NULL PRIMARY KEY,
    name TEXT COLLATE NOCASE
);

CREATE TABLE IF NOT EXISTS olympian_influence (
    deity_id TEXT NOT NULL,
    sport_id TEXT NOT NULL,
    influence REAL,
    FOREIGN KEY(deity_id) REFERENCES deities(deity_id)
    FOREIGN KEY(sport_id) REFERENCES sports(sport_id)
);


INSERT OR IGNORE INTO sports (sport_id, name) VALUES ('pugilato', 'Pugilato');
INSERT OR IGNORE INTO sports (sport_id, name) VALUES ('corsa', 'Corsa');
INSERT OR IGNORE INTO sports (sport_id, name) VALUES ('lotta', 'Lotta');
INSERT OR IGNORE INTO sports (sport_id, name) VALUES ('danza', 'Danza');
INSERT OR IGNORE INTO sports (sport_id, name) VALUES ('lancio-del-disco', 'Lancio del disco');
INSERT OR IGNORE INTO sports (sport_id, name) VALUES ('salto-in-lungo', 'Salto in lungo');
INSERT OR IGNORE INTO sports (sport_id, name) VALUES ('lancio-del-giavellotto', 'Lancio del giavellotto');


INSERT OR IGNORE INTO
    deities (deity_id, name, description, gender)
VALUES (
    'atena',
    'Atena',
    '',
    'F'
);
INSERT OR IGNORE INTO
    deities (deity_id, name, description, gender)
VALUES (
    'eros',
    'Eros',
    '',
    'M'
);
INSERT OR IGNORE INTO
    deities (deity_id, name, description, gender)
VALUES (
    'ermes',
    'Ermes',
    '',
    'M'
);
INSERT OR IGNORE INTO
    deities (deity_id, name, description, gender)
VALUES (
    'eris',
    'Eris',
    '',
    'F'
);


INSERT OR IGNORE INTO
    olympian_influence (deity_id, sport_id, influence)
VALUES ('atena', 'pugilato', 1.5);
INSERT OR IGNORE INTO
    olympian_influence (deity_id, sport_id, influence)
VALUES ('atena', 'corsa', 0.3);
INSERT OR IGNORE INTO
    olympian_influence (deity_id, sport_id, influence)
VALUES ('atena', 'lotta', 2.5);
INSERT OR IGNORE INTO
    olympian_influence (deity_id, sport_id, influence)
VALUES ('atena', 'danza', 0.1);
INSERT OR IGNORE INTO
    olympian_influence (deity_id, sport_id, influence)
VALUES ('atena', 'lancio-del-disco', 1.2);
INSERT OR IGNORE INTO
    olympian_influence (deity_id, sport_id, influence)
VALUES ('atena', 'salto-in-lungo', 0.8);
INSERT OR IGNORE INTO
    olympian_influence (deity_id, sport_id, influence)
VALUES ('atena', 'lancio-del-giavellotto', 2);

INSERT OR IGNORE INTO
    olympian_influence (deity_id, sport_id, influence)
VALUES ('eros', 'pugilato', 0.3);
INSERT OR IGNORE INTO
    olympian_influence (deity_id, sport_id, influence)
VALUES ('eros', 'corsa', 0.8);
INSERT OR IGNORE INTO
    olympian_influence (deity_id, sport_id, influence)
VALUES ('eros', 'lotta', 0.5);
INSERT OR IGNORE INTO
    olympian_influence (deity_id, sport_id, influence)
VALUES ('eros', 'danza', 2.5);
INSERT OR IGNORE INTO
    olympian_influence (deity_id, sport_id, influence)
VALUES ('eros', 'lancio-del-disco', 0.1);
INSERT OR IGNORE INTO
    olympian_influence (deity_id, sport_id, influence)
VALUES ('eros', 'salto-in-lungo', 0.5);
INSERT OR IGNORE INTO
    olympian_influence (deity_id, sport_id, influence)
VALUES ('eros', 'lancio-del-giavellotto', 0.1);

INSERT OR IGNORE INTO
    olympian_influence (deity_id, sport_id, influence)
VALUES ('ermes', 'pugilato', 0.3);
INSERT OR IGNORE INTO
    olympian_influence (deity_id, sport_id, influence)
VALUES ('ermes', 'corsa', 2.5);
INSERT OR IGNORE INTO
    olympian_influence (deity_id, sport_id, influence)
VALUES ('ermes', 'lotta', 0.2);
INSERT OR IGNORE INTO
    olympian_influence (deity_id, sport_id, influence)
VALUES ('ermes', 'danza', 1.2);
INSERT OR IGNORE INTO
    olympian_influence (deity_id, sport_id, influence)
VALUES ('ermes', 'lancio-del-disco', 0.1);
INSERT OR IGNORE INTO
    olympian_influence (deity_id, sport_id, influence)
VALUES ('ermes', 'salto-in-lungo', 1.5);
INSERT OR IGNORE INTO
    olympian_influence (deity_id, sport_id, influence)
VALUES ('ermes', 'lancio-del-giavellotto', 0.1);

-- No values for Eris since we random compute them in 0.1 < x < 0.5
-- since she make you always disappointed.
