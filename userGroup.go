package main

import(
    "sync"
)

type UserGroup struct {
    set map[*User]bool
    memberLock sync.Mutex
}

func (us *UserGroup) Add(user *User) bool {
    us.memberLock.Lock()
    _, found := us.set[user]
    us.set[user] = true
    us.memberLock.Unlock()
    return !found   //False if it existed already
}

func (us *UserGroup) Remove(user *User) {
    us.memberLock.Lock()
    delete(us.set, user)
    us.memberLock.Unlock()
}

func NewUserGroup() UserGroup {
    return UserGroup{set: make(map[*User] bool)}
  }

 