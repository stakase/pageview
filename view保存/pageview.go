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

    db, err := sql.Open("postgres", "user=xxxx password=xxxx dbname=xxxx sslmode=disable")
    checkErr(err)

    // 時間帯の取得
    t := time.Now()
    viewHour := t.Hour()

    // データの存在チェック
    query := "SELECT url FROM pageview WHERE url = $1 and looking_hour = $2"
    var url string
    err = db.QueryRow(query, path, viewHour).Scan(&url)
    fmt.Println(err)

    // データの挿入・更新
    if err == sql.ErrNoRows {
      stmt, err := db.Prepare("INSERT INTO pageview (url,looking_hour,view_count) VALUES($1,$2,$3)")
      checkErr(err)

      stmt.Exec(path, viewHour, 1)
      checkErr(err)

    } else {
      // データ件数取得
      rows, err := db.Query("SELECT * FROM pageview WHERE url = $1 and looking_hour = $2", path, viewHour)
      checkErr(err)

      updateCount := 0
      hour := 0

      for rows.Next() {
        var url string
        var lookingHour int
        var viewCount int
        err = rows.Scan(&url, &lookingHour, &viewCount)
        checkErr(err)

        updateCount = viewCount + 1
        hour = lookingHour
      }

      // データ更新処理
      stmt, err := db.Prepare("update pageview set view_count=$1 where url=$2 and looking_hour = $3")
      checkErr(err)

      res, err := stmt.Exec(updateCount, path, hour)
      checkErr(err)

      if c, err := res.RowsAffected(); err != nil {
        fmt.Println("RowsAffectedを取得できません。: %v", err)
      } else {
        fmt.Println("RowsAffected: %v", c)
      }
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
