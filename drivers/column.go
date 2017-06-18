package drivers

import (
	"strings"

	"github.com/volatiletech/sqlboiler/strmangle"
)

// Column holds information about a database column.
// Types are Go types, converted by TranslateColumnType.
type Column struct {
	Name      string
	Type      string
	DBType    string
	Default   string
	Nullable  bool
	Unique    bool
	Validated bool

	// Postgres only extension bits
	// ArrType is the underlying data type of the Postgres
	// ARRAY type. See here:
	// https://www.postgresql.org/docs/9.1/static/infoschema-element-types.html
	ArrType *string
	UDTName string

	// MySQL only bits
	// Used to get full type, ex:
	// tinyint(1) instead of tinyint
	// Used for "tinyint-as-bool" flag
	FullDBType string

	// MS SQL only bits
	// Used to indicate that the value
	// for this column is auto generated by database on insert (i.e. - timestamp (old) or rowversion (new))
	AutoGenerated bool
}

// ColumnNames of the columns.
func ColumnNames(cols []Column) []string {
	names := make([]string, len(cols))
	for i, c := range cols {
		names[i] = c.Name
	}

	return names
}

// ColumnDBTypes of the columns.
func ColumnDBTypes(cols []Column) map[string]string {
	types := map[string]string{}

	for _, c := range cols {
		types[strmangle.TitleCase(c.Name)] = c.DBType
	}

	return types
}

// FilterColumnsByAuto generates the list of columns that have autogenerated values
func FilterColumnsByAuto(auto bool, columns []Column) []Column {
	var cols []Column

	for _, c := range columns {
		if (auto && c.AutoGenerated) || (!auto && !c.AutoGenerated) {
			cols = append(cols, c)
		}
	}

	return cols
}

// FilterColumnsByDefault generates the list of columns that have default values
func FilterColumnsByDefault(defaults bool, columns []Column) []Column {
	var cols []Column

	for _, c := range columns {
		if (defaults && len(c.Default) != 0) || (!defaults && len(c.Default) == 0) {
			cols = append(cols, c)
		}
	}

	return cols
}

// FilterColumnsByEnum generates the list of columns that are enum values.
func FilterColumnsByEnum(columns []Column) []Column {
	var cols []Column

	for _, c := range columns {
		if strings.HasPrefix(c.DBType, "enum") {
			cols = append(cols, c)
		}
	}

	return cols
}
