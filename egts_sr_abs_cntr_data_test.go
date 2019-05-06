package egts

import (
	"bytes"
	"github.com/google/go-cmp/cmp"
	"testing"
)

var (
	srAbsCntrDataBytes    = []byte{0x06, 0x75, 0x1D, 0x70}
	testEgtsSrAbsCntrData = SrAbsCntrData{
		CounterNumber: 6,
		CounterValue:  7347573,
	}
)

func TestEgtsSrAbsCntrData_Encode(t *testing.T) {
	posDataBytes, err := testEgtsSrAbsCntrData.Encode()
	if err != nil {
		t.Errorf("Ошибка кодирования: %v\n", err)
	}

	if !bytes.Equal(posDataBytes, srAbsCntrDataBytes) {
		t.Errorf("Байтовые строки не совпадают: %v != %v ", posDataBytes, srAbsCntrDataBytes)
	}
}

func TestEgtsSrAbsCntrData_Decode(t *testing.T) {
	adSensData := SrAbsCntrData{}

	if err := adSensData.Decode(srAbsCntrDataBytes); err != nil {
		t.Errorf("Ошибка декадирования: %v\n", err)
	}

	if diff := cmp.Diff(adSensData, testEgtsSrAbsCntrData); diff != "" {
		t.Errorf("Записи не совпадают: (-нужно +сейчас)\n%s", diff)
	}
}

// проверяем что рекордсет работает правильно с данным типом подзаписи
func TestEgtsSrAbsCntrDataRs(t *testing.T) {
	egtsSrAbsCntrDataRDBytes := append([]byte{0x19, 0x04, 0x00}, srAbsCntrDataBytes...)
	egtsSrAbsCntrDataRD := RecordDataSet{
		RecordData{
			SubrecordType:   SrAbsCntrDataType,
			SubrecordLength: testEgtsSrAbsCntrData.Length(),
			SubrecordData:   &testEgtsSrAbsCntrData,
		},
	}
	testStruct := RecordDataSet{}

	testBytes, err := egtsSrAbsCntrDataRD.Encode()
	if err != nil {
		t.Errorf("Ошибка кодирования: %v\n", err)
	}

	if !bytes.Equal(testBytes, egtsSrAbsCntrDataRDBytes) {
		t.Errorf("Байтовые строки не совпадают: %v != %v ", testBytes, egtsSrAbsCntrDataRDBytes)
	}

	if err = testStruct.Decode(egtsSrAbsCntrDataRDBytes); err != nil {
		t.Errorf("Ошибка декадирования: %v\n", err)
	}

	if diff := cmp.Diff(egtsSrAbsCntrDataRD, testStruct); diff != "" {
		t.Errorf("Записи не совпадают: (-нужно +сейчас)\n%s", diff)
	}
}
