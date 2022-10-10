package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"myBlog/dao"
	"myBlog/model"
	"net/http"
	"sort"
)

//
func GetVisitsInfos(c *gin.Context) {

	type Item struct {
		Date  string `json:"date"`
		Count int    `json:"count"`
	}

	// 统计最新的结果，只保存7天的
	resp, err := dao.DataInterface.VisitsInfos(30)
	if err != nil {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	}

	var Items []Item
	tmp := make(map[string]int)
	for i := 0; i < len(resp); i++ {
		month := resp[i].UpdatedAt.Month()
		day := resp[i].UpdatedAt.Day()
		time := fmt.Sprintf("%d-%d", month, day)

		tmp[time] += resp[i].Count
	}

	for k, v := range tmp {
		item := Item{Date: k, Count: v}
		Items = append(Items, item)
	}

	// 从小到大排序(稳定排序)
	sort.SliceStable(Items, func(i, j int) bool {
		if Items[i].Date < Items[j].Date {
			return true
		}
		return false
	})

	c.JSON(http.StatusOK, Items)
}

// 统计站点所有的访问量
func GetTotalCount(c *gin.Context) {

	total, err := dao.DataInterface.TotalCount()
	if err != nil {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"code":   200,
		"total":  total,
	})
}
