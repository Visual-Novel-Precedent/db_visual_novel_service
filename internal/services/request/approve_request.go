package request

import (
	"db_novel_service/internal/models"
	"db_novel_service/internal/storage"
	"errors"
	"gorm.io/gorm"
	"log"
)

const (
	RequestSuperAdmin        = 0
	RequestRegistrationAdmin = 1
	RequestPublishChapter    = 2
	RequestDeleteChapter     = 3

	SuperAdminStatus    = 1
	ApprovedAdminStatus = 0

	PublishedChapterStatus = 3

	ApprovedRequestStatus = 1
)

func ApproveRequest(id int64, db *gorm.DB) error {
	allrequests, err := storage.GetAllRequests(db)

	log.Println(allrequests)

	var request models.Request

	for _, r := range allrequests {
		if r.Id == id {
			request = r
		}
	}

	log.Println("request:", request)

	if err != nil {
		return err
	}

	if request.Type == RequestSuperAdmin {
		ad, err := storage.SelectAdminWithId(db, request.RequestingAdmin)

		if err != nil {
			return errors.New("error to search admin")
		}

		ad.AdminStatus = SuperAdminStatus

		_, err = storage.UpdateAdmin(db, request.RequestingAdmin, ad)

		if err != nil {
			return errors.New("error to update admin status to superAdmin")
		}
	}

	if request.Type == RequestRegistrationAdmin {
		ad, err := storage.SelectAdminWithId(db, request.RequestingAdmin)

		if err != nil {
			return errors.New("error to search admin")
		}

		ad.AdminStatus = ApprovedAdminStatus

		log.Println(ad)

		_, err = storage.UpdateAdmin(db, request.RequestingAdmin, ad)

		if err != nil {
			return errors.New("error to update admin status to approvedAdmin")
		}
	}

	if request.Type == RequestPublishChapter {
		chapter, err := storage.SelectChapterWIthId(db, request.RequestedChapterId)

		if err != nil {
			return errors.New("error to search chapter")
		}

		chapter.Status = PublishedChapterStatus

		_, err = storage.UpdateChapter(db, chapter.Id, chapter)

		if err != nil {
			return errors.New("error to update chapter status to published")
		}
	}

	if request.Type == RequestDeleteChapter {
		chapter, err := storage.SelectChapterWIthId(db, request.RequestedChapterId)

		if err != nil {
			return errors.New("error to search chapter")
		}

		_, err = storage.DeleteChapter(db, chapter.Id)

		if err != nil {
			return errors.New("error to delete chapter")
		}
	}

	request.Status = ApprovedRequestStatus

	_, err = storage.UpdateRequest(db, request.Id, request)

	_, err = storage.DeleteRequest(db, id)

	return err
}
