package examples

import (
	"fmt"

	"github.com/brainfucker/stored"
)

// type User struct {
// 	ID       uint64
// 	Username string
// }

type User struct {
	ID       int
	Username string
	City     string
}

func (u *User) Stored(k stored.Key, v stored.Value) {
  k.IDDate(&u.ID)
	switch v.Ver(1) {
	default:
		v.String(&u.Username))
		v.String(&u.City))
	}
}

func (u *User) Pack(p stored.Pack) {
  Users.SetKey(p, u.ID)
	p.Ver(1)
  p.String(u.Username))
  p.String(u.City))
}

func (u *User) UnPack(p stored.Unpack) {
  u.ID = Users.GetKey(p)

  switch p.Ver {
  case 1:
    u.Username = p.String()
    u.City = p.String()
  }
}

const (
	DB_USERS stored.ObjectID = 1
)

//var Users stored.Object[user, user.ID]
var Users = stored.Object1Key(DB_USERS, User{}, User{}.ID)
var UsersByCity = stored.Index1Key(Users, User{}, User{}.City, false) // last bool is for unique

func main() {
	err := stored.Connect("./fdb.cluster")
	if err != nil {
		panic("connection failed")
	}

	u, err := Users.Get(23)
	fmt.Println(u, err)
	/*user := User{}
		Users = db.Object1Key(DB_USERS, user, user.ID)
		UsersUsername = db.Index1Key(DB_USERS, Users{}, user.Username)

		Users.Add(&User{Username: "Bob"})

		user, err := Users.Get(1)
		fmt.Println(user.Username, err)

	  // by index

		user,err := UsersUsername.Get("Bob")
		fmt.Println(user.ID, err)*/

}
