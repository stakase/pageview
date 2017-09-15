package main

import (
  "github.com/gin-gonic/gin"
  "fmt"
  "time"
  "database/sql"
  _ "github.com/lib/pq"
)

func main() {
  router := gin.Default()
  router.GET("/pageview", func(c *gin.Context) {
    path := c.Query("keyword")

    db, err := sql.Open("postgres", "user=xxxxx password=xxxxx dbname=xxxxx sslmode=disable")
    checkErr(err)

    // 時間帯の取得
    t := time.Now()
    viewHour := t.Hour()

    // データの存在チェック
    rows, err := db.Query("SELECT COUNT(1) FROM pageview WHERE url = $1 and looking_hour = $2", path, viewHour)
    checkErr(err)

    dataCount2 := 0

    for rows.Next() {
      var dataCount int
      err = rows.Scan(&dataCount)
      checkErr(err)

      fmt.Println(dataCount)

      dataCount2 = dataCount
    }

    // データの挿入・更新
    fmt.Println(dataCount2)
    if dataCount2 == 0 {
      stmt, err := db.Prepare("INSERT INTO pageview (url,looking_hour,view_count) VALUES($1,$2,$3)")
      checkErr(err)

      stmt.Exec(path, viewHour, 1)
      checkErr(err)

    } else {
      // データ件数取得
      rows, err := db.Query("SELECT * FROM pageview WHERE url = $1 and looking_hour = $2", path, viewHour)
      checkErr(err)

      updateCount := 0

      for rows.Next() {
        var url string
        var lookingHour int
        var viewCount int
        err = rows.Scan(&url, &lookingHour, &viewCount)
        checkErr(err)

        updateCount = viewCount + 1
      }

      // データ更新処理
      stmt, err := db.Prepare("update pageview set view_count=$1 where url=$2")
      checkErr(err)

      res, err := stmt.Exec(updateCount, path)
      checkErr(err)

      affect, err := res.RowsAffected()
      checkErr(err)

      fmt.Println(affect)

      fmt.Println("重複データあり")
    }

    db.Close()
  })
  router.Run(":8080")
}


// DB関連のエラー判定
func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
