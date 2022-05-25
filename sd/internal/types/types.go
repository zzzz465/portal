package types

type Record struct {
    Name     string         `yaml:"name" json:"name"`
    Tags     []string       `yaml:"tags" json:"tags"`
    Metadata RecordMetadata `yaml:"metadata" json:"metadata"`
}

func NewRecord() Record {
    return Record{Metadata: RecordMetadata{Labels: map[string]string{}}}
}

type RecordMetadata struct {
    DataSource string            `yaml:"dataSource" json:"dataSource"`
    Labels     map[string]string `yaml:"labels" json:"labels"`
}

func NewRecordMetadata() *RecordMetadata {
    return &RecordMetadata{
        DataSource: "",
        Labels:     make(map[string]string),
    }
}
