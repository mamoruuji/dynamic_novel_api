// Code generated by SQLBoiler 4.14.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testPositions(t *testing.T) {
	t.Parallel()

	query := Positions()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testPositionsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Position{}
	if err = randomize.Struct(seed, o, positionDBTypes, true, positionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Position struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Positions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPositionsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Position{}
	if err = randomize.Struct(seed, o, positionDBTypes, true, positionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Position struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Positions().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Positions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPositionsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Position{}
	if err = randomize.Struct(seed, o, positionDBTypes, true, positionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Position struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PositionSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Positions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPositionsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Position{}
	if err = randomize.Struct(seed, o, positionDBTypes, true, positionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Position struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := PositionExists(ctx, tx, o.PositionID)
	if err != nil {
		t.Errorf("Unable to check if Position exists: %s", err)
	}
	if !e {
		t.Errorf("Expected PositionExists to return true, but got false.")
	}
}

func testPositionsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Position{}
	if err = randomize.Struct(seed, o, positionDBTypes, true, positionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Position struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	positionFound, err := FindPosition(ctx, tx, o.PositionID)
	if err != nil {
		t.Error(err)
	}

	if positionFound == nil {
		t.Error("want a record, got nil")
	}
}

func testPositionsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Position{}
	if err = randomize.Struct(seed, o, positionDBTypes, true, positionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Position struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Positions().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testPositionsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Position{}
	if err = randomize.Struct(seed, o, positionDBTypes, true, positionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Position struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Positions().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testPositionsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	positionOne := &Position{}
	positionTwo := &Position{}
	if err = randomize.Struct(seed, positionOne, positionDBTypes, false, positionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Position struct: %s", err)
	}
	if err = randomize.Struct(seed, positionTwo, positionDBTypes, false, positionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Position struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = positionOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = positionTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Positions().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testPositionsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	positionOne := &Position{}
	positionTwo := &Position{}
	if err = randomize.Struct(seed, positionOne, positionDBTypes, false, positionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Position struct: %s", err)
	}
	if err = randomize.Struct(seed, positionTwo, positionDBTypes, false, positionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Position struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = positionOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = positionTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Positions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func positionBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Position) error {
	*o = Position{}
	return nil
}

func positionAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Position) error {
	*o = Position{}
	return nil
}

func positionAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Position) error {
	*o = Position{}
	return nil
}

func positionBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Position) error {
	*o = Position{}
	return nil
}

func positionAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Position) error {
	*o = Position{}
	return nil
}

func positionBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Position) error {
	*o = Position{}
	return nil
}

func positionAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Position) error {
	*o = Position{}
	return nil
}

func positionBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Position) error {
	*o = Position{}
	return nil
}

func positionAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Position) error {
	*o = Position{}
	return nil
}

func testPositionsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Position{}
	o := &Position{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, positionDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Position object: %s", err)
	}

	AddPositionHook(boil.BeforeInsertHook, positionBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	positionBeforeInsertHooks = []PositionHook{}

	AddPositionHook(boil.AfterInsertHook, positionAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	positionAfterInsertHooks = []PositionHook{}

	AddPositionHook(boil.AfterSelectHook, positionAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	positionAfterSelectHooks = []PositionHook{}

	AddPositionHook(boil.BeforeUpdateHook, positionBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	positionBeforeUpdateHooks = []PositionHook{}

	AddPositionHook(boil.AfterUpdateHook, positionAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	positionAfterUpdateHooks = []PositionHook{}

	AddPositionHook(boil.BeforeDeleteHook, positionBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	positionBeforeDeleteHooks = []PositionHook{}

	AddPositionHook(boil.AfterDeleteHook, positionAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	positionAfterDeleteHooks = []PositionHook{}

	AddPositionHook(boil.BeforeUpsertHook, positionBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	positionBeforeUpsertHooks = []PositionHook{}

	AddPositionHook(boil.AfterUpsertHook, positionAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	positionAfterUpsertHooks = []PositionHook{}
}

func testPositionsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Position{}
	if err = randomize.Struct(seed, o, positionDBTypes, true, positionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Position struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Positions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPositionsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Position{}
	if err = randomize.Struct(seed, o, positionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Position struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(positionColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Positions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPositionToManySections(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Position
	var b, c Section

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, positionDBTypes, true, positionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Position struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, sectionDBTypes, false, sectionColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, sectionDBTypes, false, sectionColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.PositionID, a.PositionID)
	queries.Assign(&c.PositionID, a.PositionID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Sections().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if queries.Equal(v.PositionID, b.PositionID) {
			bFound = true
		}
		if queries.Equal(v.PositionID, c.PositionID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := PositionSlice{&a}
	if err = a.L.LoadSections(ctx, tx, false, (*[]*Position)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Sections); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Sections = nil
	if err = a.L.LoadSections(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Sections); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testPositionToManyAddOpSections(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Position
	var b, c, d, e Section

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, positionDBTypes, false, strmangle.SetComplement(positionPrimaryKeyColumns, positionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Section{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, sectionDBTypes, false, strmangle.SetComplement(sectionPrimaryKeyColumns, sectionColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Section{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddSections(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.PositionID, first.PositionID) {
			t.Error("foreign key was wrong value", a.PositionID, first.PositionID)
		}
		if !queries.Equal(a.PositionID, second.PositionID) {
			t.Error("foreign key was wrong value", a.PositionID, second.PositionID)
		}

		if first.R.Position != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Position != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Sections[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Sections[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Sections().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testPositionToManySetOpSections(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Position
	var b, c, d, e Section

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, positionDBTypes, false, strmangle.SetComplement(positionPrimaryKeyColumns, positionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Section{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, sectionDBTypes, false, strmangle.SetComplement(sectionPrimaryKeyColumns, sectionColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.SetSections(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Sections().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetSections(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Sections().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.PositionID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.PositionID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.PositionID, d.PositionID) {
		t.Error("foreign key was wrong value", a.PositionID, d.PositionID)
	}
	if !queries.Equal(a.PositionID, e.PositionID) {
		t.Error("foreign key was wrong value", a.PositionID, e.PositionID)
	}

	if b.R.Position != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Position != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Position != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Position != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.Sections[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.Sections[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testPositionToManyRemoveOpSections(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Position
	var b, c, d, e Section

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, positionDBTypes, false, strmangle.SetComplement(positionPrimaryKeyColumns, positionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Section{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, sectionDBTypes, false, strmangle.SetComplement(sectionPrimaryKeyColumns, sectionColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddSections(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Sections().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveSections(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Sections().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.PositionID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.PositionID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.Position != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Position != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Position != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.Position != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.Sections) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.Sections[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.Sections[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testPositionsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Position{}
	if err = randomize.Struct(seed, o, positionDBTypes, true, positionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Position struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testPositionsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Position{}
	if err = randomize.Struct(seed, o, positionDBTypes, true, positionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Position struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PositionSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testPositionsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Position{}
	if err = randomize.Struct(seed, o, positionDBTypes, true, positionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Position struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Positions().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	positionDBTypes = map[string]string{`PositionID`: `integer`, `Name`: `character varying`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`}
	_               = bytes.MinRead
)

func testPositionsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(positionPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(positionAllColumns) == len(positionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Position{}
	if err = randomize.Struct(seed, o, positionDBTypes, true, positionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Position struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Positions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, positionDBTypes, true, positionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Position struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testPositionsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(positionAllColumns) == len(positionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Position{}
	if err = randomize.Struct(seed, o, positionDBTypes, true, positionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Position struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Positions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, positionDBTypes, true, positionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Position struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(positionAllColumns, positionPrimaryKeyColumns) {
		fields = positionAllColumns
	} else {
		fields = strmangle.SetComplement(
			positionAllColumns,
			positionPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := PositionSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testPositionsUpsert(t *testing.T) {
	t.Parallel()

	if len(positionAllColumns) == len(positionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Position{}
	if err = randomize.Struct(seed, &o, positionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Position struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Position: %s", err)
	}

	count, err := Positions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, positionDBTypes, false, positionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Position struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Position: %s", err)
	}

	count, err = Positions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
