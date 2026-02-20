package client

import (
	"fmt"
)

func ReadMove() (int, int, error){
	var r, c int
	fmt.Println("Enter Move (row col): ")
	_, err := fmt.Scan(&r,&c)
	return r, c, err
}
