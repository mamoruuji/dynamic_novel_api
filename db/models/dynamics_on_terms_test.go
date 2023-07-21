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

func testDynamicsOnTerms(t *testing.T) {
	t.Parallel()

	query := DynamicsOnTerms()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testDynamicsOnTermsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DynamicsOnTerm{}
	if err = randomize.Struct(seed, o, dynamicsOnTermDBTypes, true, dynamicsOnTermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicsOnTerm struct: %s", err)
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

	count, err := DynamicsOnTerms().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDynamicsOnTermsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DynamicsOnTerm{}
	if err = randomize.Struct(seed, o, dynamicsOnTermDBTypes, true, dynamicsOnTermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicsOnTerm struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := DynamicsOnTerms().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DynamicsOnTerms().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDynamicsOnTermsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DynamicsOnTerm{}
	if err = randomize.Struct(seed, o, dynamicsOnTermDBTypes, true, dynamicsOnTermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicsOnTerm struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DynamicsOnTermSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DynamicsOnTerms().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDynamicsOnTermsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DynamicsOnTerm{}
	if err = randomize.Struct(seed, o, dynamicsOnTermDBTypes, true, dynamicsOnTermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicsOnTerm struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := DynamicsOnTermExists(ctx, tx, o.DynamicsOnTermsID)
	if err != nil {
		t.Errorf("Unable to check if DynamicsOnTerm exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DynamicsOnTermExists to return true, but got false.")
	}
}

func testDynamicsOnTermsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DynamicsOnTerm{}
	if err = randomize.Struct(seed, o, dynamicsOnTermDBTypes, true, dynamicsOnTermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicsOnTerm struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	dynamicsOnTermFound, err := FindDynamicsOnTerm(ctx, tx, o.DynamicsOnTermsID)
	if err != nil {
		t.Error(err)
	}

	if dynamicsOnTermFound == nil {
		t.Error("want a record, got nil")
	}
}

func testDynamicsOnTermsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DynamicsOnTerm{}
	if err = randomize.Struct(seed, o, dynamicsOnTermDBTypes, true, dynamicsOnTermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicsOnTerm struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = DynamicsOnTerms().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testDynamicsOnTermsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DynamicsOnTerm{}
	if err = randomize.Struct(seed, o, dynamicsOnTermDBTypes, true, dynamicsOnTermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicsOnTerm struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := DynamicsOnTerms().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDynamicsOnTermsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dynamicsOnTermOne := &DynamicsOnTerm{}
	dynamicsOnTermTwo := &DynamicsOnTerm{}
	if err = randomize.Struct(seed, dynamicsOnTermOne, dynamicsOnTermDBTypes, false, dynamicsOnTermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicsOnTerm struct: %s", err)
	}
	if err = randomize.Struct(seed, dynamicsOnTermTwo, dynamicsOnTermDBTypes, false, dynamicsOnTermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicsOnTerm struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = dynamicsOnTermOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = dynamicsOnTermTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DynamicsOnTerms().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDynamicsOnTermsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	dynamicsOnTermOne := &DynamicsOnTerm{}
	dynamicsOnTermTwo := &DynamicsOnTerm{}
	if err = randomize.Struct(seed, dynamicsOnTermOne, dynamicsOnTermDBTypes, false, dynamicsOnTermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicsOnTerm struct: %s", err)
	}
	if err = randomize.Struct(seed, dynamicsOnTermTwo, dynamicsOnTermDBTypes, false, dynamicsOnTermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicsOnTerm struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = dynamicsOnTermOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = dynamicsOnTermTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DynamicsOnTerms().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func dynamicsOnTermBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *DynamicsOnTerm) error {
	*o = DynamicsOnTerm{}
	return nil
}

func dynamicsOnTermAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *DynamicsOnTerm) error {
	*o = DynamicsOnTerm{}
	return nil
}

func dynamicsOnTermAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *DynamicsOnTerm) error {
	*o = DynamicsOnTerm{}
	return nil
}

func dynamicsOnTermBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *DynamicsOnTerm) error {
	*o = DynamicsOnTerm{}
	return nil
}

func dynamicsOnTermAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *DynamicsOnTerm) error {
	*o = DynamicsOnTerm{}
	return nil
}

func dynamicsOnTermBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *DynamicsOnTerm) error {
	*o = DynamicsOnTerm{}
	return nil
}

func dynamicsOnTermAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *DynamicsOnTerm) error {
	*o = DynamicsOnTerm{}
	return nil
}

func dynamicsOnTermBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *DynamicsOnTerm) error {
	*o = DynamicsOnTerm{}
	return nil
}

func dynamicsOnTermAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *DynamicsOnTerm) error {
	*o = DynamicsOnTerm{}
	return nil
}

func testDynamicsOnTermsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &DynamicsOnTerm{}
	o := &DynamicsOnTerm{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, dynamicsOnTermDBTypes, false); err != nil {
		t.Errorf("Unable to randomize DynamicsOnTerm object: %s", err)
	}

	AddDynamicsOnTermHook(boil.BeforeInsertHook, dynamicsOnTermBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	dynamicsOnTermBeforeInsertHooks = []DynamicsOnTermHook{}

	AddDynamicsOnTermHook(boil.AfterInsertHook, dynamicsOnTermAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	dynamicsOnTermAfterInsertHooks = []DynamicsOnTermHook{}

	AddDynamicsOnTermHook(boil.AfterSelectHook, dynamicsOnTermAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	dynamicsOnTermAfterSelectHooks = []DynamicsOnTermHook{}

	AddDynamicsOnTermHook(boil.BeforeUpdateHook, dynamicsOnTermBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	dynamicsOnTermBeforeUpdateHooks = []DynamicsOnTermHook{}

	AddDynamicsOnTermHook(boil.AfterUpdateHook, dynamicsOnTermAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	dynamicsOnTermAfterUpdateHooks = []DynamicsOnTermHook{}

	AddDynamicsOnTermHook(boil.BeforeDeleteHook, dynamicsOnTermBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	dynamicsOnTermBeforeDeleteHooks = []DynamicsOnTermHook{}

	AddDynamicsOnTermHook(boil.AfterDeleteHook, dynamicsOnTermAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	dynamicsOnTermAfterDeleteHooks = []DynamicsOnTermHook{}

	AddDynamicsOnTermHook(boil.BeforeUpsertHook, dynamicsOnTermBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	dynamicsOnTermBeforeUpsertHooks = []DynamicsOnTermHook{}

	AddDynamicsOnTermHook(boil.AfterUpsertHook, dynamicsOnTermAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	dynamicsOnTermAfterUpsertHooks = []DynamicsOnTermHook{}
}

func testDynamicsOnTermsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DynamicsOnTerm{}
	if err = randomize.Struct(seed, o, dynamicsOnTermDBTypes, true, dynamicsOnTermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicsOnTerm struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DynamicsOnTerms().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDynamicsOnTermsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DynamicsOnTerm{}
	if err = randomize.Struct(seed, o, dynamicsOnTermDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DynamicsOnTerm struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(dynamicsOnTermColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := DynamicsOnTerms().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDynamicsOnTermToOneDynamicUsingDynamic(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local DynamicsOnTerm
	var foreign Dynamic

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, dynamicsOnTermDBTypes, false, dynamicsOnTermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicsOnTerm struct: %s", err)
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

	slice := DynamicsOnTermSlice{&local}
	if err = local.L.LoadDynamic(ctx, tx, false, (*[]*DynamicsOnTerm)(&slice), nil); err != nil {
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

func testDynamicsOnTermToOneTermUsingTerm(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local DynamicsOnTerm
	var foreign Term

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, dynamicsOnTermDBTypes, false, dynamicsOnTermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicsOnTerm struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, termDBTypes, false, termColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Term struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.TermID = foreign.TermID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Term().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.TermID != foreign.TermID {
		t.Errorf("want: %v, got %v", foreign.TermID, check.TermID)
	}

	ranAfterSelectHook := false
	AddTermHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *Term) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := DynamicsOnTermSlice{&local}
	if err = local.L.LoadTerm(ctx, tx, false, (*[]*DynamicsOnTerm)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Term == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Term = nil
	if err = local.L.LoadTerm(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Term == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testDynamicsOnTermToOneSetOpDynamicUsingDynamic(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DynamicsOnTerm
	var b, c Dynamic

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dynamicsOnTermDBTypes, false, strmangle.SetComplement(dynamicsOnTermPrimaryKeyColumns, dynamicsOnTermColumnsWithoutDefault)...); err != nil {
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

		if x.R.DynamicsOnTerms[0] != &a {
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
func testDynamicsOnTermToOneSetOpTermUsingTerm(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DynamicsOnTerm
	var b, c Term

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dynamicsOnTermDBTypes, false, strmangle.SetComplement(dynamicsOnTermPrimaryKeyColumns, dynamicsOnTermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, termDBTypes, false, strmangle.SetComplement(termPrimaryKeyColumns, termColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, termDBTypes, false, strmangle.SetComplement(termPrimaryKeyColumns, termColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Term{&b, &c} {
		err = a.SetTerm(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Term != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.DynamicsOnTerms[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.TermID != x.TermID {
			t.Error("foreign key was wrong value", a.TermID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.TermID))
		reflect.Indirect(reflect.ValueOf(&a.TermID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.TermID != x.TermID {
			t.Error("foreign key was wrong value", a.TermID, x.TermID)
		}
	}
}

func testDynamicsOnTermsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DynamicsOnTerm{}
	if err = randomize.Struct(seed, o, dynamicsOnTermDBTypes, true, dynamicsOnTermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicsOnTerm struct: %s", err)
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

func testDynamicsOnTermsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DynamicsOnTerm{}
	if err = randomize.Struct(seed, o, dynamicsOnTermDBTypes, true, dynamicsOnTermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicsOnTerm struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DynamicsOnTermSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testDynamicsOnTermsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DynamicsOnTerm{}
	if err = randomize.Struct(seed, o, dynamicsOnTermDBTypes, true, dynamicsOnTermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicsOnTerm struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DynamicsOnTerms().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	dynamicsOnTermDBTypes = map[string]string{`DynamicsOnTermsID`: `integer`, `DynamicID`: `integer`, `TermID`: `integer`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`}
	_                     = bytes.MinRead
)

func testDynamicsOnTermsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(dynamicsOnTermPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(dynamicsOnTermAllColumns) == len(dynamicsOnTermPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DynamicsOnTerm{}
	if err = randomize.Struct(seed, o, dynamicsOnTermDBTypes, true, dynamicsOnTermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicsOnTerm struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DynamicsOnTerms().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, dynamicsOnTermDBTypes, true, dynamicsOnTermPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DynamicsOnTerm struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testDynamicsOnTermsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(dynamicsOnTermAllColumns) == len(dynamicsOnTermPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DynamicsOnTerm{}
	if err = randomize.Struct(seed, o, dynamicsOnTermDBTypes, true, dynamicsOnTermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicsOnTerm struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DynamicsOnTerms().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, dynamicsOnTermDBTypes, true, dynamicsOnTermPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DynamicsOnTerm struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(dynamicsOnTermAllColumns, dynamicsOnTermPrimaryKeyColumns) {
		fields = dynamicsOnTermAllColumns
	} else {
		fields = strmangle.SetComplement(
			dynamicsOnTermAllColumns,
			dynamicsOnTermPrimaryKeyColumns,
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

	slice := DynamicsOnTermSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testDynamicsOnTermsUpsert(t *testing.T) {
	t.Parallel()

	if len(dynamicsOnTermAllColumns) == len(dynamicsOnTermPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := DynamicsOnTerm{}
	if err = randomize.Struct(seed, &o, dynamicsOnTermDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DynamicsOnTerm struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DynamicsOnTerm: %s", err)
	}

	count, err := DynamicsOnTerms().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, dynamicsOnTermDBTypes, false, dynamicsOnTermPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DynamicsOnTerm struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DynamicsOnTerm: %s", err)
	}

	count, err = DynamicsOnTerms().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
