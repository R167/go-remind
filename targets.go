package remind

type Target interface {
	ID() string
}

type User string

func (u User) ID() string {
	return string(u)
}

const Me = User("me")
