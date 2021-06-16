package gift

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

type Gift struct {
	Name    string
	Amount  int64
	Percent float64 `json:"-"`
}

var allGift = map[string]Gift{
	"coin_lv1": {
		Name:    "coin_lv1",
		Amount:  1000,
		Percent: 50,
	},
	"coin_lv2": {
		Name:    "coin_lv2",
		Amount:  2000,
		Percent: 30,
	},
	"rare_box": {
		Name:    "rare_box",
		Amount:  1,
		Percent: 12,
	},
	"epic_box": {
		Name:    "rare_box",
		Amount:  1,
		Percent: 7,
	},
	"legend_box": {
		Name:    "legend_box",
		Amount:  1,
		Percent: 1,
	},
}

func RandomGift() Gift {
	chance := rand.ExpFloat64() * 100
	fmt.Println("chance: ", chance)
	var cumulative = 0.0
	for _, value := range allGift {
		cumulative += value.Percent
		if chance <= cumulative {
			return value
		}
	}
	return allGift["coin_lv1"]
}
func (g *Gift) GiftToJsonString() string {
	fmt.Println("gift: ", g)
	b, err := json.Marshal(&g)
	if err != nil {
		fmt.Println(err)
		return "nil"
	}
	return string(b)
}
