package databases

import (
	fmt "fmt"
	time "time"

	github_com_eden_framework_sqlx "github.com/eden-framework/sqlx"
	github_com_eden_framework_sqlx_builder "github.com/eden-framework/sqlx/builder"
	github_com_eden_framework_sqlx_datatypes "github.com/eden-framework/sqlx/datatypes"
)

func (ModulePermissionApi) PrimaryKey() []string {
	return []string{
		"ID",
	}
}

func (ModulePermissionApi) Indexes() github_com_eden_framework_sqlx_builder.Indexes {
	return github_com_eden_framework_sqlx_builder.Indexes{
		"I_permission": []string{
			"PermissionID",
		},
	}
}

func (ModulePermissionApi) UniqueIndexUAPIID() string {
	return "U_api_id"
}

func (ModulePermissionApi) UniqueIndexes() github_com_eden_framework_sqlx_builder.Indexes {
	return github_com_eden_framework_sqlx_builder.Indexes{
		"U_api_id": []string{
			"ApiID",
			"DeletedAt",
		},
	}
}

func (ModulePermissionApi) Comments() map[string]string {
	return map[string]string{
		"ApiID":        "业务ID",
		"Name":         "名称",
		"PermissionID": "所属权限策略",
		"RequestKey":   "请求标识",
		"RequestPath":  "请求路径",
	}
}

var ModulePermissionApiTable *github_com_eden_framework_sqlx_builder.Table

func init() {
	ModulePermissionApiTable = Config.DB.Register(&ModulePermissionApi{})
}

type ModulePermissionApiIterator struct {
}

func (ModulePermissionApiIterator) New() interface{} {
	return &ModulePermissionApi{}
}

func (ModulePermissionApiIterator) Resolve(v interface{}) *ModulePermissionApi {
	return v.(*ModulePermissionApi)
}

func (ModulePermissionApi) TableName() string {
	return "t_module_permission_api"
}

func (ModulePermissionApi) ColDescriptions() map[string][]string {
	return map[string][]string{
		"ApiID": []string{
			"业务ID",
		},
		"Name": []string{
			"名称",
		},
		"PermissionID": []string{
			"所属权限策略",
		},
		"RequestKey": []string{
			"请求标识",
		},
		"RequestPath": []string{
			"请求路径",
		},
	}
}

func (ModulePermissionApi) FieldKeyID() string {
	return "ID"
}

func (m *ModulePermissionApi) FieldID() *github_com_eden_framework_sqlx_builder.Column {
	return ModulePermissionApiTable.F(m.FieldKeyID())
}

func (ModulePermissionApi) FieldKeyApiID() string {
	return "ApiID"
}

func (m *ModulePermissionApi) FieldApiID() *github_com_eden_framework_sqlx_builder.Column {
	return ModulePermissionApiTable.F(m.FieldKeyApiID())
}

func (ModulePermissionApi) FieldKeyName() string {
	return "Name"
}

func (m *ModulePermissionApi) FieldName() *github_com_eden_framework_sqlx_builder.Column {
	return ModulePermissionApiTable.F(m.FieldKeyName())
}

func (ModulePermissionApi) FieldKeyRequestKey() string {
	return "RequestKey"
}

func (m *ModulePermissionApi) FieldRequestKey() *github_com_eden_framework_sqlx_builder.Column {
	return ModulePermissionApiTable.F(m.FieldKeyRequestKey())
}

func (ModulePermissionApi) FieldKeyRequestPath() string {
	return "RequestPath"
}

func (m *ModulePermissionApi) FieldRequestPath() *github_com_eden_framework_sqlx_builder.Column {
	return ModulePermissionApiTable.F(m.FieldKeyRequestPath())
}

func (ModulePermissionApi) FieldKeyPermissionID() string {
	return "PermissionID"
}

func (m *ModulePermissionApi) FieldPermissionID() *github_com_eden_framework_sqlx_builder.Column {
	return ModulePermissionApiTable.F(m.FieldKeyPermissionID())
}

func (ModulePermissionApi) FieldKeyCreatedAt() string {
	return "CreatedAt"
}

func (m *ModulePermissionApi) FieldCreatedAt() *github_com_eden_framework_sqlx_builder.Column {
	return ModulePermissionApiTable.F(m.FieldKeyCreatedAt())
}

func (ModulePermissionApi) FieldKeyUpdatedAt() string {
	return "UpdatedAt"
}

func (m *ModulePermissionApi) FieldUpdatedAt() *github_com_eden_framework_sqlx_builder.Column {
	return ModulePermissionApiTable.F(m.FieldKeyUpdatedAt())
}

func (ModulePermissionApi) FieldKeyDeletedAt() string {
	return "DeletedAt"
}

func (m *ModulePermissionApi) FieldDeletedAt() *github_com_eden_framework_sqlx_builder.Column {
	return ModulePermissionApiTable.F(m.FieldKeyDeletedAt())
}

func (ModulePermissionApi) ColRelations() map[string][]string {
	return map[string][]string{}
}

func (m *ModulePermissionApi) IndexFieldNames() []string {
	return []string{
		"ApiID",
		"ID",
		"PermissionID",
	}
}

func (m *ModulePermissionApi) ConditionByStruct(db github_com_eden_framework_sqlx.DBExecutor) github_com_eden_framework_sqlx_builder.SqlCondition {
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

func (m *ModulePermissionApi) Create(db github_com_eden_framework_sqlx.DBExecutor) error {

	if m.CreatedAt.IsZero() {
		m.CreatedAt = github_com_eden_framework_sqlx_datatypes.Timestamp(time.Now())
	}

	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = github_com_eden_framework_sqlx_datatypes.Timestamp(time.Now())
	}

	_, err := db.ExecExpr(github_com_eden_framework_sqlx.InsertToDB(db, m, nil))
	return err

}

func (m *ModulePermissionApi) CreateOnDuplicateWithUpdateFields(db github_com_eden_framework_sqlx.DBExecutor, updateFields []string) error {

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

func (m *ModulePermissionApi) DeleteByStruct(db github_com_eden_framework_sqlx.DBExecutor) error {

	_, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Delete().
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(m.ConditionByStruct(db)),
				github_com_eden_framework_sqlx_builder.Comment("ModulePermissionApi.DeleteByStruct"),
			),
	)

	return err
}

func (m *ModulePermissionApi) FetchByID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_eden_framework_sqlx_builder.Select(nil).
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("ID").Eq(m.ID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("ModulePermissionApi.FetchByID"),
			),
		m,
	)

	return err
}

func (m *ModulePermissionApi) UpdateByIDWithMap(db github_com_eden_framework_sqlx.DBExecutor, fieldValues github_com_eden_framework_sqlx_builder.FieldValues) error {

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
				github_com_eden_framework_sqlx_builder.Comment("ModulePermissionApi.UpdateByIDWithMap"),
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

func (m *ModulePermissionApi) UpdateByIDWithStruct(db github_com_eden_framework_sqlx.DBExecutor, zeroFields ...string) error {

	fieldValues := github_com_eden_framework_sqlx_builder.FieldValuesFromStructByNonZero(m, zeroFields...)
	return m.UpdateByIDWithMap(db, fieldValues)

}

func (m *ModulePermissionApi) FetchByIDForUpdate(db github_com_eden_framework_sqlx.DBExecutor) error {

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
				github_com_eden_framework_sqlx_builder.Comment("ModulePermissionApi.FetchByIDForUpdate"),
			),
		m,
	)

	return err
}

func (m *ModulePermissionApi) DeleteByID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	_, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Delete().
			From(db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("ID").Eq(m.ID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("ModulePermissionApi.DeleteByID"),
			))

	return err
}

func (m *ModulePermissionApi) SoftDeleteByID(db github_com_eden_framework_sqlx.DBExecutor) error {

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
				github_com_eden_framework_sqlx_builder.Comment("ModulePermissionApi.SoftDeleteByID"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	return err

}

func (m *ModulePermissionApi) FetchByApiID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_eden_framework_sqlx_builder.Select(nil).
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("ApiID").Eq(m.ApiID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("ModulePermissionApi.FetchByApiID"),
			),
		m,
	)

	return err
}

func (m *ModulePermissionApi) UpdateByApiIDWithMap(db github_com_eden_framework_sqlx.DBExecutor, fieldValues github_com_eden_framework_sqlx_builder.FieldValues) error {

	if _, ok := fieldValues["UpdatedAt"]; !ok {
		fieldValues["UpdatedAt"] = github_com_eden_framework_sqlx_datatypes.Timestamp(time.Now())
	}

	table := db.T(m)

	result, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Update(db.T(m)).
			Where(
				github_com_eden_framework_sqlx_builder.And(
					table.F("ApiID").Eq(m.ApiID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				),
				github_com_eden_framework_sqlx_builder.Comment("ModulePermissionApi.UpdateByApiIDWithMap"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return m.FetchByApiID(db)
	}

	return nil

}

func (m *ModulePermissionApi) UpdateByApiIDWithStruct(db github_com_eden_framework_sqlx.DBExecutor, zeroFields ...string) error {

	fieldValues := github_com_eden_framework_sqlx_builder.FieldValuesFromStructByNonZero(m, zeroFields...)
	return m.UpdateByApiIDWithMap(db, fieldValues)

}

func (m *ModulePermissionApi) FetchByApiIDForUpdate(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_eden_framework_sqlx_builder.Select(nil).
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("ApiID").Eq(m.ApiID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.ForUpdate(),
				github_com_eden_framework_sqlx_builder.Comment("ModulePermissionApi.FetchByApiIDForUpdate"),
			),
		m,
	)

	return err
}

func (m *ModulePermissionApi) DeleteByApiID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	_, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Delete().
			From(db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("ApiID").Eq(m.ApiID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("ModulePermissionApi.DeleteByApiID"),
			))

	return err
}

func (m *ModulePermissionApi) SoftDeleteByApiID(db github_com_eden_framework_sqlx.DBExecutor) error {

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
					table.F("ApiID").Eq(m.ApiID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				),
				github_com_eden_framework_sqlx_builder.Comment("ModulePermissionApi.SoftDeleteByApiID"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	return err

}

func (m *ModulePermissionApi) List(db github_com_eden_framework_sqlx.DBExecutor, condition github_com_eden_framework_sqlx_builder.SqlCondition, additions ...github_com_eden_framework_sqlx_builder.Addition) ([]ModulePermissionApi, error) {

	list := make([]ModulePermissionApi, 0)

	table := db.T(m)
	_ = table

	condition = github_com_eden_framework_sqlx_builder.And(condition, table.F("DeletedAt").Eq(0))

	finalAdditions := []github_com_eden_framework_sqlx_builder.Addition{
		github_com_eden_framework_sqlx_builder.Where(condition),
		github_com_eden_framework_sqlx_builder.Comment("ModulePermissionApi.List"),
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

func (m *ModulePermissionApi) Count(db github_com_eden_framework_sqlx.DBExecutor, condition github_com_eden_framework_sqlx_builder.SqlCondition, additions ...github_com_eden_framework_sqlx_builder.Addition) (int, error) {

	count := -1

	table := db.T(m)
	_ = table

	condition = github_com_eden_framework_sqlx_builder.And(condition, table.F("DeletedAt").Eq(0))

	finalAdditions := []github_com_eden_framework_sqlx_builder.Addition{
		github_com_eden_framework_sqlx_builder.Where(condition),
		github_com_eden_framework_sqlx_builder.Comment("ModulePermissionApi.Count"),
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

func (m *ModulePermissionApi) BatchFetchByApiIDList(db github_com_eden_framework_sqlx.DBExecutor, values []uint64) ([]ModulePermissionApi, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("ApiID").In(values)

	return m.List(db, condition)

}

func (m *ModulePermissionApi) BatchFetchByIDList(db github_com_eden_framework_sqlx.DBExecutor, values []uint64) ([]ModulePermissionApi, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("ID").In(values)

	return m.List(db, condition)

}

func (m *ModulePermissionApi) BatchFetchByPermissionIDList(db github_com_eden_framework_sqlx.DBExecutor, values []uint64) ([]ModulePermissionApi, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("PermissionID").In(values)

	return m.List(db, condition)

}
