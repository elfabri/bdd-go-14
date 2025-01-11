package two

func (u User) SendMessage(message string, messageLength int) (string, bool) {
    if u.Membership.MessageCharLimit >= messageLength {
        return message, true
    }
    return "", false
}

func newUser(name string, membershipType string) User {
    if membershipType == "premium" {
        return User {
            name,
            Membership {
                "premium",
                1000,
            },
        }
    }

    return User {
        name,
        Membership {
            "standard",
            100,
        },
    }
}

// don't touch below this line

type User struct {
	Name string
	Membership
}

type Membership struct {
	Type             string
	MessageCharLimit int
}
