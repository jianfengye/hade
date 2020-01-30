package suite2

import (
	"hade/bootstrap/connection"
	"hade/controller/knowledge"
	"hade/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"net/http/httptest"
	"testing"
	"time"
)

func NormalCase1(t *testing.T) {
	daily1 := model.Daily{
		CreatedAt: time.Time{},
		ID:        1,
		Title:     "这是标题",
		Author:    "轩脉刃",
		Day:       "2019-9-1",
		Content:   `[{"title": "这是标题", "link": "http://baidu.com", "comment": "这是评论"}]`,
		UpdatedAt: time.Time{},
	}
	connection.Default.Create(daily1)
	Convey("测试正常流程", t, func() {
		router := gin.New()
		api := router.Group("/api")
		knowledge.Register(api)

		Convey("查找当前日期", func() {
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/api/knowledge/daily?day=2019-9-1", nil)
			router.ServeHTTP(w, r)
			resp := w.Result()
			res, _ := ioutil.ReadAll(resp.Body)

			sRes := models.Daily{}

			err := json.Unmarshal(res, &sRes)
			So(err, ShouldBeNil)

			Convey("检查返回值是否正确", func() {
				So(1, ShouldEqual, len(sRes.Content))
				So("这是标题", ShouldEqual, *sRes.Title)
			})
		})

		Convey("查找不存在日期", func() {
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/api/knowledge/daily?day=2019-9-2", nil)
			router.ServeHTTP(w, r)
			resp := w.Result()
			res, _ := ioutil.ReadAll(resp.Body)

			sRes := models.ErrorResponse{}

			err := json.Unmarshal(res, &sRes)
			So(err, ShouldBeNil)

			Convey("检查返回值是否正确", func() {
				So("没有当日的信息", ShouldEqual, *sRes.Msg)
			})
		})

	})

}