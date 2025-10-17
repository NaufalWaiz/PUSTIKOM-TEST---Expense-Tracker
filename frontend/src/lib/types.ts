export interface Expense {
	id: number;
	amount: number;
	description: string;
	category: string;
	created_at: string;
}

export interface ExpensePayload {
	amount: number;
	description: string;
	category: string;
}
