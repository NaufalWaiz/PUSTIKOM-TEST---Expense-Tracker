import type { Expense, ExpensePayload } from './types';

const API_BASE = '/api/expenses';

async function handleResponse(response: Response) {
	if (!response.ok) {
		let message = `Request failed with status ${response.status}`;
		try {
			const data = await response.json();
			if (data?.error) {
				message = data.error;
			}
		} catch {
			// ignore JSON parse errors and fall back to default message
		}
		throw new Error(message);
	}

	if (response.status === 204) {
		return null;
	}

	const text = await response.text();
	return text.length ? JSON.parse(text) : null;
}

export async function fetchExpenses(category?: string): Promise<Expense[]> {
	const url = new URL(API_BASE, window.location.origin);
	if (category && category !== 'All') {
		url.searchParams.set('category', category);
	}

	const response = await fetch(url.toString());
	return (await handleResponse(response)) as Expense[];
}

export async function createExpense(payload: ExpensePayload): Promise<void> {
	const response = await fetch(API_BASE, {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json'
		},
		body: JSON.stringify(payload)
	});

	await handleResponse(response);
}

export async function deleteExpense(id: number): Promise<void> {
	const response = await fetch(`${API_BASE}/${id}`, {
		method: 'DELETE'
	});

	await handleResponse(response);
}
