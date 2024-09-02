package util

func CalculateLstStatus(lst int, quantity int) string {
	switch {
	case quantity == 0:
		return "black"
	case quantity < lst:
		return "red"
	case lst <= quantity && quantity <= lst+15:
		return "yellow"
	case quantity > lst+15:
		return "green"
	default:
		return "none"
	}
}
