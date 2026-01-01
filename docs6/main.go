package main

import (
	"fmt"
	"errors"
	"time"
	"encoding/json"
)
type Product struct {
	ID int `json:"id"`
    Name string `json:"name"`
    Price int64 `json:"price"`
}

// DisCount
func (p *Product) DisCount(percent int) error {
	if percent < 0 || percent > 100 {
		return errors.New("invalid discount percentage")
	}
	p.Price = p.Price * int64(100 - percent)/100
	return nil
}

type User struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Email string `json:"email"`
	Tag []string
}


func (u *User) AddTag(tag string) {
	u.Tag = append(u.Tag,tag)
}

type Audit struct {
	CreatedAt time.Time   `json:"created_at"`
	UpdateAt time.Time   `json:"updated_at"`
}

type Employee struct {
	User
	Audit
	Title string `json:"title"`
}

type Post struct {
	ID int 		`json:"id"`
	Title string `json:"title"`
	Content string `json:"content,omitempty"`
	Author string 	`json:"author,omitempty"`
	Created time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewPost(title string, content string) (*Post,error) {
	if title == "" {
		return nil,errors.New("title cannot be empty")
	}
	return &Post{
		Title: title,
		Content: content,
		Created: time.Now(),
	}, nil
}
func demoProduct() {
	p := Product{ID:1,Name : "book",Price : 1000}
	p.DisCount(25)
	fmt.Println(p)
}

func demoUser() {
	u := User{Name: "John", Age: 30, Email: "john@example.com"}
	u.AddTag("admin")
	fmt.Println(u)
}
func demoPost() {
	post,err := NewPost("hello","content here")
	if err != nil {
		fmt.Println(err)
		return
	}
	b,_ := json.Marshal(post)
	fmt.Println(string(b))
	fmt.Println(post)
}

func main(){
	demoProduct()
	demoUser()
	demoPost()
}
