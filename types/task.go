package types

import (
	"fmt"
	"strconv"
	"time"
)

type Metadata struct {
	LastInsertedId int `json:"last_inserted_id"`
}

type Task struct {
	Id          int
	Description string
	Status      Status
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (t Task) Encode() []string {
	return []string{
		strconv.Itoa(t.Id),
		t.Description,
		fmt.Sprintf("%s", t.Status),
		t.CreatedAt.Format("02/01/2006"),
		t.UpdatedAt.Format("02/01/2006"),
	}
}

func (t *Task) Decode(content []string) {
	id, _ := strconv.ParseInt(content[0], 10, 32)
	t.Id = int(id)
	t.Description = content[1]
	t.Status.Decode(content[2])
	t.CreatedAt, _ = time.Parse("02/01/2006", content[3])
	t.UpdatedAt, _ = time.Parse("02/01/2006", content[4])
}
