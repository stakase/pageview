package main

import (
  "github.com/gin-gonic/gin"
  "net/http"
  //"fmt"
  "database/sql"
  "sort"
  _ "github.com/lib/pq"
)

type Review struct {
    Time int
    ViewCount int
}

type Reviews []Review

// ソート
func (r Reviews) Len() int {
    return len(r)
}

func (r Reviews) Swap(i, j int) {
    r[i], r[j] = r[j], r[i]
}

func (r Reviews) Less(i, j int) bool {
    return r[i].Time < r[j].Time
}


func main() {
  router := gin.Default()
  router.LoadHTMLGlob("template/*.tmpl")

  router.GET("/reviewlist", func(c *gin.Context) {
    db, err := sql.Open("postgres", "user=xxxx password=xxxx dbname=xxxx sslmode=disable")
    checkErr(err)

    // データの取得処理
    rows, err := db.Query("SELECT url FROM pageview GROUP BY url")
    checkErr(err)
    
    urlList := []string{}
    for rows.Next() {
      var url string

      err = rows.Scan(&url)
      checkErr(err)

      urlList = append(urlList, url)
    }

    c.HTML(http.StatusOK, "reviewlist.tmpl", gin.H{
      "url": urlList,
    })

    db.Close()
  })

  router.GET("/reviewdetail", func(c *gin.Context) {
    url := c.Query("url")
    
    reviews := getPv(url)  // PV数取得処理

    c.HTML(http.StatusOK, "review.tmpl", gin.H{
      "review": reviews,
    })
  })

  router.GET("/reviewall", func(c *gin.Context) {
    reviews := getPv("")  // PV数取得処理

    c.HTML(http.StatusOK, "review.tmpl", gin.H{
      "review": reviews,
    })
  })
  
  router.Run(":8080")
}


// PV数表示用処理
func getPv(url string) (Reviews) {
  // 初期化
  viewMap := map[int]int {}
  for i := 0; i < 24; i++ {
    viewMap[i] = 0;
  }

  db, err := sql.Open("postgres", "user=exam password=exam dbname=exam sslmode=disable")
  checkErr(err)
  
  var rows *sql.Rows
  
  if url =="" {
    rows, err = db.Query("SELECT looking_hour, sum(view_count) as view_count FROM pageview GROUP BY looking_hour")
    checkErr(err)

  } else {
    rows, err = db.Query("SELECT looking_hour, sum(view_count) as view_count FROM pageview WHERE url = $1 GROUP BY looking_hour", url)
    checkErr(err)
  }

  for rows.Next() {
    var lookingHour int
    var viewCount int

    err = rows.Scan(&lookingHour, &viewCount)
    checkErr(err)

    viewMap[lookingHour] = viewCount
  }
  
  var reviews Reviews
  for k, v := range viewMap {
    review := Review{k, v}
    reviews = append(reviews, review)
  }
  sort.Sort(reviews)

  return reviews
} 


// DB関連のエラー判定
func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
