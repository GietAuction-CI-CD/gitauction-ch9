package main

import "fmt"

var version string // 빌드시 ldflags 플래그를 경유해서 버전을 저장하기 위한 변수

func main(){
	fmt.Printf("Exampl %s\n", version)
}