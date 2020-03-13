package parse

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

	"github.com/mbauhardt/moneyflow/entities"
)

func ParseTags(commandline string) []entities.Tag {
	tags := make([]entities.Tag, 0)
	tags = append(tags, ParseTagsToAdd(commandline)...)
	tags = append(tags, ParseTagsToRemove(commandline)...)
	return tags
}

func ParseTagsToAdd(commandline string) []entities.Tag {
	tagRegex, _ := regexp.Compile("\\+[a-zA-Z0-9]+(\\s|$)")
	parsedTags := tagRegex.FindAllString(commandline, -1)
	tags := make([]entities.Tag, 0)
	for _, t := range parsedTags {
		parsedName := strings.Trim(t[1:len(t)], " ")
		tags = append(tags, entities.TagToAdd(parsedName))
	}
	return tags
}

func ParseTagsToRemove(commandline string) []entities.Tag {
	tagRegex, _ := regexp.Compile("\\-[a-zA-Z0-9]+(\\s|$)")
	parsedTags := tagRegex.FindAllString(commandline, -1)
	tags := make([]entities.Tag, 0)
	for _, t := range parsedTags {
		parsedName := strings.Trim(t[1:len(t)], " ")
		tags = append(tags, entities.TagToRemove(parsedName))
	}
	return tags
}

func ParseMoney(commandline string) (*entities.Money, error) {
	moneyRegex, _ := regexp.Compile("(\\-){0,1}€[0-9]+(\\s|$)")
	parsedMoneyArray := moneyRegex.FindAllString(commandline, -1)
	if len(parsedMoneyArray) == 0 {
		return nil, nil
	}
	if len(parsedMoneyArray) > 1 {
		return nil, errors.New("More than one money is detected: " + strings.Join(parsedMoneyArray, ""))
	}
	parsedMoney := strings.Replace(parsedMoneyArray[0], "€", "", 1)
	parsedValue := strings.Trim(parsedMoney[0:len(parsedMoney)], " ")
	n, err := strconv.ParseInt(parsedValue, 0, 64)
	if err != nil {
		return nil, err
	}
	return &entities.Money{Value: n}, nil
}

func ParseDescription(s string) (*entities.Description, error) {
	tags := ParseTags(s)
	m, e := ParseMoney(s)
	if e != nil {
		return nil, e
	}
	d := entities.Description{Value: s}
	for _, t := range tags {
		d.Value = strings.Trim(strings.ReplaceAll(d.Value, entities.ToString(&t), ""), " ")
	}
	if m != nil {
		d.Value = strings.Trim(strings.ReplaceAll(d.Value, entities.MoneyToString(m), ""), " ")
	}
	return &d, nil
}
