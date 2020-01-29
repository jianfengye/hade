package suite1
/*
import (
	"hade/bootstrap/connection"
	"hade/controller/comment"
	"hade/controller/tag"
	"hade/controller/topic"
	"hade/model"
	"hade/view/swagger/models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func NormalCase1(t *testing.T) {
	Convey("测试正常流程", t, func() {
		router := gin.New()
		api := router.Group("/api")
		topic.Register(api)
		tag.Register(api)
		comment.Register(api)

		Convey("查找所有标签", func() {
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/api/tag/list", nil)
			router.ServeHTTP(w, r)
			resp := w.Result()
			res, _ := ioutil.ReadAll(resp.Body)

			sRes := []*models.Tag{}

			err := json.Unmarshal(res, &sRes)
			So(err, ShouldBeNil)

			Convey("检查标签个数为3", func() {
				So(len(sRes), ShouldEqual, 3)
			})

		})

		Convey("创建一个话题，带有1个标签", func() {
			params := url.Values{
				"title": []string{"php是世界上最好的语言"},
				"content": []string{"<p>php是世界上最好的语言</p>"},
				"link": []string{"http://baidu.com"},
				"tag_ids": []string{"1"},
			}
			body := strings.NewReader(params.Encode())
			w := httptest.NewRecorder()

			r1 := httptest.NewRequest("POST", "/api/topic/create", body)
			r1.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			router.ServeHTTP(w, r1)
			resp := w.Result()
			So(resp.StatusCode, ShouldEqual, http.StatusOK)

			var sum int
			connection.Default.Model(&model.Topic{}).Count(&sum)

			var topic1 model.Topic
			connection.Default.Where("title= ?", "php是世界上最好的语言").First(&topic1)

			Convey("查找这个话题", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest("GET", "/api/topic/detail?topic_id=" + fmt.Sprint(topic1.ID), nil)
				router.ServeHTTP(w, r)
				resp := w.Result()
				res, err := ioutil.ReadAll(resp.Body)
				So(resp.StatusCode,ShouldEqual, http.StatusOK)

				So(err, ShouldBeNil)
				sRes := models.Topic{}
				err = json.Unmarshal(res, &sRes)
				So(err, ShouldBeNil)

				Convey("检查话题内容", func() {
					So(sRes.Content, ShouldEqual, topic1.Content)
				})

				Convey("检查话题标签", func() {
					So(len(sRes.Tags),ShouldEqual, 1)
				})

			})

			Convey("更新话题的内容", func() {

				params := url.Values{
					"title": []string{"php是世界上最好的语言??"},
					"content": []string{"<p>php是世界上最好的语言</p>"},
					"link": []string{""},
					"tag_ids": []string{"2"},
					"topic_id": []string{fmt.Sprint(topic1.ID)},
				}
				body := strings.NewReader(params.Encode())
				w := httptest.NewRecorder()

				r := httptest.NewRequest("POST", "/api/topic/update", body)
				r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
				router.ServeHTTP(w, r)
				resp := w.Result()
				res, err := ioutil.ReadAll(resp.Body)
				So(err, ShouldBeNil)
				t.Log(string(res))
				So(resp.StatusCode, ShouldEqual, http.StatusOK)

				Convey("查找这个话题", func() {

					w := httptest.NewRecorder()
					r := httptest.NewRequest("GET", "/api/topic/detail?topic_id=" + fmt.Sprint(topic1.ID), nil)
					router.ServeHTTP(w, r)
					resp := w.Result()
					res, err := ioutil.ReadAll(resp.Body)
					So(resp.StatusCode,ShouldEqual, http.StatusOK)

					So(err, ShouldBeNil)
					sRes := models.Topic{}
					err = json.Unmarshal(res, &sRes)
					So(err, ShouldBeNil)

					Convey("检查话题内容", func() {
						So(sRes.Content, ShouldEqual, "<p>php是世界上最好的语言</p>")
						So(sRes.Title, ShouldEqual, "php是世界上最好的语言??")
					})
				})

			})

			Convey("创建话题的评论", func() {

				params := url.Values{
					"content": []string{"<p>严重同意</p>"},
					"topic_id": []string{fmt.Sprint(topic1.ID)},
				}
				body := strings.NewReader(params.Encode())
				w := httptest.NewRecorder()

				r := httptest.NewRequest("POST", "/api/comment/create", body)
				r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
				router.ServeHTTP(w, r)
				resp := w.Result()
				So(resp.StatusCode, ShouldEqual, http.StatusOK)
				res, err := ioutil.ReadAll(resp.Body)
				So(err, ShouldBeNil)
				t.Log(string(res))

				Convey("查找话题的所有评论", func() {
					w := httptest.NewRecorder()
					r := httptest.NewRequest("GET", "/api/comment/list?topic_id=" + fmt.Sprint(topic1.ID), nil)
					router.ServeHTTP(w, r)
					resp := w.Result()
					res, err := ioutil.ReadAll(resp.Body)
					t.Log(string(res))
					So(resp.StatusCode,ShouldEqual, http.StatusOK)

					Convey("检查话题的所有评论格式", func() {
						sRes := models.Comments{}
						err = json.Unmarshal(res, &sRes)
						So(err, ShouldBeNil)
					})

				})

				var comment1 model.Comment
				connection.Default.Model(&model.Comment{}).Where("topic_id =? ", topic1.ID).First(&comment1)

				Convey("创建评论的跟评", func() {

					params := url.Values{
						"content": []string{"php是世界上最好的语言??"},
						"comment_id" : []string{fmt.Sprint(comment1.ID)},
					}
					body := strings.NewReader(params.Encode())
					w := httptest.NewRecorder()

					r := httptest.NewRequest("POST", "/api/comment/append", body)
					r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
					router.ServeHTTP(w, r)
					resp := w.Result()
					res, err := ioutil.ReadAll(resp.Body)
					So(err, ShouldBeNil)
					t.Log(string(res))
					So(resp.StatusCode, ShouldEqual, http.StatusOK)

					Convey("查找话题的所有评论", func() {
						w := httptest.NewRecorder()
						r := httptest.NewRequest("GET", "/api/comment/list?topic_id=" + fmt.Sprint(topic1.ID), nil)
						router.ServeHTTP(w, r)
						resp := w.Result()
						res, err := ioutil.ReadAll(resp.Body)
						t.Log(string(res))
						So(resp.StatusCode,ShouldEqual, http.StatusOK)

						Convey("检查话题的所有评论格式", func() {
							sRes := models.Comments{}
							err = json.Unmarshal(res, &sRes)
							So(err, ShouldBeNil)
							So(len(sRes[0].Comments), ShouldEqual, 1)
						})

					})
				})

			})
		})
	})

}

 */