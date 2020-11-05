package databases

import (
	fmt "fmt"

	github_com_eden_framework_sqlx "github.com/eden-framework/sqlx"
	github_com_eden_framework_sqlx_builder "github.com/eden-framework/sqlx/builder"
)

func (Token) PrimaryKey() []string {
	return []string{
		"ID",
	}
}

func (Token) UniqueIndexUTokenID() string {
	return "U_token_id"
}

func (Token) UniqueIndexes() github_com_eden_framework_sqlx_builder.Indexes {
	return github_com_eden_framework_sqlx_builder.Indexes{
		"U_token_id": []string{
			"TokenID",
		},
	}
}

func (Token) Comments() map[string]string {
	return map[string]string{
		"Audience": "使用方",
		"ExpireAt": "过期时间",
		"IssuedAt": "签发时间",
		"Issuer":   "签发方",
		"Subject":  "用途",
		"TokenID":  "业务ID",
	}
}

var TokenTable *github_com_eden_framework_sqlx_builder.Table

func init() {
	TokenTable = Config.DB.Register(&Token{})
}

type TokenIterator struct {
}

func (TokenIterator) New() interface{} {
	return &Token{}
}

func (TokenIterator) Resolve(v interface{}) *Token {
	return v.(*Token)
}

func (Token) TableName() string {
	return "t_token"
}

func (Token) ColDescriptions() map[string][]string {
	return map[string][]string{
		"Audience": []string{
			"使用方",
		},
		"ExpireAt": []string{
			"过期时间",
		},
		"IssuedAt": []string{
			"签发时间",
		},
		"Issuer": []string{
			"签发方",
		},
		"Subject": []string{
			"用途",
		},
		"TokenID": []string{
			"业务ID",
		},
	}
}

func (Token) FieldKeyID() string {
	return "ID"
}

func (m *Token) FieldID() *github_com_eden_framework_sqlx_builder.Column {
	return TokenTable.F(m.FieldKeyID())
}

func (Token) FieldKeyTokenID() string {
	return "TokenID"
}

func (m *Token) FieldTokenID() *github_com_eden_framework_sqlx_builder.Column {
	return TokenTable.F(m.FieldKeyTokenID())
}

func (Token) FieldKeyIssuer() string {
	return "Issuer"
}

func (m *Token) FieldIssuer() *github_com_eden_framework_sqlx_builder.Column {
	return TokenTable.F(m.FieldKeyIssuer())
}

func (Token) FieldKeySubject() string {
	return "Subject"
}

func (m *Token) FieldSubject() *github_com_eden_framework_sqlx_builder.Column {
	return TokenTable.F(m.FieldKeySubject())
}

func (Token) FieldKeyAudience() string {
	return "Audience"
}

func (m *Token) FieldAudience() *github_com_eden_framework_sqlx_builder.Column {
	return TokenTable.F(m.FieldKeyAudience())
}

func (Token) FieldKeyComment() string {
	return "Comment"
}

func (m *Token) FieldComment() *github_com_eden_framework_sqlx_builder.Column {
	return TokenTable.F(m.FieldKeyComment())
}

func (Token) FieldKeyIssuedAt() string {
	return "IssuedAt"
}

func (m *Token) FieldIssuedAt() *github_com_eden_framework_sqlx_builder.Column {
	return TokenTable.F(m.FieldKeyIssuedAt())
}

func (Token) FieldKeyExpireAt() string {
	return "ExpireAt"
}

func (m *Token) FieldExpireAt() *github_com_eden_framework_sqlx_builder.Column {
	return TokenTable.F(m.FieldKeyExpireAt())
}

func (Token) ColRelations() map[string][]string {
	return map[string][]string{}
}

func (m *Token) IndexFieldNames() []string {
	return []string{
		"ID",
		"TokenID",
	}
}

func (m *Token) ConditionByStruct(db github_com_eden_framework_sqlx.DBExecutor) github_com_eden_framework_sqlx_builder.SqlCondition {
	table := db.T(m)
	fieldValues := github_com_eden_framework_sqlx_builder.FieldValuesFromStructByNonZero(m)

	conditions := make([]github_com_eden_framework_sqlx_builder.SqlCondition, 0)

	for _, fieldName := range m.IndexFieldNames() {
		if v, exists := fieldValues[fieldName]; exists {
			conditions = append(conditions, table.F(fieldName).Eq(v))
			delete(fieldValues, fieldName)
		}
	}

	if len(conditions) == 0 {
		panic(fmt.Errorf("at least one of field for indexes has value"))
	}

	for fieldName, v := range fieldValues {
		conditions = append(conditions, table.F(fieldName).Eq(v))
	}

	condition := github_com_eden_framework_sqlx_builder.And(conditions...)

	return condition
}

func (m *Token) Create(db github_com_eden_framework_sqlx.DBExecutor) error {

	_, err := db.ExecExpr(github_com_eden_framework_sqlx.InsertToDB(db, m, nil))
	return err

}

func (m *Token) CreateOnDuplicateWithUpdateFields(db github_com_eden_framework_sqlx.DBExecutor, updateFields []string) error {

	if len(updateFields) == 0 {
		panic(fmt.Errorf("must have update fields"))
	}

	fieldValues := github_com_eden_framework_sqlx_builder.FieldValuesFromStructByNonZero(m, updateFields...)

	delete(fieldValues, "ID")

	table := db.T(m)

	cols, vals := table.ColumnsAndValuesByFieldValues(fieldValues)

	fields := make(map[string]bool, len(updateFields))
	for _, field := range updateFields {
		fields[field] = true
	}

	for _, fieldNames := range m.UniqueIndexes() {
		for _, field := range fieldNames {
			delete(fields, field)
		}
	}

	if len(fields) == 0 {
		panic(fmt.Errorf("no fields for updates"))
	}

	for field := range fieldValues {
		if !fields[field] {
			delete(fieldValues, field)
		}
	}

	additions := github_com_eden_framework_sqlx_builder.Additions{}

	switch db.Dialect().DriverName() {
	case "mysql":
		additions = append(additions, github_com_eden_framework_sqlx_builder.OnDuplicateKeyUpdate(table.AssignmentsByFieldValues(fieldValues)...))
	case "postgres":
		indexes := m.UniqueIndexes()
		fields := make([]string, 0)
		for _, fs := range indexes {
			fields = append(fields, fs...)
		}
		indexFields, _ := db.T(m).Fields(fields...)

		additions = append(additions,
			github_com_eden_framework_sqlx_builder.OnConflict(indexFields).
				DoUpdateSet(table.AssignmentsByFieldValues(fieldValues)...))
	}

	additions = append(additions, github_com_eden_framework_sqlx_builder.Comment("User.CreateOnDuplicateWithUpdateFields"))

	expr := github_com_eden_framework_sqlx_builder.Insert().Into(table, additions...).Values(cols, vals...)

	_, err := db.ExecExpr(expr)
	return err

}

func (m *Token) DeleteByStruct(db github_com_eden_framework_sqlx.DBExecutor) error {

	_, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Delete().
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(m.ConditionByStruct(db)),
				github_com_eden_framework_sqlx_builder.Comment("Token.DeleteByStruct"),
			),
	)

	return err
}

func (m *Token) FetchByID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_eden_framework_sqlx_builder.Select(nil).
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("ID").Eq(m.ID),
				)),
				github_com_eden_framework_sqlx_builder.Comment("Token.FetchByID"),
			),
		m,
	)

	return err
}

func (m *Token) UpdateByIDWithMap(db github_com_eden_framework_sqlx.DBExecutor, fieldValues github_com_eden_framework_sqlx_builder.FieldValues) error {

	table := db.T(m)

	result, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Update(db.T(m)).
			Where(
				github_com_eden_framework_sqlx_builder.And(
					table.F("ID").Eq(m.ID),
				),
				github_com_eden_framework_sqlx_builder.Comment("Token.UpdateByIDWithMap"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return m.FetchByID(db)
	}

	return nil

}

func (m *Token) UpdateByIDWithStruct(db github_com_eden_framework_sqlx.DBExecutor, zeroFields ...string) error {

	fieldValues := github_com_eden_framework_sqlx_builder.FieldValuesFromStructByNonZero(m, zeroFields...)
	return m.UpdateByIDWithMap(db, fieldValues)

}

func (m *Token) FetchByIDForUpdate(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_eden_framework_sqlx_builder.Select(nil).
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("ID").Eq(m.ID),
				)),
				github_com_eden_framework_sqlx_builder.ForUpdate(),
				github_com_eden_framework_sqlx_builder.Comment("Token.FetchByIDForUpdate"),
			),
		m,
	)

	return err
}

func (m *Token) DeleteByID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	_, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Delete().
			From(db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("ID").Eq(m.ID),
				)),
				github_com_eden_framework_sqlx_builder.Comment("Token.DeleteByID"),
			))

	return err
}

func (m *Token) FetchByTokenID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_eden_framework_sqlx_builder.Select(nil).
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("TokenID").Eq(m.TokenID),
				)),
				github_com_eden_framework_sqlx_builder.Comment("Token.FetchByTokenID"),
			),
		m,
	)

	return err
}

func (m *Token) UpdateByTokenIDWithMap(db github_com_eden_framework_sqlx.DBExecutor, fieldValues github_com_eden_framework_sqlx_builder.FieldValues) error {

	table := db.T(m)

	result, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Update(db.T(m)).
			Where(
				github_com_eden_framework_sqlx_builder.And(
					table.F("TokenID").Eq(m.TokenID),
				),
				github_com_eden_framework_sqlx_builder.Comment("Token.UpdateByTokenIDWithMap"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return m.FetchByTokenID(db)
	}

	return nil

}

func (m *Token) UpdateByTokenIDWithStruct(db github_com_eden_framework_sqlx.DBExecutor, zeroFields ...string) error {

	fieldValues := github_com_eden_framework_sqlx_builder.FieldValuesFromStructByNonZero(m, zeroFields...)
	return m.UpdateByTokenIDWithMap(db, fieldValues)

}

func (m *Token) FetchByTokenIDForUpdate(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_eden_framework_sqlx_builder.Select(nil).
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("TokenID").Eq(m.TokenID),
				)),
				github_com_eden_framework_sqlx_builder.ForUpdate(),
				github_com_eden_framework_sqlx_builder.Comment("Token.FetchByTokenIDForUpdate"),
			),
		m,
	)

	return err
}

func (m *Token) DeleteByTokenID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	_, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Delete().
			From(db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("TokenID").Eq(m.TokenID),
				)),
				github_com_eden_framework_sqlx_builder.Comment("Token.DeleteByTokenID"),
			))

	return err
}

func (m *Token) List(db github_com_eden_framework_sqlx.DBExecutor, condition github_com_eden_framework_sqlx_builder.SqlCondition, additions ...github_com_eden_framework_sqlx_builder.Addition) ([]Token, error) {

	list := make([]Token, 0)

	table := db.T(m)
	_ = table

	finalAdditions := []github_com_eden_framework_sqlx_builder.Addition{
		github_com_eden_framework_sqlx_builder.Where(condition),
		github_com_eden_framework_sqlx_builder.Comment("Token.List"),
	}

	if len(additions) > 0 {
		finalAdditions = append(finalAdditions, additions...)
	}

	err := db.QueryExprAndScan(
		github_com_eden_framework_sqlx_builder.Select(nil).
			From(db.T(m), finalAdditions...),
		&list,
	)

	return list, err

}

func (m *Token) Count(db github_com_eden_framework_sqlx.DBExecutor, condition github_com_eden_framework_sqlx_builder.SqlCondition, additions ...github_com_eden_framework_sqlx_builder.Addition) (int, error) {

	count := -1

	table := db.T(m)
	_ = table

	finalAdditions := []github_com_eden_framework_sqlx_builder.Addition{
		github_com_eden_framework_sqlx_builder.Where(condition),
		github_com_eden_framework_sqlx_builder.Comment("Token.Count"),
	}

	if len(additions) > 0 {
		finalAdditions = append(finalAdditions, additions...)
	}

	err := db.QueryExprAndScan(
		github_com_eden_framework_sqlx_builder.Select(
			github_com_eden_framework_sqlx_builder.Count(),
		).
			From(db.T(m), finalAdditions...),
		&count,
	)

	return count, err

}

func (m *Token) BatchFetchByIDList(db github_com_eden_framework_sqlx.DBExecutor, values []uint64) ([]Token, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("ID").In(values)

	return m.List(db, condition)

}

func (m *Token) BatchFetchByTokenIDList(db github_com_eden_framework_sqlx.DBExecutor, values []uint64) ([]Token, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("TokenID").In(values)

	return m.List(db, condition)

}
