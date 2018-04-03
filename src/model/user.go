package model

import (
	"fmt"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

/*
id: number; // 인덱스 고유값
    username: string;  // 유져ID
    password: string;
    confirmPassword: string; // DB생성 필요 없음
    nickname: string;
    email: string;
*/
type User struct {
	ID       bson.ObjectId `bson:"_id,omitempty" json:"_id"`
	UserNo   int           `bson:"user_no" json:"user_no"`
	Password string        `bson:"password" json:"password"`
	Nickname string        `bson:"nickName" json:"nickName"`
	Email    string        `bson:"email" json:"email"`
	Token    string        `bson:"token" json:"token"`
	Image    string        `bson:"image" json:"image"`
	CreateAt time.Time     `bson:"createAt" json:"createAt"`
	UpdateAt time.Time     `bson:"updateAt" json:"updateAt"`
}

// CreateUser is 유저 생성
func (u *User) CreateUser(ucol, icol *mgo.Collection) error {
	identity := Identity{Name: "user"}

	err := identity.Increment(icol)
	if err != nil {
		return err
	}
	u.UserNo = identity.Count
	u.CreateAt = time.Now()
	u.UpdateAt = time.Now()
	return ucol.Insert(u)
}

// UpdateUser is 유저 업데이트
func (u *User) UpdateUser(c *mgo.Collection) error {
	fmt.Println(u)
	return c.Update(bson.M{"user_no": u.UserNo},
		bson.M{"$set": bson.M{"password": u.Password, "userName": u.UserName,
			"nickName": u.Nickname, "email": u.Email, "updateAt": time.Now()}})
}

// DeleteUser is 유저 삭제
func (u *User) DeleteUser(c *mgo.Collection) error {
	fmt.Println(u)
	return c.Remove(bson.M{"user_no": u.UserNo})
}
