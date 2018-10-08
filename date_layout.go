package paratils

// DateLayouter ...
type DateLayouter interface {
	String() string
}

type dateLayout string

// String ...
func (d dateLayout) String() string {
	return string(d)
}

// YYYYMMDD ...
var YYYYMMDD dateLayout = "2006-01-02"

// ISO8601 ...
var ISO8601 dateLayout = "2006-01-02T15:04:05-0700"

// DateLayout ...
func DateLayout(dl DateLayouter) string {
	return dl.String()
}
