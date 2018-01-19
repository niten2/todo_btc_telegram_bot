package models

import (
  // "fmt"
  // "gopkg.in/mgo.v2"
  // "crypto/sha256"
  "gopkg.in/mgo.v2/bson"
  "app-telegram/db"
)

type User struct {
  // ID bson.ObjectId `json:"id" bson:"_id,omitempty"`
  ID string `json:"id" bson:"_id,omitempty"`
  IdTelegramm string `json:"id_telegramm"`
  Name string `json:"name"`
}

func (u *User) AsMap() map[string]string {
  user := make(map[string]string)

  user["id"] = u.ID
  user["id_telegramm"] = u.IdTelegramm

  return user
}

func CreateUser(Name string, IdTelegramm string) User {
  user_collection := db.Db.C("users")

  user_document := User{Name: Name, IdTelegramm: IdTelegramm}

  user_collection.Insert(user_document)

  return user_document
}

func FindUser(Name string) User {
  user_collection := db.Db.C("users")

  user := User{}
  err := user_collection.Find(bson.M{"name": Name}).One(&user)

  if err != nil {
    panic(err.Error())
  }

  return user
}

// func UserLogin(email string, password string) (user *User, err error) {
//   user = &User{}
//   encryptedPassword := userEncrypedPassword(password)
//   db := db.Conn.Where("email = ? AND encrypted_password = ?", email, encryptedPassword).First(user).Scan(user)
//   err = db.Error
//   if err != nil {
//     return nil, err
//   }
//   return
// }

// func userEncrypedPassword(password string) (encryptedPassword string) {
//   saltedPassword := fmt.Sprintf("%s_%s", password, beego.AppConfig.String("salt"))
//   encryptedPassword = fmt.Sprintf("%x", sha256.Sum256([]byte(saltedPassword)))
//   return
// }

// func AllUsers() (users []*User, err error) {
//   res := db.Conn.Debug().Order("first_name asc, last_name asc").Find(&users)
//   err = res.Error
//   if err != nil {
//     beego.BeeLogger.Error("Error finding users: %v", res.Error.Error())
//   }
//   return
// }

// func FindUser(id string) (*User, error) {
//   user := &User{ID: id}
//   res := db.Conn.Debug().First(&user, "id=?", id)
//   err := res.Error
//   if err != nil {
//     beego.BeeLogger.Error("Error finding user with id %s: %v", id, res.Error.Error())
//   }
//   return user, err
// }

// func CreateUser(firstName string, lastName string, email string, password string) (*User, error) {
//   user := &User{FirstName: firstName, LastName: lastName, Email: email, EncryptedPassword: userEncrypedPassword(password)}
//   res := db.Conn.Debug().Create(user)
//   err := res.Error
//   return user, err
// }

// func UpdateUser(id string, firstName *string, lastName *string, email *string, password *string) (*User, error) {
//   user := &User{}
//   res := db.Conn.Debug().First(user, "id=?", id)
//   if res.Error != nil {
//     beego.BeeLogger.Error("Error finding user with id %s: %v", id, res.Error.Error())
//   }
//   if firstName != nil {
//     user.FirstName = *firstName
//   }
//   if lastName != nil {
//     user.LastName = *lastName
//   }
//   if email != nil {
//     user.Email = *email
//   }
//   if password != nil {
//     user.EncryptedPassword = userEncrypedPassword(*password)
//   }
//   res = db.Conn.Debug().Save(user)
//   err := res.Error
//   return user, err
// }

// func DeleteUser(id string) (User, error) {
//   user := User{ID: id}
//   res := db.Conn.Debug().Delete(&user)
//   err := res.Error
//   return user, err
// }
