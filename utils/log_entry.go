package utils

type LogEntry struct {
	Target string `bson:"target,omitempty" json:"target,omitempty"`
	Type   string `bson:"type,omitempty" json:"type,omitempty"`
	Action string `bson:"action,omitempty" json:"action,omitempty"`
	Actor  string `bson:"actor,omitempty" json:"actor,omitempty"`
	Data   any    `bson:"data,omitempty" json:"data,omitempty"`
}
