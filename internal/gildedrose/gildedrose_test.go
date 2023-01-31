//go:build unit
// +build unit

package gildedrose_test

import (
	"testing"

	"github.com/ppicom/kata-gilded-rose-refactoring/internal/gildedrose"
)

func Test_Foo(t *testing.T) {
	var items = []*gildedrose.Item{
		{"foo", 0, 0},
	}

	gildedrose.UpdateQuality(items)

	if items[0].Name != "fixme" {
		t.Errorf("Name: Expected %s but got %s ", "fixme", items[0].Name)
	}
}
