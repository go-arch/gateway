package utils

import (

	"fmt"
)

func Fatal(err error){
	if err != nil {
		fmt.Println(err)
	}
}

