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

// OrganizationsMember is an object representing the database table.
type OrganizationsMember struct {
	ID             int64  `boil:"id" json:"id" toml:"id" yaml:"id"`
	OrganizationId int64  `boil:"organizationId" json:"organizationId" toml:"organizationId" yaml:"organizationId"`
	UserId         string `boil:"userId" json:"userId" toml:"userId" yaml:"userId"`

	R *organizationsMemberR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L organizationsMemberL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var OrganizationsMemberColumns = struct {
	ID             string
	OrganizationId string
	UserId         string
}{
	ID:             "id",
	OrganizationId: "organizationId",
	UserId:         "userId",
}

var OrganizationsMemberTableColumns = struct {
	ID             string
	OrganizationId string
	UserId         string
}{
	ID:             "organizations_members.id",
	OrganizationId: "organizations_members.organizationId",
	UserId:         "organizations_members.userId",
}

// Generated where

var OrganizationsMemberWhere = struct {
	ID             whereHelperint64
	OrganizationId whereHelperint64
	UserId         whereHelperstring
}{
	ID:             whereHelperint64{field: "`organizations_members`.`id`"},
	OrganizationId: whereHelperint64{field: "`organizations_members`.`organizationId`"},
	UserId:         whereHelperstring{field: "`organizations_members`.`userId`"},
}

// OrganizationsMemberRels is where relationship names are stored.
var OrganizationsMemberRels = struct {
}{}

// organizationsMemberR is where relationships are stored.
type organizationsMemberR struct {
}

// NewStruct creates a new relationship struct
func (*organizationsMemberR) NewStruct() *organizationsMemberR {
	return &organizationsMemberR{}
}

// organizationsMemberL is where Load methods for each relationship are stored.
type organizationsMemberL struct{}

var (
	organizationsMemberAllColumns            = []string{"id", "organizationId", "userId"}
	organizationsMemberColumnsWithoutDefault = []string{"organizationId", "userId"}
	organizationsMemberColumnsWithDefault    = []string{"id"}
	organizationsMemberPrimaryKeyColumns     = []string{"id"}
	organizationsMemberGeneratedColumns      = []string{}
)

type (
	// OrganizationsMemberSlice is an alias for a slice of pointers to OrganizationsMember.
	// This should almost always be used instead of []OrganizationsMember.
	OrganizationsMemberSlice []*OrganizationsMember
	// OrganizationsMemberHook is the signature for custom OrganizationsMember hook methods
	OrganizationsMemberHook func(context.Context, boil.ContextExecutor, *OrganizationsMember) error

	organizationsMemberQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	organizationsMemberType                 = reflect.TypeOf(&OrganizationsMember{})
	organizationsMemberMapping              = queries.MakeStructMapping(organizationsMemberType)
	organizationsMemberPrimaryKeyMapping, _ = queries.BindMapping(organizationsMemberType, organizationsMemberMapping, organizationsMemberPrimaryKeyColumns)
	organizationsMemberInsertCacheMut       sync.RWMutex
	organizationsMemberInsertCache          = make(map[string]insertCache)
	organizationsMemberUpdateCacheMut       sync.RWMutex
	organizationsMemberUpdateCache          = make(map[string]updateCache)
	organizationsMemberUpsertCacheMut       sync.RWMutex
	organizationsMemberUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var organizationsMemberAfterSelectHooks []OrganizationsMemberHook

var organizationsMemberBeforeInsertHooks []OrganizationsMemberHook
var organizationsMemberAfterInsertHooks []OrganizationsMemberHook

var organizationsMemberBeforeUpdateHooks []OrganizationsMemberHook
var organizationsMemberAfterUpdateHooks []OrganizationsMemberHook

var organizationsMemberBeforeDeleteHooks []OrganizationsMemberHook
var organizationsMemberAfterDeleteHooks []OrganizationsMemberHook

var organizationsMemberBeforeUpsertHooks []OrganizationsMemberHook
var organizationsMemberAfterUpsertHooks []OrganizationsMemberHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *OrganizationsMember) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range organizationsMemberAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *OrganizationsMember) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range organizationsMemberBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *OrganizationsMember) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range organizationsMemberAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *OrganizationsMember) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range organizationsMemberBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *OrganizationsMember) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range organizationsMemberAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *OrganizationsMember) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range organizationsMemberBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *OrganizationsMember) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range organizationsMemberAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *OrganizationsMember) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range organizationsMemberBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *OrganizationsMember) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range organizationsMemberAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddOrganizationsMemberHook registers your hook function for all future operations.
func AddOrganizationsMemberHook(hookPoint boil.HookPoint, organizationsMemberHook OrganizationsMemberHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		organizationsMemberAfterSelectHooks = append(organizationsMemberAfterSelectHooks, organizationsMemberHook)
	case boil.BeforeInsertHook:
		organizationsMemberBeforeInsertHooks = append(organizationsMemberBeforeInsertHooks, organizationsMemberHook)
	case boil.AfterInsertHook:
		organizationsMemberAfterInsertHooks = append(organizationsMemberAfterInsertHooks, organizationsMemberHook)
	case boil.BeforeUpdateHook:
		organizationsMemberBeforeUpdateHooks = append(organizationsMemberBeforeUpdateHooks, organizationsMemberHook)
	case boil.AfterUpdateHook:
		organizationsMemberAfterUpdateHooks = append(organizationsMemberAfterUpdateHooks, organizationsMemberHook)
	case boil.BeforeDeleteHook:
		organizationsMemberBeforeDeleteHooks = append(organizationsMemberBeforeDeleteHooks, organizationsMemberHook)
	case boil.AfterDeleteHook:
		organizationsMemberAfterDeleteHooks = append(organizationsMemberAfterDeleteHooks, organizationsMemberHook)
	case boil.BeforeUpsertHook:
		organizationsMemberBeforeUpsertHooks = append(organizationsMemberBeforeUpsertHooks, organizationsMemberHook)
	case boil.AfterUpsertHook:
		organizationsMemberAfterUpsertHooks = append(organizationsMemberAfterUpsertHooks, organizationsMemberHook)
	}
}

// One returns a single organizationsMember record from the query.
func (q organizationsMemberQuery) One(ctx context.Context, exec boil.ContextExecutor) (*OrganizationsMember, error) {
	o := &OrganizationsMember{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for organizations_members")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all OrganizationsMember records from the query.
func (q organizationsMemberQuery) All(ctx context.Context, exec boil.ContextExecutor) (OrganizationsMemberSlice, error) {
	var o []*OrganizationsMember

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to OrganizationsMember slice")
	}

	if len(organizationsMemberAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all OrganizationsMember records in the query.
func (q organizationsMemberQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count organizations_members rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q organizationsMemberQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if organizations_members exists")
	}

	return count > 0, nil
}

// OrganizationsMembers retrieves all the records using an executor.
func OrganizationsMembers(mods ...qm.QueryMod) organizationsMemberQuery {
	mods = append(mods, qm.From("`organizations_members`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`organizations_members`.*"})
	}

	return organizationsMemberQuery{q}
}

// FindOrganizationsMember retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindOrganizationsMember(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*OrganizationsMember, error) {
	organizationsMemberObj := &OrganizationsMember{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `organizations_members` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, organizationsMemberObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from organizations_members")
	}

	if err = organizationsMemberObj.doAfterSelectHooks(ctx, exec); err != nil {
		return organizationsMemberObj, err
	}

	return organizationsMemberObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *OrganizationsMember) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no organizations_members provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(organizationsMemberColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	organizationsMemberInsertCacheMut.RLock()
	cache, cached := organizationsMemberInsertCache[key]
	organizationsMemberInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			organizationsMemberAllColumns,
			organizationsMemberColumnsWithDefault,
			organizationsMemberColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(organizationsMemberType, organizationsMemberMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(organizationsMemberType, organizationsMemberMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `organizations_members` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `organizations_members` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `organizations_members` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, organizationsMemberPrimaryKeyColumns))
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
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into organizations_members")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == organizationsMemberMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for organizations_members")
	}

CacheNoHooks:
	if !cached {
		organizationsMemberInsertCacheMut.Lock()
		organizationsMemberInsertCache[key] = cache
		organizationsMemberInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the OrganizationsMember.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *OrganizationsMember) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	organizationsMemberUpdateCacheMut.RLock()
	cache, cached := organizationsMemberUpdateCache[key]
	organizationsMemberUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			organizationsMemberAllColumns,
			organizationsMemberPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update organizations_members, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `organizations_members` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, organizationsMemberPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(organizationsMemberType, organizationsMemberMapping, append(wl, organizationsMemberPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update organizations_members row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for organizations_members")
	}

	if !cached {
		organizationsMemberUpdateCacheMut.Lock()
		organizationsMemberUpdateCache[key] = cache
		organizationsMemberUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q organizationsMemberQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for organizations_members")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for organizations_members")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o OrganizationsMemberSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), organizationsMemberPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `organizations_members` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, organizationsMemberPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in organizationsMember slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all organizationsMember")
	}
	return rowsAff, nil
}

var mySQLOrganizationsMemberUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *OrganizationsMember) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no organizations_members provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(organizationsMemberColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLOrganizationsMemberUniqueColumns, o)

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

	organizationsMemberUpsertCacheMut.RLock()
	cache, cached := organizationsMemberUpsertCache[key]
	organizationsMemberUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			organizationsMemberAllColumns,
			organizationsMemberColumnsWithDefault,
			organizationsMemberColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			organizationsMemberAllColumns,
			organizationsMemberPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert organizations_members, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`organizations_members`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `organizations_members` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(organizationsMemberType, organizationsMemberMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(organizationsMemberType, organizationsMemberMapping, ret)
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
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for organizations_members")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == organizationsMemberMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(organizationsMemberType, organizationsMemberMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for organizations_members")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for organizations_members")
	}

CacheNoHooks:
	if !cached {
		organizationsMemberUpsertCacheMut.Lock()
		organizationsMemberUpsertCache[key] = cache
		organizationsMemberUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single OrganizationsMember record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *OrganizationsMember) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no OrganizationsMember provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), organizationsMemberPrimaryKeyMapping)
	sql := "DELETE FROM `organizations_members` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from organizations_members")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for organizations_members")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q organizationsMemberQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no organizationsMemberQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from organizations_members")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for organizations_members")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o OrganizationsMemberSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(organizationsMemberBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), organizationsMemberPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `organizations_members` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, organizationsMemberPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from organizationsMember slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for organizations_members")
	}

	if len(organizationsMemberAfterDeleteHooks) != 0 {
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
func (o *OrganizationsMember) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindOrganizationsMember(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OrganizationsMemberSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := OrganizationsMemberSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), organizationsMemberPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `organizations_members`.* FROM `organizations_members` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, organizationsMemberPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in OrganizationsMemberSlice")
	}

	*o = slice

	return nil
}

// OrganizationsMemberExists checks if the OrganizationsMember row exists.
func OrganizationsMemberExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `organizations_members` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if organizations_members exists")
	}

	return exists, nil
}

// Exists checks if the OrganizationsMember row exists.
func (o *OrganizationsMember) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return OrganizationsMemberExists(ctx, exec, o.ID)
}