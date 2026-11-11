'use client';

import React, { useState } from 'react';
import { SearchBar } from '../components/SearchBar';
import { ProductGrid } from '../components/ProductGrid';
import { ProductDocument, UserEvent } from '../types/api';

export default function Home() {
  const [products, setProducts] = useState<ProductDocument[]>([]);
  const [loading, setLoading] = useState(false);

  const handleSearch = async (query: string) => {
    setLoading(true);
    try {
      const res = await fetch(`/api/search?q=${encodeURIComponent(query)}`);
      if (res.ok) {
        const data = await res.json();
        const items = data.root?.children?.map((child: any) => child.fields) || [];
        setProducts(items);
      }
    } catch (err) {
      console.error('Search error:', err);
    } finally {
      setLoading(false);
    }
  };

  const handleProductClick = async (productId: string) => {
    const event: UserEvent = {
      eventId: `evt-${Date.now()}`,
      userId: 'usr-demo',
      action: 'click',
      productId,
      timestamp: Math.floor(Date.now() / 1000),
    };

    try {
      await fetch('/api/events', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(event),
      });
    } catch (err) {
      console.error('Event dispatch error:', err);
    }
  };

  return (
    <main style={{ maxWidth: '1000px', margin: '2rem auto', padding: '0 1rem', fontFamily: 'sans-serif' }}>
      <h1>Real-Time AI Search & Recommendations</h1>
      <p style={{ color: '#666', marginBottom: '1.5rem' }}>
        Powered by Golang, Apache Kafka, Vespa.ai, and Next.js
      </p>
      <SearchBar onSearch={handleSearch} />
      {loading ? <p>Loading results...</p> : <ProductGrid products={products} onProductClick={handleProductClick} />}
    </main>
  );
}
