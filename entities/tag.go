package entities

type TagModifier string

const (
	add    TagModifier = "+"
	remove TagModifier = "-"
)

type Tag struct {
	modifier TagModifier
	name     string
}

func TagToAdd(name string) Tag {
	return Tag{modifier: add, name: name}
}

func TagToRemove(name string) Tag {
	return Tag{modifier: remove, name: name}
}

func Equals(t1, t2 []Tag) bool {
	if len(t1) != len(t2) {
		return false
	}
	for i, v := range t1 {
		if v.modifier != t2[i].modifier {
			return false
		}
		if v.name != t2[i].name {
			return false
		}
	}
	return true
}
