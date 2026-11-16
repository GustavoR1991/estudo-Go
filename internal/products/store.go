package products

import "sync"

type Product struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Category    string   `json:"category"`
	Price       float64  `json:"price"`
	OldPrice    *float64 `json:"oldPrice,omitempty"`
	Image       string   `json:"image"`
	Stock       int      `json:"stock"`
	Featured    bool     `json:"featured"`
}

var (
	mu       = sync.RWMutex{}
	products = []Product{
		{ID: 1, Name: "Fone Wave Pro", Description: "Cancelamento de ruído e áudio imersivo para o seu dia.", Category: "Áudio", Price: 429.90, OldPrice: floatPtr(499.90), Image: "https://images.unsplash.com/photo-1505740420928-5e560c06d30e?auto=format&fit=crop&w=800&q=80", Stock: 12, Featured: true},
		{ID: 2, Name: "Tênis Urban Move", Description: "Conforto leve para acompanhar todos os seus passos.", Category: "Moda", Price: 289.90, Image: "https://images.unsplash.com/photo-1542291026-7eec264c27ff?auto=format&fit=crop&w=800&q=80", Stock: 8, Featured: true},
		{ID: 3, Name: "Mochila Nomad", Description: "Design resistente com espaço para trabalho e viagens.", Category: "Acessórios", Price: 199.90, Image: "https://images.unsplash.com/photo-1553062407-98eeb64c6a62?auto=format&fit=crop&w=800&q=80", Stock: 15, Featured: false},
		{ID: 4, Name: "Smartwatch Pulse", Description: "Monitore sua rotina com estilo e praticidade.", Category: "Tecnologia", Price: 599.90, OldPrice: floatPtr(699.90), Image: "https://images.unsplash.com/photo-1523275335684-37898b6baf30?auto=format&fit=crop&w=800&q=80", Stock: 5, Featured: true},
		{ID: 5, Name: "Câmera Instant Mini", Description: "Registre momentos especiais e revele na hora.", Category: "Tecnologia", Price: 349.90, Image: "https://images.unsplash.com/photo-1526170375885-4d8ecf77b99f?auto=format&fit=crop&w=800&q=80", Stock: 0, Featured: false},
		{ID: 6, Name: "Garrafa Thermo", Description: "Bebida na temperatura certa por muito mais tempo.", Category: "Casa", Price: 89.90, Image: "https://images.unsplash.com/photo-1602143407151-7111542de6e8?auto=format&fit=crop&w=800&q=80", Stock: 24, Featured: false},
	}
	nextID = 7
)

func floatPtr(value float64) *float64 { return &value }

func All() []Product {
	mu.RLock()
	defer mu.RUnlock()
	result := make([]Product, len(products))
	copy(result, products)
	return result
}

func Add(product Product) Product {
	mu.Lock()
	defer mu.Unlock()
	product.ID = nextID
	nextID++
	products = append(products, product)
	return product
}

func GetByID(id int) (Product, bool) {
	mu.RLock()
	defer mu.RUnlock()
	for _, product := range products {
		if product.ID == id {
			return product, true
		}
	}
	return Product{}, false
}

func UpdateByID(id int, product Product) (Product, bool) {
	mu.Lock()
	defer mu.Unlock()
	for i := range products {
		if products[i].ID == id {
			product.ID = id
			products[i] = product
			return product, true
		}
	}
	return Product{}, false
}

func DeleteByID(id int) bool {
	mu.Lock()
	defer mu.Unlock()
	for i := range products {
		if products[i].ID == id {
			products = append(products[:i], products[i+1:]...)
			return true
		}
	}
	return false
}
