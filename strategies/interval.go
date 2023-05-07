package strategies

import (
	"errors"
	"time"

	"github.com/saniales/golang-crypto-trading-bot/environment"
	"github.com/saniales/golang-crypto-trading-bot/exchanges"
)

// IntervalStrategy is an interval based strategy.
type IntervalStrategy struct {
	Model    StrategyModel
	Interval time.Duration
}

// Name returns the name of the strategy.
func (is IntervalStrategy) Name() string {
	return is.Model.Name
}

// String returns a string representation of the object.
func (is IntervalStrategy) String() string {
	return is.Name()
}

// Apply executes Cyclically the On Update, basing on provided interval.
func (is IntervalStrategy) Apply(wrappers []exchanges.ExchangeWrapper, markets []*environment.Market) {
	var err error

	hasSetupFunc := is.Model.Setup != nil
	hasTearDownFunc := is.Model.TearDown != nil
	hasUpdateFunc := is.Model.OnUpdate != nil
	hasErrorFunc := is.Model.OnError != nil

	if hasSetupFunc {
		err = is.Model.Setup(wrappers, markets)
		if err != nil && hasErrorFunc {
			is.Model.OnError(err)
		}
	}

	if !hasUpdateFunc {
		_err := errors.New("OnUpdate func cannot be empty")
		if hasErrorFunc {
			is.Model.OnError(_err)
		} else {
			panic(_err)
		}
	}
	for err == nil {
		err = is.Model.OnUpdate(wrappers, markets)
		if err != nil && hasErrorFunc {
			is.Model.OnError(err)
		}
		time.Sleep(is.Interval)
	}
	if hasTearDownFunc {
		err = is.Model.TearDown(wrappers, markets)
		if err != nil && hasErrorFunc {
			is.Model.OnError(err)
		}
	}
}
