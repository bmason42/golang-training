package pkg

type Base struct {
	Version  string
	DataType string
}
type SampleStruct struct {
	Base
	ExportData   string
	internalData string
}

func NewSampleStruct(iData string, eData string) *SampleStruct {
	ret := new(SampleStruct)
	ret.internalData = iData
	ret.ExportData = eData
	ret.Version = "1.0"
	ret.DataType = "MyType"
	return ret
}

func (t *SampleStruct) GetInternal() string {
	return t.internalData
}
