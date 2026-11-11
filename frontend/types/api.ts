export interface ProductDocument {
  id: string;
  title: string;
  description: string;
  category: string;
  price: number;
  embedding?: number[];
}

export interface UserEvent {
  eventId: string;
  userId: string;
  action: 'view' | 'click' | 'purchase';
  productId: string;
  timestamp?: number;
}

export interface VespaSearchResponse<T> {
  root: {
    id: string;
    relevance: number;
    fields?: {
      totalCount: number;
    };
    children?: Array<{
      id: string;
      relevance: number;
      fields: T;
    }>;
  };
}
