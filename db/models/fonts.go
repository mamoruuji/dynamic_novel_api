// Code generated by SQLBoiler 4.15.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

// Font is an object representing the database table.
type Font struct {
	FontID    int32     `boil:"font_id" json:"font_id" toml:"font_id" yaml:"font_id"`
	Name      string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *fontR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L fontL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var FontColumns = struct {
	FontID    string
	Name      string
	CreatedAt string
	UpdatedAt string
}{
	FontID:    "font_id",
	Name:      "name",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

var FontTableColumns = struct {
	FontID    string
	Name      string
	CreatedAt string
	UpdatedAt string
}{
	FontID:    "fonts.font_id",
	Name:      "fonts.name",
	CreatedAt: "fonts.created_at",
	UpdatedAt: "fonts.updated_at",
}

// Generated where

var FontWhere = struct {
	FontID    whereHelperint32
	Name      whereHelperstring
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
}{
	FontID:    whereHelperint32{field: "\"fonts\".\"font_id\""},
	Name:      whereHelperstring{field: "\"fonts\".\"name\""},
	CreatedAt: whereHelpertime_Time{field: "\"fonts\".\"created_at\""},
	UpdatedAt: whereHelpertime_Time{field: "\"fonts\".\"updated_at\""},
}

// FontRels is where relationship names are stored.
var FontRels = struct {
	Sections string
}{
	Sections: "Sections",
}

// fontR is where relationships are stored.
type fontR struct {
	Sections SectionSlice `boil:"Sections" json:"Sections" toml:"Sections" yaml:"Sections"`
}

// NewStruct creates a new relationship struct
func (*fontR) NewStruct() *fontR {
	return &fontR{}
}

func (r *fontR) GetSections() SectionSlice {
	if r == nil {
		return nil
	}
	return r.Sections
}

// fontL is where Load methods for each relationship are stored.
type fontL struct{}

var (
	fontAllColumns            = []string{"font_id", "name", "created_at", "updated_at"}
	fontColumnsWithoutDefault = []string{"name", "updated_at"}
	fontColumnsWithDefault    = []string{"font_id", "created_at"}
	fontPrimaryKeyColumns     = []string{"font_id"}
	fontGeneratedColumns      = []string{}
)

type (
	// FontSlice is an alias for a slice of pointers to Font.
	// This should almost always be used instead of []Font.
	FontSlice []*Font
	// FontHook is the signature for custom Font hook methods
	FontHook func(context.Context, boil.ContextExecutor, *Font) error

	fontQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	fontType                 = reflect.TypeOf(&Font{})
	fontMapping              = queries.MakeStructMapping(fontType)
	fontPrimaryKeyMapping, _ = queries.BindMapping(fontType, fontMapping, fontPrimaryKeyColumns)
	fontInsertCacheMut       sync.RWMutex
	fontInsertCache          = make(map[string]insertCache)
	fontUpdateCacheMut       sync.RWMutex
	fontUpdateCache          = make(map[string]updateCache)
	fontUpsertCacheMut       sync.RWMutex
	fontUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var fontAfterSelectHooks []FontHook

var fontBeforeInsertHooks []FontHook
var fontAfterInsertHooks []FontHook

var fontBeforeUpdateHooks []FontHook
var fontAfterUpdateHooks []FontHook

var fontBeforeDeleteHooks []FontHook
var fontAfterDeleteHooks []FontHook

var fontBeforeUpsertHooks []FontHook
var fontAfterUpsertHooks []FontHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Font) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fontAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Font) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fontBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Font) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fontAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Font) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fontBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Font) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fontAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Font) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fontBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Font) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fontAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Font) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fontBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Font) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fontAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddFontHook registers your hook function for all future operations.
func AddFontHook(hookPoint boil.HookPoint, fontHook FontHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		fontAfterSelectHooks = append(fontAfterSelectHooks, fontHook)
	case boil.BeforeInsertHook:
		fontBeforeInsertHooks = append(fontBeforeInsertHooks, fontHook)
	case boil.AfterInsertHook:
		fontAfterInsertHooks = append(fontAfterInsertHooks, fontHook)
	case boil.BeforeUpdateHook:
		fontBeforeUpdateHooks = append(fontBeforeUpdateHooks, fontHook)
	case boil.AfterUpdateHook:
		fontAfterUpdateHooks = append(fontAfterUpdateHooks, fontHook)
	case boil.BeforeDeleteHook:
		fontBeforeDeleteHooks = append(fontBeforeDeleteHooks, fontHook)
	case boil.AfterDeleteHook:
		fontAfterDeleteHooks = append(fontAfterDeleteHooks, fontHook)
	case boil.BeforeUpsertHook:
		fontBeforeUpsertHooks = append(fontBeforeUpsertHooks, fontHook)
	case boil.AfterUpsertHook:
		fontAfterUpsertHooks = append(fontAfterUpsertHooks, fontHook)
	}
}

// One returns a single font record from the query.
func (q fontQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Font, error) {
	o := &Font{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for fonts")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Font records from the query.
func (q fontQuery) All(ctx context.Context, exec boil.ContextExecutor) (FontSlice, error) {
	var o []*Font

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Font slice")
	}

	if len(fontAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Font records in the query.
func (q fontQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count fonts rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q fontQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if fonts exists")
	}

	return count > 0, nil
}

// Sections retrieves all the section's Sections with an executor.
func (o *Font) Sections(mods ...qm.QueryMod) sectionQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"sections\".\"font_id\"=?", o.FontID),
	)

	return Sections(queryMods...)
}

// LoadSections allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (fontL) LoadSections(ctx context.Context, e boil.ContextExecutor, singular bool, maybeFont interface{}, mods queries.Applicator) error {
	var slice []*Font
	var object *Font

	if singular {
		var ok bool
		object, ok = maybeFont.(*Font)
		if !ok {
			object = new(Font)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeFont)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeFont))
			}
		}
	} else {
		s, ok := maybeFont.(*[]*Font)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeFont)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeFont))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &fontR{}
		}
		args = append(args, object.FontID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &fontR{}
			}

			for _, a := range args {
				if a == obj.FontID {
					continue Outer
				}
			}

			args = append(args, obj.FontID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`sections`),
		qm.WhereIn(`sections.font_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load sections")
	}

	var resultSlice []*Section
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice sections")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on sections")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for sections")
	}

	if len(sectionAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Sections = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &sectionR{}
			}
			foreign.R.Font = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.FontID == foreign.FontID {
				local.R.Sections = append(local.R.Sections, foreign)
				if foreign.R == nil {
					foreign.R = &sectionR{}
				}
				foreign.R.Font = local
				break
			}
		}
	}

	return nil
}

// AddSections adds the given related objects to the existing relationships
// of the font, optionally inserting them as new records.
// Appends related to o.R.Sections.
// Sets related.R.Font appropriately.
func (o *Font) AddSections(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Section) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.FontID = o.FontID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"sections\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"font_id"}),
				strmangle.WhereClause("\"", "\"", 2, sectionPrimaryKeyColumns),
			)
			values := []interface{}{o.FontID, rel.SectionID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.FontID = o.FontID
		}
	}

	if o.R == nil {
		o.R = &fontR{
			Sections: related,
		}
	} else {
		o.R.Sections = append(o.R.Sections, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &sectionR{
				Font: o,
			}
		} else {
			rel.R.Font = o
		}
	}
	return nil
}

// Fonts retrieves all the records using an executor.
func Fonts(mods ...qm.QueryMod) fontQuery {
	mods = append(mods, qm.From("\"fonts\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"fonts\".*"})
	}

	return fontQuery{q}
}

// FindFont retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindFont(ctx context.Context, exec boil.ContextExecutor, fontID int32, selectCols ...string) (*Font, error) {
	fontObj := &Font{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"fonts\" where \"font_id\"=$1", sel,
	)

	q := queries.Raw(query, fontID)

	err := q.Bind(ctx, exec, fontObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from fonts")
	}

	if err = fontObj.doAfterSelectHooks(ctx, exec); err != nil {
		return fontObj, err
	}

	return fontObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Font) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no fonts provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(fontColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	fontInsertCacheMut.RLock()
	cache, cached := fontInsertCache[key]
	fontInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			fontAllColumns,
			fontColumnsWithDefault,
			fontColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(fontType, fontMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(fontType, fontMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"fonts\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"fonts\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
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

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into fonts")
	}

	if !cached {
		fontInsertCacheMut.Lock()
		fontInsertCache[key] = cache
		fontInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Font.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Font) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	fontUpdateCacheMut.RLock()
	cache, cached := fontUpdateCache[key]
	fontUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			fontAllColumns,
			fontPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update fonts, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"fonts\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, fontPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(fontType, fontMapping, append(wl, fontPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update fonts row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for fonts")
	}

	if !cached {
		fontUpdateCacheMut.Lock()
		fontUpdateCache[key] = cache
		fontUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q fontQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for fonts")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for fonts")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o FontSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), fontPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"fonts\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, fontPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in font slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all font")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Font) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no fonts provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(fontColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
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
	key := buf.String()
	strmangle.PutBuffer(buf)

	fontUpsertCacheMut.RLock()
	cache, cached := fontUpsertCache[key]
	fontUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			fontAllColumns,
			fontColumnsWithDefault,
			fontColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			fontAllColumns,
			fontPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert fonts, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(fontPrimaryKeyColumns))
			copy(conflict, fontPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"fonts\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(fontType, fontMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(fontType, fontMapping, ret)
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
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert fonts")
	}

	if !cached {
		fontUpsertCacheMut.Lock()
		fontUpsertCache[key] = cache
		fontUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Font record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Font) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Font provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), fontPrimaryKeyMapping)
	sql := "DELETE FROM \"fonts\" WHERE \"font_id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from fonts")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for fonts")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q fontQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no fontQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from fonts")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for fonts")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o FontSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(fontBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), fontPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"fonts\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, fontPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from font slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for fonts")
	}

	if len(fontAfterDeleteHooks) != 0 {
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
func (o *Font) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindFont(ctx, exec, o.FontID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FontSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := FontSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), fontPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"fonts\".* FROM \"fonts\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, fontPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in FontSlice")
	}

	*o = slice

	return nil
}

// FontExists checks if the Font row exists.
func FontExists(ctx context.Context, exec boil.ContextExecutor, fontID int32) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"fonts\" where \"font_id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, fontID)
	}
	row := exec.QueryRowContext(ctx, sql, fontID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if fonts exists")
	}

	return exists, nil
}

// Exists checks if the Font row exists.
func (o *Font) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return FontExists(ctx, exec, o.FontID)
}
