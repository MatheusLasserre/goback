// Code generated by SQLBoiler 4.14.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// ExerciseModel is an object representing the database table.
type ExerciseModel struct {
	ID             string `boil:"id" json:"id" toml:"id" yaml:"id"`
	Task           string `boil:"task" json:"task" toml:"task" yaml:"task"`
	Concluded      bool   `boil:"concluded" json:"concluded" toml:"concluded" yaml:"concluded"`
	RoutineModelId string `boil:"routineModelId" json:"routineModelId" toml:"routineModelId" yaml:"routineModelId"`

	R *exerciseModelR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L exerciseModelL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ExerciseModelColumns = struct {
	ID             string
	Task           string
	Concluded      string
	RoutineModelId string
}{
	ID:             "id",
	Task:           "task",
	Concluded:      "concluded",
	RoutineModelId: "routineModelId",
}

var ExerciseModelTableColumns = struct {
	ID             string
	Task           string
	Concluded      string
	RoutineModelId string
}{
	ID:             "ExerciseModel.id",
	Task:           "ExerciseModel.task",
	Concluded:      "ExerciseModel.concluded",
	RoutineModelId: "ExerciseModel.routineModelId",
}

// Generated where

var ExerciseModelWhere = struct {
	ID             whereHelperstring
	Task           whereHelperstring
	Concluded      whereHelperbool
	RoutineModelId whereHelperstring
}{
	ID:             whereHelperstring{field: "`ExerciseModel`.`id`"},
	Task:           whereHelperstring{field: "`ExerciseModel`.`task`"},
	Concluded:      whereHelperbool{field: "`ExerciseModel`.`concluded`"},
	RoutineModelId: whereHelperstring{field: "`ExerciseModel`.`routineModelId`"},
}

// ExerciseModelRels is where relationship names are stored.
var ExerciseModelRels = struct {
}{}

// exerciseModelR is where relationships are stored.
type exerciseModelR struct {
}

// NewStruct creates a new relationship struct
func (*exerciseModelR) NewStruct() *exerciseModelR {
	return &exerciseModelR{}
}

// exerciseModelL is where Load methods for each relationship are stored.
type exerciseModelL struct{}

var (
	exerciseModelAllColumns            = []string{"id", "task", "concluded", "routineModelId"}
	exerciseModelColumnsWithoutDefault = []string{"id", "task", "routineModelId"}
	exerciseModelColumnsWithDefault    = []string{"concluded"}
	exerciseModelPrimaryKeyColumns     = []string{"id"}
	exerciseModelGeneratedColumns      = []string{}
)

type (
	// ExerciseModelSlice is an alias for a slice of pointers to ExerciseModel.
	// This should almost always be used instead of []ExerciseModel.
	ExerciseModelSlice []*ExerciseModel
	// ExerciseModelHook is the signature for custom ExerciseModel hook methods
	ExerciseModelHook func(context.Context, boil.ContextExecutor, *ExerciseModel) error

	exerciseModelQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	exerciseModelType                 = reflect.TypeOf(&ExerciseModel{})
	exerciseModelMapping              = queries.MakeStructMapping(exerciseModelType)
	exerciseModelPrimaryKeyMapping, _ = queries.BindMapping(exerciseModelType, exerciseModelMapping, exerciseModelPrimaryKeyColumns)
	exerciseModelInsertCacheMut       sync.RWMutex
	exerciseModelInsertCache          = make(map[string]insertCache)
	exerciseModelUpdateCacheMut       sync.RWMutex
	exerciseModelUpdateCache          = make(map[string]updateCache)
	exerciseModelUpsertCacheMut       sync.RWMutex
	exerciseModelUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var exerciseModelAfterSelectHooks []ExerciseModelHook

var exerciseModelBeforeInsertHooks []ExerciseModelHook
var exerciseModelAfterInsertHooks []ExerciseModelHook

var exerciseModelBeforeUpdateHooks []ExerciseModelHook
var exerciseModelAfterUpdateHooks []ExerciseModelHook

var exerciseModelBeforeDeleteHooks []ExerciseModelHook
var exerciseModelAfterDeleteHooks []ExerciseModelHook

var exerciseModelBeforeUpsertHooks []ExerciseModelHook
var exerciseModelAfterUpsertHooks []ExerciseModelHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ExerciseModel) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range exerciseModelAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ExerciseModel) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range exerciseModelBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ExerciseModel) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range exerciseModelAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ExerciseModel) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range exerciseModelBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ExerciseModel) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range exerciseModelAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ExerciseModel) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range exerciseModelBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ExerciseModel) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range exerciseModelAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ExerciseModel) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range exerciseModelBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ExerciseModel) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range exerciseModelAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddExerciseModelHook registers your hook function for all future operations.
func AddExerciseModelHook(hookPoint boil.HookPoint, exerciseModelHook ExerciseModelHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		exerciseModelAfterSelectHooks = append(exerciseModelAfterSelectHooks, exerciseModelHook)
	case boil.BeforeInsertHook:
		exerciseModelBeforeInsertHooks = append(exerciseModelBeforeInsertHooks, exerciseModelHook)
	case boil.AfterInsertHook:
		exerciseModelAfterInsertHooks = append(exerciseModelAfterInsertHooks, exerciseModelHook)
	case boil.BeforeUpdateHook:
		exerciseModelBeforeUpdateHooks = append(exerciseModelBeforeUpdateHooks, exerciseModelHook)
	case boil.AfterUpdateHook:
		exerciseModelAfterUpdateHooks = append(exerciseModelAfterUpdateHooks, exerciseModelHook)
	case boil.BeforeDeleteHook:
		exerciseModelBeforeDeleteHooks = append(exerciseModelBeforeDeleteHooks, exerciseModelHook)
	case boil.AfterDeleteHook:
		exerciseModelAfterDeleteHooks = append(exerciseModelAfterDeleteHooks, exerciseModelHook)
	case boil.BeforeUpsertHook:
		exerciseModelBeforeUpsertHooks = append(exerciseModelBeforeUpsertHooks, exerciseModelHook)
	case boil.AfterUpsertHook:
		exerciseModelAfterUpsertHooks = append(exerciseModelAfterUpsertHooks, exerciseModelHook)
	}
}

// One returns a single exerciseModel record from the query.
func (q exerciseModelQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ExerciseModel, error) {
	o := &ExerciseModel{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for ExerciseModel")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all ExerciseModel records from the query.
func (q exerciseModelQuery) All(ctx context.Context, exec boil.ContextExecutor) (ExerciseModelSlice, error) {
	var o []*ExerciseModel

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ExerciseModel slice")
	}

	if len(exerciseModelAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all ExerciseModel records in the query.
func (q exerciseModelQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count ExerciseModel rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q exerciseModelQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if ExerciseModel exists")
	}

	return count > 0, nil
}

// ExerciseModels retrieves all the records using an executor.
func ExerciseModels(mods ...qm.QueryMod) exerciseModelQuery {
	mods = append(mods, qm.From("`ExerciseModel`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`ExerciseModel`.*"})
	}

	return exerciseModelQuery{q}
}

// FindExerciseModel retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindExerciseModel(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*ExerciseModel, error) {
	exerciseModelObj := &ExerciseModel{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `ExerciseModel` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, exerciseModelObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from ExerciseModel")
	}

	if err = exerciseModelObj.doAfterSelectHooks(ctx, exec); err != nil {
		return exerciseModelObj, err
	}

	return exerciseModelObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ExerciseModel) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no ExerciseModel provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(exerciseModelColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	exerciseModelInsertCacheMut.RLock()
	cache, cached := exerciseModelInsertCache[key]
	exerciseModelInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			exerciseModelAllColumns,
			exerciseModelColumnsWithDefault,
			exerciseModelColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(exerciseModelType, exerciseModelMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(exerciseModelType, exerciseModelMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `ExerciseModel` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `ExerciseModel` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `ExerciseModel` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, exerciseModelPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into ExerciseModel")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for ExerciseModel")
	}

CacheNoHooks:
	if !cached {
		exerciseModelInsertCacheMut.Lock()
		exerciseModelInsertCache[key] = cache
		exerciseModelInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the ExerciseModel.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ExerciseModel) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	exerciseModelUpdateCacheMut.RLock()
	cache, cached := exerciseModelUpdateCache[key]
	exerciseModelUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			exerciseModelAllColumns,
			exerciseModelPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update ExerciseModel, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `ExerciseModel` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, exerciseModelPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(exerciseModelType, exerciseModelMapping, append(wl, exerciseModelPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update ExerciseModel row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for ExerciseModel")
	}

	if !cached {
		exerciseModelUpdateCacheMut.Lock()
		exerciseModelUpdateCache[key] = cache
		exerciseModelUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q exerciseModelQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for ExerciseModel")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for ExerciseModel")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ExerciseModelSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), exerciseModelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `ExerciseModel` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, exerciseModelPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in exerciseModel slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all exerciseModel")
	}
	return rowsAff, nil
}

var mySQLExerciseModelUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ExerciseModel) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no ExerciseModel provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(exerciseModelColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLExerciseModelUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	exerciseModelUpsertCacheMut.RLock()
	cache, cached := exerciseModelUpsertCache[key]
	exerciseModelUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			exerciseModelAllColumns,
			exerciseModelColumnsWithDefault,
			exerciseModelColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			exerciseModelAllColumns,
			exerciseModelPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert ExerciseModel, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`ExerciseModel`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `ExerciseModel` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(exerciseModelType, exerciseModelMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(exerciseModelType, exerciseModelMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for ExerciseModel")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(exerciseModelType, exerciseModelMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for ExerciseModel")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for ExerciseModel")
	}

CacheNoHooks:
	if !cached {
		exerciseModelUpsertCacheMut.Lock()
		exerciseModelUpsertCache[key] = cache
		exerciseModelUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single ExerciseModel record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ExerciseModel) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no ExerciseModel provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), exerciseModelPrimaryKeyMapping)
	sql := "DELETE FROM `ExerciseModel` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from ExerciseModel")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for ExerciseModel")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q exerciseModelQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no exerciseModelQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from ExerciseModel")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for ExerciseModel")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ExerciseModelSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(exerciseModelBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), exerciseModelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `ExerciseModel` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, exerciseModelPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from exerciseModel slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for ExerciseModel")
	}

	if len(exerciseModelAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *ExerciseModel) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindExerciseModel(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ExerciseModelSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ExerciseModelSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), exerciseModelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `ExerciseModel`.* FROM `ExerciseModel` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, exerciseModelPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ExerciseModelSlice")
	}

	*o = slice

	return nil
}

// ExerciseModelExists checks if the ExerciseModel row exists.
func ExerciseModelExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `ExerciseModel` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if ExerciseModel exists")
	}

	return exists, nil
}

// Exists checks if the ExerciseModel row exists.
func (o *ExerciseModel) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return ExerciseModelExists(ctx, exec, o.ID)
}
