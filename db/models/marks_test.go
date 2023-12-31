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

func testMarks(t *testing.T) {
	t.Parallel()

	query := Marks()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testMarksDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Mark{}
	if err = randomize.Struct(seed, o, markDBTypes, true, markColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Mark struct: %s", err)
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

	count, err := Marks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMarksQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Mark{}
	if err = randomize.Struct(seed, o, markDBTypes, true, markColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Mark struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Marks().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Marks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMarksSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Mark{}
	if err = randomize.Struct(seed, o, markDBTypes, true, markColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Mark struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := MarkSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Marks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMarksExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Mark{}
	if err = randomize.Struct(seed, o, markDBTypes, true, markColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Mark struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := MarkExists(ctx, tx, o.MarkID)
	if err != nil {
		t.Errorf("Unable to check if Mark exists: %s", err)
	}
	if !e {
		t.Errorf("Expected MarkExists to return true, but got false.")
	}
}

func testMarksFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Mark{}
	if err = randomize.Struct(seed, o, markDBTypes, true, markColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Mark struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	markFound, err := FindMark(ctx, tx, o.MarkID)
	if err != nil {
		t.Error(err)
	}

	if markFound == nil {
		t.Error("want a record, got nil")
	}
}

func testMarksBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Mark{}
	if err = randomize.Struct(seed, o, markDBTypes, true, markColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Mark struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Marks().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testMarksOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Mark{}
	if err = randomize.Struct(seed, o, markDBTypes, true, markColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Mark struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Marks().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testMarksAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	markOne := &Mark{}
	markTwo := &Mark{}
	if err = randomize.Struct(seed, markOne, markDBTypes, false, markColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Mark struct: %s", err)
	}
	if err = randomize.Struct(seed, markTwo, markDBTypes, false, markColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Mark struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = markOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = markTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Marks().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testMarksCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	markOne := &Mark{}
	markTwo := &Mark{}
	if err = randomize.Struct(seed, markOne, markDBTypes, false, markColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Mark struct: %s", err)
	}
	if err = randomize.Struct(seed, markTwo, markDBTypes, false, markColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Mark struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = markOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = markTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Marks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func markBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Mark) error {
	*o = Mark{}
	return nil
}

func markAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Mark) error {
	*o = Mark{}
	return nil
}

func markAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Mark) error {
	*o = Mark{}
	return nil
}

func markBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Mark) error {
	*o = Mark{}
	return nil
}

func markAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Mark) error {
	*o = Mark{}
	return nil
}

func markBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Mark) error {
	*o = Mark{}
	return nil
}

func markAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Mark) error {
	*o = Mark{}
	return nil
}

func markBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Mark) error {
	*o = Mark{}
	return nil
}

func markAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Mark) error {
	*o = Mark{}
	return nil
}

func testMarksHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Mark{}
	o := &Mark{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, markDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Mark object: %s", err)
	}

	AddMarkHook(boil.BeforeInsertHook, markBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	markBeforeInsertHooks = []MarkHook{}

	AddMarkHook(boil.AfterInsertHook, markAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	markAfterInsertHooks = []MarkHook{}

	AddMarkHook(boil.AfterSelectHook, markAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	markAfterSelectHooks = []MarkHook{}

	AddMarkHook(boil.BeforeUpdateHook, markBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	markBeforeUpdateHooks = []MarkHook{}

	AddMarkHook(boil.AfterUpdateHook, markAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	markAfterUpdateHooks = []MarkHook{}

	AddMarkHook(boil.BeforeDeleteHook, markBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	markBeforeDeleteHooks = []MarkHook{}

	AddMarkHook(boil.AfterDeleteHook, markAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	markAfterDeleteHooks = []MarkHook{}

	AddMarkHook(boil.BeforeUpsertHook, markBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	markBeforeUpsertHooks = []MarkHook{}

	AddMarkHook(boil.AfterUpsertHook, markAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	markAfterUpsertHooks = []MarkHook{}
}

func testMarksInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Mark{}
	if err = randomize.Struct(seed, o, markDBTypes, true, markColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Mark struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Marks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMarksInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Mark{}
	if err = randomize.Struct(seed, o, markDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Mark struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(markColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Marks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMarkToOneDynamicUsingDynamic(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Mark
	var foreign Dynamic

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, markDBTypes, false, markColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Mark struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, dynamicDBTypes, false, dynamicColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Dynamic struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.DynamicID = foreign.DynamicID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Dynamic().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.DynamicID != foreign.DynamicID {
		t.Errorf("want: %v, got %v", foreign.DynamicID, check.DynamicID)
	}

	ranAfterSelectHook := false
	AddDynamicHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *Dynamic) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := MarkSlice{&local}
	if err = local.L.LoadDynamic(ctx, tx, false, (*[]*Mark)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Dynamic == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Dynamic = nil
	if err = local.L.LoadDynamic(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Dynamic == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testMarkToOneUserUsingUser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Mark
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, markDBTypes, false, markColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Mark struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.UserID = foreign.UserID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.User().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.UserID != foreign.UserID {
		t.Errorf("want: %v, got %v", foreign.UserID, check.UserID)
	}

	ranAfterSelectHook := false
	AddUserHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *User) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := MarkSlice{&local}
	if err = local.L.LoadUser(ctx, tx, false, (*[]*Mark)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.User = nil
	if err = local.L.LoadUser(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testMarkToOneSetOpDynamicUsingDynamic(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Mark
	var b, c Dynamic

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, markDBTypes, false, strmangle.SetComplement(markPrimaryKeyColumns, markColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, dynamicDBTypes, false, strmangle.SetComplement(dynamicPrimaryKeyColumns, dynamicColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, dynamicDBTypes, false, strmangle.SetComplement(dynamicPrimaryKeyColumns, dynamicColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Dynamic{&b, &c} {
		err = a.SetDynamic(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Dynamic != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Marks[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.DynamicID != x.DynamicID {
			t.Error("foreign key was wrong value", a.DynamicID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.DynamicID))
		reflect.Indirect(reflect.ValueOf(&a.DynamicID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.DynamicID != x.DynamicID {
			t.Error("foreign key was wrong value", a.DynamicID, x.DynamicID)
		}
	}
}
func testMarkToOneSetOpUserUsingUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Mark
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, markDBTypes, false, strmangle.SetComplement(markPrimaryKeyColumns, markColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*User{&b, &c} {
		err = a.SetUser(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.User != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Marks[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.UserID != x.UserID {
			t.Error("foreign key was wrong value", a.UserID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.UserID))
		reflect.Indirect(reflect.ValueOf(&a.UserID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.UserID != x.UserID {
			t.Error("foreign key was wrong value", a.UserID, x.UserID)
		}
	}
}

func testMarksReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Mark{}
	if err = randomize.Struct(seed, o, markDBTypes, true, markColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Mark struct: %s", err)
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

func testMarksReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Mark{}
	if err = randomize.Struct(seed, o, markDBTypes, true, markColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Mark struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := MarkSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testMarksSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Mark{}
	if err = randomize.Struct(seed, o, markDBTypes, true, markColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Mark struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Marks().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	markDBTypes = map[string]string{`MarkID`: `integer`, `UserID`: `text`, `DynamicID`: `integer`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`}
	_           = bytes.MinRead
)

func testMarksUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(markPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(markAllColumns) == len(markPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Mark{}
	if err = randomize.Struct(seed, o, markDBTypes, true, markColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Mark struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Marks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, markDBTypes, true, markPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Mark struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testMarksSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(markAllColumns) == len(markPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Mark{}
	if err = randomize.Struct(seed, o, markDBTypes, true, markColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Mark struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Marks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, markDBTypes, true, markPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Mark struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(markAllColumns, markPrimaryKeyColumns) {
		fields = markAllColumns
	} else {
		fields = strmangle.SetComplement(
			markAllColumns,
			markPrimaryKeyColumns,
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

	slice := MarkSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testMarksUpsert(t *testing.T) {
	t.Parallel()

	if len(markAllColumns) == len(markPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Mark{}
	if err = randomize.Struct(seed, &o, markDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Mark struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Mark: %s", err)
	}

	count, err := Marks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, markDBTypes, false, markPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Mark struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Mark: %s", err)
	}

	count, err = Marks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
