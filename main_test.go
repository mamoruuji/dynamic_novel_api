package main

import (
	_ "github.com/lib/pq"
)

// func TestGetDynamicsHandler(t *testing.T) {
// 	// t.Parallel()
// 	db, err := sql.Open("postgres", "host=db port=5432 user=postgres password=pass dbname=dynamic_novel sslmode=disable")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	boil.SetDB(db)

// 	mux := handler.server()
// 	server := httptest.NewUnstartedServer(mux)
// 	server.EnableHTTP2 = true
// 	server.StartTLS()
// 	t.Cleanup(server.Close)

// 	// TODOリストのテストデータを準備
// 	cases := []*models.Dynamic{
// 		{
// 			DynamicID: 1,
// 			Title:     "hoge",
// 			Overview:  "hoge1",
// 			UserID:    "userId1",
// 			Published: true,
// 		},
// 		{
// 			DynamicID: 2,
// 			Title:     "hogehoge",
// 			Overview:  "hoge2",
// 			UserID:    "userId2",
// 			Published: false,
// 		},
// 		{
// 			DynamicID: 3,
// 			Title:     "hogehogehoge",
// 			Overview:  "hoge3",
// 			UserID:    "userId3",
// 			Published: true,
// 		},
// 	}

// 	// テストデータをデータベースに挿入
// 	for _, dynamic := range cases {
// 		err := dynamic.Insert(context.Background(), db, boil.Infer())
// 		if err != nil {
// 			t.Fatalf("failed to insert test data: %v", err)
// 		}
// 	}

// 	// テスト関数の最後にテストデータを削除
// 	defer func() {
// 		for _, dynamic := range cases {
// 			_, err := dynamic.Delete(context.Background(), db)
// 			if err != nil {
// 				t.Fatalf("failed to delete test data: %v", err)
// 			}
// 		}
// 	}()

// 	for _, c := range cases {
// 		c := c
// 		t.Run(c.Title, func(t *testing.T) {
// 			// t.Parallel()
// 			client := dynamicv1connect.NewDynamicServiceClient(
// 				server.Client(),
// 				server.URL,
// 			)

// 			res, err := client.ListDynamics(context.Background(), connect.NewRequest(&dynamicv1.ListDynamicsRequest{}))
// 			if err != nil {
// 				t.Error(err)
// 			}

// 			result := res.Msg.Dynamics

// 			// テストデータと取得結果を比較
// 			if len(result) != len(cases) {
// 				t.Fatalf("unexpected number of todos, got: %d, want: %d", len(result), len(cases))
// 			}
// 			for i := range result {
// 				if result[i].DynamicId != cases[i].DynamicID || result[i].Title != cases[i].Title {
// 					t.Fatalf("mismatched todo data, got: %v, want: %v", result[i], cases[i])
// 				}
// 			}

// 			if res.Msg.GetDynamics() != nil {
// 				t.Errorf("dynamics got: %s, want: %s", res.Msg.GetDynamics(), c.Title)
// 			}
// 		})
// 	}
// }
