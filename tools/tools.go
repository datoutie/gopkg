package tools

//查看值是否在数组中
func InArray(needle interface{}, haystack []interface{}) bool {
	for _, item := range haystack {
		if item == needle {
			return true
		}
	}
	return false
}

// 使用ArrayColumn提取列
func ArrayColumn(data []map[string]interface{}, columnName string) []interface{} {
	var columnValues []interface{}
	for _, row := range data {
		if value, found := row[columnName]; found {
			columnValues = append(columnValues, value)
		} else {
			columnValues = append(columnValues, nil) // 或者根据需要设置默认值
		}
	}
	return columnValues
}

// 使用ArrayColumnToMap列作为键，列作为值
func ArrayColumnToMap(data []map[string]interface{}, keyColumn, valueColumn string) map[interface{}]interface{} {
	result := make(map[interface{}]interface{})
	for _, row := range data {
		key, keyExists := row[keyColumn]
		value, valueExists := row[valueColumn]

		if keyExists && valueExists {
			result[key] = value
		}
	}
	return result
}
