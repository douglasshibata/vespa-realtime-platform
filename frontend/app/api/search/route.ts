import { NextRequest, NextResponse } from 'next/server';
import { ProductDocument, VespaSearchResponse } from '../../../types/api';

const VESPA_SEARCH_URL = process.env.VESPA_SEARCH_URL || 'http://localhost:8080/search/';

export async function GET(request: NextRequest) {
  const { searchParams } = new URL(request.url);
  const query = searchParams.get('q') || '';
  const vectorStr = searchParams.get('vector');

  try {
    const vespaParams = new URLSearchParams();
    vespaParams.append('ranking', 'hybrid_ranking');

    if (vectorStr) {
      const vector = JSON.parse(vectorStr);
      vespaParams.append(
        'yql',
        `select * from product where ([{"targetNumHits": 10}]nearestNeighbor(embedding, user_embedding)) or userQuery();`
      );
      vespaParams.append('input.query(user_embedding)', JSON.stringify(vector));
    } else {
      vespaParams.append('yql', 'select * from product where userQuery();');
    }

    vespaParams.append('query', query);

    const res = await fetch(`${VESPA_SEARCH_URL}?${vespaParams.toString()}`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
      next: { revalidate: 10 },
    });

    if (!res.ok) {
      throw new Error(`Vespa search returned HTTP status ${res.status}`);
    }

    const data: VespaSearchResponse<ProductDocument> = await res.json();
    return NextResponse.json(data);
  } catch (error: any) {
    return NextResponse.json(
      { error: 'Failed to query Vespa search engine', details: error.message },
      { status: 500 }
    );
  }
}
