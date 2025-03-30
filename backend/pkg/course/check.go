package course

import (
	"context"

	"github.com/pecet3/logger"
)

func (c Course) CheckIfIsPlace(courseID int64) bool {
	cc, err := c.d.GetCourse(context.Background(), courseID)
	if err != nil {
		logger.Error(err)
		return false
	}
	rp, err := c.d.GetNumberOfPersonsInCourse(context.Background(), courseID)
	if err != nil {
		logger.Error(err)
		return false
	}
	if cc.MaxPersons > rp {
		return true
	}
	return false
}
