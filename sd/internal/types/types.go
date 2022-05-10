package types

type Record struct {
    Name     string         `yaml:"name"`
    Metadata RecordMetadata `yaml:"metadata"`
}

func NewRecord() Record {
    return Record{Metadata: RecordMetadata{Tags: map[string]string{}}}
}

type RecordMetadata struct {
    DataSource string            `yaml:"dataSource"`
    Tags       map[string]string `yaml:"tags"`
}

func NewRecordMetadata() *RecordMetadata {
    return &RecordMetadata{
        DataSource: "",
        Tags:       make(map[string]string),
    }
}
