create schema if not exists public;

create table if not exists public.item
(
    id          serial not null primary key,
    item_id     bigint,
    name        varchar(10),
    seller_id   bigint,
    description text
);

create table if not exists public.sku
(
    id          serial not null primary key,
    sku_id      bigint,
    item_id     bigint,
    name        varchar(32),
    description text
);

insert into public.item(item_id, name, seller_id, description)
values (1009276444, 'mbp-16-pro', 100600, 'macbook pro 16 inch'),
       (1673422, 'ipad-1', 10033, 'iPad pro 1');

insert into public.sku(sku_id, item_id, name, description)
values (115597, 1009276444, 'grey macbook pro m1', 'Super power macbook with 16 inch display'),
       (115598, 1009276444, 'grey macbook pro m2', 'Small power macbook with 13 inch display'),
       (2255, 1673422, 'ipad 14 inch', 'luxury ipad for best issues'),
       (2256, 1673422, 'ipad 11 inch', 'small ipad for business');
