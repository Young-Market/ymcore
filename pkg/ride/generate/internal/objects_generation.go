package internal

import (
	"fmt"
	"sort"
	"strings"
)

func getType(types []string) string {
	if len(types) == 1 {
		return types[0]
	}
	return "rideType"
}

func GenerateObjects(fn string) {
	s, err := parseConfig()
	if err != nil {
		panic(err)
	}
	cd := NewCoder("ride")
	cd.Import("strings")
	cd.Import("github.com/pkg/errors")

	for _, act := range s.Actions {
		// Struct Implementation
		cd.Line("type ride%s struct {", act.StructName)
		for _, field := range act.Fields {
			cd.Line("%s %s", field.Name, getType(field.Types))
		}
		cd.Line("}")
		cd.Line("")

		// Constructor
		arguments := make([]string, len(act.Fields))
		for i, field := range act.Fields {
			arguments[i] = fmt.Sprintf("%s %s", field.Name, getType(field.Types))
		}
		cd.Line("func newRide%s(%s) ride%s {", act.StructName, strings.Join(arguments, ", "), act.StructName)
		cd.Line("return ride%s{", act.StructName)
		for _, field := range act.Fields {
			cd.Line("%s: %s,", field.Name, field.Name)
		}
		cd.Line("}")
		cd.Line("}")
		cd.Line("")

		// instanceOf method
		cd.Line("func (o ride%s) instanceOf() string {", act.StructName)
		cd.Line("return %sTypeName", act.Name)
		cd.Line("}")
		cd.Line("")

		// qe method
		cd.Line("func (o ride%s) eq(other rideType) bool {", act.StructName)
		cd.Line("if oo, ok := other.(ride%s); ok {", act.StructName)
		for _, field := range act.Fields {
			cd.Line("if !o.%s.eq(oo.%s) {", field.Name, field.Name)
			cd.Line("return false")
			cd.Line("}")
		}
		cd.Line("return true")
		cd.Line("}")
		cd.Line("return false")
		cd.Line("}")
		cd.Line("")

		// get method
		cd.Line("func (o ride%s) get(prop string) (rideType, error) {", act.StructName)
		cd.Line("switch prop {")
		cd.Line("case instanceField:")
		cd.Line("return rideString(%sTypeName), nil", act.Name)
		for _, field := range act.Fields {
			cd.Line("case %sField:", field.Name)
			cd.Line("return o.%s, nil", field.Name)
		}
		cd.Line("default:")
		cd.Line("return nil, errors.Errorf(\"type '%%s' has no property '%%s'\", o.instanceOf(), prop)")
		cd.Line("}")
		cd.Line("}")
		cd.Line("")

		//copy method
		for i, field := range act.Fields {
			arguments[i] = fmt.Sprintf("o.%s", field.Name)
		}
		cd.Line("func (o ride%s) copy() rideType {", act.StructName)
		cd.Line("return newRide%s(%s)", act.StructName, strings.Join(arguments, ", "))
		cd.Line("}")
		cd.Line("")

		// lines method
		cd.Line("func (o ride%s) lines() []string {", act.StructName)
		cd.Line("r := make([]string, 0, %d)", len(act.Fields)+2)
		cd.Line("r = append(r, %sTypeName + \"(\")", act.Name)
		sort.SliceStable(act.Fields, func(i, j int) bool {
			return act.Fields[i].Order < act.Fields[j].Order
		})
		for _, field := range act.Fields {
			if field.Order != -1 {
				cd.Line("r = append(r, fieldLines(%sField, o.%s.lines())...)", field.Name, field.Name)
			}
		}
		cd.Line("r = append(r, \")\")")
		cd.Line("return r")
		cd.Line("}")
		cd.Line("")

		// String method
		cd.Line("func (o ride%s) String() string {", act.StructName)
		cd.Line("return strings.Join(o.lines(), \"\\n\")")
		cd.Line("}")
		cd.Line("")

		// SetProofs (only for transactions)
		if act.SetProofs {
			cd.Line("func (o ride%s) setProofs(proofs rideList) rideProven {", act.StructName)
			cd.Line("o.proofs = proofs")
			cd.Line("return o")
			cd.Line("}")
			cd.Line("")
			cd.Line("func (o ride%s) getProofs() rideList {", act.StructName)
			cd.Line("return o.proofs")
			cd.Line("}")
			cd.Line("")
		}
	}
	// ResetProofs (only for transactions)
	cd.Line("func resetProofs(obj rideType) error {")
	cd.Line("switch tx := obj.(type) {")
	cd.Line("case rideProven:")
	cd.Line("tx.setProofs(rideList{})")
	cd.Line("default:")
	cd.Line("return errors.Errorf(\"type '%%s' is not tx\", obj.instanceOf())")
	cd.Line("}")
	cd.Line("return nil")
	cd.Line("}")
	cd.Line("")

	if err := cd.Save(fn); err != nil {
		panic(err)
	}
}
