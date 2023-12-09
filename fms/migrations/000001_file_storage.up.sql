create table if not exists storages
(
    id                  serial primary key,
    addr                text   not null,
    initial_space_bytes bigint not null
);

create table if not exists files
(
    id   bigserial primary key,
    name varchar(200) not null unique,
    size bigint       not null
);

create table if not exists fileparts
(
    part_no         int2   not null,
    file            bigint not null references files (id),
    storage         int    not null references storages (id),
    size            bigint not null,
    path_in_storage text   not null,
    is_stored       bool   not null default false,

    primary key (file, part_no)
);

-- используем hash index по полю fileparts.storage для запросов group by storage. Дешевле, чем b-tree
create index if not exists ix_fp_storage on fileparts using hash (storage);

-- индекс по полюс fileparts.file для быстрого поиска всех частей файла
create index if not exists ix_fp_file on fileparts using hash (file);