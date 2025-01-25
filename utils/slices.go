package utils

// SliceToMatrix 将slice按照N个一组 , 拆分为[][]T
func SliceToMatrix[T any](list []T, batchSize int) [][]T {
	var matrix [][]T
	length := len(list)
	for i := 0; i < length; i += batchSize {
		end := i + batchSize
		if end > length {
			end = length
		}
		matrix = append(matrix, list[i:end])
	}
	return matrix
}
