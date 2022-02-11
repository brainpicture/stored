package main

/*import "fmt"

type Users struct {
	ID       uint64
	Username string
}

func (u *User) Key(k stored.Key) {
	 k.Add(u.ID)
}

func (u *User) Write(v stored.Writer) {
   v.Version(1)
	 v.Add(u.Username)
}

func (u *User) Read(v stored.Reader) {
	switch v.Version {
	case 1:
	  v.Scan(&u.Username)
	}
}


func (u *User) GetByUsername(i stored.Index) {
	i.Add(u.ID)
	stored.ByIndex(u)
}

var Users stored.Object

const (
	DB_USERS stored.ObjectID = 1
)

func main() {
	db := stored.New("./fdb.cluster")
	Users = db.Object(DB_USERS, Users{}, stored.PrimaryInt64)
	UsersUsername = db.Index(DB_USERS, Users{}, Users.Username)

	db.Connect()

	Users.Add(&User{Username: "Bob"})

	user, err := Users.Get(1)
	fmt.Println(user.Username, err)

  // by index

	user,err := UsersUsername.Get("Bob")
	fmt.Println(user.ID, err)

}
*/
