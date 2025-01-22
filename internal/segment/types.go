package segment

// SegmentOutput holds raw data from a segment (plain text, plus flags).
// The theme uses this to colorize/wrap text.
type SegmentOutput struct {
    Name   string
    Text   string
    IsRoot bool
    IsSSH  bool
}
