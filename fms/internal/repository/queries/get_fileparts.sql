select fp.part_no as part_id,
       s.addr     as storage_addr,
       fp.size,
       fp.path_in_storage,
       fp.is_stored
from fileparts fp
         join public.files f on f.id = fp.file
         join public.storages s on s.id = fp.storage
where f.name = $1