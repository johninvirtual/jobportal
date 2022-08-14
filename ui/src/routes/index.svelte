<script context="module" lang="ts">
	import { PUBLIC_BACKEND_URL } from '$env/static/public';
	import type { Load } from '@sveltejs/kit';

	export const load: Load = async ({ fetch }) => {
		const url = `${PUBLIC_BACKEND_URL || 'http://localhost:3000'}/job`;

		const response = await fetch(url);
		return {
			status: response.status,
			props: {
				jobs: response.ok && (await response.json())
			}
		};
	};
</script>

<script lang="ts">
	import { goto, invalidate } from '$app/navigation';

	export let jobs: any[];

	let addOpen = false;

	let formData = {
		title: '',
		description: ''
	};

	const handleSubmit = () => {
		fetch(`${PUBLIC_BACKEND_URL || 'http://localhost:3000'}/job`, {
			method: 'POST',
			body: JSON.stringify(formData),
			headers: {
				'Content-Type': 'application/json'
			}
		}).then(() => {
			addOpen = false;
			formData = {
				title: '',
				description: ''
			};

			// reload the page
			invalidate(() => true);
		});
	};
</script>

<main class="container">
	<nav>
		<ul>
			<li><h1 style="margin-bottom: 0;">Job Portal</h1></li>
		</ul>
		<ul>
			<li><a href="/">Mechanical</a></li>
			<li><a href="/">IT</a></li>
			<li><button class="outline" on:click={() => (addOpen = true)}>Add</button></li>
		</ul>
	</nav>

	<button>Apply Now</button>

	<section class="grid">
		{#each jobs as job}
			<article>
				<header>
					{job.title}
				</header>
				<p style="white-space: pre-line">
					{job.description}
				</p>
				<footer style="display: flex; justify-content: center;">
					<button style="width: 200px">Apply</button>
				</footer>
			</article>
		{/each}
	</section>

	<dialog open={addOpen}>
		<article>
			<header>
				<i class="close" on:click={() => (addOpen = false)} style="cursor: pointer;" />
				Add Job
			</header>
			<form on:submit|preventDefault={handleSubmit}>
				<label for="title">
					Title
					<input
						type="text"
						id="title"
						name="title"
						placeholder="Job title"
						bind:value={formData.title}
						required
					/>
				</label>

				<!-- Markup example 2: input is after label -->
				<label for="description">description</label>
				<textarea
					type="text"
					id="description"
					name="description"
					placeholder="Enter job description"
					bind:value={formData.description}
					required
				/>
				<small>We'll never share your email with anyone else.</small>

				<!-- Button -->
				<button type="submit">Submit</button>
			</form>
		</article>
	</dialog>
</main>

<style>
	@media (min-width: 992px) {
		.grid {
			grid-template-columns: 1fr 1fr;
		}
	}
</style>
