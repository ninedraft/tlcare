package core

import (
	"time"

	"github.com/asdine/storm/q"
	"github.com/ninedraft/tlcare/pkg/models/advicemodel"
)

func (care *Care) GetAdvice() (advicemodel.Advice, error) {
	tx, err := care.db.Begin(true)
	if err != nil {
		return advicemodel.Advice{}, err
	}
	defer tx.Rollback()
	var advice advicemodel.Advice
	var now = time.Now()
	query := tx.Select(q.Lt("last_retrieved", now.AddDate(0, 0, -7)))
	if err := query.First(&advice); err != nil {
		return advicemodel.Advice{}, err
	}
	advice.LastRetrieved = time.Now()
	if err := tx.UpdateField(advice, "last_retrieved", now); err != nil {
		return advicemodel.Advice{}, err
	}
	return advice, tx.Commit()
}

func (care *Care) Like(adviceId string) (uint64, error) {
	tx, err := care.db.Begin(true)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()
	var advice advicemodel.Advice
	if err := tx.One("id", adviceId, &advice); err != nil {
		return 0, err
	}
	if err := tx.UpdateField(advicemodel.Advice{ID: adviceId}, "likes", advice.Likes+1); err != nil {
		return 0, err
	}
	return 0, tx.Commit()
}

func (care *Care) TopAdvices(n uint) ([]advicemodel.Advice, error) {
	tx, err := care.db.Begin(true)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	var advices = make([]advicemodel.Advice, 0, 32)
	if err := tx.Select().Limit(int(n)).OrderBy("likes").Find(&advices); err != nil {
		return nil, err
	}
	return nil, tx.Commit()
}

func (care *Care) AddAdvice(txt string) (advicemodel.Advice, error) {
	advice := advicemodel.FromText(txt)
	tx, err := care.db.Begin(true)
	if err != nil {
		return advicemodel.Advice{}, err
	}
	defer tx.Rollback()
	if err := tx.Save(advice); err != nil {
		return advicemodel.Advice{}, err
	}
	return advicemodel.Advice{}, tx.Commit()
}
