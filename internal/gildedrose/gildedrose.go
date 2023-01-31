package gildedrose

import "strings"

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {

		if items[i].Name == "Sulfuras, Hand of Ragnaros" {
			transformSulfuras(items[i])
			continue
		}

		if items[i].Name == "Aged Brie" {
			transformBrie(items[i])
			continue
		}

		if items[i].Name == "Backstage passes to a TAFKAL80ETC concert" {
			transformBasckstagePasses(items[i])
			continue
		}

		if strings.HasPrefix(items[i].Name, "Conjured") {
			transformConjured(items[i])
			continue
		}

		transformRegularItem(items[i])

	}
}

func transformBasckstagePasses(item *Item) {
	var qualityDelta int = 1
	switch {
	case item.SellIn < 1:
		qualityDelta = -item.Quality
	case item.SellIn <= 5:
		qualityDelta = 3
	case item.SellIn <= 10:
		qualityDelta = 2
	}
	item.Quality += qualityDelta
	item.SellIn -= 1
}

func transformSulfuras(item *Item) {
	item.SellIn -= 1
}

func transformBrie(item *Item) {
	if item.Quality < 50 {
		item.Quality += 1
	}
	item.SellIn -= 1
}

func transformRegularItem(item *Item) {
	if item.Quality > 0 {
		var qualityDecrease = 1
		if item.SellIn < 1 {
			qualityDecrease = 2
		}
		item.Quality -= qualityDecrease
	}

	item.SellIn -= 1
}

func transformConjured(item *Item) {
	if item.Quality > 0 {
		var qualityDecrease = 2
		if item.SellIn < 1 {
			qualityDecrease = 4
		}
		item.Quality -= qualityDecrease
	}

	item.SellIn -= 1
}
