package main

type UserSet struct {
    set map[*User]bool
}

func (set *UserSet) Add(user *User) bool {
    _, found := set.set[user]
    set.set[user] = true
    return !found   //False if it existed already
}

func NewUserSet() UserSet {
    return UserSet{make(map[*User] bool)}
  }