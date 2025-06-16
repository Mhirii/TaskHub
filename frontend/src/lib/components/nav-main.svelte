<script lang="ts">
	import * as Collapsible from "$lib/components/ui/collapsible/index.js";
	import * as Sidebar from "$lib/components/ui/sidebar/index.js";
	import type { Project } from "$lib/types";
	import { NotepadTextIcon, UsersIcon } from "@lucide/svelte";
	import ChevronRightIcon from "@lucide/svelte/icons/chevron-right";

	let { projects }: { projects: Project[] } = $props();

	type MenuItem = {
		title: string;
		url: string;
		icon?: any;
		isActive: boolean;
		items?: Array<{
			title: string;
			url: string;
		}>;
	};

	const items: MenuItem[] = [
		{
			title: "Projects",
			url: "/app/projects",
			icon: NotepadTextIcon,
			isActive: false,
			items:
				projects?.length > 0
					? projects.map((p) => ({
							title: p.name,
							url: `/app/projects/${p.id}`,
						}))
					: undefined,
		},
		{
			title: "Users",
			url: "/app/users",
			icon: UsersIcon,
			isActive: false,
		},
	];
</script>

<Sidebar.Group>
	<Sidebar.GroupLabel>Platform</Sidebar.GroupLabel>
	<Sidebar.Menu>
		{#if !!items && items.length > 0}
			{#each items as item (item.title)}
				<Collapsible.Root open={item.isActive} class="group/collapsible">
					{#snippet child({ props })}
						<Sidebar.MenuItem {...props}>
							{#if !!item.items}
								<Collapsible.Trigger>
									{#snippet child({ props })}
										<Sidebar.MenuButton {...props} tooltipContent={item.title}>
											{#if item.icon}
												<a href={item.url}>
													<item.icon class=" h-5 w-5" />
												</a>
											{/if}
											<a href={item.url}><span>{item.title}</span></a>
											<ChevronRightIcon
												class="ml-auto transition-transform duration-200 group-data-[state=open]/collapsible:rotate-90"
											/>
										</Sidebar.MenuButton>
									{/snippet}
								</Collapsible.Trigger>
							{:else}
								<Sidebar.MenuButton {...props} class="w-full h-full">
									{#snippet child({ props })}
										<a href={item.url}>
											<Sidebar.MenuButton
												{...props}
												tooltipContent={item.title}
												class="w-full h-full"
											>
												{#if item.icon}
													<item.icon class=" h-5 w-5" />
												{/if}
												<a href={item.url}><span>{item.title}</span></a>
											</Sidebar.MenuButton>
										</a>
									{/snippet}
								</Sidebar.MenuButton>
							{/if}
							<Collapsible.Content>
								<Sidebar.MenuSub>
									{#if !!item.items && item.items.length > 0}
										{#each item.items ?? [] as subItem (subItem.title)}
											<Sidebar.MenuSubItem>
												<Sidebar.MenuSubButton>
													{#snippet child({ props })}
														<a href={subItem.url} {...props}>
															<span>{subItem.title}</span>
														</a>
													{/snippet}
												</Sidebar.MenuSubButton>
											</Sidebar.MenuSubItem>
										{/each}
									{/if}
								</Sidebar.MenuSub>
							</Collapsible.Content>
						</Sidebar.MenuItem>
					{/snippet}
				</Collapsible.Root>
			{/each}
		{/if}
	</Sidebar.Menu>
</Sidebar.Group>
