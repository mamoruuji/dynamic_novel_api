// Code generated by SQLBoiler 4.14.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("PrismaMigrations", testPrismaMigrations)
	t.Run("Bubbles", testBubbles)
	t.Run("ChapersOnTerms", testChapersOnTerms)
	t.Run("Chapters", testChapters)
	t.Run("Dynamics", testDynamics)
	t.Run("DynamicsOnTerms", testDynamicsOnTerms)
	t.Run("Folders", testFolders)
	t.Run("Fonts", testFonts)
	t.Run("Images", testImages)
	t.Run("Impressions", testImpressions)
	t.Run("Marks", testMarks)
	t.Run("Pages", testPages)
	t.Run("Positions", testPositions)
	t.Run("Sections", testSections)
	t.Run("Terms", testTerms)
	t.Run("Types", testTypes)
	t.Run("Users", testUsers)
}

func TestDelete(t *testing.T) {
	t.Run("PrismaMigrations", testPrismaMigrationsDelete)
	t.Run("Bubbles", testBubblesDelete)
	t.Run("ChapersOnTerms", testChapersOnTermsDelete)
	t.Run("Chapters", testChaptersDelete)
	t.Run("Dynamics", testDynamicsDelete)
	t.Run("DynamicsOnTerms", testDynamicsOnTermsDelete)
	t.Run("Folders", testFoldersDelete)
	t.Run("Fonts", testFontsDelete)
	t.Run("Images", testImagesDelete)
	t.Run("Impressions", testImpressionsDelete)
	t.Run("Marks", testMarksDelete)
	t.Run("Pages", testPagesDelete)
	t.Run("Positions", testPositionsDelete)
	t.Run("Sections", testSectionsDelete)
	t.Run("Terms", testTermsDelete)
	t.Run("Types", testTypesDelete)
	t.Run("Users", testUsersDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("PrismaMigrations", testPrismaMigrationsQueryDeleteAll)
	t.Run("Bubbles", testBubblesQueryDeleteAll)
	t.Run("ChapersOnTerms", testChapersOnTermsQueryDeleteAll)
	t.Run("Chapters", testChaptersQueryDeleteAll)
	t.Run("Dynamics", testDynamicsQueryDeleteAll)
	t.Run("DynamicsOnTerms", testDynamicsOnTermsQueryDeleteAll)
	t.Run("Folders", testFoldersQueryDeleteAll)
	t.Run("Fonts", testFontsQueryDeleteAll)
	t.Run("Images", testImagesQueryDeleteAll)
	t.Run("Impressions", testImpressionsQueryDeleteAll)
	t.Run("Marks", testMarksQueryDeleteAll)
	t.Run("Pages", testPagesQueryDeleteAll)
	t.Run("Positions", testPositionsQueryDeleteAll)
	t.Run("Sections", testSectionsQueryDeleteAll)
	t.Run("Terms", testTermsQueryDeleteAll)
	t.Run("Types", testTypesQueryDeleteAll)
	t.Run("Users", testUsersQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("PrismaMigrations", testPrismaMigrationsSliceDeleteAll)
	t.Run("Bubbles", testBubblesSliceDeleteAll)
	t.Run("ChapersOnTerms", testChapersOnTermsSliceDeleteAll)
	t.Run("Chapters", testChaptersSliceDeleteAll)
	t.Run("Dynamics", testDynamicsSliceDeleteAll)
	t.Run("DynamicsOnTerms", testDynamicsOnTermsSliceDeleteAll)
	t.Run("Folders", testFoldersSliceDeleteAll)
	t.Run("Fonts", testFontsSliceDeleteAll)
	t.Run("Images", testImagesSliceDeleteAll)
	t.Run("Impressions", testImpressionsSliceDeleteAll)
	t.Run("Marks", testMarksSliceDeleteAll)
	t.Run("Pages", testPagesSliceDeleteAll)
	t.Run("Positions", testPositionsSliceDeleteAll)
	t.Run("Sections", testSectionsSliceDeleteAll)
	t.Run("Terms", testTermsSliceDeleteAll)
	t.Run("Types", testTypesSliceDeleteAll)
	t.Run("Users", testUsersSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("PrismaMigrations", testPrismaMigrationsExists)
	t.Run("Bubbles", testBubblesExists)
	t.Run("ChapersOnTerms", testChapersOnTermsExists)
	t.Run("Chapters", testChaptersExists)
	t.Run("Dynamics", testDynamicsExists)
	t.Run("DynamicsOnTerms", testDynamicsOnTermsExists)
	t.Run("Folders", testFoldersExists)
	t.Run("Fonts", testFontsExists)
	t.Run("Images", testImagesExists)
	t.Run("Impressions", testImpressionsExists)
	t.Run("Marks", testMarksExists)
	t.Run("Pages", testPagesExists)
	t.Run("Positions", testPositionsExists)
	t.Run("Sections", testSectionsExists)
	t.Run("Terms", testTermsExists)
	t.Run("Types", testTypesExists)
	t.Run("Users", testUsersExists)
}

func TestFind(t *testing.T) {
	t.Run("PrismaMigrations", testPrismaMigrationsFind)
	t.Run("Bubbles", testBubblesFind)
	t.Run("ChapersOnTerms", testChapersOnTermsFind)
	t.Run("Chapters", testChaptersFind)
	t.Run("Dynamics", testDynamicsFind)
	t.Run("DynamicsOnTerms", testDynamicsOnTermsFind)
	t.Run("Folders", testFoldersFind)
	t.Run("Fonts", testFontsFind)
	t.Run("Images", testImagesFind)
	t.Run("Impressions", testImpressionsFind)
	t.Run("Marks", testMarksFind)
	t.Run("Pages", testPagesFind)
	t.Run("Positions", testPositionsFind)
	t.Run("Sections", testSectionsFind)
	t.Run("Terms", testTermsFind)
	t.Run("Types", testTypesFind)
	t.Run("Users", testUsersFind)
}

func TestBind(t *testing.T) {
	t.Run("PrismaMigrations", testPrismaMigrationsBind)
	t.Run("Bubbles", testBubblesBind)
	t.Run("ChapersOnTerms", testChapersOnTermsBind)
	t.Run("Chapters", testChaptersBind)
	t.Run("Dynamics", testDynamicsBind)
	t.Run("DynamicsOnTerms", testDynamicsOnTermsBind)
	t.Run("Folders", testFoldersBind)
	t.Run("Fonts", testFontsBind)
	t.Run("Images", testImagesBind)
	t.Run("Impressions", testImpressionsBind)
	t.Run("Marks", testMarksBind)
	t.Run("Pages", testPagesBind)
	t.Run("Positions", testPositionsBind)
	t.Run("Sections", testSectionsBind)
	t.Run("Terms", testTermsBind)
	t.Run("Types", testTypesBind)
	t.Run("Users", testUsersBind)
}

func TestOne(t *testing.T) {
	t.Run("PrismaMigrations", testPrismaMigrationsOne)
	t.Run("Bubbles", testBubblesOne)
	t.Run("ChapersOnTerms", testChapersOnTermsOne)
	t.Run("Chapters", testChaptersOne)
	t.Run("Dynamics", testDynamicsOne)
	t.Run("DynamicsOnTerms", testDynamicsOnTermsOne)
	t.Run("Folders", testFoldersOne)
	t.Run("Fonts", testFontsOne)
	t.Run("Images", testImagesOne)
	t.Run("Impressions", testImpressionsOne)
	t.Run("Marks", testMarksOne)
	t.Run("Pages", testPagesOne)
	t.Run("Positions", testPositionsOne)
	t.Run("Sections", testSectionsOne)
	t.Run("Terms", testTermsOne)
	t.Run("Types", testTypesOne)
	t.Run("Users", testUsersOne)
}

func TestAll(t *testing.T) {
	t.Run("PrismaMigrations", testPrismaMigrationsAll)
	t.Run("Bubbles", testBubblesAll)
	t.Run("ChapersOnTerms", testChapersOnTermsAll)
	t.Run("Chapters", testChaptersAll)
	t.Run("Dynamics", testDynamicsAll)
	t.Run("DynamicsOnTerms", testDynamicsOnTermsAll)
	t.Run("Folders", testFoldersAll)
	t.Run("Fonts", testFontsAll)
	t.Run("Images", testImagesAll)
	t.Run("Impressions", testImpressionsAll)
	t.Run("Marks", testMarksAll)
	t.Run("Pages", testPagesAll)
	t.Run("Positions", testPositionsAll)
	t.Run("Sections", testSectionsAll)
	t.Run("Terms", testTermsAll)
	t.Run("Types", testTypesAll)
	t.Run("Users", testUsersAll)
}

func TestCount(t *testing.T) {
	t.Run("PrismaMigrations", testPrismaMigrationsCount)
	t.Run("Bubbles", testBubblesCount)
	t.Run("ChapersOnTerms", testChapersOnTermsCount)
	t.Run("Chapters", testChaptersCount)
	t.Run("Dynamics", testDynamicsCount)
	t.Run("DynamicsOnTerms", testDynamicsOnTermsCount)
	t.Run("Folders", testFoldersCount)
	t.Run("Fonts", testFontsCount)
	t.Run("Images", testImagesCount)
	t.Run("Impressions", testImpressionsCount)
	t.Run("Marks", testMarksCount)
	t.Run("Pages", testPagesCount)
	t.Run("Positions", testPositionsCount)
	t.Run("Sections", testSectionsCount)
	t.Run("Terms", testTermsCount)
	t.Run("Types", testTypesCount)
	t.Run("Users", testUsersCount)
}

func TestHooks(t *testing.T) {
	t.Run("PrismaMigrations", testPrismaMigrationsHooks)
	t.Run("Bubbles", testBubblesHooks)
	t.Run("ChapersOnTerms", testChapersOnTermsHooks)
	t.Run("Chapters", testChaptersHooks)
	t.Run("Dynamics", testDynamicsHooks)
	t.Run("DynamicsOnTerms", testDynamicsOnTermsHooks)
	t.Run("Folders", testFoldersHooks)
	t.Run("Fonts", testFontsHooks)
	t.Run("Images", testImagesHooks)
	t.Run("Impressions", testImpressionsHooks)
	t.Run("Marks", testMarksHooks)
	t.Run("Pages", testPagesHooks)
	t.Run("Positions", testPositionsHooks)
	t.Run("Sections", testSectionsHooks)
	t.Run("Terms", testTermsHooks)
	t.Run("Types", testTypesHooks)
	t.Run("Users", testUsersHooks)
}

func TestInsert(t *testing.T) {
	t.Run("PrismaMigrations", testPrismaMigrationsInsert)
	t.Run("PrismaMigrations", testPrismaMigrationsInsertWhitelist)
	t.Run("Bubbles", testBubblesInsert)
	t.Run("Bubbles", testBubblesInsertWhitelist)
	t.Run("ChapersOnTerms", testChapersOnTermsInsert)
	t.Run("ChapersOnTerms", testChapersOnTermsInsertWhitelist)
	t.Run("Chapters", testChaptersInsert)
	t.Run("Chapters", testChaptersInsertWhitelist)
	t.Run("Dynamics", testDynamicsInsert)
	t.Run("Dynamics", testDynamicsInsertWhitelist)
	t.Run("DynamicsOnTerms", testDynamicsOnTermsInsert)
	t.Run("DynamicsOnTerms", testDynamicsOnTermsInsertWhitelist)
	t.Run("Folders", testFoldersInsert)
	t.Run("Folders", testFoldersInsertWhitelist)
	t.Run("Fonts", testFontsInsert)
	t.Run("Fonts", testFontsInsertWhitelist)
	t.Run("Images", testImagesInsert)
	t.Run("Images", testImagesInsertWhitelist)
	t.Run("Impressions", testImpressionsInsert)
	t.Run("Impressions", testImpressionsInsertWhitelist)
	t.Run("Marks", testMarksInsert)
	t.Run("Marks", testMarksInsertWhitelist)
	t.Run("Pages", testPagesInsert)
	t.Run("Pages", testPagesInsertWhitelist)
	t.Run("Positions", testPositionsInsert)
	t.Run("Positions", testPositionsInsertWhitelist)
	t.Run("Sections", testSectionsInsert)
	t.Run("Sections", testSectionsInsertWhitelist)
	t.Run("Terms", testTermsInsert)
	t.Run("Terms", testTermsInsertWhitelist)
	t.Run("Types", testTypesInsert)
	t.Run("Types", testTypesInsertWhitelist)
	t.Run("Users", testUsersInsert)
	t.Run("Users", testUsersInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("ChapersOnTermToChapterUsingChaper", testChapersOnTermToOneChapterUsingChaper)
	t.Run("ChapersOnTermToTermUsingTerm", testChapersOnTermToOneTermUsingTerm)
	t.Run("ChapterToDynamicUsingDynamic", testChapterToOneDynamicUsingDynamic)
	t.Run("DynamicToUserUsingUser", testDynamicToOneUserUsingUser)
	t.Run("DynamicsOnTermToDynamicUsingDynamic", testDynamicsOnTermToOneDynamicUsingDynamic)
	t.Run("DynamicsOnTermToTermUsingTerm", testDynamicsOnTermToOneTermUsingTerm)
	t.Run("FolderToFolderUsingParentIdFolder", testFolderToOneFolderUsingParentIdFolder)
	t.Run("FolderToUserUsingUser", testFolderToOneUserUsingUser)
	t.Run("ImageToFolderUsingFolder", testImageToOneFolderUsingFolder)
	t.Run("ImageToUserUsingUser", testImageToOneUserUsingUser)
	t.Run("ImpressionToDynamicUsingDynamic", testImpressionToOneDynamicUsingDynamic)
	t.Run("ImpressionToUserUsingUser", testImpressionToOneUserUsingUser)
	t.Run("MarkToDynamicUsingDynamic", testMarkToOneDynamicUsingDynamic)
	t.Run("MarkToUserUsingUser", testMarkToOneUserUsingUser)
	t.Run("PageToChapterUsingChapter", testPageToOneChapterUsingChapter)
	t.Run("SectionToBubbleUsingBubble", testSectionToOneBubbleUsingBubble)
	t.Run("SectionToFontUsingFont", testSectionToOneFontUsingFont)
	t.Run("SectionToImageUsingImage", testSectionToOneImageUsingImage)
	t.Run("SectionToPageUsingPage", testSectionToOnePageUsingPage)
	t.Run("SectionToPositionUsingPosition", testSectionToOnePositionUsingPosition)
	t.Run("SectionToTermUsingTerm", testSectionToOneTermUsingTerm)
	t.Run("SectionToTypeUsingType", testSectionToOneTypeUsingType)
	t.Run("TermToImageUsingImage", testTermToOneImageUsingImage)
	t.Run("TermToUserUsingUser", testTermToOneUserUsingUser)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("BubbleToSections", testBubbleToManySections)
	t.Run("ChapterToChaperChapersOnTerms", testChapterToManyChaperChapersOnTerms)
	t.Run("ChapterToPages", testChapterToManyPages)
	t.Run("DynamicToChapters", testDynamicToManyChapters)
	t.Run("DynamicToDynamicsOnTerms", testDynamicToManyDynamicsOnTerms)
	t.Run("DynamicToImpressions", testDynamicToManyImpressions)
	t.Run("DynamicToMarks", testDynamicToManyMarks)
	t.Run("FolderToParentIdFolders", testFolderToManyParentIdFolders)
	t.Run("FolderToImages", testFolderToManyImages)
	t.Run("FontToSections", testFontToManySections)
	t.Run("ImageToSections", testImageToManySections)
	t.Run("ImageToTerms", testImageToManyTerms)
	t.Run("PageToSections", testPageToManySections)
	t.Run("PositionToSections", testPositionToManySections)
	t.Run("TermToChapersOnTerms", testTermToManyChapersOnTerms)
	t.Run("TermToDynamicsOnTerms", testTermToManyDynamicsOnTerms)
	t.Run("TermToSections", testTermToManySections)
	t.Run("TypeToSections", testTypeToManySections)
	t.Run("UserToDynamics", testUserToManyDynamics)
	t.Run("UserToFolders", testUserToManyFolders)
	t.Run("UserToImages", testUserToManyImages)
	t.Run("UserToImpressions", testUserToManyImpressions)
	t.Run("UserToMarks", testUserToManyMarks)
	t.Run("UserToTerms", testUserToManyTerms)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("ChapersOnTermToChapterUsingChaperChapersOnTerms", testChapersOnTermToOneSetOpChapterUsingChaper)
	t.Run("ChapersOnTermToTermUsingChapersOnTerms", testChapersOnTermToOneSetOpTermUsingTerm)
	t.Run("ChapterToDynamicUsingChapters", testChapterToOneSetOpDynamicUsingDynamic)
	t.Run("DynamicToUserUsingDynamics", testDynamicToOneSetOpUserUsingUser)
	t.Run("DynamicsOnTermToDynamicUsingDynamicsOnTerms", testDynamicsOnTermToOneSetOpDynamicUsingDynamic)
	t.Run("DynamicsOnTermToTermUsingDynamicsOnTerms", testDynamicsOnTermToOneSetOpTermUsingTerm)
	t.Run("FolderToFolderUsingParentIdFolders", testFolderToOneSetOpFolderUsingParentIdFolder)
	t.Run("FolderToUserUsingFolders", testFolderToOneSetOpUserUsingUser)
	t.Run("ImageToFolderUsingImages", testImageToOneSetOpFolderUsingFolder)
	t.Run("ImageToUserUsingImages", testImageToOneSetOpUserUsingUser)
	t.Run("ImpressionToDynamicUsingImpressions", testImpressionToOneSetOpDynamicUsingDynamic)
	t.Run("ImpressionToUserUsingImpressions", testImpressionToOneSetOpUserUsingUser)
	t.Run("MarkToDynamicUsingMarks", testMarkToOneSetOpDynamicUsingDynamic)
	t.Run("MarkToUserUsingMarks", testMarkToOneSetOpUserUsingUser)
	t.Run("PageToChapterUsingPages", testPageToOneSetOpChapterUsingChapter)
	t.Run("SectionToBubbleUsingSections", testSectionToOneSetOpBubbleUsingBubble)
	t.Run("SectionToFontUsingSections", testSectionToOneSetOpFontUsingFont)
	t.Run("SectionToImageUsingSections", testSectionToOneSetOpImageUsingImage)
	t.Run("SectionToPageUsingSections", testSectionToOneSetOpPageUsingPage)
	t.Run("SectionToPositionUsingSections", testSectionToOneSetOpPositionUsingPosition)
	t.Run("SectionToTermUsingSections", testSectionToOneSetOpTermUsingTerm)
	t.Run("SectionToTypeUsingSections", testSectionToOneSetOpTypeUsingType)
	t.Run("TermToImageUsingTerms", testTermToOneSetOpImageUsingImage)
	t.Run("TermToUserUsingTerms", testTermToOneSetOpUserUsingUser)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {
	t.Run("FolderToFolderUsingParentIdFolders", testFolderToOneRemoveOpFolderUsingParentIdFolder)
	t.Run("ImageToFolderUsingImages", testImageToOneRemoveOpFolderUsingFolder)
	t.Run("SectionToBubbleUsingSections", testSectionToOneRemoveOpBubbleUsingBubble)
	t.Run("SectionToFontUsingSections", testSectionToOneRemoveOpFontUsingFont)
	t.Run("SectionToImageUsingSections", testSectionToOneRemoveOpImageUsingImage)
	t.Run("SectionToPositionUsingSections", testSectionToOneRemoveOpPositionUsingPosition)
	t.Run("SectionToTermUsingSections", testSectionToOneRemoveOpTermUsingTerm)
	t.Run("SectionToTypeUsingSections", testSectionToOneRemoveOpTypeUsingType)
	t.Run("TermToImageUsingTerms", testTermToOneRemoveOpImageUsingImage)
}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("BubbleToSections", testBubbleToManyAddOpSections)
	t.Run("ChapterToChaperChapersOnTerms", testChapterToManyAddOpChaperChapersOnTerms)
	t.Run("ChapterToPages", testChapterToManyAddOpPages)
	t.Run("DynamicToChapters", testDynamicToManyAddOpChapters)
	t.Run("DynamicToDynamicsOnTerms", testDynamicToManyAddOpDynamicsOnTerms)
	t.Run("DynamicToImpressions", testDynamicToManyAddOpImpressions)
	t.Run("DynamicToMarks", testDynamicToManyAddOpMarks)
	t.Run("FolderToParentIdFolders", testFolderToManyAddOpParentIdFolders)
	t.Run("FolderToImages", testFolderToManyAddOpImages)
	t.Run("FontToSections", testFontToManyAddOpSections)
	t.Run("ImageToSections", testImageToManyAddOpSections)
	t.Run("ImageToTerms", testImageToManyAddOpTerms)
	t.Run("PageToSections", testPageToManyAddOpSections)
	t.Run("PositionToSections", testPositionToManyAddOpSections)
	t.Run("TermToChapersOnTerms", testTermToManyAddOpChapersOnTerms)
	t.Run("TermToDynamicsOnTerms", testTermToManyAddOpDynamicsOnTerms)
	t.Run("TermToSections", testTermToManyAddOpSections)
	t.Run("TypeToSections", testTypeToManyAddOpSections)
	t.Run("UserToDynamics", testUserToManyAddOpDynamics)
	t.Run("UserToFolders", testUserToManyAddOpFolders)
	t.Run("UserToImages", testUserToManyAddOpImages)
	t.Run("UserToImpressions", testUserToManyAddOpImpressions)
	t.Run("UserToMarks", testUserToManyAddOpMarks)
	t.Run("UserToTerms", testUserToManyAddOpTerms)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {
	t.Run("BubbleToSections", testBubbleToManySetOpSections)
	t.Run("FolderToParentIdFolders", testFolderToManySetOpParentIdFolders)
	t.Run("FolderToImages", testFolderToManySetOpImages)
	t.Run("FontToSections", testFontToManySetOpSections)
	t.Run("ImageToSections", testImageToManySetOpSections)
	t.Run("ImageToTerms", testImageToManySetOpTerms)
	t.Run("PositionToSections", testPositionToManySetOpSections)
	t.Run("TermToSections", testTermToManySetOpSections)
	t.Run("TypeToSections", testTypeToManySetOpSections)
}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {
	t.Run("BubbleToSections", testBubbleToManyRemoveOpSections)
	t.Run("FolderToParentIdFolders", testFolderToManyRemoveOpParentIdFolders)
	t.Run("FolderToImages", testFolderToManyRemoveOpImages)
	t.Run("FontToSections", testFontToManyRemoveOpSections)
	t.Run("ImageToSections", testImageToManyRemoveOpSections)
	t.Run("ImageToTerms", testImageToManyRemoveOpTerms)
	t.Run("PositionToSections", testPositionToManyRemoveOpSections)
	t.Run("TermToSections", testTermToManyRemoveOpSections)
	t.Run("TypeToSections", testTypeToManyRemoveOpSections)
}

func TestReload(t *testing.T) {
	t.Run("PrismaMigrations", testPrismaMigrationsReload)
	t.Run("Bubbles", testBubblesReload)
	t.Run("ChapersOnTerms", testChapersOnTermsReload)
	t.Run("Chapters", testChaptersReload)
	t.Run("Dynamics", testDynamicsReload)
	t.Run("DynamicsOnTerms", testDynamicsOnTermsReload)
	t.Run("Folders", testFoldersReload)
	t.Run("Fonts", testFontsReload)
	t.Run("Images", testImagesReload)
	t.Run("Impressions", testImpressionsReload)
	t.Run("Marks", testMarksReload)
	t.Run("Pages", testPagesReload)
	t.Run("Positions", testPositionsReload)
	t.Run("Sections", testSectionsReload)
	t.Run("Terms", testTermsReload)
	t.Run("Types", testTypesReload)
	t.Run("Users", testUsersReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("PrismaMigrations", testPrismaMigrationsReloadAll)
	t.Run("Bubbles", testBubblesReloadAll)
	t.Run("ChapersOnTerms", testChapersOnTermsReloadAll)
	t.Run("Chapters", testChaptersReloadAll)
	t.Run("Dynamics", testDynamicsReloadAll)
	t.Run("DynamicsOnTerms", testDynamicsOnTermsReloadAll)
	t.Run("Folders", testFoldersReloadAll)
	t.Run("Fonts", testFontsReloadAll)
	t.Run("Images", testImagesReloadAll)
	t.Run("Impressions", testImpressionsReloadAll)
	t.Run("Marks", testMarksReloadAll)
	t.Run("Pages", testPagesReloadAll)
	t.Run("Positions", testPositionsReloadAll)
	t.Run("Sections", testSectionsReloadAll)
	t.Run("Terms", testTermsReloadAll)
	t.Run("Types", testTypesReloadAll)
	t.Run("Users", testUsersReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("PrismaMigrations", testPrismaMigrationsSelect)
	t.Run("Bubbles", testBubblesSelect)
	t.Run("ChapersOnTerms", testChapersOnTermsSelect)
	t.Run("Chapters", testChaptersSelect)
	t.Run("Dynamics", testDynamicsSelect)
	t.Run("DynamicsOnTerms", testDynamicsOnTermsSelect)
	t.Run("Folders", testFoldersSelect)
	t.Run("Fonts", testFontsSelect)
	t.Run("Images", testImagesSelect)
	t.Run("Impressions", testImpressionsSelect)
	t.Run("Marks", testMarksSelect)
	t.Run("Pages", testPagesSelect)
	t.Run("Positions", testPositionsSelect)
	t.Run("Sections", testSectionsSelect)
	t.Run("Terms", testTermsSelect)
	t.Run("Types", testTypesSelect)
	t.Run("Users", testUsersSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("PrismaMigrations", testPrismaMigrationsUpdate)
	t.Run("Bubbles", testBubblesUpdate)
	t.Run("ChapersOnTerms", testChapersOnTermsUpdate)
	t.Run("Chapters", testChaptersUpdate)
	t.Run("Dynamics", testDynamicsUpdate)
	t.Run("DynamicsOnTerms", testDynamicsOnTermsUpdate)
	t.Run("Folders", testFoldersUpdate)
	t.Run("Fonts", testFontsUpdate)
	t.Run("Images", testImagesUpdate)
	t.Run("Impressions", testImpressionsUpdate)
	t.Run("Marks", testMarksUpdate)
	t.Run("Pages", testPagesUpdate)
	t.Run("Positions", testPositionsUpdate)
	t.Run("Sections", testSectionsUpdate)
	t.Run("Terms", testTermsUpdate)
	t.Run("Types", testTypesUpdate)
	t.Run("Users", testUsersUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("PrismaMigrations", testPrismaMigrationsSliceUpdateAll)
	t.Run("Bubbles", testBubblesSliceUpdateAll)
	t.Run("ChapersOnTerms", testChapersOnTermsSliceUpdateAll)
	t.Run("Chapters", testChaptersSliceUpdateAll)
	t.Run("Dynamics", testDynamicsSliceUpdateAll)
	t.Run("DynamicsOnTerms", testDynamicsOnTermsSliceUpdateAll)
	t.Run("Folders", testFoldersSliceUpdateAll)
	t.Run("Fonts", testFontsSliceUpdateAll)
	t.Run("Images", testImagesSliceUpdateAll)
	t.Run("Impressions", testImpressionsSliceUpdateAll)
	t.Run("Marks", testMarksSliceUpdateAll)
	t.Run("Pages", testPagesSliceUpdateAll)
	t.Run("Positions", testPositionsSliceUpdateAll)
	t.Run("Sections", testSectionsSliceUpdateAll)
	t.Run("Terms", testTermsSliceUpdateAll)
	t.Run("Types", testTypesSliceUpdateAll)
	t.Run("Users", testUsersSliceUpdateAll)
}
