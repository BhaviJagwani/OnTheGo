package main

import(
    "sync"
)

type UserGroup struct {
    set map[*User]bool
    memberLock sync.Mutex
}

func (us *UserGroup) Add(user *User) bool {
    _, found := us.set[user]
    us.set[user] = true
    return !found   //False if it existed already
}

func (us *UserGroup) Remove(user *User) {
    delete(us.set, user)
}

func NewUserGroup() UserGroup {
    return UserGroup{set: make(map[*User] bool)}
  }

 