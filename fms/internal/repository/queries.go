package repository

import _ "embed"

var (
	//go:embed queries/insert_storage.sql
	insertStorageQuery string

	//go:embed queries/get_storages.sql
	selectStoragesQuery string

	//go:embed queries/insert_file.sql
	insertFileQuery string

	//go:embed queries/get_fileparts.sql
	getFilePartsQuery string

	//go:embed queries/mark_part_stored.sql
	markPartStoredQuery string

	//go:embed queries/delete_fileparts.sql
	deleteFilePartsQuery string

	//go:embed queries/delete_file.sql
	deleteFileQuery string
)
