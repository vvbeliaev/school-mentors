<script lang="ts">
	//@ts-ignore
	import { pwaInfo } from 'virtual:pwa-info';

	import '$lib/shared/pb/pb-hook';
	import favicon from '$lib/shared/assets/favicon_io/favicon.ico';
	import { ThemeLoad, PortalHost } from '$lib';
	import './layout.css';

	let { children } = $props();

	const webManifestLink = $derived(pwaInfo ? pwaInfo.webManifest.linkTag : '');
</script>

<svelte:head>
	{@html webManifestLink}
	<link rel="icon" href={favicon} />
	<ThemeLoad />
</svelte:head>

{@render children()}

<PortalHost />
