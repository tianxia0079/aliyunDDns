package dns_jobs

import (
	"github.com/GUAIK-ORG/go-snowflake/snowflake"
	"github.com/gogf/gf/frame/g"
	"github.com/golang/glog"
)

type Entity struct {
	Id      int    `orm:"id"    json:"id"`           //
	Cron    string `orm:"cron"    json:"cron"`       //
	Comment string `orm:"comment"    json:"comment"` //
	Domain  string `orm:"domain"    json:"domain"`   //
	State   int    `orm:"state" json:"state"`        //
}

func DeleteOneJob(id string) string {
	db := g.DB()
	_, err := db.Delete("dns_jobs", "id=?", id)
	if err != nil {
		glog.Error(err)
		return "fail"
	}
	return "success"
}
func QueryOneJob(id string) *Entity {
	var one *Entity
	db := g.DB()
	record, err := db.Table("dns_jobs").Where("id", id).FindOne()
	if err != nil {
		glog.Error(err)
	}
	//Struct 一定要传入指针
	record.Struct(&one)
	return one
}
func AllToStart() []*Entity {
	var list []*Entity
	db := g.DB()
	all, _ := db.Table("dns_jobs").Where("state=", 1).Order("id desc").All()
	err := all.Structs(&list)
	if err != nil {
		glog.Error(err)
	}
	return list
}
func QueryJobs(page, limit int) ([]*Entity, int) {
	var list []*Entity
	db := g.DB()
	count, _ := db.Table("dns_jobs").Count()

	all, _ := db.Table("dns_jobs").Order("id desc").Limit(limit*(page-1), limit).All()
	err := all.Structs(&list)
	if err != nil {
		glog.Error(err)
	}
	return list, count
}
func AddJob(domain, comment, cron string) (string, string) {
	db := g.DB()
	count, err1 := db.Table("dns_jobs").FindCount("domain=?", domain)
	if count > 0 {
		return "false", "该域名已经存在"
	}
	if err1 != nil {
		return "false", err1.Error()
	}
	s, err := snowflake.NewSnowflake(int64(1), int64(1))
	if err != nil {
		glog.Error(err)
	}
	id := s.NextVal()
	_, err2 := db.Insert("dns_jobs", g.Map{"id": id, "domain": domain, "comment": comment, "cron": cron})

	if err2 == nil {
		return "success", ""
	} else {
		return "false", err2.Error()
	}
}
func EditJob(id, domain, comment, cron string) string {
	db := g.DB()
	_, err := db.Update("dns_jobs", g.Map{"domain": domain, "comment": comment, "cron": cron}, "id=?", id)

	if err == nil {
		return "success"
	} else {
		return "false"
	}
}
func UpdateState(id string, state int) string {
	db := g.DB()
	_, err := db.Update("dns_jobs", g.Map{"state": state}, "id=?", id)

	if err == nil {
		return "success"
	} else {
		return "false"
	}
}
