package main

// gorm原生SQL
// gorm的SQL生成器

type Ceshi struct { //默认加s，所以对应表ceshis
	Id      int    `json:"id"` //在结构体变为json返回的时候，自动用id替换Id
	Punk    int    `json:"punk"`
	Bigname string `json:"bigname"`
}

func main() {

}
