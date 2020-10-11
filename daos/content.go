package daos

import (
	"github.com/tombiers/estuary-backend/models"
)

// GetContentByUUID returns content with the given uuid
func GetContentByUUID(uuid string) models.ContentDB {
	var dbContent models.ContentDB
	db.Where("UUID = ?", uuid).First(&dbContent)
	return dbContent
}

// CreateContent creates a new content
func CreateContent(content models.ContentDB) models.ContentDB {
	db.Create(&content)
	return content
}

// UpdateContent updates the content with the given uuid to the given data and return the updated content
func UpdateContent(uuid string, content models.ContentDB) models.ContentDB {
	var dbContent models.ContentDB
	db.Where("UUID = ?", uuid).First(&dbContent).Updates(content)
	return dbContent
}

// DeleteContent delete the content with the given uuid
func DeleteContent(uuid string) {
	var dbContent models.ContentDB
	db.Where("UUID = ?", uuid).First(&dbContent).Delete(&dbContent)
}

// GetContentsFromWorkshop returns all contents belonging to the workshop with the given UUID
func GetContentsFromWorkshop(workshopUUID string) []models.ContentDB {
	dbContent := []models.ContentDB{}
	db.Where("workshop_UUID = ?", workshopUUID).Find(&dbContent)
	return dbContent
}