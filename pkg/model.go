package pkg

type SampleStruct struct {
	ExportData   string
	internalData string
}

func NewSampleStruct(iData string, eData string) *SampleStruct {
	ret := new(SampleStruct)
	ret.internalData = iData
	ret.ExportData = eData
	return ret
}
