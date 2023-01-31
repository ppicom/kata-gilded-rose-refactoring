//go:build unit
// +build unit

package gildedrose_test

import (
	"testing"

	"github.com/ppicom/kata-gilded-rose-refactoring/internal/gildedrose"
	"github.com/stretchr/testify/suite"
)

type updateQualityShould struct {
	suite.Suite
}

func (s *updateQualityShould) Test_decrease_sellin_date_of_regular_items_by_one() {
	var items = []*gildedrose.Item{
		{Name: "+5 Dexterity Vest", SellIn: 10, Quality: 20},
	}

	gildedrose.UpdateQuality(items)

	s.Equal(9, items[0].SellIn)
}

func (s *updateQualityShould) Test_decrease_quality_of_regular_items_by_one() {
	var items = []*gildedrose.Item{
		{Name: "+5 Dexterity Vest", SellIn: 10, Quality: 20},
	}

	gildedrose.UpdateQuality(items)

	s.Equal(19, items[0].Quality)
}

func (s *updateQualityShould) Test_decrease_quality_twice_as_fast_after_sellin_date_has_passed() {
	items := []*gildedrose.Item{
		{Name: "+5 Dexterity Vest", SellIn: 0, Quality: 20},
	}

	gildedrose.UpdateQuality(items)

	s.Equal(18, items[0].Quality)
}

func (s *updateQualityShould) Test_not_decrease_quality_to_negative_values() {
	items := []*gildedrose.Item{
		{Name: "+5 Dexterity Vest", SellIn: 10, Quality: 0},
	}

	gildedrose.UpdateQuality(items)

	s.Equal(0, items[0].Quality)
}

func (s *updateQualityShould) Test_increase_quality_of_aged_brie() {
	items := []*gildedrose.Item{
		{Name: "Aged Brie", SellIn: 2, Quality: 0},
	}

	gildedrose.UpdateQuality(items)

	s.Equal(1, items[0].Quality)
}

func (s *updateQualityShould) Test_not_increase_quality_above_50() {
	items := []*gildedrose.Item{
		{Name: "Aged Brie", SellIn: 2, Quality: 50},
	}

	gildedrose.UpdateQuality(items)

	s.Equal(50, items[0].Quality)
}

func (s *updateQualityShould) Test_not_decrease_sulfuras_quality() {
	items := []*gildedrose.Item{
		{Name: "Sulfuras, Hand of Ragnaros", SellIn: -1, Quality: 80},
	}

	gildedrose.UpdateQuality(items)

	s.Equal(80, items[0].Quality)
}

func (s *updateQualityShould) Test_increase_backstage_passes_quality() {
	items := []*gildedrose.Item{
		{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 15, Quality: 20},
		{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 10, Quality: 20},
		{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 5, Quality: 20},
	}

	gildedrose.UpdateQuality(items)

	s.Equal(21, items[0].Quality)
	s.Equal(22, items[1].Quality)
	s.Equal(23, items[2].Quality)
}

func (s *updateQualityShould) Test_set_quality_of_concerts_to_0_after_sellin_date() {
	items := []*gildedrose.Item{
		{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 0, Quality: 20},
	}

	gildedrose.UpdateQuality(items)

	s.Equal(0, items[0].Quality)
}

func (s *updateQualityShould) Test_decrease_quality_of_conjured_items_twice_as_fast() {
	items := []*gildedrose.Item{
		{Name: "Conjured Mana Cake", SellIn: 3, Quality: 6},
		{Name: "Conjured Mana Cake", SellIn: 0, Quality: 6},
	}

	gildedrose.UpdateQuality(items)

	s.Equal(4, items[0].Quality)
	s.Equal(2, items[1].Quality)
}

func TestRegularItemsShould(t *testing.T) {
	suite.Run(t, new(updateQualityShould))
}
