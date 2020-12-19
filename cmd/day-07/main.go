package main

import (
	"aoc-2020/internal"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type BagContains struct {
	quantity int
	bag      *Bag
}
type Bag struct {
	name     string
	children []*BagContains
}

type Bags []*Bag

func (r *Bags) findOrCreateBag(name string) *Bag {
	for _, bag := range *r {
		if bag.name == name {
			return bag
		}
	}

	newBag := &Bag{
		name: name,
	}

	*r = append(*r, newBag)

	return newBag
}

func (b *Bag) String() string {
	children := make([]string, 0)
	for _, child := range b.children {
		children = append(children, fmt.Sprintf("%d %s", child.quantity, child.bag.name))
	}
	return fmt.Sprintf("%s bag contains %s", b.name, strings.Join(children, ";"))
}

func (b *Bag) fitsBag(name string) bool {
	for _, child := range b.children {
		if child.bag.name == name || child.bag.fitsBag(name) {
			return true
		}
	}
	return false
}

func (b *Bag) countInsideBags() (count int) {
	for _, child := range b.children {
		count += child.quantity + (child.quantity * child.bag.countInsideBags())
	}
	return
}

var test = "light red bags contain 1 bright white bag, 2 muted yellow bags.\ndark orange bags contain 3 bright white bags, 4 muted yellow bags.\nbright white bags contain 1 shiny gold bag.\nmuted yellow bags contain 2 shiny gold bags, 9 faded blue bags.\nshiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.\ndark olive bags contain 3 faded blue bags, 4 dotted black bags.\nvibrant plum bags contain 5 faded blue bags, 6 dotted black bags.\nfaded blue bags contain no other bags.\ndotted black bags contain no other bags."

func main() {
	input := internal.ReadFileLines("cmd/day-07/input.txt")
	//input := strings.Split(test, "\n")

	bags := make(Bags, 0)
	regexInnerBags, _ := regexp.Compile("(\\d)\\s?([\\w|\\s]+)\\sbags?")

	for _, rawBags := range input {
		parts := strings.Split(rawBags, " bags contain ")
		outerBagName := strings.TrimSpace(parts[0])
		outerBag := bags.findOrCreateBag(outerBagName)
		matches := regexInnerBags.FindAllStringSubmatch(parts[1], -1)
		for _, innerMatch := range matches {
			innerBagQuantity, err := strconv.Atoi(innerMatch[1])
			if err != nil {
				log.Fatal(err.Error())
			}
			innerBagName := strings.TrimSpace(innerMatch[2])
			innerBag := bags.findOrCreateBag(innerBagName)
			outerBag.children = append(outerBag.children, &BagContains{
				quantity: innerBagQuantity,
				bag:      innerBag,
			})
		}
	}

	bagsContaining := make([]*Bag, 0)

	for _, bag := range bags {

		fmt.Println(bag)

		if bag.fitsBag("shiny gold") {
			bagsContaining = append(bagsContaining, bag)
		}
	}
	fmt.Println(len(bagsContaining))

	shinyGoldBag := bags.findOrCreateBag("shiny gold")
	fmt.Println(shinyGoldBag.countInsideBags())
}
