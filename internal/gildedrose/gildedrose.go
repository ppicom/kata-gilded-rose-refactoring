package gildedrose

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {

		if items[i].Name == "Sulfuras, Hand of Ragnaros" {
			items[i].SellIn -= 1
			continue
		}

		if items[i].Name == "Aged Brie" {
			if items[i].Quality < 50 {
				items[i].Quality += 1
			}
			items[i].SellIn -= 1
			continue
		}

		if items[i].Name == "Backstage passes to a TAFKAL80ETC concert" {
			var qualityDelta int = 1
			switch {
			case items[i].SellIn < 1:
				qualityDelta = -items[i].Quality
			case items[i].SellIn <= 5:
				qualityDelta = 3
			case items[i].SellIn <= 10:
				qualityDelta = 2
			}
			items[i].Quality += qualityDelta
			items[i].SellIn -= 1
			continue
		}

		items[i].SellIn -= 1

		if items[i].Quality > 0 {
			var qualityDecrease = 1
			if items[i].SellIn < 1 {
				qualityDecrease = 2
			}
			items[i].Quality -= qualityDecrease
		}
	}
}
