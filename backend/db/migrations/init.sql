CREATE TABLE public.users
(
    id              serial                                   PRIMARY KEY,
    email           character varying                         NOT NULL,
    phone_number    character varying                         NOT NULL, 
    firstname    character varying   NOT NULL,
    lastname     character varying NOT NULL,
    password        character varying                         NOT NULL,
    role            text                                      NOT NULL,
    about_me  character varying DEFAULT ''::character varying NOT NULL,
    link_to_instagram character varying,
    rating          float                     DEFAULT 0     NOT NULL,
    shop_id integer,
    created         timestamp without time zone DEFAULT now() NOT NULL,
    updated         timestamp without time zone DEFAULT now() NOT NULL,
    CONSTRAINT fk_shops
        FOREIGN KEY (shop_id) REFERENCES shop(id)
);

INSERT INTO users (email, phone_number, username, firstname, lastname, password, role, about_me, link_to_instagram, rating) VALUES ("Mdidara@quirduck.khs", "77017345566", "Didara", "Mamyrova","$2a$10$1JwL9V/KDXep5cYqNGpwJ.g2yStQrvPkw5xrCbopsu2APSjGdpH7K", "client", "Hi, this is me!","https://instagram/azaza", 10);
INSERT INTO users (email, phone_number, username, firstname, lastname, password, role, about_me, link_to_instagram, rating) VALUES ("KLeya@gmail.com", "77017341111", "Leya", "Kim","$2a$10$AmWDAtt4TirVpFytlOdesuRl2cyF7z4X3sWQetBoa/yYQM/Nlu7Ei", "shop", "Hi, this is me!","https://instagram/azaza", 10);

CREATE TABLE address
(
    id serial PRIMARY KEY,
    country character varying                         NOT NULL,
    region character varying                         NOT NULL,
    city character varying                         NOT NULL,
    street_name character varying                         NOT NULL,
    street_number character varying                         NOT NULL,
    building character varying                         NOT NULL,
    postal_code character varying                         NOT NULL,
    flat_number character varying,
    floor character varying
);

CREATE TABLE category
(
    id serial PRIMARY KEY,
    name character varying NOT NULL
);

CREATE TABLE shop
(
    id              serial                                   PRIMARY KEY,
    user_id integer NOT NULL,
    address_id integer NOT NULL,
    logo character varying DEFAULT NULL, 
    description  character varying DEFAULT ''::character varying NOT NULL,
    link_to_instagram character varying,
    CONSTRAINT fk_address
        FOREIGN KEY (address_id) REFERENCES address(id),
    CONSTRAINT fk_users
        FOREIGN KEY (user_id) REFERENCES users(id)
);

INSERT INTO shop ()
VALUES ("Mdidara@quirduck.khs", "77017345566", "Didara", "Mamyrova","$2a$10$1JwL9V/KDXep5cYqNGpwJ.g2yStQrvPkw5xrCbopsu2APSjGdpH7K", "client", "Hi, this is me!","https://instagram/azaza", 10);

CREATE TABLE product
(
    id serial PRIMARY KEY,
    shop_id integer NOT NULL,
    name    character varying   NOT NULL,
    description    character varying   NOT NULL,
    sizes text[], 
    colors integer array,
    weight float, -- in kg
    price float NOT NULL, -- in dollars?
    rating float -- average from comments
    category_id int NOT NULL,
    created         timestamp without time zone DEFAULT now() NOT NULL,
    updated         timestamp without time zone DEFAULT now() NOT NULL
    CONSTRAINT fk_shop
        FOREIGN KEY (shop_id) REFERENCES shop(id),
    CONSTRAINT fk_category
        FOREIGN KEY (category_id) REFERENCES category(id),
);

CREATE TABLE public.comment
(
    id         integer                                   NOT NULL,
    text       character varying                         NOT NULL,
    score integer NOT NULL, -- from 1 to 5
    created    timestamp without time zone DEFAULT now() NOT NULL,
    updated    timestamp without time zone DEFAULT now() NOT NULL,
    user_id   integer,
    product_id integer,
    CONSTRAINT fk_product
        FOREIGN KEY (product_id) REFERENCES product(id),
    CONSTRAINT fk_users_1
        FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE public.orders
(
    id serial PRIMARY KEY,
    status integer not null,
    client_id integer,
    shop_id integer,
    product_id integer,
    created_at    timestamp without time zone DEFAULT now() NOT NULL,
    payed_at timestamp without time zone,
    packed_at timestamp without time zone,
    delivered_at timestamp without time zone,
    cancel_reason text,
    CONSTRAINT fk_client_id
        FOREIGN KEY (client_id) REFERENCES users(id),
    CONSTRAINT fk_shop_id
        FOREIGN KEY (shop_id) REFERENCES shop(id),
    CONSTRAINT fk_product_id
        FOREIGN KEY (product_id) REFERENCES product(id),
)
