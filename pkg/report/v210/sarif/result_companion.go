package sarif

// NewRuleResult ...
func NewRuleResult(ruleID string) *Result {
	return NewResult().WithRuleID(ruleID)
}
