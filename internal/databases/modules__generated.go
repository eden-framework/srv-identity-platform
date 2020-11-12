package databases

import (
	fmt "fmt"
	time "time"

	github_com_eden_framework_sqlx "github.com/eden-framework/sqlx"
	github_com_eden_framework_sqlx_builder "github.com/eden-framework/sqlx/builder"
	github_com_eden_framework_sqlx_datatypes "github.com/eden-framework/sqlx/datatypes"
)

func (Modules) PrimaryKey() []string {
	return []string{
		"ID",
	}
}

func (Modules) Indexes() github_com_eden_framework_sqlx_builder.Indexes {
	return github_com_eden_framework_sqlx_builder.Indexes{
		"I_service": []string{
			"ServiceID",
		},
	}
}

func (Modules) UniqueIndexUModuleID() string {
	return "U_module_id"
}

func (Modules) UniqueIndexes() github_com_eden_framework_sqlx_builder.Indexes {
	return github_com_eden_framework_sqlx_builder.Indexes{
		"U_module_id": []string{
			"ModuleID",
			"DeletedAt",
		},
	}
}

func (Modules) Comments() map[string]string {
	return map[string]string{
		"Comment":   "描述",
		"ModuleID":  "业务ID",
		"ModuleKey": "模块标识",
		"Name":      "模块名称",
		"ServiceID": "所属服务",
	}
}

var ModulesTable *github_com_eden_framework_sqlx_builder.Table

func init() {
	ModulesTable = Config.DB.Register(&Modules{})
}

type ModulesIterator struct {
}

func (ModulesIterator) New() interface{} {
	return &Modules{}
}

func (ModulesIterator) Resolve(v interface{}) *Modules {
	return v.(*Modules)
}

func (Modules) TableName() string {
	return "t_modules"
}

func (Modules) ColDescriptions() map[string][]string {
	return map[string][]string{
		"Comment": []string{
			"描述",
		},
		"ModuleID": []string{
			"业务ID",
		},
		"ModuleKey": []string{
			"模块标识",
		},
		"Name": []string{
			"模块名称",
		},
		"ServiceID": []string{
			"所属服务",
		},
	}
}

func (Modules) FieldKeyID() string {
	return "ID"
}

func (m *Modules) FieldID() *github_com_eden_framework_sqlx_builder.Column {
	return ModulesTable.F(m.FieldKeyID())
}

func (Modules) FieldKeyModuleID() string {
	return "ModuleID"
}

func (m *Modules) FieldModuleID() *github_com_eden_framework_sqlx_builder.Column {
	return ModulesTable.F(m.FieldKeyModuleID())
}

func (Modules) FieldKeyModuleKey() string {
	return "ModuleKey"
}

func (m *Modules) FieldModuleKey() *github_com_eden_framework_sqlx_builder.Column {
	return ModulesTable.F(m.FieldKeyModuleKey())
}

func (Modules) FieldKeyName() string {
	return "Name"
}

func (m *Modules) FieldName() *github_com_eden_framework_sqlx_builder.Column {
	return ModulesTable.F(m.FieldKeyName())
}

func (Modules) FieldKeyComment() string {
	return "Comment"
}

func (m *Modules) FieldComment() *github_com_eden_framework_sqlx_builder.Column {
	return ModulesTable.F(m.FieldKeyComment())
}

func (Modules) FieldKeyServiceID() string {
	return "ServiceID"
}

func (m *Modules) FieldServiceID() *github_com_eden_framework_sqlx_builder.Column {
	return ModulesTable.F(m.FieldKeyServiceID())
}

func (Modules) FieldKeyCreatedAt() string {
	return "CreatedAt"
}

func (m *Modules) FieldCreatedAt() *github_com_eden_framework_sqlx_builder.Column {
	return ModulesTable.F(m.FieldKeyCreatedAt())
}

func (Modules) FieldKeyUpdatedAt() string {
	return "UpdatedAt"
}

func (m *Modules) FieldUpdatedAt() *github_com_eden_framework_sqlx_builder.Column {
	return ModulesTable.F(m.FieldKeyUpdatedAt())
}

func (Modules) FieldKeyDeletedAt() string {
	return "DeletedAt"
}

func (m *Modules) FieldDeletedAt() *github_com_eden_framework_sqlx_builder.Column {
	return ModulesTable.F(m.FieldKeyDeletedAt())
}

func (Modules) ColRelations() map[string][]string {
	return map[string][]string{}
}

func (m *Modules) IndexFieldNames() []string {
	return []string{
		"ID",
		"ModuleID",
		"ServiceID",
	}
}

func (m *Modules) ConditionByStruct(db github_com_eden_framework_sqlx.DBExecutor) github_com_eden_framework_sqlx_builder.SqlCondition {
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

	condition = github_com_eden_framework_sqlx_builder.And(condition, table.F("DeletedAt").Eq(0))
	return condition
}

func (m *Modules) Create(db github_com_eden_framework_sqlx.DBExecutor) error {

	if m.CreatedAt.IsZero() {
		m.CreatedAt = github_com_eden_framework_sqlx_datatypes.Timestamp(time.Now())
	}

	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = github_com_eden_framework_sqlx_datatypes.Timestamp(time.Now())
	}

	_, err := db.ExecExpr(github_com_eden_framework_sqlx.InsertToDB(db, m, nil))
	return err

}

func (m *Modules) CreateOnDuplicateWithUpdateFields(db github_com_eden_framework_sqlx.DBExecutor, updateFields []string) error {

	if len(updateFields) == 0 {
		panic(fmt.Errorf("must have update fields"))
	}

	if m.CreatedAt.IsZero() {
		m.CreatedAt = github_com_eden_framework_sqlx_datatypes.Timestamp(time.Now())
	}

	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = github_com_eden_framework_sqlx_datatypes.Timestamp(time.Now())
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

func (m *Modules) DeleteByStruct(db github_com_eden_framework_sqlx.DBExecutor) error {

	_, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Delete().
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(m.ConditionByStruct(db)),
				github_com_eden_framework_sqlx_builder.Comment("Modules.DeleteByStruct"),
			),
	)

	return err
}

func (m *Modules) FetchByID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_eden_framework_sqlx_builder.Select(nil).
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("ID").Eq(m.ID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("Modules.FetchByID"),
			),
		m,
	)

	return err
}

func (m *Modules) UpdateByIDWithMap(db github_com_eden_framework_sqlx.DBExecutor, fieldValues github_com_eden_framework_sqlx_builder.FieldValues) error {

	if _, ok := fieldValues["UpdatedAt"]; !ok {
		fieldValues["UpdatedAt"] = github_com_eden_framework_sqlx_datatypes.Timestamp(time.Now())
	}

	table := db.T(m)

	result, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Update(db.T(m)).
			Where(
				github_com_eden_framework_sqlx_builder.And(
					table.F("ID").Eq(m.ID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				),
				github_com_eden_framework_sqlx_builder.Comment("Modules.UpdateByIDWithMap"),
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

func (m *Modules) UpdateByIDWithStruct(db github_com_eden_framework_sqlx.DBExecutor, zeroFields ...string) error {

	fieldValues := github_com_eden_framework_sqlx_builder.FieldValuesFromStructByNonZero(m, zeroFields...)
	return m.UpdateByIDWithMap(db, fieldValues)

}

func (m *Modules) FetchByIDForUpdate(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_eden_framework_sqlx_builder.Select(nil).
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("ID").Eq(m.ID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.ForUpdate(),
				github_com_eden_framework_sqlx_builder.Comment("Modules.FetchByIDForUpdate"),
			),
		m,
	)

	return err
}

func (m *Modules) DeleteByID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	_, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Delete().
			From(db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("ID").Eq(m.ID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("Modules.DeleteByID"),
			))

	return err
}

func (m *Modules) SoftDeleteByID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	fieldValues := github_com_eden_framework_sqlx_builder.FieldValues{}
	if _, ok := fieldValues["DeletedAt"]; !ok {
		fieldValues["DeletedAt"] = github_com_eden_framework_sqlx_datatypes.Timestamp(time.Now())
	}

	if _, ok := fieldValues["UpdatedAt"]; !ok {
		fieldValues["UpdatedAt"] = github_com_eden_framework_sqlx_datatypes.Timestamp(time.Now())
	}

	_, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Update(db.T(m)).
			Where(
				github_com_eden_framework_sqlx_builder.And(
					table.F("ID").Eq(m.ID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				),
				github_com_eden_framework_sqlx_builder.Comment("Modules.SoftDeleteByID"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	return err

}

func (m *Modules) FetchByModuleID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_eden_framework_sqlx_builder.Select(nil).
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("ModuleID").Eq(m.ModuleID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("Modules.FetchByModuleID"),
			),
		m,
	)

	return err
}

func (m *Modules) UpdateByModuleIDWithMap(db github_com_eden_framework_sqlx.DBExecutor, fieldValues github_com_eden_framework_sqlx_builder.FieldValues) error {

	if _, ok := fieldValues["UpdatedAt"]; !ok {
		fieldValues["UpdatedAt"] = github_com_eden_framework_sqlx_datatypes.Timestamp(time.Now())
	}

	table := db.T(m)

	result, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Update(db.T(m)).
			Where(
				github_com_eden_framework_sqlx_builder.And(
					table.F("ModuleID").Eq(m.ModuleID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				),
				github_com_eden_framework_sqlx_builder.Comment("Modules.UpdateByModuleIDWithMap"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return m.FetchByModuleID(db)
	}

	return nil

}

func (m *Modules) UpdateByModuleIDWithStruct(db github_com_eden_framework_sqlx.DBExecutor, zeroFields ...string) error {

	fieldValues := github_com_eden_framework_sqlx_builder.FieldValuesFromStructByNonZero(m, zeroFields...)
	return m.UpdateByModuleIDWithMap(db, fieldValues)

}

func (m *Modules) FetchByModuleIDForUpdate(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_eden_framework_sqlx_builder.Select(nil).
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("ModuleID").Eq(m.ModuleID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.ForUpdate(),
				github_com_eden_framework_sqlx_builder.Comment("Modules.FetchByModuleIDForUpdate"),
			),
		m,
	)

	return err
}

func (m *Modules) DeleteByModuleID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	_, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Delete().
			From(db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("ModuleID").Eq(m.ModuleID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("Modules.DeleteByModuleID"),
			))

	return err
}

func (m *Modules) SoftDeleteByModuleID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	fieldValues := github_com_eden_framework_sqlx_builder.FieldValues{}
	if _, ok := fieldValues["DeletedAt"]; !ok {
		fieldValues["DeletedAt"] = github_com_eden_framework_sqlx_datatypes.Timestamp(time.Now())
	}

	if _, ok := fieldValues["UpdatedAt"]; !ok {
		fieldValues["UpdatedAt"] = github_com_eden_framework_sqlx_datatypes.Timestamp(time.Now())
	}

	_, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Update(db.T(m)).
			Where(
				github_com_eden_framework_sqlx_builder.And(
					table.F("ModuleID").Eq(m.ModuleID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				),
				github_com_eden_framework_sqlx_builder.Comment("Modules.SoftDeleteByModuleID"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	return err

}

func (m *Modules) List(db github_com_eden_framework_sqlx.DBExecutor, condition github_com_eden_framework_sqlx_builder.SqlCondition, additions ...github_com_eden_framework_sqlx_builder.Addition) ([]Modules, error) {

	list := make([]Modules, 0)

	table := db.T(m)
	_ = table

	condition = github_com_eden_framework_sqlx_builder.And(condition, table.F("DeletedAt").Eq(0))

	finalAdditions := []github_com_eden_framework_sqlx_builder.Addition{
		github_com_eden_framework_sqlx_builder.Where(condition),
		github_com_eden_framework_sqlx_builder.Comment("Modules.List"),
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

func (m *Modules) Count(db github_com_eden_framework_sqlx.DBExecutor, condition github_com_eden_framework_sqlx_builder.SqlCondition, additions ...github_com_eden_framework_sqlx_builder.Addition) (int, error) {

	count := -1

	table := db.T(m)
	_ = table

	condition = github_com_eden_framework_sqlx_builder.And(condition, table.F("DeletedAt").Eq(0))

	finalAdditions := []github_com_eden_framework_sqlx_builder.Addition{
		github_com_eden_framework_sqlx_builder.Where(condition),
		github_com_eden_framework_sqlx_builder.Comment("Modules.Count"),
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

func (m *Modules) BatchFetchByIDList(db github_com_eden_framework_sqlx.DBExecutor, values []uint64) ([]Modules, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("ID").In(values)

	return m.List(db, condition)

}

func (m *Modules) BatchFetchByModuleIDList(db github_com_eden_framework_sqlx.DBExecutor, values []uint64) ([]Modules, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("ModuleID").In(values)

	return m.List(db, condition)

}

func (m *Modules) BatchFetchByServiceIDList(db github_com_eden_framework_sqlx.DBExecutor, values []uint64) ([]Modules, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("ServiceID").In(values)

	return m.List(db, condition)

}
