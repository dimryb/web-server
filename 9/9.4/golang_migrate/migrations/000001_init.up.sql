create table users(
    id serial primary key,
    name varchar(50) not null,
    email varchar(50) not null
);

insert into users(name, email) values
    ('ivan', 'ivan@mail.ru'),
    ('andrey', 'andrey@gmail.com'),
    ('john', 'john@gmail.com'),
    ('slava', 'slava@exemple.com'),
    ('alex', 'alex@testserver')
;