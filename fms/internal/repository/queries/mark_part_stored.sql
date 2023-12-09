update fileparts fp
set is_stored = true
from files f
where f.name = $1    -- ищем файл, откуда кусок
  and f.id = fp.file -- берем только части этого файла
  and fp.part_no = $2 -- только нужную часть