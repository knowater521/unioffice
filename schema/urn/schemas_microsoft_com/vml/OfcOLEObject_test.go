// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io website.

package vml_test

import (
	"encoding/xml"
	"testing"

	"github.com/unidoc/unioffice/schema/urn/schemas_microsoft_com/vml"
)

func TestOfcOLEObjectConstructor(t *testing.T) {
	v := vml.NewOfcOLEObject()
	if v == nil {
		t.Errorf("vml.NewOfcOLEObject must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.OfcOLEObject should validate: %s", err)
	}
}

func TestOfcOLEObjectMarshalUnmarshal(t *testing.T) {
	v := vml.NewOfcOLEObject()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewOfcOLEObject()
	xml.Unmarshal(buf, v2)
}
