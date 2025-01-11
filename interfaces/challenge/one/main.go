package main

type Formatter interface {
    Format() string
}

type PlainText struct {
    message string
}

type Bold struct {
    message string
}

type Code struct {
    message string
}

func (msg PlainText) Format() string {
    return msg.message
}

func (msg Bold) Format() string {
    return "**"+msg.message+"**"
}

func (msg Code) Format() string {
    return "`"+msg.message+"`"
}

// Don't Touch below this line

func SendMessage(formatter Formatter) string {
	return formatter.Format() // Adjusted to call Format without an argument
}
