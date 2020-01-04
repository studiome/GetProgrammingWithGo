package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type coordinate struct {
	d, m, s float64
	h       rune
}

func (c coordinate) String() string {
	return fmt.Sprintf("%v˚%v'%.1f\"%c", c.d, c.m, c.s, c.h)
}

func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1.0
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

type location struct {
	Lat  coordinate `json:"latitude"`
	Long coordinate `json:"longitude"`
}

func (c coordinate) MarshalJSON() ([]byte, error) {
	cc := struct {
		Decimal float64 `json:"decimal"`
		DMS     string  `json:"dms"`
		D       float64 `json:"degree"`
		M       float64 `json:"minutes"`
		S       float64 `json:"seconds"`
		H       rune    `json:"hemisphere"`
	}{
		Decimal: c.decimal(),
		DMS:     c.String(),
		D:       c.d,
		M:       c.m,
		S:       c.s,
		H:       c.h,
	}

	return json.Marshal(cc)
}

func main() {
	elysium := location{
		Lat:  coordinate{4, 30, 0.0, 'N'},
		Long: coordinate{135, 54, 0.0, 'E'},
	}
	b, err := json.MarshalIndent(elysium, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}
