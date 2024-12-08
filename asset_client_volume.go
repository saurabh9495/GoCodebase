package main

import "fmt"

func main() {

	result := asset_client_volume(1, 1, 10)
	fmt.Println(result)
}

func asset_client_volume(asset_id int, client_id int, volume int) bool {

	assets := []map[string]interface{}{
		{
			"asset_id":             1,
			"asset_name":           "Apple Inc.",
			"asset_description":    "Technology company",
			"current_price":        125.5,
			"total_available_count": 1000000,
		},
		{
			"asset_id":             2,
			"asset_name":           "Tesla Inc.",
			"asset_description":    "Electric vehicle and clean energy company",
			"current_price":        650.75,
			"total_available_count": 500000,
		},
		{
			"asset_id":             3,
			"asset_name":           "Amazon.com Inc.",
			"asset_description":    "E-commerce and cloud energy company",
			"current_price":        3250.0,
			"total_available_count": 200000,
		},
	}

	clients := []map[string]interface{}{
		{
			"client_id":        1,
			"client_name":      "John Doe",
			"contact_details":  "john.doe@example.com",
			"account_balance":  10000.0,
		},
		{
			"client_id":        2,
			"client_name":      "Jane Smith",
			"contact_details":  "jane.smith@example.com",
			"account_balance":  15000.0,
		},
		{
			"client_id":        3,
			"client_name":      "Bob Johnson",
			"contact_details":  "bob.johnson@example.com",
			"account_balance":  20000.0,
		},
	}

	asset_exists := false
	var current_price float64

	client_exists := false
	var account_balance float64

	for _, asset := range assets {
		if asset_id == asset["asset_id"].(int) {
			asset_exists = true
			current_price = asset["current_price"].(float64)
			break
		}
	}

	for _, client := range clients {
		if client_id == client["client_id"].(int) {
			client_exists = true
			account_balance = client["account_balance"].(float64)
			break
		}
	}

	if asset_exists && client_exists {
		balance_needed := float64(volume) * current_price
		if balance_needed <= account_balance {
			return true
		}
	}

	return false
}
