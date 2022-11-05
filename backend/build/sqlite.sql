CREATE TABLE IF NOT EXISTS users (
    id integer PRIMARY KEY AUTOINCREMENT,
    email text UNIQUE NOT NULL,
    password text NOT NULL,
    role text NOT NULL
);

CREATE TABLE IF NOT EXISTS tokens (
    id integer PRIMARY KEY AUTOINCREMENT,
    user_id integer NOT NULL,
    refresh_token text NOT NULL,
    access_token text NOT NULL,
    refresh_expire bigint NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- CREATE TABLE IF NOT EXISTS panel (
--     id integer PRIMARY KEY AUTOINCREMENT,
--     device text NOT NULL,
--     status int NOT NULL
-- );

-- INSERT INTO panel (device, status) VALUES ('lights', 1);
-- INSERT INTO panel (device, status) VALUES ('radar', 1);

INSERT INTO users (email, password, role) VALUES ("Mdidara@quirduck.khs", "$2a$10$1JwL9V/KDXep5cYqNGpwJ.g2yStQrvPkw5xrCbopsu2APSjGdpH7K", "client");
INSERT INTO users (email, password, role) VALUES ("KLeya@gmail.com", "$2a$10$AmWDAtt4TirVpFytlOdesuRl2cyF7z4X3sWQetBoa/yYQM/Nlu7Ei", "shop");