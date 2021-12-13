package main

import (
	"fmt"

	"github.com/go-playground/validator"
)

func main() {
	mystruct := MyStruct{
		ID: uint16(1),
		IPList: []string{
			"192.168.0.0",
			"192.168.0.1",
			"192.168.0.2",
			"2001:0db8:85a3:0000:0000:8a2e:0370:7334",
			"2001:DB8::8:800:200C:417A",
		},
	}

	validate := validator.New()
	if err := validate.Struct(mystruct); err != nil {
		fmt.Println(err)
		return
	}

}

type MyStruct struct {
	ID     uint16   `validate:"required"`
	IPList []string `validate:"unique,dive,ip_addr"`
}
