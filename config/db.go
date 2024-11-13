package config

import (
	"context"
	"database/sql"
	"time"

	// _ "github.com/davecgh/go-spew/spew"

	. "github.com/mamoruuji/dynamic_novel_api/db/models"
	dynamicv1 "github.com/mamoruuji/dynamic_novel_api/gen/proto/dynamic/v1"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var db boil.ContextExecutor

func SetDB(database *sql.DB) {
	db = database
	boil.SetDB(db)
}

func GetDB() boil.ContextExecutor {
	return db
}

func ConvertToPostgresTimestamp(date string) time.Time {
	layout := "2006/01/02"
	datetime, _ := time.Parse(layout, date)
	return datetime
}

func SetImageData(image *Image) *dynamicv1.ImageData {
	var pbImage *dynamicv1.ImageData
	if image != nil {
		imagePath := image.Path + image.Name
		pbImage = &dynamicv1.ImageData{
			ImageId:   image.ImageID,
			Name:      image.Name,
			ImagePath: imagePath,
			Type:      image.R.TypeOfImage.Name,
		}
	}
	return pbImage
}

func SetImagesData(images []*Image) []*dynamicv1.ImageData {
	var pbImages []*dynamicv1.ImageData
	for _, image := range images {
		pbImage := SetImageData(image)
		pbImages = append(pbImages, pbImage)
	}
	return pbImages
}

func SetTermData(terms []*Term) []*dynamicv1.TermData {
	var pbTerms []*dynamicv1.TermData
	for _, term := range terms {
		// pbTermImage := SetImageData(term.R.Image)
		pbTerm := &dynamicv1.TermData{
			TermId: term.TermID,
			Name:   term.Name,
			Text:   term.Text,
			Order:  term.Order,
			// Image:  pbTermImage,
			ImageUrl: term.ImageURL,
		}
		pbTerms = append(pbTerms, pbTerm)
	}

	return pbTerms
}

func SetTagData(dynamicId int32, ctx context.Context, db boil.ContextExecutor) []*dynamicv1.TagData {
	modifiers := []QueryMod{
		Load(DynamicsOnTagRels.Tag),
		DynamicsOnTagWhere.DynamicID.EQ(dynamicId),
	}
	DynamicsOnTags, _ := DynamicsOnTags(modifiers...).All(ctx, db)

	// if err != nil {
	// 	log.Printf("failed to get tags: %v", err)
	// 	return nil, err
	// }
	var pbTags []*dynamicv1.TagData
	for _, dynamicsOnTag := range DynamicsOnTags {

		pbTag := &dynamicv1.TagData{
			TagId: dynamicsOnTag.TagID,
			Name:  dynamicsOnTag.R.Tag.Name,
		}
		pbTags = append(pbTags, pbTag)
	}
	return pbTags
}