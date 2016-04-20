package routes

import (
	"io"
	"fmt"
	"time"
	"crypto/md5"
)

type Token struct{
	Id string
}

var token Token

func (t *Token) CreateToken() {
	h := md5.New()
	word :="eto_pirka(^^)(--)"
	io.WriteString(h,word+time.Now().String())
	t.Id = fmt.Sprintf("%x",h.Sum(nil))
}
