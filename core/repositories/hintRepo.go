package repositories

import (
	"context"

	"go-fiber-app/configs"
	"go-fiber-app/core/dto"
	"go-fiber-app/core/entities"
	"go-fiber-app/core/models"
	"go-fiber-app/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type HintRepository struct {
	// cache                   		pkg.Cache
	db                         *mongo.Database
	CollectionHintTransaction  *mongo.Collection
	CollectionHintRegistration *mongo.Collection
	// collectionHintBudget       *mongo.Collection
}

func NewHintRepository(db *mongo.Database) entities.HintRepository {

	return &HintRepository{
		db:                         db,
		CollectionHintTransaction:  db.Collection(configs.CollectionHintTransaction),
		CollectionHintRegistration: db.Collection(configs.CollectionHintRegistration),
		// collectionHintBudget:       db.Collection(configs.collectionHintBudget),
	}
}

func (r *HintRepository) QueryHintTransaction(filter dto.HintfilterReq) (results []models.Transaction, err error) {
	filterMap := utils.ConvertFilterToMap(filter)
	filterBson, err := utils.FilterHintTransaction(filterMap)

	// fmt.Println("filterBson:", filterBson)

	// fmt.Printf("DEBUG: filter[\"area\"] = %v (type: %T)\n", filterBson["area"], filterBson["area"])

	if err != nil {
		return nil, err
	}

	pipeline := []bson.M{
		{
			"$match": filterBson,
		},
		{
			"$project": bson.M{
				"hcode":    1,
				"hospital": 1,
				"_id":      0,
			},
		},
		// {
		// 	"$limit": 10,
		// },
		// {
		// 	"$sort": bson.M{
		// 		"hcode": -1,
		// 	},
		// },
	}

	// fmt.Println("Pipeline:", pipeline)

	cursor, err := r.CollectionHintTransaction.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return nil, err
	}

	if err := cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}
	// fmt.Println(results)
	return results, nil
}

func (r *HintRepository) QueryHintRegistration(filter dto.HintfilterReq) (results []models.Register, err error) {
	// สร้าง filter จาก utils
	filterMap := utils.ConvertFilterToMap(filter)
	filterBson, err := utils.FilterHintRegister(filterMap)
	if err != nil {
		return nil, err
	}
	size := 1000

	// Pipeline รวม filter เข้าไป
	pipeline := []bson.M{
		{
			"$match": filterBson, // ใช้ filter ที่ได้จาก utils
		},
		{
			"$project": bson.M{
				"hcode":    1,
				"hospital": 1,
				"_id":      0,
			},
		},
		// {
		// "$sort": bson.M{
		// 	"hcode": -1,
		// },
		// },
		{
			"$limit": size,
		},
	}

	cursor, err := r.CollectionHintTransaction.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	if err := cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}
	return results, nil
}
