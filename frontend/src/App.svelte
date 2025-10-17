<script lang="ts">
	import { onMount } from 'svelte';
	import { createExpense, deleteExpense, fetchExpenses } from './lib/api';
	import type { Expense } from './lib/types';

	const CATEGORIES = ['Food', 'Transport', 'Shopping', 'Other'];

	let expenses: Expense[] = [];
	let loading = false;
	let submitting = false;
	let errorMessage = '';
	let filter = 'All';
	let form = {
		amount: '',
		description: '',
		category: CATEGORIES[0]
	};
	let formErrors = {
		amount: '',
		description: '',
		category: ''
	};

	const currencyFormatter = new Intl.NumberFormat('id-ID', {
		style: 'currency',
		currency: 'IDR',
		minimumFractionDigits: 0
	});

	onMount(() => {
		loadExpenses();
	});

	async function loadExpenses() {
		loading = true;
		errorMessage = '';
		try {
			expenses = await fetchExpenses(filter);
		} catch (error) {
			errorMessage = error instanceof Error ? error.message : 'Gagal memuat data';
		} finally {
			loading = false;
		}
	}

	function validateForm() {
		formErrors = {
			amount: '',
			description: '',
			category: ''
		};
		let valid = true;
		const amountValue = Number(form.amount);
		if (!form.amount || Number.isNaN(amountValue) || amountValue <= 0) {
			formErrors.amount = 'Masukkan nominal yang valid (> 0)';
			valid = false;
		}
		if (!form.description.trim() || form.description.trim().length < 3) {
			formErrors.description = 'Deskripsi minimal 3 karakter';
			valid = false;
		}
		if (!form.category) {
			formErrors.category = 'Pilih kategori';
			valid = false;
		}
		return valid;
	}

	async function handleSubmit(event: Event) {
		event.preventDefault();
		if (!validateForm()) return;

		submitting = true;
		errorMessage = '';
		try {
			await createExpense({
				amount: Number(form.amount),
				description: form.description.trim(),
				category: form.category
			});
			form = {
				amount: '',
				description: '',
				category: form.category
			};
			await loadExpenses();
		} catch (error) {
			errorMessage = error instanceof Error ? error.message : 'Gagal menambah expense';
		} finally {
			submitting = false;
		}
	}

	async function handleDelete(id: number) {
		if (!confirm('Hapus expense ini?')) return;
		errorMessage = '';
		try {
			await deleteExpense(id);
			await loadExpenses();
		} catch (error) {
			errorMessage = error instanceof Error ? error.message : 'Gagal menghapus expense';
		}
	}

	function formatAmount(value: number) {
		return currencyFormatter.format(value);
	}

	function formatDate(value: string) {
		const date = new Date(value);
		return Number.isNaN(date.getTime()) ? '' : date.toLocaleString('id-ID');
	}

	async function handleFilterChange(event: Event) {
		const value = (event.target as HTMLSelectElement).value;
		filter = value;
		await loadExpenses();
	}

	$: filteredExpenses =
		filter === 'All' ? expenses : expenses.filter((expense) => expense.category === filter);
	$: totalAmount = filteredExpenses.reduce((sum, expense) => sum + expense.amount, 0);
	$: if (!CATEGORIES.includes(form.category)) {
		form.category = CATEGORIES[0];
	}
</script>

<main class="container">
	<header>
		<h1>PUSTIKOM - Expense Tracker</h1>
		<p>Aplikasi pencatatan pengeluaran! Jangan lupa track Pengeluaran Harianmu 😉</p>
	</header>

	{#if errorMessage}
		<div class="alert alert-error">
			<span>{errorMessage}</span>
		</div>
	{/if}

	<section class="card">
		<h2>Tambah Pengeluaran</h2>
		<form on:submit|preventDefault={handleSubmit} novalidate>
			<div class="field">
				<label for="amount">Nominal</label>
				<input
					id="amount"
					type="number"
					min="1"
					step="100"
					placeholder="Masukkan nominal"
					bind:value={form.amount}
				/>
				{#if formErrors.amount}
					<p class="error">{formErrors.amount}</p>
				{/if}
			</div>

			<div class="field">
				<label for="description">Deskripsi</label>
				<input
					id="description"
					type="text"
					placeholder="Misal: Makan siang"
					bind:value={form.description}
				/>
				{#if formErrors.description}
					<p class="error">{formErrors.description}</p>
				{/if}
			</div>

			<div class="field">
				<label for="category">Kategori</label>
				<select id="category" bind:value={form.category}>
					{#each CATEGORIES as category}
						<option value={category}>{category}</option>
					{/each}
				</select>
				{#if formErrors.category}
					<p class="error">{formErrors.category}</p>
				{/if}
			</div> 

			<button type="submit" class="btn" disabled={submitting}>
				{submitting ? 'Menyimpan...' : 'Simpan'}
			</button>
		</form>
	</section>

	<section class="card">
		<div class="controls">
			<div class="control">
				<label for="filter">Filter Kategori</label>
				<select id="filter" on:change={handleFilterChange} bind:value={filter}>
					<option value="All">All</option>
					{#each CATEGORIES as category}
						<option value={category}>{category}</option>
					{/each}
				</select>
			</div>
			<div class="control total">
				<span>Total</span>
				<strong>{formatAmount(totalAmount)}</strong>
			</div>
		</div>

		{#if loading}
			<p>Memuat data...</p>
		{:else if filteredExpenses.length === 0}
			<p>Belum ada pengeluaran pada kategori ini.</p>
		{:else}
			<ul class="expense-list">
				{#each filteredExpenses as expense (expense.id)}
					<li class="expense-item">
						<div class="expense-main">
							<h3>{expense.description}</h3>
							<p class="meta">
								<span>{expense.category}</span>
								<span>{formatDate(expense.created_at)}</span>
							</p>
						</div>
						<div class="expense-actions">
							<strong>{formatAmount(expense.amount)}</strong>
							<button class="btn-delete" on:click={() => handleDelete(expense.id)}>Hapus</button>
						</div>
					</li>
				{/each}
			</ul>
		{/if}
	</section>
</main>

<style>
	:global(body) {
		background: #f1f5f9;
	}

	.container {
		max-width: 960px;
		margin: 0 auto;
		padding: 2.5rem 1.5rem 4rem;
		display: flex;
		flex-direction: column;
		gap: 1.5rem;
	}

	header h1 {
		margin: 0;
		font-size: 2.5rem;
		color: #111827;
	}

	header p {
		font-size: 20px;
		margin: 0.2rem 0 0;
		color: #4b5563;
	}

	.card {
		background: #fff;
		border-radius: 1rem;
		padding: 1.75rem;
		box-shadow: 0 10px 30px rgba(15, 23, 42, 0.1);
	}

	h2 {
		margin: 0 0 1.25rem;
		color: #1f2937;
	}

	form {
		display: grid;
		grid-template-columns: repeat(auto-fit, minmax(220px, 1fr));
		gap: 1rem 1.25rem;
		align-items: end;
	}

	.field {
		display: flex;
		flex-direction: column;
		gap: 0.45rem;
	}

	label {
		font-weight: 600;
		color: #1f2937;
	}

	input,
	select {
		padding: 0.65rem 0.75rem;
		border: 1px solid #d1d5db;
		border-radius: 0.65rem;
		font-size: 0.95rem;
		transition: border-color 0.2s, box-shadow 0.2s;
	}

	input:focus,
	select:focus {
		outline: none;
		border-color: #2563eb;
		box-shadow: 0 0 0 3px rgba(37, 99, 235, 0.15);
	}

	.error {
		margin: 0;
		font-size: 0.8rem;
		color: #dc2626;
	}

	.btn {
		background: #2563eb;
		color: #fff;
		padding: 0.85rem 1.5rem;
		border: none;
		border-radius: 0.75rem;
		font-weight: 600;
		transition: background 0.2s ease;
	}

	.btn:disabled {
		background: #93c5fd;
		cursor: not-allowed;
	}

	.btn:not(:disabled):hover {
		background: #1d4ed8;
	}

	.controls {
		display: flex;
		flex-wrap: wrap;
		gap: 1rem;
		align-items: center;
		justify-content: space-between;
		margin-bottom: 1.5rem;
	}

	.control {
		display: flex;
		flex-direction: column;
		gap: 0.35rem;
		color: #1f2937;
	}

	.total strong {
		font-size: 1.4rem;
		color: #16a34a;
	}

	.expense-list {
		list-style: none;
		padding: 0;
		margin: 0;
		display: flex;
		flex-direction: column;
		gap: 0.75rem;
	}

	.expense-item {
		display: flex;
		justify-content: space-between;
		align-items: center;
		padding: 1rem 1.25rem;
		border-radius: 0.85rem;
		background: #f8fafc;
	}

	.expense-main h3 {
		margin: 0 0 0.25rem;
		font-size: 1.1rem;
		color: #0f172a;
	}

	.meta {
		margin: 0;
		color: #6b7280;
		display: flex;
		gap: 1rem;
		font-size: 0.85rem;
	}

	.expense-actions {
		display: flex;
		align-items: center;
		gap: 1rem;
	}

	.btn-delete {
		background: #ef4444;
		color: #fff;
		border: none;
		padding: 0.5rem 1rem;
		border-radius: 0.65rem;
		font-weight: 600;
		transition: background 0.2s ease;
	}

	.btn-delete:hover {
		background: #dc2626;
	}

	.alert {
		padding: 0.85rem 1.25rem;
		border-radius: 0.75rem;
		font-weight: 500;
	}

	.alert-error {
		background: #fee2e2;
		color: #b91c1c;
		border: 1px solid #fecaca;
	}

	@media (max-width: 640px) {
		form {
			grid-template-columns: 1fr;
		}

		.expense-item {
			flex-direction: column;
			align-items: flex-start;
			gap: 0.75rem;
		}

		.expense-actions {
			width: 100%;
			justify-content: space-between;
		}
	}
</style>
