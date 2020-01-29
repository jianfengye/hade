package suite1

/*
func CaseTopic(t *testing.T) {
	Convey("测试topic相关流程", t, func() {
		router := gin.New()
		api := router.Group("/api")
		topic.Register(api)
		tag.Register(api)
		comment.Register(api)

		// 创建20个话题
		for i := 1; i <= 20; i++ {
			topic := model.Topic{
				Category:  "",
				Content:   "tester" + fmt.Sprint(i),
				CreatedAt: time.Now(),
				LikeCount: rand.Int63() + 1,
				Link:      "",
				Score:     0,
				Source:    "",
				Status:    model.TOPIC_STATUS_REVIEWED,
				Title:     "tester title " + fmt.Sprint(i),
				UpdatedAt: time.Now(),
				UserID:    1,
			}
			connection.Default.Save(&topic)
		}

		Convey("测试获取分页数据", func() {
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/api/topic/list?page=1&size=10", nil)
			router.ServeHTTP(w, r)
			resp := w.Result()
			res, err := ioutil.ReadAll(resp.Body)
			So(resp.StatusCode,ShouldEqual, http.StatusOK)

			Convey("检查话题的格式", func() {
				sRes := models.GetTopicListOKBody{}
				err = json.Unmarshal(res, &sRes)
				So(err, ShouldBeNil)

				So(len(sRes.Topics), ShouldEqual, 0)
				So(sRes.Pager.TotalPage, ShouldEqual, 2)
				So(sRes.Topics[0].LikeCount, ShouldBeGreaterThan, 0)
			})
		})
	})

}

 */