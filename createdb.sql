CREATE DATABASE "L0WB" WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'Russian_Russia.1251';
ALTER DATABASE "L0WB" OWNER TO postgres;
\connect "L0WB"

CREATE TABLE "Deliverys" (
    id integer NOT NULL,
    name text NOT NULL,
    phone text,
    zip text NOT NULL,
    city text NOT NULL,
    address text NOT NULL,
    region text,
    email text
);

CREATE TABLE "Items" (
    id integer NOT NULL,
    chrt_id bigint,
    track_number text,
    price integer NOT NULL,
    rid text NOT NULL,
    name text NOT NULL,
    sale integer NOT NULL,
    size text NOT NULL,
    total_price integer NOT NULL,
    nm_id integer NOT NULL,
    brand text,
    status integer NOT NULL
);

CREATE TABLE "Orders" (
    id integer NOT NULL,
    order_uid text NOT NULL,
    track_number text,
    entry text,
    delivery_id integer NOT NULL,
    payment_id integer NOT NULL,
    locale text,
    internal_signature text,
    customer_id text NOT NULL,
    delivery_service text,
    shardkey text,
    sm_id integer,
    date_created text DEFAULT CAST(NOW() AS text) NOT NULL,
    oof_shard text
);

CREATE TABLE "Payments" (
    id integer NOT NULL,
    transaction text NOT NULL,
    request_id text,
    currency text NOT NULL,
    provider text NOT NULL,
    amount integer,
    payment_dt bigint NOT NULL,
    bank text,
    delivery_cost integer DEFAULT 0 NOT NULL,
    goods_total integer,
    custom_fee integer
);

INSERT INTO "Deliverys" (id, name, phone, zip, city, address, region, email)
VALUES (1, 'Test Testov', '+9720000000', '2639809', 'Kiryat Mozkin', 'Ploshad Mira 15', 'Kraiot', 'test@gmail.com');


INSERT INTO "Items" (id, chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status)
VALUES (1, 9934930, 'WBILMTESTTRACK', 453, 'ab4219087a764ae0btest', 'Mascaras', 30, '0', 317, 2389212, 'Vivienne Sabo', 202);


INSERT INTO "Orders" (order_uid, track_number, entry, delivery_id, payment_id, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard, id)
VALUES ('b563feb7b2b84b6test', 'WBILMTESTTRACK', 'WBIL', 1, 1, 'en', '', 'test', 'meest', '9', 99, '2022-07-22 17:42:39.5555+03:0', '1', 1);


INSERT INTO "Payments" (id, transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee)
VALUES (1, 'b563feb7b2b84b6test', '', 'USD', 'wbpay', 1817, 1637907727, 'alpha', 1500, 317, 0);