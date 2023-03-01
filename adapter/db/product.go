package db

import (
	"database/sql"	
	"github.com/anlopes123/hexagonal/application"	 
)

type ProductDb struct {
	db *sql.DB
}

func NewProductDb(db *sql.DB) (*ProductDb) {
	return &ProductDb{db: db}
}

func (p *ProductDb) Get(id string) (application.ProductInterface, error) {
	var product application.Product
	stmt, err := p.db.Prepare("Select id, name, price, status from products where id = ?");
	if err != nil {
		return nil, err
	}
	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func(p *ProductDb) Save(product application.ProductInterface) (application.ProductInterface, error) {
	var rows int
	p.db.QueryRow("Select id from products where id=?", product.GetID()).Scan(&rows)
	if rows == 0 {
		product, err := p.create(product)
		if err != nil {
			return nil, err
		}
		return product, nil
	} else {
		product, err := p.update(product)
		if err != nil {
			return nil, err
		}
		return product, nil
	}
}

func(p *ProductDb) create(product application.ProductInterface) (application.ProductInterface, error) {
	stmt, err := p.db.Prepare(`INSERT INTO products(id, name, price, status) values(?, ?, ?, ?)`)
	if err != nil {
		return nil, err;
	}
	_, err = stmt.Exec(
		product.GetID(), 
		product.GetName(),
		product.GetPrice(), 
		product.GetStatus(),
	)
	if err != nil {
		return nil, err
	}
	err = stmt.Close()
	if err != nil {
		return nil, err
	}
	return product, nil
}

func(p *ProductDb) update(product application.ProductInterface) (application.ProductInterface, error) {
	_, err := p.db.Exec(`update products set name=?, price=?, status=? where id=?`, 
		product.GetName(), product.GetPrice(), product.GetStatus(), product.GetID())
	
	if err != nil {
		return nil, err
	}
	return product, nil
} 