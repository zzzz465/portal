package types

type Record struct {
	Key      string
	Name     string
	Metadata RecordMetadata
}

type RecordMetadata struct {
	DataSource string
	Tags       map[string]string
}

func NewRecordMetadata() *RecordMetadata {
	return &RecordMetadata{
		DataSource: "",
		Tags:       make(map[string]string),
	}
}
