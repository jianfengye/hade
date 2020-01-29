package adapter

import (
	"hade/model"
	"hade/view/swagger/models"
	"reflect"
	"testing"
	"time"
)

func TestToComments(t *testing.T) {
	type args struct {
		commentsWithUser []model.Comment
	}

	comment1 := model.Comment{
		Content:   "comment1",
		CreatedAt: time.Time{},
		ID:        1,
		LeftID:    0,
		LikeCount: 0,
		ParentID:  0,
		RightID:   0,
		Status:    model.COMMENT_STATUS_REVIEWED,
		TopicID:   1,
		UpdatedAt: time.Time{},
		UserID:    1,
		User: model.User{
			CreatedAt: time.Time{},
			Email:     "test1@gmail.com",
			ID:        1,
			Name:      "tester",
			Password:  "",
			UpdatedAt: time.Time{},
		},
	}
	comment2 := model.Comment{
		Content:   "comment2",
		CreatedAt: time.Time{},
		ID:        2,
		LeftID:    0,
		LikeCount: 0,
		ParentID:  1,
		RightID:   0,
		Status:    model.COMMENT_STATUS_REVIEWED,
		TopicID:   1,
		UpdatedAt: time.Time{},
		UserID:    2,
		User: model.User{
			CreatedAt: time.Time{},
			Email:     "",
			ID:        2,
			Name:      "tester2",
			Password:  "",
			UpdatedAt: time.Time{},
		},
	}
	args1 := args{commentsWithUser: []model.Comment{comment1, comment2}}

	tests := []struct {
		name    string
		args    args
		want    *models.Comments
		wantErr bool
	}{
		{
			name: "测试1",
			args: args1,
			want: nil,
			wantErr: false,
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToComments(tt.args.commentsWithUser)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToComments() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToComments() = %v, want %v", got, tt.want)
			}
		})
	}
}
