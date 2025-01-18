create table users (
    id serial primary key,
    name text not null,
    surname text not null,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now(),
    deleted_at timestamp with time zone
);

insert into users(name, surname) values ('John', 'Doe');
insert into users(name, surname) values ('Jane', 'Doe');
insert into users(name, surname) values ('Alice', 'Smith');