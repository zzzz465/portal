package types

type Record struct {
    Name     string         `yaml:"name"`
    Tags     []string       `yaml:"tags"`
    Metadata RecordMetadata `yaml:"metadata"`
}

func NewRecord() Record {
    return Record{Metadata: RecordMetadata{Labels: map[string]string{}}}
}

type RecordMetadata struct {
    DataSource string            `yaml:"dataSource"`
    Labels     map[string]string `yaml:"labels"`
}

func NewRecordMetadata() *RecordMetadata {
    return &RecordMetadata{
        DataSource: "",
        Labels:     make(map[string]string),
    }
}
