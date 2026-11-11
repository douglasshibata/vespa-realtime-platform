import { NextRequest, NextResponse } from 'next/server';
import { UserEvent } from '../../../types/api';

const GO_BACKEND_URL = process.env.GO_BACKEND_URL || 'http://localhost:8080/api/v1/events';

export async function POST(request: NextRequest) {
  try {
    const body: UserEvent = await request.json();

    const res = await fetch(GO_BACKEND_URL, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(body),
    });

    if (!res.ok) {
      throw new Error(`Backend returned status ${res.status}`);
    }

    const data = await res.json();
    return NextResponse.json(data);
  } catch (error: any) {
    return NextResponse.json(
      { error: 'Failed to dispatch event to Go backend', details: error.message },
      { status: 500 }
    );
  }
}
