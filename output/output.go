package output

type OutputProvider interface {
	FlushData(customerMap map[string]int64, filePath string) error
}
