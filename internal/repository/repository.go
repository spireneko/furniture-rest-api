package repository

import (
	"encoding/json"
	"log"
	"os"

	"github.com/spireneko/furniture-rest-api/internal/model"
)

type JSONDB struct {
	Path          string
	FurnitureJSON FurnitureJSON
}

type FurnitureJSON struct {
	LastID         int64             `json:"last_id"`
	FurnitureArray []model.Furniture `json:"furniture_array"`
}

func NewJSONDB(path string) JSONDB {
	file, err := os.OpenFile(path, os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if err := file.Close(); err != nil {
		log.Fatal(err)
	}

	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	furnJSON := new(FurnitureJSON)
	if err := json.Unmarshal(content, furnJSON); err != nil {
		log.Println("Database is empty")
		return JSONDB{
			Path:          path,
			FurnitureJSON: FurnitureJSON{LastID: 0, FurnitureArray: []model.Furniture{}},
		}
	}

	return JSONDB{
		Path:          path,
		FurnitureJSON: *furnJSON,
	}
}

func Create(data *model.Furniture, db *JSONDB) error {
	db.FurnitureJSON.LastID += 1
	data.ID = db.FurnitureJSON.LastID
	db.FurnitureJSON.FurnitureArray = append(db.FurnitureJSON.FurnitureArray, *data)

	return updateDB(db)
}

func Get(id int64, db *JSONDB) *model.Furniture {
	for _, furniture := range db.FurnitureJSON.FurnitureArray {
		if furniture.ID == id {
			return &furniture
		}
	}

	return nil
}

func Update(id int64, db *JSONDB, newFurniture *model.Furniture) error {
	for index, furniture := range db.FurnitureJSON.FurnitureArray {
		if furniture.ID == id {
			newFurniture.ID = db.FurnitureJSON.FurnitureArray[index].ID
			db.FurnitureJSON.FurnitureArray[index] = *newFurniture

			return updateDB(db)
		}
	}

	return nil
}

func Patch(id int64, db *JSONDB, newFurniture *model.Furniture) error {
	for index, furniture := range db.FurnitureJSON.FurnitureArray {
		if furniture.ID == id {
			if len(newFurniture.Name) > 0 {
				db.FurnitureJSON.FurnitureArray[index].Name = newFurniture.Name
			}
			if len(newFurniture.Fabricator) > 0 {
				db.FurnitureJSON.FurnitureArray[index].Fabricator = newFurniture.Fabricator
			}
			if newFurniture.Height > 0 {
				db.FurnitureJSON.FurnitureArray[index].Height = newFurniture.Height
			}
			if newFurniture.Width > 0 {
				db.FurnitureJSON.FurnitureArray[index].Width = newFurniture.Width
			}
			if newFurniture.Length > 0 {
				db.FurnitureJSON.FurnitureArray[index].Length = newFurniture.Length
			}

			return updateDB(db)
		}
	}

	return nil
}

func Delete(id int64, db *JSONDB) error {
	for index, furniture := range db.FurnitureJSON.FurnitureArray {
		if furniture.ID == id {
			db.FurnitureJSON.FurnitureArray = append(db.FurnitureJSON.FurnitureArray[:index], db.FurnitureJSON.FurnitureArray[index+1:]...)

			return updateDB(db)
		}
	}

	return nil
}

func updateDB(db *JSONDB) error {
	text, err := json.Marshal(db.FurnitureJSON)
	if err != nil {
		return err
	}

	if err := os.WriteFile(db.Path, text, 0644); err != nil {
		return err
	}

	return nil
}
