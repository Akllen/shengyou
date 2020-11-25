/*
 * @Author: 黄河
 * @Date: 2020-11-25 09:35:50
 * @LastEditTime: 2020-11-25 09:41:33
 * @LastEditors: 黄河
 */
package validate

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func DemoTest() {
	fmt.Println("Demo Test")
}

func CheckSruct(data interface{}) error {
	if err := validator.New().Struct(data); err != nil {
		return err
	}
	return nil
}
