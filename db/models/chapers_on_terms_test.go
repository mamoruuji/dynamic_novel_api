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

func testChapersOnTerms(t *testing.T) {
	t.Parallel()

	query := ChapersOnTerms()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testChapersOnTermsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ChapersOnTerm{}
	if err = randomize.Struct(seed, o, chapersOnTermDBTypes, true, chapersOnTermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ChapersOnTerm struct: %s", err)
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

	count, err := ChapersOnTerms().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testChapersOnTermsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ChapersOnTerm{}
	if err = randomize.Struct(seed, o, chapersOnTermDBTypes, true, chapersOnTermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ChapersOnTerm struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := ChapersOnTerms().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ChapersOnTerms().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testChapersOnTermsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ChapersOnTerm{}
	if err = randomize.Struct(seed, o, chapersOnTermDBTypes, true, chapersOnTermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ChapersOnTerm struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ChapersOnTermSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ChapersOnTerms().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testChapersOnTermsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ChapersOnTerm{}
	if err = randomize.Struct(seed, o, chapersOnTermDBTypes, true, chapersOnTermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ChapersOnTerm struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ChapersOnTermExists(ctx, tx, o.ChapersOnTermsID)
	if err != nil {
		t.Errorf("Unable to check if ChapersOnTerm exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ChapersOnTermExists to return true, but got false.")
	}
}

func testChapersOnTermsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ChapersOnTerm{}
	if err = randomize.Struct(seed, o, chapersOnTermDBTypes, true, chapersOnTermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ChapersOnTerm struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	chapersOnTermFound, err := FindChapersOnTerm(ctx, tx, o.ChapersOnTermsID)
	if err != nil {
		t.Error(err)
	}

	if chapersOnTermFound == nil {
		t.Error("want a record, got nil")
	}
}

func testChapersOnTermsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ChapersOnTerm{}
	if err = randomize.Struct(seed, o, chapersOnTermDBTypes, true, chapersOnTermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ChapersOnTerm struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = ChapersOnTerms().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testChapersOnTermsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ChapersOnTerm{}
	if err = randomize.Struct(seed, o, chapersOnTermDBTypes, true, chapersOnTermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ChapersOnTerm struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := ChapersOnTerms().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testChapersOnTermsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	chapersOnTermOne := &ChapersOnTerm{}
	chapersOnTermTwo := &ChapersOnTerm{}
	if err = randomize.Struct(seed, chapersOnTermOne, chapersOnTermDBTypes, false, chapersOnTermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ChapersOnTerm struct: %s", err)
	}
	if err = randomize.Struct(seed, chapersOnTermTwo, chapersOnTermDBTypes, false, chapersOnTermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ChapersOnTerm struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = chapersOnTermOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = chapersOnTermTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ChapersOnTerms().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testChapersOnTermsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	chapersOnTermOne := &ChapersOnTerm{}
	chapersOnTermTwo := &ChapersOnTerm{}
	if err = randomize.Struct(seed, chapersOnTermOne, chapersOnTermDBTypes, false, chapersOnTermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ChapersOnTerm struct: %s", err)
	}
	if err = randomize.Struct(seed, chapersOnTermTwo, chapersOnTermDBTypes, false, chapersOnTermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ChapersOnTerm struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = chapersOnTermOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = chapersOnTermTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ChapersOnTerms().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func chapersOnTermBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *ChapersOnTerm) error {
	*o = ChapersOnTerm{}
	return nil
}

func chapersOnTermAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *ChapersOnTerm) error {
	*o = ChapersOnTerm{}
	return nil
}

func chapersOnTermAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *ChapersOnTerm) error {
	*o = ChapersOnTerm{}
	return nil
}

func chapersOnTermBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ChapersOnTerm) error {
	*o = ChapersOnTerm{}
	return nil
}

func chapersOnTermAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ChapersOnTerm) error {
	*o = ChapersOnTerm{}
	return nil
}

func chapersOnTermBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ChapersOnTerm) error {
	*o = ChapersOnTerm{}
	return nil
}

func chapersOnTermAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ChapersOnTerm) error {
	*o = ChapersOnTerm{}
	return nil
}

func chapersOnTermBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ChapersOnTerm) error {
	*o = ChapersOnTerm{}
	return nil
}

func chapersOnTermAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ChapersOnTerm) error {
	*o = ChapersOnTerm{}
	return nil
}

func testChapersOnTermsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &ChapersOnTerm{}
	o := &ChapersOnTerm{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, chapersOnTermDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ChapersOnTerm object: %s", err)
	}

	AddChapersOnTermHook(boil.BeforeInsertHook, chapersOnTermBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	chapersOnTermBeforeInsertHooks = []ChapersOnTermHook{}

	AddChapersOnTermHook(boil.AfterInsertHook, chapersOnTermAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	chapersOnTermAfterInsertHooks = []ChapersOnTermHook{}

	AddChapersOnTermHook(boil.AfterSelectHook, chapersOnTermAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	chapersOnTermAfterSelectHooks = []ChapersOnTermHook{}

	AddChapersOnTermHook(boil.BeforeUpdateHook, chapersOnTermBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	chapersOnTermBeforeUpdateHooks = []ChapersOnTermHook{}

	AddChapersOnTermHook(boil.AfterUpdateHook, chapersOnTermAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	chapersOnTermAfterUpdateHooks = []ChapersOnTermHook{}

	AddChapersOnTermHook(boil.BeforeDeleteHook, chapersOnTermBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	chapersOnTermBeforeDeleteHooks = []ChapersOnTermHook{}

	AddChapersOnTermHook(boil.AfterDeleteHook, chapersOnTermAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	chapersOnTermAfterDeleteHooks = []ChapersOnTermHook{}

	AddChapersOnTermHook(boil.BeforeUpsertHook, chapersOnTermBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	chapersOnTermBeforeUpsertHooks = []ChapersOnTermHook{}

	AddChapersOnTermHook(boil.AfterUpsertHook, chapersOnTermAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	chapersOnTermAfterUpsertHooks = []ChapersOnTermHook{}
}

func testChapersOnTermsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ChapersOnTerm{}
	if err = randomize.Struct(seed, o, chapersOnTermDBTypes, true, chapersOnTermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ChapersOnTerm struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ChapersOnTerms().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testChapersOnTermsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ChapersOnTerm{}
	if err = randomize.Struct(seed, o, chapersOnTermDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ChapersOnTerm struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(chapersOnTermColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := ChapersOnTerms().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testChapersOnTermToOneChapterUsingChaper(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local ChapersOnTerm
	var foreign Chapter

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, chapersOnTermDBTypes, false, chapersOnTermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ChapersOnTerm struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, chapterDBTypes, false, chapterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Chapter struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.ChapersID = foreign.ChapterID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Chaper().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ChapterID != foreign.ChapterID {
		t.Errorf("want: %v, got %v", foreign.ChapterID, check.ChapterID)
	}

	ranAfterSelectHook := false
	AddChapterHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *Chapter) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := ChapersOnTermSlice{&local}
	if err = local.L.LoadChaper(ctx, tx, false, (*[]*ChapersOnTerm)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Chaper == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Chaper = nil
	if err = local.L.LoadChaper(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Chaper == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testChapersOnTermToOneTermUsingTerm(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local ChapersOnTerm
	var foreign Term

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, chapersOnTermDBTypes, false, chapersOnTermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ChapersOnTerm struct: %s", err)
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

	slice := ChapersOnTermSlice{&local}
	if err = local.L.LoadTerm(ctx, tx, false, (*[]*ChapersOnTerm)(&slice), nil); err != nil {
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

func testChapersOnTermToOneSetOpChapterUsingChaper(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ChapersOnTerm
	var b, c Chapter

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, chapersOnTermDBTypes, false, strmangle.SetComplement(chapersOnTermPrimaryKeyColumns, chapersOnTermColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, chapterDBTypes, false, strmangle.SetComplement(chapterPrimaryKeyColumns, chapterColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, chapterDBTypes, false, strmangle.SetComplement(chapterPrimaryKeyColumns, chapterColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Chapter{&b, &c} {
		err = a.SetChaper(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Chaper != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ChaperChapersOnTerms[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ChapersID != x.ChapterID {
			t.Error("foreign key was wrong value", a.ChapersID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ChapersID))
		reflect.Indirect(reflect.ValueOf(&a.ChapersID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.ChapersID != x.ChapterID {
			t.Error("foreign key was wrong value", a.ChapersID, x.ChapterID)
		}
	}
}
func testChapersOnTermToOneSetOpTermUsingTerm(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ChapersOnTerm
	var b, c Term

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, chapersOnTermDBTypes, false, strmangle.SetComplement(chapersOnTermPrimaryKeyColumns, chapersOnTermColumnsWithoutDefault)...); err != nil {
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

		if x.R.ChapersOnTerms[0] != &a {
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

func testChapersOnTermsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ChapersOnTerm{}
	if err = randomize.Struct(seed, o, chapersOnTermDBTypes, true, chapersOnTermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ChapersOnTerm struct: %s", err)
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

func testChapersOnTermsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ChapersOnTerm{}
	if err = randomize.Struct(seed, o, chapersOnTermDBTypes, true, chapersOnTermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ChapersOnTerm struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ChapersOnTermSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testChapersOnTermsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ChapersOnTerm{}
	if err = randomize.Struct(seed, o, chapersOnTermDBTypes, true, chapersOnTermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ChapersOnTerm struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ChapersOnTerms().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	chapersOnTermDBTypes = map[string]string{`ChapersOnTermsID`: `integer`, `ChapersID`: `integer`, `TermID`: `integer`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`}
	_                    = bytes.MinRead
)

func testChapersOnTermsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(chapersOnTermPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(chapersOnTermAllColumns) == len(chapersOnTermPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ChapersOnTerm{}
	if err = randomize.Struct(seed, o, chapersOnTermDBTypes, true, chapersOnTermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ChapersOnTerm struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ChapersOnTerms().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, chapersOnTermDBTypes, true, chapersOnTermPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ChapersOnTerm struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testChapersOnTermsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(chapersOnTermAllColumns) == len(chapersOnTermPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ChapersOnTerm{}
	if err = randomize.Struct(seed, o, chapersOnTermDBTypes, true, chapersOnTermColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ChapersOnTerm struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ChapersOnTerms().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, chapersOnTermDBTypes, true, chapersOnTermPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ChapersOnTerm struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(chapersOnTermAllColumns, chapersOnTermPrimaryKeyColumns) {
		fields = chapersOnTermAllColumns
	} else {
		fields = strmangle.SetComplement(
			chapersOnTermAllColumns,
			chapersOnTermPrimaryKeyColumns,
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

	slice := ChapersOnTermSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testChapersOnTermsUpsert(t *testing.T) {
	t.Parallel()

	if len(chapersOnTermAllColumns) == len(chapersOnTermPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := ChapersOnTerm{}
	if err = randomize.Struct(seed, &o, chapersOnTermDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ChapersOnTerm struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ChapersOnTerm: %s", err)
	}

	count, err := ChapersOnTerms().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, chapersOnTermDBTypes, false, chapersOnTermPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ChapersOnTerm struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ChapersOnTerm: %s", err)
	}

	count, err = ChapersOnTerms().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
