package main

type notification interface {
	importance() int
}

type directMessage struct {
	senderUsername string
	messageContent string
	priorityLevel  int
	isUrgent       bool
}

type groupMessage struct {
	groupName      string
	messageContent string
	priorityLevel  int
}

type systemAlert struct {
	alertCode      string
	messageContent string
}

func (msg directMessage) importance() int {
    if msg.isUrgent {
        return 50
    }
    return msg.priorityLevel
}

func (msg groupMessage) importance() int {
    return msg.priorityLevel
}

func (msg systemAlert) importance() int {
    return 100
}

func processNotification(n notification) (string, int) {
    switch msg := n.(type) {
    case directMessage:
        return msg.senderUsername, msg.importance()
    case groupMessage:
        return msg.groupName, msg.importance()
    case systemAlert:
        return msg.alertCode, msg.importance()
    default:
        return "", 0
    }
}
