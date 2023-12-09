insert into storages(addr, initial_space_bytes)
values ($1, $2)
returning id