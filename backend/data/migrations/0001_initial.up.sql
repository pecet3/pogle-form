
create table if not exists persons (
    id integer primary key autoincrement,
    email text unique,
    full_name text not null,
    created_at timestamp default current_timestamp
);