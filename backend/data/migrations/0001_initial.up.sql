create table if not exists courses (
    id integer primary key autoincrement,
    name text not null,
    max_persons integer not null,
    created_at timestamp default current_timestamp

);


create table if not exists chosen_reserved_courses (
    id integer primary key autoincrement,
    person_id integer not null,
    course_id integer not null,
    foreign key (course_id) references courses(id),
    foreign key (person_id) references persons(id)
);

create table if not exists persons (
    id integer primary key autoincrement,
    email text unique,
    full_name text not null,
    course_id integer not null,
    created_at timestamp default current_timestamp,
    foreign key (course_id) references courses(id)
);

insert into courses (name, max_persons) values ('test', 2);