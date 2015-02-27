package obj_lib

type Baser interface  {
	BaseFunc() string
}

type Childer interface {
	Baser
	ChildFunc() string
}

type Base struct{ name string }

func NewBase(s string) Base {
	if s == "" {
		s = "I am Base"
	}
	return Base{s}
}

func (base Base)GetName() string{
	return base.name
}

type Child struct { base Base
	                  child_name string
}

func NewChild(base Base,s string) Child {
	if s == "" {
		s = "I am Child"
	}
	return Child{base,s}
}

func (child Child)GetFullName() string {
	return child.base.name + " " + child.child_name
}

