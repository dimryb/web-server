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

create table photos(
    id serial primary key,
    user_id int not null references users(id) on delete cascade,
    filename varchar(1024),
    width int not null,
    height int not null,
    created_at timestamp with time zone default current_timestamp
);

insert into photos(user_id, filename, width, height) values
    (1, 'cat.jpg', 1920, 1080),
    (1, 'dog.jpg', 1924, 1080),
    (2, 'pine.jpg', 120, 720),
    (2, 'banana.jpg', 120, 720),
    (2, 'tomato.jpg', 120, 720),
    (3, 'parrot.jpg', 800, 600),
    (3, 'fish.jpg', 800, 600)
;