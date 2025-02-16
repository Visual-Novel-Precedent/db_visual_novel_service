package storage

import (
	"database/sql"
	"db_novel_service/internal/models"
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	"time"
)

func RegisterChapter(db *gorm.DB, chapter models.Chapter) (int64, error) {
	result := db.Create(&chapter)
	if result.RowsAffected == 0 {
		return 0, errors.New("chapter not created")
	}
	return result.RowsAffected, nil
}

//func SelectChapterWIthId(db *gorm.DB, id int64) (models.Chapter, error) {
//	var chapter models.Chapter
//	result := db.First(&chapter, "id = ?", id)
//	if result.RowsAffected == 0 {
//		return models.Chapter{}, errors.New("chapter data not found")
//	}
//	return chapter, nil
//}

func SelectChapterWIthId(db *gorm.DB, id int64) (models.Chapter, error) {
	var chapter models.Chapter
	query := `
        SELECT id, name, start_node,
               CAST(nodes AS TEXT) as nodes_raw,
               CAST(characters AS TEXT) as characters_raw,
               status,
               CAST(updated_at AS TEXT) as updated_at_raw,
               author
        FROM chapters
        WHERE id = $1
        LIMIT 1
    `

	row := db.Raw(query, id).Row()
	if err := row.Err(); err != nil {
		return models.Chapter{}, err
	}

	var (
		nodesRaw      string
		charactersRaw string
		updatedAtRaw  string
	)

	err := row.Scan(
		&chapter.Id,
		&chapter.Name,
		&chapter.StartNode,
		&nodesRaw,
		&charactersRaw,
		&chapter.Status,
		&updatedAtRaw,
		&chapter.Author,
	)

	log.Println(chapter)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.Chapter{}, errors.New("chapter data not found")
		}
		return models.Chapter{}, err
	}

	// Десериализация JSON полей
	if err := json.Unmarshal([]byte(nodesRaw), &chapter.Nodes); err != nil {
		return models.Chapter{}, fmt.Errorf("failed to unmarshal nodes: %w", err)
	}

	if err := json.Unmarshal([]byte(charactersRaw), &chapter.Characters); err != nil {
		return models.Chapter{}, fmt.Errorf("failed to unmarshal characters: %w", err)
	}

	// Специальная обработка UpdatedAt
	var updatedAtMap map[time.Time]int64
	if err := json.Unmarshal([]byte(updatedAtRaw), &updatedAtMap); err != nil {
		return models.Chapter{}, fmt.Errorf("failed to unmarshal updated_at: %w", err)
	}
	chapter.UpdatedAt = updatedAtMap

	return chapter, nil
}

func UpdateChapter(db *gorm.DB, id int64, newChapter models.Chapter) (models.Chapter, error) {
	var chapter models.Chapter

	nodesJSON, err := json.Marshal(newChapter.Nodes)
	charactersJSON, err := json.Marshal(newChapter.Characters)
	updatedInfoJSON, err := json.Marshal(newChapter.UpdatedAt)

	if err != nil {
		return models.Chapter{}, err
	}

	result := db.Model(&chapter).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"name":       newChapter.Name,
			"StartNode":  newChapter.StartNode,
			"Nodes":      json.RawMessage(nodesJSON),
			"Characters": json.RawMessage(charactersJSON),
			"Status":     newChapter.Status,
			"UpdatedAt":  json.RawMessage(updatedInfoJSON),
		})

	if result.RowsAffected == 0 {
		return models.Chapter{}, errors.New("chapter data not update")
	}

	return chapter, nil
}

func DeleteChapter(db *gorm.DB, id int64) (int64, error) {
	var deletedChapter models.Chapter
	result := db.Where("id = ?", id).Delete(&deletedChapter)
	if result.RowsAffected == 0 {
		return 0, errors.New("payment data not update")
	}
	return result.RowsAffected, nil
}

func GetChaptersForAdmin(db *gorm.DB) ([]models.Chapter, error) {
	var chapters []models.Chapter
	result := db.Find(&chapters)

	if result.Error != nil {
		return nil, result.Error
	}

	return chapters, nil
}

func FindPublishedChapters(db *gorm.DB) ([]models.Chapter, error) {
	var chapters []models.Chapter
	query := `
        SELECT 
            id,
            name,
            start_node,
            CAST(nodes AS TEXT) as nodes_raw,
            CAST(characters AS TEXT) as characters_raw,
            status,
            CAST(updated_at AS TEXT) as updated_at_raw,
            author
        FROM chapters
        WHERE status = $1
    `

	rows, err := db.Raw(query, 3).Rows()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch published chapters: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			chapter       models.Chapter
			nodesRaw      string
			charactersRaw string
			updatedAtRaw  string
		)

		err = rows.Scan(
			&chapter.Id,
			&chapter.Name,
			&chapter.StartNode,
			&nodesRaw,
			&charactersRaw,
			&chapter.Status,
			&updatedAtRaw,
			&chapter.Author,
		)

		if err != nil {
			return nil, fmt.Errorf("failed to scan chapter: %w", err)
		}

		// Десериализация JSON полей
		if err := json.Unmarshal([]byte(nodesRaw), &chapter.Nodes); err != nil {
			return nil, fmt.Errorf("failed to unmarshal nodes: %w", err)
		}

		if err := json.Unmarshal([]byte(charactersRaw), &chapter.Characters); err != nil {
			return nil, fmt.Errorf("failed to unmarshal characters: %w", err)
		}

		// Обработка UpdatedAt
		var updatedAtMap map[time.Time]int64
		if err := json.Unmarshal([]byte(updatedAtRaw), &updatedAtMap); err != nil {
			return nil, fmt.Errorf("failed to unmarshal updated_at: %w", err)
		}
		chapter.UpdatedAt = updatedAtMap

		chapters = append(chapters, chapter)
	}

	return chapters, nil
}
