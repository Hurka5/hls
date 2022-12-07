package main

func filterItems(items []Item) []Item {

	var filtered []Item

	for _, item := range items {

		if !args.All && item.isHidden() {
			continue
		}

		filtered = append(filtered, item)
	}

	return filtered
}
