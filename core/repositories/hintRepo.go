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
	"go.mongodb.org/mongo-driver/mongo/options"
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

// CountHintRegistration ใช้นับจำนวน document ที่ตรงกับ filter
func (r *HintRepository) CountHintRegister(filter dto.HintfilterReq) (int64, error) {
	filterMap := utils.ConvertFilterToMap(filter)
	filterBson, err := utils.FilterHintRegister(filterMap)
	if err != nil {
		return 0, err
	}

	return r.CollectionHintRegistration.CountDocuments(context.TODO(), filterBson)
}

// FindHintRegistrationBatch ดึงข้อมูลแบบแบ่งหน้า (batch) สำหรับ export excel
func (r *HintRepository) FindHintRegistrationBatch(filter dto.HintfilterReq, batchIndex int, batchSize int64) ([]models.Register, error) {
	filterMap := utils.ConvertFilterToMap(filter)
	filterBson, err := utils.FilterHintRegister(filterMap)
	if err != nil {
		return nil, err
	}

	findOptions := options.Find()
	findOptions.Skip = ptrInt64(int64(batchIndex) * batchSize)
	findOptions.Limit = &batchSize

	cursor, err := r.CollectionHintRegistration.Find(context.TODO(), filterBson, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var results []models.Register
	if err := cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	return results, nil
}

// ptrInt64 เป็น helper สำหรับสร้าง pointer ของ int64
func ptrInt64(i int64) *int64 {
	return &i
}

func (r *HintRepository) FindAllHintRegistration(filter dto.HintfilterReq) (result []models.Register, err error) {
	filterMap := utils.ConvertFilterToMap(filter)
	filterBson, err := utils.FilterHintRegister(filterMap)
	if err != nil {
		return nil, err
	}

	// size := 100000
	pipeline := []bson.M{
		{
			"$match": filterBson,
		},
		{
			"$project": bson.M{
				"pid":                1,
				"dob":                1,
				"age":                1,
				"sex":                1,
				"fname":              1,
				"lname":              1,
				"title":              1,
				"register_date":      1,
				"code_hospital_main": 1,
				"hospital_main":      1,
				"code_hospital_sub":  1,
				"hospital_sub":       1,
				"province_main":      1,
				"change_right_date":  1,
				"change_right_memo":  1,
			},
		},
		// {
		// "$sort": bson.M{
		// 	"hcode": -1,
		// },
		// },
		// {
		// 	"$limit": size,
		// },
	}

	cursor, err := r.CollectionHintRegistration.Aggregate(context.TODO(), pipeline)

	// cursor, err := r.CollectionHintRegistration.Find(context.TODO(), filterBson)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var results []models.Register
	if err := cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	return results, nil
}

func (r *HintRepository) FindDataHIVTemplate(filter dto.HintfilterReq) (result []models.Transaction, err error) {
	filterMap := utils.ConvertFilterToMap(filter)
	// c, _ := json.MarshalIndent(filterMap, "", "  ")
	// fmt.Println(string(c))

	filterBson, err := utils.FilterHintTransaction(filterMap)
	if err != nil {
		return nil, err
	}

	codeQuery := bson.M{
		"diagnosis.code_disease": bson.M{
			"$in": []string{"B20", "B21", "B22", "B23", "B24", "Z21", "Z113", "Z114", "Z206", "Z717", "B171", "B182"},
		},
		"status": bson.M{
			"$in": []string{"settled", "approved"},
		},
	}

	// รวม query ทั้งสอง
	for k, v := range codeQuery {
		filterBson[k] = v
	}

	var size = 1

	// b, _ := json.MarshalIndent(codeQuery, "", "  ")
	// fmt.Println(string(b))
	// c, _ := json.MarshalIndent(filterBson, "", "  ")
	// fmt.Println(string(c))

	pipeline := mongo.Pipeline{
		bson.D{{Key: "$match", Value: filterBson}},
		bson.D{{Key: "$project", Value: bson.M{
			"_id":           0, // ปิด _id
			"hcode":         1,
			"hospital":      1,
			"service_type":  1,
			"hn":            1,
			"an":            bson.M{"$ifNull": bson.A{"$an", "-"}},
			"patient_pid":   1,
			"fullname":      1,
			"sex":           1,
			"patient_dob":   1,
			"create_at":     1,
			"submit_amount": 1,
			"dru":           1,
			"adp":           1,
		}}},
		bson.D{{Key: "$limit", Value: size}},
	}

	cursor, err := r.CollectionHintTransaction.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var results []models.Transaction
	if err := cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	return results, nil
}
