// Code generated by gnorm, DO NOT EDIT!

package tables

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
	"github.com/sillydong/gnorm/database/drivers/mysql/gnorm"
)

// Row represents a row from 'TABLES'.
type Row struct {
	TableCatalog   string         // TABLE_CATALOG
	TableSchema    string         // TABLE_SCHEMA
	TableName      string         // TABLE_NAME
	TableType      string         // TABLE_TYPE
	Engine         sql.NullString // ENGINE
	Version        sql.NullInt64  // VERSION
	RowFormat      sql.NullString // ROW_FORMAT
	TableRows      sql.NullInt64  // TABLE_ROWS
	AvgRowLength   sql.NullInt64  // AVG_ROW_LENGTH
	DataLength     sql.NullInt64  // DATA_LENGTH
	MaxDataLength  sql.NullInt64  // MAX_DATA_LENGTH
	IndexLength    sql.NullInt64  // INDEX_LENGTH
	DataFree       sql.NullInt64  // DATA_FREE
	AutoIncrement  sql.NullInt64  // AUTO_INCREMENT
	CreateTime     mysql.NullTime // CREATE_TIME
	UpdateTime     mysql.NullTime // UPDATE_TIME
	CheckTime      mysql.NullTime // CHECK_TIME
	TableCollation sql.NullString // TABLE_COLLATION
	Checksum       sql.NullInt64  // CHECKSUM
	CreateOptions  sql.NullString // CREATE_OPTIONS
	TableComment   string         // TABLE_COMMENT
}

// Field values for every column in Tables.
var (
	TableCatalogCol   gnorm.StringField        = "TABLE_CATALOG"
	TableSchemaCol    gnorm.StringField        = "TABLE_SCHEMA"
	TableNameCol      gnorm.StringField        = "TABLE_NAME"
	TableTypeCol      gnorm.StringField        = "TABLE_TYPE"
	EngineCol         gnorm.SqlNullStringField = "ENGINE"
	VersionCol        gnorm.SqlNullInt64Field  = "VERSION"
	RowFormatCol      gnorm.SqlNullStringField = "ROW_FORMAT"
	TableRowsCol      gnorm.SqlNullInt64Field  = "TABLE_ROWS"
	AvgRowLengthCol   gnorm.SqlNullInt64Field  = "AVG_ROW_LENGTH"
	DataLengthCol     gnorm.SqlNullInt64Field  = "DATA_LENGTH"
	MaxDataLengthCol  gnorm.SqlNullInt64Field  = "MAX_DATA_LENGTH"
	IndexLengthCol    gnorm.SqlNullInt64Field  = "INDEX_LENGTH"
	DataFreeCol       gnorm.SqlNullInt64Field  = "DATA_FREE"
	AutoIncrementCol  gnorm.SqlNullInt64Field  = "AUTO_INCREMENT"
	CreateTimeCol     gnorm.MysqlNullTimeField = "CREATE_TIME"
	UpdateTimeCol     gnorm.MysqlNullTimeField = "UPDATE_TIME"
	CheckTimeCol      gnorm.MysqlNullTimeField = "CHECK_TIME"
	TableCollationCol gnorm.SqlNullStringField = "TABLE_COLLATION"
	ChecksumCol       gnorm.SqlNullInt64Field  = "CHECKSUM"
	CreateOptionsCol  gnorm.SqlNullStringField = "CREATE_OPTIONS"
	TableCommentCol   gnorm.StringField        = "TABLE_COMMENT"
)

// Query retrieves rows from 'TABLES' as a slice of Row.
func Query(db gnorm.DB, where gnorm.WhereClause) ([]*Row, error) {
	const origsqlstr = `SELECT 
		TABLE_CATALOG, TABLE_SCHEMA, TABLE_NAME, TABLE_TYPE, ENGINE, VERSION, ROW_FORMAT, TABLE_ROWS, AVG_ROW_LENGTH, DATA_LENGTH, MAX_DATA_LENGTH, INDEX_LENGTH, DATA_FREE, AUTO_INCREMENT, CREATE_TIME, UPDATE_TIME, CHECK_TIME, TABLE_COLLATION, CHECKSUM, CREATE_OPTIONS, TABLE_COMMENT
		FROM information_schema.TABLES WHERE (`

	sqlstr := origsqlstr + where.String() + ") "

	var vals []*Row
	q, err := db.Query(sqlstr, where.Values()...)
	if err != nil {
		return nil, err
	}
	for q.Next() {
		r := Row{}

		err = q.Scan(&r.TableCatalog, &r.TableSchema, &r.TableName, &r.TableType, &r.Engine, &r.Version, &r.RowFormat, &r.TableRows, &r.AvgRowLength, &r.DataLength, &r.MaxDataLength, &r.IndexLength, &r.DataFree, &r.AutoIncrement, &r.CreateTime, &r.UpdateTime, &r.CheckTime, &r.TableCollation, &r.Checksum, &r.CreateOptions, &r.TableComment)
		if err != nil {
			return nil, err
		}

		vals = append(vals, &r)
	}
	return vals, nil
}
