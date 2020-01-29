package adapter

import (
	"hade/model"
	"hade/util"
	"hade/view/swagger/models"
	"github.com/jianfengye/collection"
)

// 转换成commentSummary结构
func ToCommentSummary(comment model.Comment, user *model.User) *models.CommentSummary {
	t := comment.CreatedAt.Format(util.TIME_ISO_NO_T)
	out := &models.CommentSummary{
		Content:   &comment.Content,
		CreatedAt: &t,
		ID:        &comment.ID,
		LikeCount: &comment.LikeCount,
		ParentID:  &comment.ParentID,
	}
	if user == nil {
		out.User = nil
	} else {
		out.User = ToUserSummary(*user)
	}

	return out
}


func ToComments(commentsWithUser []model.Comment) (*models.Comments, error) {
	commentColl := collection.NewObjCollection(commentsWithUser)
	level1Coll := commentColl.Filter(func(obj interface{}, index int) bool {
		comment := obj.(model.Comment)
		if comment.ParentID == 0 {
			return true
		}
		return false
	})

	out := make([]*models.CommentDetail, level1Coll.Count())
	for i := 0; i < level1Coll.Count(); i++ {
		item, err := level1Coll.Index(i).ToInterface()
		if err != nil {
			return nil, err
		}
		comment := item.(model.Comment)
		ut := ToUserSummary(comment.User)
		t1 := comment.CreatedAt.Format(util.TIME_ISO_NO_T)
		out[i] = &models.CommentDetail{
			Comments:  nil,
			Content:   &comment.Content,
			CreatedAt: &t1,
			ID:        &comment.ID,
			LikeCount: &comment.LikeCount,
			ParentID:  &comment.ParentID,
			User: ut,
		}

		// 补充二级评论
		level2Coll := commentColl.Filter(func(obj interface{}, index int) bool {
			t := obj.(model.Comment)
			if t.ParentID == comment.ID {
				return true
			}
			return false
		})
		if level2Coll.Count() == 0 {
			continue
		}
		level2Coll.SetCompare(func(a interface{}, b interface{}) int {
			aModel := a.(model.Comment)
			bModel := b.(model.Comment)
			aTime := aModel.CreatedAt
			bTime := bModel.CreatedAt
			if  aTime.Before(bTime) {
				return 1
			}
			if aTime.Equal(bTime) {
				return 0
			}
			return -1
		}).SortDesc()

		out[i].Comments = make([]*models.CommentSummary, level2Coll.Count())
		for j := 0; j < level2Coll.Count(); j++{
			t, err := level2Coll.Index(j).ToInterface()
			if err != nil {
				return nil, err
			}
			ts := t.(model.Comment)
			out[i].Comments[j] = ToCommentSummary(ts, &ts.User)
		}
	}
	o := models.Comments(out)
	return &o, nil
}