package example

import (
	query "github.com/infobloxopen/atlas-app-toolkit/v2/query"
	options "github.com/infobloxopen/protoc-gen-atlas-query-validate/options"
)

var ExampleMethodsRequireFilteringValidation = map[string]map[string]options.FilteringOption{
	"/example.TestService/List": map[string]options.FilteringOption{
		"custom_search_2":                options.FilteringOption{Deny: []options.QueryValidate_FilterOperator{options.QueryValidate_GT, options.QueryValidate_GE, options.QueryValidate_LT, options.QueryValidate_LE, options.QueryValidate_IN, options.QueryValidate_IEQ}, ValueType: options.QueryValidate_STRING},
		"custom_search.city":             options.FilteringOption{Deny: []options.QueryValidate_FilterOperator{options.QueryValidate_MATCH, options.QueryValidate_GT, options.QueryValidate_GE, options.QueryValidate_LT, options.QueryValidate_LE, options.QueryValidate_IN, options.QueryValidate_IEQ}, ValueType: options.QueryValidate_STRING},
		"custom_search.country":          options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"list_of_addresses.city":         options.FilteringOption{Deny: []options.QueryValidate_FilterOperator{options.QueryValidate_MATCH, options.QueryValidate_GT, options.QueryValidate_GE, options.QueryValidate_LT, options.QueryValidate_LE, options.QueryValidate_IN, options.QueryValidate_IEQ}, ValueType: options.QueryValidate_STRING},
		"list_of_addresses.country":      options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"user_friend.custom_search_2":    options.FilteringOption{Deny: []options.QueryValidate_FilterOperator{options.QueryValidate_GT, options.QueryValidate_GE, options.QueryValidate_LT, options.QueryValidate_LE, options.QueryValidate_IN, options.QueryValidate_IEQ}, ValueType: options.QueryValidate_STRING},
		"user_friend.first_name":         options.FilteringOption{Deny: []options.QueryValidate_FilterOperator{options.QueryValidate_GT, options.QueryValidate_GE, options.QueryValidate_LT, options.QueryValidate_LE, options.QueryValidate_IN, options.QueryValidate_IEQ}, ValueType: options.QueryValidate_STRING},
		"user_friend.weight":             options.FilteringOption{Deny: []options.QueryValidate_FilterOperator{options.QueryValidate_LE}, ValueType: options.QueryValidate_NUMBER},
		"user_friend.on_vacation":        options.FilteringOption{ValueType: options.QueryValidate_BOOL},
		"user_friend.speciality":         options.FilteringOption{Deny: []options.QueryValidate_FilterOperator{options.QueryValidate_EQ, options.QueryValidate_GT, options.QueryValidate_GE, options.QueryValidate_LT, options.QueryValidate_LE, options.QueryValidate_IN, options.QueryValidate_IEQ}, ValueType: options.QueryValidate_STRING},
		"user_friend.comment":            options.FilteringOption{Deny: []options.QueryValidate_FilterOperator{options.QueryValidate_IN}, ValueType: options.QueryValidate_STRING},
		"user_friend.last_name":          options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"user_friend.id":                 options.FilteringOption{Deny: []options.QueryValidate_FilterOperator{options.QueryValidate_ALL}, ValueType: options.QueryValidate_STRING},
		"user_friend.array":              options.FilteringOption{Deny: []options.QueryValidate_FilterOperator{options.QueryValidate_ALL}, ValueType: options.QueryValidate_DEFAULT},
		"user_friend.custom_type_string": options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"user_friend.company":            options.FilteringOption{Deny: []options.QueryValidate_FilterOperator{options.QueryValidate_IEQ}, ValueType: options.QueryValidate_STRING},
		"user_friend.nationality":        options.FilteringOption{Deny: []options.QueryValidate_FilterOperator{options.QueryValidate_IN}, ValueType: options.QueryValidate_STRING},
		"user_friend.boolean_field":      options.FilteringOption{ValueType: options.QueryValidate_BOOL},
		"first_name":                     options.FilteringOption{Deny: []options.QueryValidate_FilterOperator{options.QueryValidate_GT, options.QueryValidate_GE, options.QueryValidate_LT, options.QueryValidate_LE, options.QueryValidate_IN, options.QueryValidate_IEQ}, ValueType: options.QueryValidate_STRING},
		"weight":                         options.FilteringOption{Deny: []options.QueryValidate_FilterOperator{options.QueryValidate_LE}, ValueType: options.QueryValidate_NUMBER},
		"on_vacation":                    options.FilteringOption{ValueType: options.QueryValidate_BOOL},
		"speciality":                     options.FilteringOption{Deny: []options.QueryValidate_FilterOperator{options.QueryValidate_EQ, options.QueryValidate_GT, options.QueryValidate_GE, options.QueryValidate_LT, options.QueryValidate_LE, options.QueryValidate_IN, options.QueryValidate_IEQ}, ValueType: options.QueryValidate_STRING},
		"comment":                        options.FilteringOption{Deny: []options.QueryValidate_FilterOperator{options.QueryValidate_IN}, ValueType: options.QueryValidate_STRING},
		"last_name":                      options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"id":                             options.FilteringOption{Deny: []options.QueryValidate_FilterOperator{options.QueryValidate_ALL}, ValueType: options.QueryValidate_STRING},
		"array":                          options.FilteringOption{Deny: []options.QueryValidate_FilterOperator{options.QueryValidate_ALL}, ValueType: options.QueryValidate_DEFAULT},
		"custom_type.name":               options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"custom_type_string":             options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"home_address.city":              options.FilteringOption{Deny: []options.QueryValidate_FilterOperator{options.QueryValidate_MATCH, options.QueryValidate_GT, options.QueryValidate_GE, options.QueryValidate_LT, options.QueryValidate_LE, options.QueryValidate_IN, options.QueryValidate_IEQ}, ValueType: options.QueryValidate_STRING},
		"home_address.country":           options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"work_address":                   options.FilteringOption{Deny: []options.QueryValidate_FilterOperator{options.QueryValidate_ALL}, ValueType: options.QueryValidate_DEFAULT},
		"company":                        options.FilteringOption{Deny: []options.QueryValidate_FilterOperator{options.QueryValidate_IEQ}, ValueType: options.QueryValidate_STRING},
		"nationality":                    options.FilteringOption{Deny: []options.QueryValidate_FilterOperator{options.QueryValidate_IN}, ValueType: options.QueryValidate_STRING},
		"boolean_field":                  options.FilteringOption{ValueType: options.QueryValidate_BOOL},
		"non_object":                     options.FilteringOption{Deny: []options.QueryValidate_FilterOperator{options.QueryValidate_ALL}, ValueType: options.QueryValidate_DEFAULT},
	},
}
var ExampleMethodsRequireSortingValidation = map[string][]string{
	"/example.TestService/List": []string{
		"first_name",
		"weight",
		"comment",
		"last_name",
		"id",
		"custom_type.name",
		"custom_type_string",
		"home_address.country",
		"company",
		"nationality",
		"boolean_field",
	},
	"/example.TestService/Read": []string{
		"first_name",
		"weight",
		"comment",
		"last_name",
		"id",
		"custom_type.name",
		"custom_type_string",
		"home_address.country",
		"company",
		"nationality",
		"boolean_field",
	},
}
var ExampleMethodsRequireFieldSelectionValidation = map[string][]string{
	"/example.TestService/List": {
		"list_of_addresses.city",
		"list_of_addresses.country",
		"list_of_addresses",
		"first_name",
		"weight",
		"on_vacation",
		"speciality",
		"comment",
		"last_name",
		"id",
		"array",
		"custom_type.name",
		"custom_type",
		"custom_type_string",
		"home_address.city",
		"home_address.country",
		"home_address",
		"work_address.city",
		"work_address.country",
		"work_address",
		"company",
		"nationality",
		"boolean_field",
		"non_object",
	},
	"/example.TestService/Read": {
		"list_of_addresses.city",
		"list_of_addresses.country",
		"list_of_addresses",
		"first_name",
		"weight",
		"on_vacation",
		"speciality",
		"comment",
		"last_name",
		"id",
		"array",
		"custom_type.name",
		"custom_type",
		"custom_type_string",
		"home_address.city",
		"home_address.country",
		"home_address",
		"work_address.city",
		"work_address.country",
		"work_address",
		"company",
		"nationality",
		"boolean_field",
		"non_object",
	},
}

func ExampleValidateFiltering(methodName string, f *query.Filtering) error {
	info, ok := ExampleMethodsRequireFilteringValidation[methodName]
	if !ok {
		return nil
	}
	return options.ValidateFiltering(f, info)
}
func ExampleValidateSorting(methodName string, s *query.Sorting) error {
	info, ok := ExampleMethodsRequireSortingValidation[methodName]
	if !ok {
		return nil
	}
	return options.ValidateSorting(s, info)
}
func ExampleValidateFieldSelection(methodName string, s *query.FieldSelection) error {
	info, ok := ExampleMethodsRequireFieldSelectionValidation[methodName]
	if !ok {
		return nil
	}
	return options.ValidateFieldSelection(s, info)
}
