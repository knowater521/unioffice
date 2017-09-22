// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetml_test

import (
	"encoding/xml"
	"testing"

	"baliance.com/gooxml/schema/soo/spreadsheetml"
)

func TestCT_MergeCellConstructor(t *testing.T) {
	v := spreadsheetml.NewCT_MergeCell()
	if v == nil {
		t.Errorf("spreadsheetml.NewCT_MergeCell must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed spreadsheetml.CT_MergeCell should validate: %s", err)
	}
}

func TestCT_MergeCellMarshalUnmarshal(t *testing.T) {
	v := spreadsheetml.NewCT_MergeCell()
	buf, _ := xml.Marshal(v)
	v2 := spreadsheetml.NewCT_MergeCell()
	xml.Unmarshal(buf, v2)
}