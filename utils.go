package main

type UserSet struct {
    set map[*User]bool
}

func (us *UserSet) Add(user *User) bool {
    _, found := us.set[user]
    us.set[user] = true
    return !found   //False if it existed already
}

func (us *UserSet) Remove(user *User) {
    delete(us.set, user)
}

func NewUserSet() UserSet {
    return UserSet{make(map[*User] bool)}
  }