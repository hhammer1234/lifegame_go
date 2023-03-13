package main

import (
	"fmt"
	"math/rand"
	"time" //time.Sleep(8 * time.Second)
)

const whentrue string = "⏹" //기호 설정
const whenfalse string = "_"
const squaresize int = 30

type celltype [][]bool

func cellsprint(cells *celltype) {
	for i, _ := range *cells {
		for _, v := range (*cells)[i] {
			if v {
				fmt.Print(whentrue, " ")
			} else {
				fmt.Print(whenfalse, " ")
			}
		}
		fmt.Printf("\n")
	}
}
func nearcount(cells *celltype, x int, y int) int {
	rst := 0
	for i := -1; i <= 1; i++ { //x에 더할 값 -1,0,1
		for j := -1; j <= 1; j++ { //y에 더할 값 -1,0,1
			xx := x + i
			yy := y + j
			if xx == x && yy == y {
				//자기자신 제외
			} else if xx >= 0 &&
				xx < len(*cells) &&
				yy >= 0 &&
				y+j < len(*cells) { //인덱스 밖으로 튀는 것 방지
				if (*cells)[xx][yy] { //x,y에 더한 값의 t, f 여부로 카운트
					rst++
				}
			}
		}
	}
	return rst
}

func cellsturn(cells *celltype) {
	for i := 0; i < len(*cells); i++ {
		for j := 0; j < len((*cells)[i]); j++ {
			//fmt.Print(nearcount(cells, i, j), " ") //셀별 nearcount 출력
			now := nearcount(cells, i, j)
			if now == 3 {
				(*cells)[i][j] = true
			} else if now == 2 || now == 3 {
				//아무것도 안함
			} else if now <= 1 {
				(*cells)[i][j] = false
			} else if now >= 4 {
				(*cells)[i][j] = false
			}
		}
		//fmt.Println() //셀별 nearcount 출력 줄바꿈
	}
}

func main() {
	cells := make(celltype, squaresize) //cells 슬라이스 선언
	for i, _ := range cells {
		cells[i] = make([]bool, squaresize) //cells 2차원 선언
	}
	fmt.Print("랜덤 여부를 입력하세요. 0=초기설정 1=랜덤:")
	randomswitch := false
	fmt.Scanln(&randomswitch)
	if randomswitch {
		for i, _ := range cells {
			for j, _ := range cells[i] {
				a := rand.Intn(1000)
				if a > 900 {
					cells[i][j] = true
				} else {
					cells[i][j] = false
				}

			}
		}
	} else {
		//초기 사용자 설정
		cells[7][2] = true
		cells[7][3] = true
		cells[7][5] = true
		cells[7][7] = true
		cells[8][2] = true
		cells[8][3] = true
		cells[8][4] = true
		cells[9][5] = true
		cells[9][7] = true

	}
	cellsprint(&cells)
	for i := 0; i < 10000; i++ {
		fmt.Printf("\x1bc") //터미널 비우기
		cellsprint(&cells)  //예쁘게 출력
		cellsturn(&cells)   //한턴
		time.Sleep(100 * time.Millisecond)
	}
}
