package example

import (
	"testing"

	"github.com/infobloxopen/atlas-app-toolkit/v2/query"
)

func TestValidateFiltering(t *testing.T) {
	tests := []struct {
		Query string
		Err   bool
	}{
		{`first_name~"Sam"`, false},
		{`first_name=="Sam"`, false},
		{`weight==1`, false},
		{`comment=="comment1"`, false},
		{`speciality~"spec"`, false},
		{`last_name=="Smith"`, false},
		{`last_name~"Smith"`, false},
		{`id=="some_id"`, true},
		{`first_name<"Sam"`, true},
		{`weight<=1`, true},
		{`first_name<"Jan"`, true},
		{`speciality=="spec"`, true},
		{`unknown_field=="unk"`, true},
		{`id=="some_id"`, true},
		{`array=="array"`, true},
		{`custom_type=="custom_type"`, true},
		{`custom_type_string~"custom_type"`, false},
		{`home_address.city=="city"`, false},
		{`home_address.city~"city"`, true},
		{`home_address.country~"country"`, false},
		{`work_address.city=="city"`, true},
		{`home_address.country := "country"`, false},
		{`company := "company"`, true},
		{`home_address.country in ["country", "another_country"]`, false},
		{`nationality in ["American", "Сhinese"]`, true},
		{`weight in [10, 20, 30]`, false},
		{`comment in [10, 20, 30]`, true},
		{`boolean_field in ["true", "1", "t", "True", "TRUE", "false", "0", "f", "False", "FALSE"]`, false},
		{`boolean_field in ["tRuE"]`, true},
		{`boolean_field=="True"`, false},
		{`boolean_field=="Blah"`, true},
		{`boolean_field:="True"`, true},
		{`first_name~".*Sam"`, false},
		{`first_name~"*Sam"`, true},
		{`first_name~"*Sam" and last_name~"*Smith"`, true},
		{`first_name~"Sam" and last_name~"*Smith"`, true},
	}

	for _, test := range tests {
		f, err := query.ParseFiltering(test.Query)
		if err != nil {
			t.Fatalf("Invalid filtering data '%s'", test.Query)
		}
		err = ExampleValidateFiltering("/example.TestService/List", f)
		if err != nil {
			if test.Err == false {
				t.Errorf("Unexpected error for %s query: %s", test.Query, err)
			}
		} else {
			if test.Err == true {
				t.Errorf("Expected error for %s query, but got no error", test.Query)
			}
		}
	}
}

func TestValidateSorting(t *testing.T) {
	tests := []struct {
		Query string
		Err   bool
	}{
		{`first_name, weight`, false},
		{`comment`, false},
		{`on_vacation`, true},
		{`speciality`, true},
		{`unknown_field`, true},
		{`array`, true},
		{`custom_type`, true},
		{`custom_type_string`, false},
		{`home_address.city`, true},
		{`home_address.country`, false},
		{`work_address.city`, true},
	}

	for _, test := range tests {
		s, err := query.ParseSorting(test.Query)
		if err != nil {
			t.Fatalf("Invalid sorting data '%s'", test.Query)
		}
		err = ExampleValidateSorting("/example.TestService/List", s)
		if err != nil {
			if test.Err == false {
				t.Errorf("Unexpected error for %s query: %s", test.Query, err)
			}
		} else {
			if test.Err == true {
				t.Errorf("Expected error for %s query, but got no error", test.Query)
			}
		}
	}
}

func TestValidateFieldSelection(t *testing.T) {
	tests := []struct {
		Query string
		Err   bool
	}{
		{`first_name`, false},
		{`weight`, false},
		{`on_vacation`, false},
		{`speciality`, false},
		{`comment`, false},
		{`last_name`, false},
		{`id`, false},
		{`array`, false},
		{`custom_type`, false},
		{`custom_type_string`, false},
		{`home_address`, false},
		{`home_address.city`, false},
		{`home_address.country`, false},
		{`work_address`, false},
		{`work_address.city`, false},
		{`work_address.country`, false},
		{`first_name,weight,on_vacation`, false},
		{`unknown_field`, true},
		{`work_address.unknown_field`, true},
		{`last_name.value`, true},
		{`first_name,unknown_field`, true},
	}

	for _, test := range tests {
		fs := query.ParseFieldSelection(test.Query)
		err := ExampleValidateFieldSelection("/example.TestService/List", fs)
		if err != nil {
			if test.Err == false {
				t.Errorf("Unexpected error for %s query: %s", test.Query, err)
			}
		} else {
			if test.Err == true {
				t.Errorf("Expected error for %s query, but got no error", test.Query)
			}
		}
	}
}
