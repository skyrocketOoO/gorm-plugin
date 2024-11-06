package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"test/internal/model"

	"gorm.io/gorm"
)

// GenerateStruct generates a struct definition with `interface{}` types
func GenerateStruct(structType reflect.Type) string {
	var structFields []string

	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		fieldType := field.Type

		// Handle gorm.Model fields specifically
		if fieldType == reflect.TypeOf(gorm.Model{}) {
			structFields = append(structFields, "\tID interface{}")
			structFields = append(structFields, "\tCreatedAt interface{}")
			structFields = append(structFields, "\tUpdatedAt interface{}")
			structFields = append(structFields, "\tDeletedAt interface{}")
			continue
		}
		if field.Anonymous {
			structFields = append(structFields, genEmbedded(field)...)
		}
		if field.Type.Kind() == reflect.Struct {
			continue
		}
		fieldName := field.Name

		// Create field definition with interface{} type
		structFields = append(structFields, fmt.Sprintf("\t%s interface{}", fieldName))
	}

	// Join fields to form the struct body
	structBody := strings.Join(structFields, "\n")
	return fmt.Sprintf("type %sQ struct {\n%s\n}", structType.Name(), structBody)
}

func genEmbedded(structType reflect.StructField) []string {
	var structFields []string
	for i := 0; i < structType.Type.NumField(); i++ {
		field := structType.Type.Field(i)
		// fmt.Println(field.Name)
		if field.Anonymous {
			structFields = append(structFields, genEmbedded(field)...)
		}
		if field.Type.Kind() == reflect.Struct {
			continue
		}
		fieldName := field.Name

		// Create field definition with interface{} type
		structFields = append(structFields, fmt.Sprintf("\t%s interface{}", fieldName))
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
