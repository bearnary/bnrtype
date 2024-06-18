package becustomtype

import (
	"encoding/xml"
	"fmt"
	"math"
)

type Float64TwoDecimal float64

func (mf Float64TwoDecimal) MarshalJSON() ([]byte, error) {

	val := mf.twoDecimalString()
	return []byte(val), nil
}

func (mf Float64TwoDecimal) Float64() float64 {
	return float64(mf)
}

func (mf Float64TwoDecimal) String() string {
	return fmt.Sprintf("%.2f", mf)
}

func (mf Float64TwoDecimal) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	val := mf.twoDecimalString()
	err := e.EncodeElement(val, start)
	return err
}

func (mf Float64TwoDecimal) CentValue() int64 {
	v := mf.Float64() * 100
	cent := int64(v)
	return cent
}

func (mf Float64TwoDecimal) twoDecimalString() string {
	v := float64(mf)

	isMinus := false
	if v < 0 {
		isMinus = true
		v *= -1
	}
	x := math.Round((v + 0.001) * 100)
	dv := math.Mod(x, 100)
	w, _ := math.Modf(x / 100)
	dvStr := fmt.Sprintf("%v", dv)
	if len(dvStr) < 2 {
		dvStr = "0" + dvStr
	}
	val := fmt.Sprintf(`%.0f.%v`, math.Trunc(w), dvStr)
	if isMinus {
		val = "-" + val
	}
	return val
}

func (mf Float64TwoDecimal) ValueWithPercent(percent float64) Float64TwoDecimal {
	val := mf.Float64() * percent / 100
	return Float64TwoDecimal(val)
}
