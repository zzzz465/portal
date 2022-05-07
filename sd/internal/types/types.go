package types

type Record struct {
	IP       string
	Host     string
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
