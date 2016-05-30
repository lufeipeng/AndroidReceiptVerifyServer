package models

import (
    "testing"
    "encoding/xml"
    "fmt"
)

func TestXYZ(t *testing.T) {
	var billResponse models.BillResponse
	billResponse.Hret  = "1";
	billResponse.Message  = "successful";
	
	resultBytes, _ := xml.MarshalIndent(&billResponse, "  ", "  ");
	fmt.Printf(xml.Header);
	fmt.Printf(string(resultBytes));
}
