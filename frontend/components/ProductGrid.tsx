'use client';

import React from 'react';
import { ProductDocument } from '../types/api';

interface ProductGridProps {
  products: ProductDocument[];
  onProductClick: (productId: string) => void;
}

export const ProductGrid: React.FC<ProductGridProps> = ({ products, onProductClick }) => {
  if (!products || products.length === 0) {
    return <p style={{ color: '#666' }}>No products found.</p>;
  }

  return (
    <div
      style={{
        display: 'grid',
        gridTemplateColumns: 'repeat(auto-fill, minmax(250px, 1fr))',
        gap: '1rem',
      }}
    >
      {products.map((p) => (
        <div
          key={p.id}
          onClick={() => onProductClick(p.id)}
          style={{
            border: '1px solid #e0e0e0',
            borderRadius: '8px',
            padding: '1rem',
            cursor: 'pointer',
            backgroundColor: '#fafafa',
            transition: 'transform 0.2s, box-shadow 0.2s',
          }}
        >
          <h3 style={{ margin: '0 0 0.5rem 0', fontSize: '1.2rem' }}>{p.title}</h3>
          <p style={{ fontSize: '0.9rem', color: '#555', margin: '0 0 0.5rem 0' }}>{p.description}</p>
          <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
            <span style={{ fontWeight: 'bold', color: '#0070f3' }}>${p.price.toFixed(2)}</span>
            <span style={{ fontSize: '0.8rem', color: '#888' }}>{p.category}</span>
          </div>
        </div>
      ))}
    </div>
  );
};
