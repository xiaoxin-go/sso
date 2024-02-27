package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"sso/conf"
	"sso/database"
	"sso/model"
	"strings"
	"time"
)

func testFirstThenRedis() {
	log.Println("测试直接查询用户和通过redis查询性能差异")
	s := time.Now()
	for i := 0; i < 999; i++ {
		u := model.TUser{}
		u.FirstById(3)
	}
	end := time.Now()

	s1 := time.Now()
	for i := 0; i < 999; i++ {
		//user := model.TUser{}
		//e := user.QueryById(3)
		//fmt.Println("user------>", user, e, user.CreatedAt, user.UpdatedAt)
		(&model.TUser{}).QueryById(3)
	}
	end1 := time.Now()
	log.Println("直接从数据库取100次", end.Sub(s).Milliseconds())
	log.Println("不通过反射从redis读取", end1.Sub(s1).Milliseconds())
}

func insertUser1000() {
	// 插入1000次用户信息
	users := make([]*model.TUser, 0)
	for i := 1; i <= 1000; i++ {
		users = append(users, &model.TUser{
			Username: fmt.Sprintf("test-%d", i),
			NameCn:   fmt.Sprintf("测试%d号", i),
			Email:    fmt.Sprintf("test%d@qq.com", i),
			Password: fmt.Sprintf("password%d", i),
		})
	}
	database.DB.Create(users)
}
func toUpper(name string) string {
	nameLetters := strings.Split(name, "_")
	results := make([]string, 0)
	for _, str := range nameLetters {
		results = append(results, strings.ToUpper(str[:1])+str[1:])
	}
	return strings.Join(results, "")
}

type U1 struct {
	Age int
}

type U2 struct {
	Age uint32
}

func testTo() {
	u1 := U1{Age: -10}
	bs, _ := json.Marshal(u1)
	u2 := U2{}
	e := json.Unmarshal(bs, &u2)
	fmt.Println(e)
	fmt.Println(u2, u1)
}

func testTime() {
	e := database.R.HMSet("test1", map[string]any{"username": "test1", "enabled": 1, "created_at": time.Now().Format(time.DateTime)}).Err()
	fmt.Println("e", e)
	fields := []string{"username", "enabled", "created_at"}
	results, e := database.R.HMGet("test1", fields...).Result()
	fmt.Println(results, e)
	user11 := &model.TUser{}
	for _, v := range fields {
		te := reflect.ValueOf(user11)
		te = te.Elem()
		fe := te.FieldByName(toUpper(v))
		if fe.IsValid() && fe.CanSet() {
			fmt.Println(fe.Type().Name(), fe.Type().String(), fe.CanSet(), fe.Kind())
			switch fe.Type().Name() {
			case "Time":
				t, _ := time.Parse(time.DateTime, results[2].(string))
				fe.Set(reflect.ValueOf(t))
			}
		}
	}
	fmt.Println(user11)
}
func main() {
	testTo()
	return
	conf.InitConfig()
	database.InitDB()
	database.InitRedis()

	var user2 model.TUser
	database.DB.Where("username = ?", "dfafdsfs").First(&user2)
	fmt.Println(database.DB.Updates(&user2).Error)

	return
	p := map[string]any{
		"name":        "netops",
		"name_cn":     "网络自动化",
		"description": "网络自动化111",
		"enabled":     1,
	}
	p1 := &model.TPlatform{
		Description: "sso",
	}
	bs, _ := json.Marshal(&p)
	json.Unmarshal(bs, p1)
	platform := &model.TPlatform{
		Name: "sso",
	}
	if e := database.DB.Model(platform).Where("id = ?", 1).Updates(p1).Error; e != nil {
		fmt.Println(e)
		return
	}
	fmt.Println(platform)
	fmt.Println(p1)
	return

	user := model.TUser{}
	//insertUser1000()
	testFirstThenRedis()
	return

	if err := user.FirstById(3); err != nil {
		fmt.Println(err)
		return
	}
	user1 := user
	if e := database.DB.Model(&user).Updates(map[string]any{"password": "123"}).Error; e != nil {
		fmt.Println(e)
		return
	}
	fmt.Println(user)
	fmt.Println(user1)

}
