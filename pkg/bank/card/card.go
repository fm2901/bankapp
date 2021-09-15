package card

import "bank/pkg/bank/types"

func PaymentSources(cards []types.Card) []types.PaymentSource {
	var result []types.PaymentSource

	for _, card := range cards {
		
		if !card.Active {
			continue
		}

		if card.Balance < 1 {
			continue
		}

		result = append(result, card)
	}

	return result
}