package repair

type Evaluator struct{}

func (e *Evaluator) Evaluate(risk RiskLevel) string {
    switch risk {
    case Low:
        return "OBSERVE"
    case Medium:
        return "SIMULATE"
    case High:
        return "BLOCK"
    case Dangerous:
        return "HUMAN_ONLY"
    default:
        return "BLOCK"
    }
}
