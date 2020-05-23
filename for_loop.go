package main

import "fmt"

func main(){
	i:=1
	for i <= 3  {
		fmt.Println(i)
		i = i + 1
	}
	for j:= 7; j<=9; j++{
		fmt.Println(j)
	}
	for n:=10; n<=20; n++{
		if n%2==0 {
			continue
		}
		fmt.Println(n)
		}

}
