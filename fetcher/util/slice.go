package util

func ChunkSlice[T any](slice []T, chunkSize int) [][]T {
	var chunks [][]T
	for {
		if len(slice) == 0 {
			break
		}
		// necessary check to avoid slicing beyond
		// slice capacity
		if len(slice) < chunkSize {
			chunkSize = len(slice)
		}
		chunks = append(chunks, slice[0:chunkSize])
		slice = slice[chunkSize:]
	}
	return chunks
}

func Slice(from, to uint64) []uint64 {
	var chunks []uint64
	var i uint64 = 0
	for ; i <= to-from; i++ {
		chunks = append(chunks, from+i)
	}
	return chunks
}
