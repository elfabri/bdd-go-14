package one

type User struct {
	Name string
    Membership
}

type Membership struct {
    Type                string
    MessageCharLimit    int
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
