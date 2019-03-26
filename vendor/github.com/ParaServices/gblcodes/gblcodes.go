package gblcodes

// AppCode ...
type AppCode interface {
	Message() string
	Code() string
}

type appCode string

var messages = make(map[AppCode]string)

// Message returns Message string from message
// map for particular AppCode
func (a appCode) Message() string {
	return messages[a]
}

// Code returns AppCode title string
func (a appCode) Code() string {
	return string(a)
}

func init() {
	// Add Messages to messagesMaps, any new message maps must be
	// addded to be available to global callers, add new codes
	// and messages to respective files
	messageMaps := []map[AppCode]string{crudMessages, httpMessages}
	for _, msgMap := range messageMaps {
		for k, v := range msgMap {
			messages[k] = v
		}
	}
}
