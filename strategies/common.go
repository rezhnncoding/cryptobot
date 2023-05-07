package strategies

import (
	"fmt"
	"sync"

	"github.com/saniales/golang-crypto-trading-bot/environment"
	"github.com/saniales/golang-crypto-trading-bot/exchanges"
)

var available map[string]Strategy //mapped name -> strategy
var appliedTactics []Tactic

// Strategy represents a generic strategy.
type Strategy interface {
	Name() string                                             // Name returns the name of the strategy.
	Apply([]exchanges.ExchangeWrapper, []*environment.Market) // Apply applies the strategy when called, using the specified wrapper.
}

// StrategyFunc represents a standard function binded to a strategy model execution.
//
//	Can define a Setup, TearDown and Update behaviour.
type StrategyFunc func([]exchanges.ExchangeWrapper, []*environment.Market) error

// StrategyModel represents a strategy model used by strategies.
type StrategyModel struct {
	Name     string
	Setup    StrategyFunc
	TearDown StrategyFunc
	OnUpdate StrategyFunc
	OnError  func(error)
}

// Tactic represents the effective appliance of a strategy.
type Tactic struct {
	Markets  []*environment.Market
	Strategy Strategy
}

// Execute executes effectively a tactic.
func (t *Tactic) Execute(wrappers []exchanges.ExchangeWrapper) {
	t.Strategy.Apply(wrappers, t.Markets)
}

func init() {
	available = make(map[string]Strategy)
}

// AddCustomStrategy adds a strategy to the available set.
func AddCustomStrategy(s Strategy) {
	available[s.Name()] = s
}

// MatchWithMarkets matches a strategy with the markets.
func MatchWithMarkets(strategyName string, markets []*environment.Market) error {
	s, exists := available[strategyName]
	if !exists {
		return fmt.Errorf("Strategy %s does not exist, cannot bind to markets %v", strategyName, markets)
	}
	appliedTactics = append(appliedTactics, Tactic{
		Markets:  markets,
		Strategy: s,
	})
	return nil
}

// ApplyAllStrategies applies all matched strategies concurrently.
func ApplyAllStrategies(wrappers []exchanges.ExchangeWrapper) {
	var wg sync.WaitGroup
	wg.Add(len(appliedTactics))
	for _, t := range appliedTactics {
		go func(wrappers []exchanges.ExchangeWrapper, t Tactic, wg *sync.WaitGroup) {
			defer wg.Done()
			t.Execute(wrappers)
		}(wrappers, t, &wg)
	}
	wg.Wait()
}
