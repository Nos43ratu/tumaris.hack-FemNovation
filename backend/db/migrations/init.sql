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

INSERT INTO address (country, region, city, street_name, street_number, building, postal_code, flat_number, floor) VALUES ('Kazakhstan', 'Astana', 'Astana', 'Mangilik-el', '1', '1/2', '010000', '12', '12');


CREATE TABLE shop
(
    id              serial                                   PRIMARY KEY,
    address_id integer NOT NULL,
    logo character varying DEFAULT NULL, 
    description  character varying DEFAULT ''::character varying NOT NULL,
    link_to_instagram character varying,
    CONSTRAINT fk_address
        FOREIGN KEY (address_id) REFERENCES address(id)
);

INSERT INTO shop (address_id, logo, description, link_to_instagram) VALUES (1, NULL, 'Super shop', 'https://instagram/azaza');

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

INSERT INTO users (email, phone_number, firstname, lastname, password, role, about_me, link_to_instagram, rating) VALUES ('Mdidara@quirduck.khs', '77017345566', 'Didara', 'Mamyrova','$2a$10$1JwL9V/KDXep5cYqNGpwJ.g2yStQrvPkw5xrCbopsu2APSjGdpH7K', 'client', 'Hi, this is me!','https://instagram/azaza', 10);
INSERT INTO users (email, phone_number, firstname, lastname, password, role, about_me, link_to_instagram, rating, shop_id) VALUES ('KLeya@gmail.com', '77017341111', 'Leya', 'Kim','$2a$10$AmWDAtt4TirVpFytlOdesuRl2cyF7z4X3sWQetBoa/yYQM/Nlu7Ei', 'shop', 'Hi, this is me!','https://instagram/azaza', 10, 1);


CREATE TABLE category
(
    id serial PRIMARY KEY,
    name character varying NOT NULL
);

INSERT INTO category (name) values ('soap');
INSERT INTO category (name) values ('candle');
INSERT INTO category (name) values ('hat');
INSERT INTO category (name) values ('bath bomb');


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
    rating float, -- average from comments
    category_id integer NOT NULL,
    created         timestamp without time zone DEFAULT now() NOT NULL,
    updated         timestamp without time zone DEFAULT now() NOT NULL,
    CONSTRAINT fk_shop
        FOREIGN KEY (shop_id) REFERENCES shop(id),
    CONSTRAINT fk_category
        FOREIGN KEY (category_id) REFERENCES category(id)
);

INSERT INTO product (shop_id, name, description, sizes, colors, weight, price, category_id) VALUES(1, 'soap', 'the best soap', '{"small", "medium", "large"}', '{0, 1, 2}', 1.32, 20.25, 1);

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
        FOREIGN KEY (product_id) REFERENCES product(id)
);

