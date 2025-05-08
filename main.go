package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/xuri/excelize/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Global MongoDB client
var mongoClient *mongo.Client

type Item struct {
	Name  string  `bson:"name" json:"name"`
	Price float64 `bson:"price" json:"price"`
}

// Register struct
type Register struct {
	Batch            string    `bson:"batch" json:"batch"`
	AreaMain         int       `bson:"area_main" json:"area_main"`
	AreaSub          int       `bson:"area_sub" json:"area_sub"`
	CodeHospitalMain string    `bson:"code_hospital_main" json:"code_hospital_main"`
	CodeHospitalSub  string    `bson:"code_hospital_sub" json:"code_hospital_sub"`
	HospitalMain     string    `bson:"hospital_main" json:"hospital_main"`
	HospitalSub      string    `bson:"hospital_sub" json:"hospital_sub"`
	ProvinceMain     string    `bson:"province_main" json:"province_main"`
	ProvinceSub      string    `bson:"province_sub" json:"province_sub"`
	Pid              string    `bson:"pid" json:"pid"`
	Dob              time.Time `bson:"dob" json:"dob"`
	Sex              string    `bson:"sex" json:"sex"`
	Title            string    `bson:"title" json:"title"`
	Fname            string    `bson:"fname" json:"fname"`
	Lname            string    `bson:"lname" json:"lname"`
	Fullname         string    `bson:"fullname" json:"fullname"`
	RegisterDate     time.Time `bson:"register_date" json:"register_date"`
	Status           string    `bson:"status" json:"status"`
	TypeHospitalMain string    `bson:"type_hospital_main" json:"type_hospital_main"`
	ChangeRightDate  time.Time `bson:"change_right_date" json:"change_right_date"`
	ChangeRightMemo  string    `bson:"change_right_memo" json:"change_right_memo"`
	CreatedAt        time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt        time.Time `bson:"updated_at" json:"updated_at"`
}

func initMongo() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI is not set in .env")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal("MongoDB connection error:", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("MongoDB ping error:", err)
	}

	mongoClient = client
	log.Println("✅ Connected to MongoDB!")
}

func main() {
	initMongo()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	app := fiber.New()

	// 🟢 API: Get all registers
	app.Get("/registers", func(c *fiber.Ctx) error {
		collection := mongoClient.Database("myappdb").Collection("registers")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// ✅ เพิ่ม limit 10000
		findOptions := options.Find()
		findOptions.SetLimit(50000)

		cursor, err := collection.Find(ctx, bson.D{}, findOptions)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": fmt.Sprintf("Failed to fetch registers: %v", err),
			})
		}
		defer cursor.Close(ctx)

		var registers []Register
		if err := cursor.All(ctx, &registers); err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": fmt.Sprintf("Failed to decode registers: %v", err),
			})
		}

		return c.JSON(registers)
	})

	app.Post("/items", func(c *fiber.Ctx) error {
		var item Item
		// Parse JSON body -> struct
		if err := c.BodyParser(&item); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		// ตรวจสอบข้อมูล (optional)
		if item.Name == "" || item.Price <= 0 {
			return c.Status(400).JSON(fiber.Map{
				"error": "Name and Price are required and must be valid",
			})
		}

		// Insert ลง MongoDB
		collection := mongoClient.Database("myappdb").Collection("items")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		res, err := collection.InsertOne(ctx, bson.M{
			"name":       item.Name,
			"price":      item.Price,
			"created_at": time.Now(),
		})
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": fmt.Sprintf("Failed to insert item: %v", err),
			})
		}

		return c.Status(201).JSON(fiber.Map{
			"message": "Item created successfully",
			"item_id": res.InsertedID,
			"item":    item,
		})
	})

	app.Get("/hello", func(c *fiber.Ctx) error {
		// สร้างไฟล์ Excel ใหม่
		f := excelize.NewFile()

		// เพิ่มข้อมูลลงใน sheet
		sheet := "Sheet1"
		index, err := f.NewSheet(sheet) // แก้ไขการรับค่าจากฟังก์ชันให้ถูกต้อง

		if err != nil {
			log.Fatal(err)
		}

		// เขียนข้อมูลลงในเซลล์
		f.SetCellValue(sheet, "A1", "ID")
		f.SetCellValue(sheet, "B1", "Name")
		f.SetCellValue(sheet, "C1", "Price")

		// ตั้งค่า sheet เป็น active sheet
		f.SetActiveSheet(index)

		// บันทึกไฟล์
		if err := f.SaveAs("example.xlsx"); err != nil {
			log.Fatal(err)
		}

		fmt.Println("Excel file created successfully")
		return c.SendString("Hello World")
	})

	log.Printf("🚀 Server starting on port %s...\n", port)
	log.Fatal(app.Listen(":" + port))
}
