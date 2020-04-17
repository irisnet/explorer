package document

import "testing"

func TestExStaticValidatorMonth_GetValidatorStaticByMonth(t *testing.T) {
	res, err := new(ExStaticValidatorMonth).GetValidatorStaticByMonth("2020.04", "fva186qhtc62cf6ejlt3erw6zk28mgw8ne7grhmyfn")
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(res)
}
