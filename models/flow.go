package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Flow entity
type Flow struct {
	ID                  int       `gorm:"Column:id;PRIMARY_KEY;AUTO_INCREMENT" json:"id"`
	Name                string    `gorm:"Column:name;NOT NULL" json:"name"`
	TriggerType         string    `gorm:"Column:triggerType;NOT NULL" json:"triggerType"`
	CronExpression      string    `gorm:"Column:cronExpression;type:varchar(50)" json:"cron"`
	CronTimeZone        string    `gorm:"Column:cronTimeZone;type:varchar(50)" json:"cronTimeZone"`
	State               string    `gorm:"Column:state;size:50" json:"state"`
	Token               string    `gorm:"Column:token;size:255" json:"-"`
	NodeCount           int       `gorm:"Column:nodeCount" json:"nodeCount"`
	LastExecutedAt      time.Time `gorm:"Column:lastExecutedAt" json:"lastExecutedAt"`
	LastExecutedSummary string    `gorm:"Column:lastExecutedSummary;type:text" json:"lastExecutedSummary"`
	RunAt               string    `gorm:"Column:runAt;type:text" json:"runAt"` // 在哪个worker上被执行
	HostedOn            string    `gorm:"Column:hostedOn" json:"hostedOn"`     // 在哪里触发，slave
	Pointer             string    `gorm:"Column:pointer" json:"-"`             // 创建后的cronJob在内存里的ID

	Owner int    `gorm:"Column:owner" json:"owner"`
	Tags  string `gorm:"Column:tags;type:text" json:"tags"`

	Comment   string    `gorm:"Column:comment;type:text" json:"comment"`
	CreatedAt time.Time `gorm:"Column:createdAt" json:"createdAt"`
	UpdatedAt time.Time `gorm:"Column:updatedAt" json:"updatedAt"`
}

// // FlowView is the model-view of Flow
// type FlowView struct {
// }

// TableName of Flow entity for ORM
func (Flow) TableName() string {
	return "flow"
}

// Create a new Flow
func (f *Flow) Create() *gorm.DB {
	return DatabaseEngine.Create(f)
}

// List flows
func (f *Flow) List() []Flow {
	flows := make([]Flow, 0)
	DatabaseEngine.Find(&flows)
	return flows
}
