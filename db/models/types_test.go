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

func testTypes(t *testing.T) {
	t.Parallel()

	query := Types()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testTypesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Type{}
	if err = randomize.Struct(seed, o, typeDBTypes, true, typeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Type struct: %s", err)
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

	count, err := Types().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTypesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Type{}
	if err = randomize.Struct(seed, o, typeDBTypes, true, typeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Type struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Types().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Types().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTypesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Type{}
	if err = randomize.Struct(seed, o, typeDBTypes, true, typeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Type struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := TypeSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Types().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTypesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Type{}
	if err = randomize.Struct(seed, o, typeDBTypes, true, typeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Type struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := TypeExists(ctx, tx, o.TypeID)
	if err != nil {
		t.Errorf("Unable to check if Type exists: %s", err)
	}
	if !e {
		t.Errorf("Expected TypeExists to return true, but got false.")
	}
}

func testTypesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Type{}
	if err = randomize.Struct(seed, o, typeDBTypes, true, typeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Type struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	typeFound, err := FindType(ctx, tx, o.TypeID)
	if err != nil {
		t.Error(err)
	}

	if typeFound == nil {
		t.Error("want a record, got nil")
	}
}

func testTypesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Type{}
	if err = randomize.Struct(seed, o, typeDBTypes, true, typeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Type struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Types().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testTypesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Type{}
	if err = randomize.Struct(seed, o, typeDBTypes, true, typeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Type struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Types().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testTypesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	typeOne := &Type{}
	typeTwo := &Type{}
	if err = randomize.Struct(seed, typeOne, typeDBTypes, false, typeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Type struct: %s", err)
	}
	if err = randomize.Struct(seed, typeTwo, typeDBTypes, false, typeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Type struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = typeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = typeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Types().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testTypesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	typeOne := &Type{}
	typeTwo := &Type{}
	if err = randomize.Struct(seed, typeOne, typeDBTypes, false, typeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Type struct: %s", err)
	}
	if err = randomize.Struct(seed, typeTwo, typeDBTypes, false, typeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Type struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = typeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = typeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Types().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func typeBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Type) error {
	*o = Type{}
	return nil
}

func typeAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Type) error {
	*o = Type{}
	return nil
}

func typeAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Type) error {
	*o = Type{}
	return nil
}

func typeBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Type) error {
	*o = Type{}
	return nil
}

func typeAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Type) error {
	*o = Type{}
	return nil
}

func typeBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Type) error {
	*o = Type{}
	return nil
}

func typeAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Type) error {
	*o = Type{}
	return nil
}

func typeBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Type) error {
	*o = Type{}
	return nil
}

func typeAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Type) error {
	*o = Type{}
	return nil
}

func testTypesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Type{}
	o := &Type{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, typeDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Type object: %s", err)
	}

	AddTypeHook(boil.BeforeInsertHook, typeBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	typeBeforeInsertHooks = []TypeHook{}

	AddTypeHook(boil.AfterInsertHook, typeAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	typeAfterInsertHooks = []TypeHook{}

	AddTypeHook(boil.AfterSelectHook, typeAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	typeAfterSelectHooks = []TypeHook{}

	AddTypeHook(boil.BeforeUpdateHook, typeBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	typeBeforeUpdateHooks = []TypeHook{}

	AddTypeHook(boil.AfterUpdateHook, typeAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	typeAfterUpdateHooks = []TypeHook{}

	AddTypeHook(boil.BeforeDeleteHook, typeBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	typeBeforeDeleteHooks = []TypeHook{}

	AddTypeHook(boil.AfterDeleteHook, typeAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	typeAfterDeleteHooks = []TypeHook{}

	AddTypeHook(boil.BeforeUpsertHook, typeBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	typeBeforeUpsertHooks = []TypeHook{}

	AddTypeHook(boil.AfterUpsertHook, typeAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	typeAfterUpsertHooks = []TypeHook{}
}

func testTypesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Type{}
	if err = randomize.Struct(seed, o, typeDBTypes, true, typeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Type struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Types().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTypesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Type{}
	if err = randomize.Struct(seed, o, typeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Type struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(typeColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Types().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTypeToManySections(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Type
	var b, c Section

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, typeDBTypes, true, typeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Type struct: %s", err)
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

	queries.Assign(&b.TypeID, a.TypeID)
	queries.Assign(&c.TypeID, a.TypeID)
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
		if queries.Equal(v.TypeID, b.TypeID) {
			bFound = true
		}
		if queries.Equal(v.TypeID, c.TypeID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := TypeSlice{&a}
	if err = a.L.LoadSections(ctx, tx, false, (*[]*Type)(&slice), nil); err != nil {
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

func testTypeToManyAddOpSections(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Type
	var b, c, d, e Section

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, typeDBTypes, false, strmangle.SetComplement(typePrimaryKeyColumns, typeColumnsWithoutDefault)...); err != nil {
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

		if !queries.Equal(a.TypeID, first.TypeID) {
			t.Error("foreign key was wrong value", a.TypeID, first.TypeID)
		}
		if !queries.Equal(a.TypeID, second.TypeID) {
			t.Error("foreign key was wrong value", a.TypeID, second.TypeID)
		}

		if first.R.Type != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Type != &a {
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

func testTypeToManySetOpSections(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Type
	var b, c, d, e Section

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, typeDBTypes, false, strmangle.SetComplement(typePrimaryKeyColumns, typeColumnsWithoutDefault)...); err != nil {
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

	if !queries.IsValuerNil(b.TypeID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.TypeID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.TypeID, d.TypeID) {
		t.Error("foreign key was wrong value", a.TypeID, d.TypeID)
	}
	if !queries.Equal(a.TypeID, e.TypeID) {
		t.Error("foreign key was wrong value", a.TypeID, e.TypeID)
	}

	if b.R.Type != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Type != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Type != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Type != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.Sections[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.Sections[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testTypeToManyRemoveOpSections(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Type
	var b, c, d, e Section

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, typeDBTypes, false, strmangle.SetComplement(typePrimaryKeyColumns, typeColumnsWithoutDefault)...); err != nil {
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

	if !queries.IsValuerNil(b.TypeID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.TypeID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.Type != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Type != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Type != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.Type != &a {
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

func testTypesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Type{}
	if err = randomize.Struct(seed, o, typeDBTypes, true, typeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Type struct: %s", err)
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

func testTypesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Type{}
	if err = randomize.Struct(seed, o, typeDBTypes, true, typeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Type struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := TypeSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testTypesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Type{}
	if err = randomize.Struct(seed, o, typeDBTypes, true, typeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Type struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Types().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	typeDBTypes = map[string]string{`TypeID`: `integer`, `Name`: `character varying`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`}
	_           = bytes.MinRead
)

func testTypesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(typePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(typeAllColumns) == len(typePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Type{}
	if err = randomize.Struct(seed, o, typeDBTypes, true, typeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Type struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Types().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, typeDBTypes, true, typePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Type struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testTypesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(typeAllColumns) == len(typePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Type{}
	if err = randomize.Struct(seed, o, typeDBTypes, true, typeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Type struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Types().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, typeDBTypes, true, typePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Type struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(typeAllColumns, typePrimaryKeyColumns) {
		fields = typeAllColumns
	} else {
		fields = strmangle.SetComplement(
			typeAllColumns,
			typePrimaryKeyColumns,
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

	slice := TypeSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testTypesUpsert(t *testing.T) {
	t.Parallel()

	if len(typeAllColumns) == len(typePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Type{}
	if err = randomize.Struct(seed, &o, typeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Type struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Type: %s", err)
	}

	count, err := Types().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, typeDBTypes, false, typePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Type struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Type: %s", err)
	}

	count, err = Types().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
