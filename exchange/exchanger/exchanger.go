package exchanger

import "strconv"

type Exchanger interface {
	Exchange()
	String() string
}

// StringPair
type StringPair struct {
	first,second string
}

func NewStringPair(s1 string,s2 string) *StringPair {
	return &StringPair{s1,s2}
}

func (spair *StringPair) Exchange() {
	temp := spair.first
	spair.first = spair.second
	spair.second = temp
}

func (spair *StringPair) String() string {
	return spair.first + "," + spair.second
}

// IntPair
type IntPair struct {
	first_i,second_i int
}

func NewIntPair(i1 int,i2 int) *IntPair {
	return &IntPair{i1,i2}
}

func (ipair *IntPair) Exchange() {
	temp := ipair.first_i
	ipair.first_i = ipair.second_i
	ipair.second_i = temp
}

func (ipair *IntPair) String() string {
	first_s := strconv.Itoa(ipair.first_i)
	second_s := strconv.Itoa(ipair.second_i)
	return first_s + "," + second_s

}

// FloatPair
type FloatPair struct {
	first_f,second_f float64
}

func NewFloatPair(f1 float64, f2 float64) *FloatPair {
	return &FloatPair{f1,f2}
}

func (fpair *FloatPair) Exchange()  {
	temp := fpair.first_f
	fpair.first_f = fpair.second_f
	fpair.second_f = temp
}

func (fpair *FloatPair) String() string {
	first_s := strconv.FormatFloat(fpair.first_f,'f',4,64)
	second_s := strconv.FormatFloat(fpair.second_f,'f',4,64)
	return first_s + "," + second_s
}
	
