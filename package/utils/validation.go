package utils

import "sistem_kasir_dan_database_toko/package/constant"

func ValidateFile(extension string) bool {
	return constant.ALLOWED_EXTENSIONS[extension]
}
