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

// RoutineModel is an object representing the database table.
type RoutineModel struct {
	ID        string    `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt time.Time `boil:"createdAt" json:"createdAt" toml:"createdAt" yaml:"createdAt"`
	UpdatedAt time.Time `boil:"updatedAt" json:"updatedAt" toml:"updatedAt" yaml:"updatedAt"`
	UserId    string    `boil:"userId" json:"userId" toml:"userId" yaml:"userId"`

	R *routineModelR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L routineModelL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var RoutineModelColumns = struct {
	ID        string
	CreatedAt string
	UpdatedAt string
	UserId    string
}{
	ID:        "id",
	CreatedAt: "createdAt",
	UpdatedAt: "updatedAt",
	UserId:    "userId",
}

var RoutineModelTableColumns = struct {
	ID        string
	CreatedAt string
	UpdatedAt string
	UserId    string
}{
	ID:        "RoutineModel.id",
	CreatedAt: "RoutineModel.createdAt",
	UpdatedAt: "RoutineModel.updatedAt",
	UserId:    "RoutineModel.userId",
}

// Generated where

var RoutineModelWhere = struct {
	ID        whereHelperstring
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
	UserId    whereHelperstring
}{
	ID:        whereHelperstring{field: "`RoutineModel`.`id`"},
	CreatedAt: whereHelpertime_Time{field: "`RoutineModel`.`createdAt`"},
	UpdatedAt: whereHelpertime_Time{field: "`RoutineModel`.`updatedAt`"},
	UserId:    whereHelperstring{field: "`RoutineModel`.`userId`"},
}

// RoutineModelRels is where relationship names are stored.
var RoutineModelRels = struct {
}{}

// routineModelR is where relationships are stored.
type routineModelR struct {
}

// NewStruct creates a new relationship struct
func (*routineModelR) NewStruct() *routineModelR {
	return &routineModelR{}
}

// routineModelL is where Load methods for each relationship are stored.
type routineModelL struct{}

var (
	routineModelAllColumns            = []string{"id", "createdAt", "updatedAt", "userId"}
	routineModelColumnsWithoutDefault = []string{"id", "updatedAt", "userId"}
	routineModelColumnsWithDefault    = []string{"createdAt"}
	routineModelPrimaryKeyColumns     = []string{"id"}
	routineModelGeneratedColumns      = []string{}
)

type (
	// RoutineModelSlice is an alias for a slice of pointers to RoutineModel.
	// This should almost always be used instead of []RoutineModel.
	RoutineModelSlice []*RoutineModel
	// RoutineModelHook is the signature for custom RoutineModel hook methods
	RoutineModelHook func(context.Context, boil.ContextExecutor, *RoutineModel) error

	routineModelQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	routineModelType                 = reflect.TypeOf(&RoutineModel{})
	routineModelMapping              = queries.MakeStructMapping(routineModelType)
	routineModelPrimaryKeyMapping, _ = queries.BindMapping(routineModelType, routineModelMapping, routineModelPrimaryKeyColumns)
	routineModelInsertCacheMut       sync.RWMutex
	routineModelInsertCache          = make(map[string]insertCache)
	routineModelUpdateCacheMut       sync.RWMutex
	routineModelUpdateCache          = make(map[string]updateCache)
	routineModelUpsertCacheMut       sync.RWMutex
	routineModelUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var routineModelAfterSelectHooks []RoutineModelHook

var routineModelBeforeInsertHooks []RoutineModelHook
var routineModelAfterInsertHooks []RoutineModelHook

var routineModelBeforeUpdateHooks []RoutineModelHook
var routineModelAfterUpdateHooks []RoutineModelHook

var routineModelBeforeDeleteHooks []RoutineModelHook
var routineModelAfterDeleteHooks []RoutineModelHook

var routineModelBeforeUpsertHooks []RoutineModelHook
var routineModelAfterUpsertHooks []RoutineModelHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *RoutineModel) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range routineModelAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *RoutineModel) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range routineModelBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *RoutineModel) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range routineModelAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *RoutineModel) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range routineModelBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *RoutineModel) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range routineModelAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *RoutineModel) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range routineModelBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *RoutineModel) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range routineModelAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *RoutineModel) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range routineModelBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *RoutineModel) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range routineModelAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddRoutineModelHook registers your hook function for all future operations.
func AddRoutineModelHook(hookPoint boil.HookPoint, routineModelHook RoutineModelHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		routineModelAfterSelectHooks = append(routineModelAfterSelectHooks, routineModelHook)
	case boil.BeforeInsertHook:
		routineModelBeforeInsertHooks = append(routineModelBeforeInsertHooks, routineModelHook)
	case boil.AfterInsertHook:
		routineModelAfterInsertHooks = append(routineModelAfterInsertHooks, routineModelHook)
	case boil.BeforeUpdateHook:
		routineModelBeforeUpdateHooks = append(routineModelBeforeUpdateHooks, routineModelHook)
	case boil.AfterUpdateHook:
		routineModelAfterUpdateHooks = append(routineModelAfterUpdateHooks, routineModelHook)
	case boil.BeforeDeleteHook:
		routineModelBeforeDeleteHooks = append(routineModelBeforeDeleteHooks, routineModelHook)
	case boil.AfterDeleteHook:
		routineModelAfterDeleteHooks = append(routineModelAfterDeleteHooks, routineModelHook)
	case boil.BeforeUpsertHook:
		routineModelBeforeUpsertHooks = append(routineModelBeforeUpsertHooks, routineModelHook)
	case boil.AfterUpsertHook:
		routineModelAfterUpsertHooks = append(routineModelAfterUpsertHooks, routineModelHook)
	}
}

// One returns a single routineModel record from the query.
func (q routineModelQuery) One(ctx context.Context, exec boil.ContextExecutor) (*RoutineModel, error) {
	o := &RoutineModel{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for RoutineModel")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all RoutineModel records from the query.
func (q routineModelQuery) All(ctx context.Context, exec boil.ContextExecutor) (RoutineModelSlice, error) {
	var o []*RoutineModel

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to RoutineModel slice")
	}

	if len(routineModelAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all RoutineModel records in the query.
func (q routineModelQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count RoutineModel rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q routineModelQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if RoutineModel exists")
	}

	return count > 0, nil
}

// RoutineModels retrieves all the records using an executor.
func RoutineModels(mods ...qm.QueryMod) routineModelQuery {
	mods = append(mods, qm.From("`RoutineModel`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`RoutineModel`.*"})
	}

	return routineModelQuery{q}
}

// FindRoutineModel retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindRoutineModel(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*RoutineModel, error) {
	routineModelObj := &RoutineModel{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `RoutineModel` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, routineModelObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from RoutineModel")
	}

	if err = routineModelObj.doAfterSelectHooks(ctx, exec); err != nil {
		return routineModelObj, err
	}

	return routineModelObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *RoutineModel) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no RoutineModel provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(routineModelColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	routineModelInsertCacheMut.RLock()
	cache, cached := routineModelInsertCache[key]
	routineModelInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			routineModelAllColumns,
			routineModelColumnsWithDefault,
			routineModelColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(routineModelType, routineModelMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(routineModelType, routineModelMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `RoutineModel` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `RoutineModel` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `RoutineModel` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, routineModelPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into RoutineModel")
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
		return errors.Wrap(err, "models: unable to populate default values for RoutineModel")
	}

CacheNoHooks:
	if !cached {
		routineModelInsertCacheMut.Lock()
		routineModelInsertCache[key] = cache
		routineModelInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the RoutineModel.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *RoutineModel) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	routineModelUpdateCacheMut.RLock()
	cache, cached := routineModelUpdateCache[key]
	routineModelUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			routineModelAllColumns,
			routineModelPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update RoutineModel, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `RoutineModel` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, routineModelPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(routineModelType, routineModelMapping, append(wl, routineModelPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update RoutineModel row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for RoutineModel")
	}

	if !cached {
		routineModelUpdateCacheMut.Lock()
		routineModelUpdateCache[key] = cache
		routineModelUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q routineModelQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for RoutineModel")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for RoutineModel")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o RoutineModelSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), routineModelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `RoutineModel` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, routineModelPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in routineModel slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all routineModel")
	}
	return rowsAff, nil
}

var mySQLRoutineModelUniqueColumns = []string{
	"id",
	"userId",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *RoutineModel) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no RoutineModel provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(routineModelColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLRoutineModelUniqueColumns, o)

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

	routineModelUpsertCacheMut.RLock()
	cache, cached := routineModelUpsertCache[key]
	routineModelUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			routineModelAllColumns,
			routineModelColumnsWithDefault,
			routineModelColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			routineModelAllColumns,
			routineModelPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert RoutineModel, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`RoutineModel`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `RoutineModel` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(routineModelType, routineModelMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(routineModelType, routineModelMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for RoutineModel")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(routineModelType, routineModelMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for RoutineModel")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for RoutineModel")
	}

CacheNoHooks:
	if !cached {
		routineModelUpsertCacheMut.Lock()
		routineModelUpsertCache[key] = cache
		routineModelUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single RoutineModel record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *RoutineModel) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no RoutineModel provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), routineModelPrimaryKeyMapping)
	sql := "DELETE FROM `RoutineModel` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from RoutineModel")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for RoutineModel")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q routineModelQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no routineModelQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from RoutineModel")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for RoutineModel")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o RoutineModelSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(routineModelBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), routineModelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `RoutineModel` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, routineModelPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from routineModel slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for RoutineModel")
	}

	if len(routineModelAfterDeleteHooks) != 0 {
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
func (o *RoutineModel) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindRoutineModel(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *RoutineModelSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := RoutineModelSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), routineModelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `RoutineModel`.* FROM `RoutineModel` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, routineModelPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in RoutineModelSlice")
	}

	*o = slice

	return nil
}

// RoutineModelExists checks if the RoutineModel row exists.
func RoutineModelExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `RoutineModel` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if RoutineModel exists")
	}

	return exists, nil
}

// Exists checks if the RoutineModel row exists.
func (o *RoutineModel) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return RoutineModelExists(ctx, exec, o.ID)
}
