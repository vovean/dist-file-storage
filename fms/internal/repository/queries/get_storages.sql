select s.id,
       s.addr,
       s.initial_space_bytes - coalesce(fp.occupied_space, 0) as free_space_bytes
from storages s
         left join (select storage, sum(size) as occupied_space
               from fileparts
               group by storage) fp on s.id = fp.storage