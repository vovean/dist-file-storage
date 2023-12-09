delete
from fileparts
where file in (select id from files where name = $1)