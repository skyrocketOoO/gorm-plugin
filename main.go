package main

import (
	"fmt"
	"os"
	"reflect"

	"test/internal/model"

	"gorm.io/gorm"
)

func GenerateStruct(structType reflect.Type) string {
	var structFields []string

	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		fieldType := field.Type
		fieldName := field.Name

		// Handle gorm.Model fields specifically
		if fieldType == reflect.TypeOf(gorm.Model{}) {
			structFields = append(structFields, "ID")
			structFields = append(structFields, "CreatedAt")
			structFields = append(structFields, "UpdatedAt")
			structFields = append(structFields, "DeletedAt")
			continue
		}
		if field.Anonymous {
			tag := field.Tag
			tagVal := DecodeTag(tag.Get("gorm"))
			if !InSlice(tagVal.List, "embedded") {
				continue
			}
			if prefix, ok := tagVal.Map["embeddedPrefix"]; ok {
				for _, embeddedField := range genEmbedded(field) {
					structFields = append(structFields, fmt.Sprintf("%s%s", prefix, embeddedField))
				}
			}
			continue
		}
		if field.Type.Kind() == reflect.Struct {
			tag := field.Tag
			tagVal := DecodeTag(tag.Get("gorm"))
			if !InSlice(tagVal.List, "embedded") {
				continue
			}
			if prefix, ok := tagVal.Map["embeddedPrefix"]; ok {
				for _, embeddedField := range genEmbedded(field) {
					structFields = append(structFields, fmt.Sprintf("%s%s", prefix, embeddedField))
				}
			}
			continue
		}

		structFields = append(structFields, fieldName)
	}

	structBody := ""
	for _, field := range structFields {
		structBody += fmt.Sprintf("\t%s any\n", field)
	}
	return fmt.Sprintf("type %sQ struct {\n%s}", structType.Name(), structBody)
}

func genEmbedded(structType reflect.StructField) []string {
	var structFields []string
	for i := 0; i < structType.Type.NumField(); i++ {
		field := structType.Type.Field(i)
		fieldName := field.Name
		if field.Anonymous {
			tag := field.Tag
			tagVal := DecodeTag(tag.Get("gorm"))
			if !InSlice(tagVal.List, "embedded") {
				continue
			}
			if prefix, ok := tagVal.Map["embeddedPrefix"]; ok {
				for _, embeddedField := range genEmbedded(field) {
					structFields = append(structFields, fmt.Sprintf("%s%s", prefix, embeddedField))
				}
			}
			continue
		}
		if field.Type.Kind() == reflect.Struct {
			tag := field.Tag
			tagVal := DecodeTag(tag.Get("gorm"))
			if !InSlice(tagVal.List, "embedded") {
				continue
			}
			if prefix, ok := tagVal.Map["embeddedPrefix"]; ok {
				for _, embeddedField := range genEmbedded(field) {
					structFields = append(structFields, fmt.Sprintf("%s%s", prefix, embeddedField))
				}
			}
			continue
		}

		// Create field definition with interface{} type
		structFields = append(structFields, fieldName)
	}
	return structFields
}

func main() {
	// Get the type of the Account struct
	accountType := reflect.TypeOf(model.Account{})

	// Generate the new struct definition
	accountQ := GenerateStruct(accountType)

	// Create and write to the output file
	file, err := os.Create("account_query.go")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Write package declaration and struct to file
	file.WriteString("package main\n\n")
	file.WriteString(accountQ)

	fmt.Println("account_query.go generated successfully.")
}
