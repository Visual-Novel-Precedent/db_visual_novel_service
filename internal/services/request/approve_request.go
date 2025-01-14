package request

import (
	"db_novel_service/internal/storage"
	"errors"
	"gorm.io/gorm"
)

const (
	RequestSuperAdmin        = 0
	RequestRegistrationAdmin = 1
	RequestPublishChapter    = 2
	RequestDeleteChapter     = 3

	SuperAdminStatus    = 1
	ApprovedAdminStatus = 0

	PublishedChapterStatus = 2

	ApprovedRequestStatus = 1
)

func ApproveRequest(id int64, db *gorm.DB) error {
	request, err := storage.SelectRequestWIthId(db, id)

	if err != nil {
		return err
	}

	if request.Type == RequestSuperAdmin {
		ad, err := storage.SelectAdminWIthId(db, request.RequestingAdmin)

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
		ad, err := storage.SelectAdminWIthId(db, request.RequestingAdmin)

		if err != nil {
			return errors.New("error to search admin")
		}

		ad.AdminStatus = ApprovedAdminStatus

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

	return err
}
