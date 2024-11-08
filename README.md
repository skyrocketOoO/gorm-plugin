# gorm-enhance-plugin

# Target
- Solve the hard code when use db.Where(field = ?) to avoid runtime error and support global modify
- When use db.Where(&struct{a: 0}) where a is int, it will be igonre because of zero value

## TODO
- Parse column with various case
- Output address customize
- Generate same structure(only set scalar type to interface{}) for where condition 
https://wiki.goframe.org/pages/viewpage.action?pageId=7296196