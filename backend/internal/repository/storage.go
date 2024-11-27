package repository

//
//import (
//	"github.com/loloneme/pr12/backend/internal/entities"
//	"strings"
//)
//
//type Storage struct {
//	Drinks map[int64]*entities.Drink
//}
//
//func ConnectToStorage() *Storage {
//	return &Storage{Drinks: map[int64]*entities.Drink{
//		1: {
//			ID:          1,
//			ImageURL:    "https://github.com/loloneme/images/blob/main/cappuccino.png?raw=true",
//			Name:        "Капучино",
//			Description: "Кофейный напиток итальянской кухни на основе эспрессо с добавлением молока и молочной пенки. В нём сохраняется баланс: чувствуется вкус эспрессо, но он не преобладает над вкусом молока.",
//			Compound:   strings.Join([]string{"Эспрессо", "Молоко", "Молочная пена", "Корица"}, " ") ,
//			Cold:        false,
//			Hot:         true,
//			Prices:      map[string]int{"250": 180, "350": 230, "500": 280},
//			IsFavorite:  false,
//		},
//		2: {
//			ID:          2,
//			ImageURL:    "https://github.com/loloneme/images/blob/main/latte.png?raw=true",
//			Name:        "Латте",
//			Description: "Латте — самый большой и самый молочный кофейный напиток на основе эспрессо. Подойдёт для тех, кто не любит яркий вкус кофе. Его готовят в чашку объёмом 250–300 мл и используют одну порцию эспрессо. Остальной объем напитка составляет молоко и немного молочной пены.",
//			Compound:    []string{"Эспрессо", "Молоко", "Молочная пена", "Лед (по желанию)"},
//			Cold:        true,
//			Hot:         true,
//			Prices:      map[string]int{"350": 200, "500": 300},
//			IsFavorite:  false,
//		},
//		3: {
//			ID:          3,
//			ImageURL:    "https://github.com/loloneme/images/blob/main/bumble.png?raw=true",
//			Name:        "Бамбл",
//			Description: "Все ингредиенты наливают слоями: сначала тягучий сироп, потом яркий сок и два шота эспрессо. У него яркий и сложный вкус: сначала чувствуется сладость сиропа, затем — цитрусовая кислинка и легкая кофейная горечь в конце.",
//			Compound:    []string{"Двойной Эспрессо", "Апельсиновый фреш", "Карамельный сироп", "Лед"},
//			Cold:        true,
//			Hot:         false,
//			Prices:      map[string]int{"350": 300, "500": 400},
//			IsFavorite:  true,
//		},
//		4: {
//			ID:          4,
//			ImageURL:    "https://github.com/loloneme/images/blob/main/caramel_frappe.png?raw=true",
//			Name:        "Карамельный Фраппе",
//			Description: "Фраппе - напиток на основе эспрессо, с добавлением молока или сливочной основы, льдом и сладким сиропом. Все это взбивается в блендере в результате чего получается пюреобразная масса, которая очень хорошо освежает летом. Украшается взбитыми сливками",
//			Compound:    []string{"Двойной Эспрессо", "Молоко", "Взбитые сливки", "Сливки", "Карамельный сироп", "Лед"},
//			Cold:        true,
//			Hot:         false,
//			Prices:      map[string]int{"500": 450},
//			IsFavorite:  false,
//		},
//		5: {
//			ID:          5,
//			ImageURL:    "https://github.com/loloneme/images/blob/main/matcha_latte.png?raw=true",
//			Name:        "Матча-Латте",
//			Description: "Это напиток, приготовленный из японского порошкового чая под названием матча, или маття. В чай добавляется молоко, как в латте, и сверху молочная пена. На вкус напоминает сладкий зеленый чай со сливками. Можно приготовить как горячим, так и холодным.",
//			Compound:    []string{"Матча", "Вода", "Молоко", "Лед (по желанию)"},
//			Cold:        true,
//			Hot:         true,
//			Prices:      map[string]int{"250": 200, "350": 250, "500": 350},
//			IsFavorite:  true,
//		},
//		6: {
//			ID:          6,
//			ImageURL:    "https://github.com/loloneme/images/blob/main/lavanda_raf.png?raw=true",
//			Name:        "Лавандовый раф",
//			Description: "Популярный российский кофейный напиток, который готовят на основе эспрессо с добавлением ванильного и простого сахара, а также сливок. Раф-кофе ценят за особенный сливочно-ванильный вкус, насыщенный кофейный аромат и уникальную мягкость. Многие потребители называют «Раф» самым нежным из всех кофейных напитков. Лавандовая его разновидность отличается своим тонким сливочным вкусом с оригинальным цветочным оттенком лаванды.",
//			Compound:    []string{"Эспрессо", "10% Сливки", "Ванильный сахар", "Сахар", "Лавандовый сироп"},
//			Cold:        false,
//			Hot:         true,
//			Prices:      map[string]int{"350": 230, "500": 330},
//			IsFavorite:  false,
//		},
//		7: {
//			ID:          7,
//			ImageURL:    "https://github.com/loloneme/images/blob/main/hot_chocolate.png?raw=true",
//			Name:        "Горячий шоколад",
//			Description: "Горячий шоколад — густой напиток из горького или молочного шоколада. По сравнению с какао у него более интенсивный вкус с бархатной текстурой.",
//			Compound:    []string{"Молочный шоколад", "20% Сливки", "Молоко", "Корица", "Взбитые сливки", "Какао"},
//			Cold:        false,
//			Hot:         true,
//			Prices:      map[string]int{"500": 250},
//			IsFavorite:  false,
//		},
//		8: {
//			ID:          8,
//			ImageURL:    "https://github.com/loloneme/images/blob/main/glintwein.png?raw=true",
//			Name:        "Б/А глинтвейн",
//			Description: "Глинтвейн — подогретое до определенной температуры вино, соединенное с пряностями, специями и фруктами. В зависимости от конкретной рецептуры, напиток может успокаивать или, наоборот, тонизировать. Его часто используют как средство профилактики простудных заболеваний",
//			Compound:    []string{"Вода", "Сахар", "Апельсин", "Виноградный сок", "Корица", "Специи", "Яблоко"},
//			Cold:        false,
//			Hot:         true,
//			Prices:      map[string]int{"350": 270, "500": 370},
//			IsFavorite:  false,
//		},
//		9: {
//			ID:          9,
//			ImageURL:    "https://github.com/loloneme/images/blob/main/milkshake.png?raw=true",
//			Name:        "Милкшейк",
//			Description: "Шоколадный милкшейк — это густой, освежающий напиток, приготовленный из молока, мороженого и шоколадного сиропа. Его текстура кремовая и воздушная, а вкус — насыщенный, бархатисто-шоколадный с приятной сладостью. Каждый глоток дарит ощущение десертного удовольствия, оставляя на языке мягкое, сладкое шоколадное послевкусие. Украшается взбитыми сливками",
//			Compound:    []string{"Шоколадное мороженое", "Молоко", "Шоколадный сироп", "Взбитые сливки"},
//			Cold:        true,
//			Hot:         false,
//			Prices:      map[string]int{"350": 260, "500": 330},
//			IsFavorite:  false,
//		},
//	},
//	}
//}
