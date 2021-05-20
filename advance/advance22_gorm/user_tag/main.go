package main

import (
	"log"

	"github.com/ForrestSu/go_learn/advance/advance22_gorm/dao"
	"github.com/scylladb/go-set/uset"
)

func QueryUsers() {

	/**
	select t_tag.category_tag_id,
	 group_concat(distinct t_user_tags.user_id) as user_ids
	from t_user_tags
	left join t_tag on t_user_tags.tag_id = t_tag.id
	-- where t_user_tags.tag_id in (1,2, 8,9,18,)
	group by t_tag.category_tag_id
	*/
	var relations []UserTag
	dao.GetDB().Table("t_user_tags").
		Select("t_tag.category_tag_id, t_user_tags.user_id, t_user_tags.tag_id").
		Joins("left join t_tag on t_user_tags.tag_id = t_tag.id").
		// Where("t_user_tags.tag_id IN ?", tagIDs).
		Order("t_tag.category_tag_id").
		Scan(&relations)

	Solve(relations)
	// 思路二: 根据传入的标签,查询这些标签下的所有用户 (按category_id 排序)
	// 1 然后把同一 category_id 下的用户ID,  合并+去重
	// 2 把不同category_id 用户组做 mapreduce,  最后获取用户组中计数值 == category_id总数的用户
	// 3 获取userID 列表，走正常的 in 查询规则

	//var sets []*strset.Set
	//for _, group := range groups {
	//	sets = append(sets, group.UserIDs.Data)
	//}
	//var selectUserIDs = strset.Intersection(sets...)
	//
	//log.Printf("select userID = %+v\n", selectUserIDs.List())
}

func Solve(records []UserTag) {
	if len(records) <= 0 {
		return
	}
	// add first set
	var prevID = records[0].CategoryTagID
	// var total = len(records)
	var curSet = uset.New()
	var sets = []*uset.Set{curSet}
	for _, r := range records {
		if r.CategoryTagID == prevID {
			curSet.Add(r.UserID)
		} else {
			prevID = r.CategoryTagID
			curSet = uset.New(r.UserID)
			sets = append(sets, curSet)
		}
	}
	// debug log
	for _, s := range sets {
		log.Println(s.List())
	}
	// 计算交集
	var selectUserIDSet = uset.Intersection(sets...)
	var userIDs = selectUserIDSet.List()
	log.Println(">>> ", userIDs)
}

func main() {
	QueryUsers()

}
