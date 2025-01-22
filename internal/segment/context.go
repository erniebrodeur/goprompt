package segment

// Context holds any data the Manager wants to share with segments.
// For example, the current working directory, whether the user is root, etc.
type Context struct {
	Pwd    string
	IsRoot bool
	IsSSH  bool
}
