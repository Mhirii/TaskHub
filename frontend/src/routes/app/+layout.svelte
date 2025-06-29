<script lang="ts">
	import AppSidebar from "$lib/components/app-sidebar.svelte";
	import * as Breadcrumb from "$lib/components/ui/breadcrumb/index.js";
	import { Separator } from "$lib/components/ui/separator/index.js";
	import * as Sidebar from "$lib/components/ui/sidebar/index.js";
	import { page } from "$app/state";
	import { type PageData } from "./$types";
	import type { Project } from "$lib/types";
	import type { Snippet } from "svelte";

	let {
		children,
		data,
	}: {
		children: Snippet<[]>;
		data: PageData & { projects: Promise<Project[]> };
	} = $props();

	const buildBreads = () => {
		const paths = (page.url.pathname || "/").split("/");
		return paths.slice(1).map((path, i) => {
			const isLast = i === paths.length - 2;
			return {
				href: paths.slice(0, i + 2).join("/"),
				name: path,
				isLast,
			};
		});
	};
	let breads = $derived(buildBreads());
</script>

<Sidebar.Provider>
	{#await data.projects}
		<span></span>
	{:then projects}
		<AppSidebar {projects} user={data.user} />
	{/await}
	<Sidebar.Inset>
		<header
			class="group-has-data-[collapsible=icon]/sidebar-wrapper:h-12 flex h-16 shrink-0 items-center gap-2 transition-[width,height] ease-linear"
		>
			<div class="flex items-center gap-2 px-4">
				<Sidebar.Trigger class="-ml-1" />
				<Separator
					orientation="vertical"
					class="mr-2 data-[orientation=vertical]:h-4"
				/>
				<Breadcrumb.Root>
					<Breadcrumb.List>
						{#each breads as { href, name, isLast }}
							<Breadcrumb.Item class="hidden md:block">
								<Breadcrumb.Link {href} class="capitalize">
									{name}
								</Breadcrumb.Link>
							</Breadcrumb.Item>
							{#if !isLast}
								<Breadcrumb.Separator class="hidden md:block" />
							{/if}
						{/each}
					</Breadcrumb.List>
				</Breadcrumb.Root>
			</div>
		</header>
		<div class="flex flex-1 flex-col gap-4 p-4 lg:px-12 pt-0">
			{@render children()}
		</div>
	</Sidebar.Inset>
</Sidebar.Provider>
