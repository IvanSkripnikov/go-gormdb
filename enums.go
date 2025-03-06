package gormdb

const (
	CreateTable = iota
	DropTable
	AddColumn
	AlterColumn
	DropColumn
)

var MigrationTypes = map[int]string{
	CreateTable: "Create table",
	DropTable:   "Drop table",
	AddColumn:   "Add column",
	AlterColumn: "Alter column",
	DropColumn:  "Drop column",
}
