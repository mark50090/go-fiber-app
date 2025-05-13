package repositories

import (
	"context"
	"crypto/sha256"
	"go-fiber-app/configs"
	"go-fiber-app/core/dto"
	"go-fiber-app/core/entities"
	"go-fiber-app/core/models"
	"go-fiber-app/utils"
	"sync"

	"github.com/bytedance/sonic"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type unitCostRepo struct {
	// cache                               pkg.Cache
	db                                  *mongo.Database
	collectionFDHTransaction            *mongo.Collection
	collectionSummaryUnitCost           *mongo.Collection
	collectionSummaryUnitCostHospital   *mongo.Collection
	collectionSummaryUnitCostInsclDiary *mongo.Collection
	collectionICD                       *mongo.Collection
	CollectionFDHInscl                  *mongo.Collection
	CollectionItem                      *mongo.Collection
}

func NewUnitCostRepository(db *mongo.Database) entities.UnitCostRepository {

	return &unitCostRepo{
		// cache:                               cache,
		db:                                  db,
		collectionFDHTransaction:            db.Collection(configs.CollectionFDHTransaction),
		collectionSummaryUnitCost:           db.Collection(configs.CollectionSummaryUnitCostDiary),
		collectionSummaryUnitCostHospital:   db.Collection(configs.CollectionSummaryUnitCostHospitalDiary),
		collectionSummaryUnitCostInsclDiary: db.Collection(configs.CollectionSummaryUnitCostInsclDiary),
		collectionICD:                       db.Collection(configs.CollectionICD),
		CollectionFDHInscl:                  db.Collection(configs.CollectionFDHInscl),
		CollectionItem:                      db.Collection(configs.CollectionItem),
	}
}

func (r *unitCostRepo) QueryUnitCostSummary(filter dto.QueryParser, QueryPage dto.QueryPage, querySearch dto.QuerySearch) (results []models.UnitCostSummary, total int64, err error) {

	filterBson, err := utils.ConvertStructToBsonFilter(filter)
	if err != nil {
		return nil, 0, err
	}
	if querySearch.Search != nil && *querySearch.Search != "" {
		filterBson = append(filterBson, bson.E{Key: "icd", Value: bson.M{"$regex": "^" + *querySearch.Search, "$options": "i"}})
	}

	skip := (QueryPage.Page - 1) * QueryPage.Limit
	limit := QueryPage.Limit

	wg := sync.WaitGroup{}

	collection := r.CollectionSelecter(filter)

	group := bson.M{
		"_id": "$icd",
	}
	wg.Add(1)
	go func(t *int64) {
		defer wg.Done()
		pipeline := []bson.M{
			{"$match": filterBson},
			{"$group": group},
			{"$count": "total"},
		}

		total := struct {
			Total int64 `bson:"total"`
		}{Total: 0}

		cursor, err := r.db.Collection(collection).Aggregate(context.TODO(), pipeline)
		if err != nil {
			return
		}

		for cursor.Next(context.TODO()) {
			if err := cursor.Decode(&total); err != nil {
				return
			}
			*t = total.Total
			return
		}
	}(&total)

	group["icd"] = bson.M{"$first": "$icd"}
	group["total"] = bson.M{"$sum": "$total"}
	group["len"] = bson.M{"$sum": "$len"}

	pipeline := []bson.M{
		{"$match": filterBson},
		{"$group": group},
		{"$project": bson.M{
			"icd": 1,
			"avg": bson.M{
				"$divide": bson.A{"$total", "$len"},
			},
		}},
		{"$sort": bson.M{"avg": -1}},
		{"$skip": skip},
		{"$limit": limit},
		{"$lookup": bson.M{
			"from":         r.collectionICD.Name(),
			"localField":   "_id",
			"foreignField": "code",
			"as":           "icd_info",
		}},
		{"$project": bson.M{
			"icd": 1,
			"avg": 1,
			"name_en": bson.M{
				"$first": "$icd_info.description",
			},
			"name_th": bson.M{
				"$first": "$icd_info.description_th",
			},
		}},
	}

	err = r.aggregateWithCache(collection, pipeline, &results)
	if err != nil {
		return nil, 0, err
	}

	wg.Wait()
	return

}

func (r *unitCostRepo) aggregateWithCache(collection string, pipeline interface{}, result interface{}) (err error) {

	q, _ := sonic.Marshal(pipeline)
	hashed := sha256.New()
	hashed.Write(q)
	// cacheKey := fmt.Sprintf("cache:query:%s:%x:aggregateWithCache", collection, hashed.Sum(nil))

	// err = r.cache.GetWithUnmarshal(cacheKey, result)
	// if err == nil {
	// 	return
	// }

	cursor, err := r.db.Collection(collection).Aggregate(context.TODO(), pipeline)
	if err != nil {
		return err
	}

	if err := cursor.All(context.TODO(), result); err != nil {
		return err
	}

	// err = r.cache.SetWithMarshal(cacheKey, result)
	// if err != nil {
	// 	return err
	// }

	return nil
}

func (r *unitCostRepo) CollectionSelecter(filter dto.QueryParser) string {

	collection := r.collectionSummaryUnitCost.Name()

	if filter.Inscl != nil && *filter.Inscl != "" {
		collection = r.collectionSummaryUnitCostInsclDiary.Name()
	}

	return collection
}

func (r *unitCostRepo) QueryInsclDetail() (results []dto.InsclDetailResponse, err error) {

	pipeline := []bson.M{
		{
			"$group": bson.M{
				"_id":    "$code",
				"detail": bson.M{"$first": "$detail"},
			},
		},
		{
			"$sort": bson.M{"_id": 1},
		},
	}

	cursor, err := r.CollectionFDHInscl.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return nil, err
	}

	if err := cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}
	return results, nil

}

func (r *unitCostRepo) QueryItemData() (results []dto.ItemRespone, err error) {

	pipeline := []bson.M{
		{
			"$project": bson.M{
				"name":  1,
				"price": 1,
				"_id":   0,
			},
		},
		{
			"$sort": bson.M{
				"price": -1,
			},
		},
	}

	cursor, err := r.CollectionItem.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return nil, err
	}

	if err := cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}
	return results, nil

}
