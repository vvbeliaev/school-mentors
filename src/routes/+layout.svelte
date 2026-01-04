<script lang="ts">
	import '$lib/shared/pb/pb-hook';
	import favicon from '$lib/shared/assets/favicon.svg';
	import { ThemeLoad } from '$lib';
	import './layout.css';

	let { children } = $props();
</script>

<svelte:head
	><link rel="icon" href={favicon} />
	<ThemeLoad />
</svelte:head>
{@render children()}
