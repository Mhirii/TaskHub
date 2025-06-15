<script lang="ts">
	import * as Collapsible from "$lib/components/ui/collapsible/index.js";
	import * as Sidebar from "$lib/components/ui/sidebar/index.js";
	import type { Project } from "$lib/types";
	import { FrameIcon, UsersIcon } from "@lucide/svelte";
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
			icon: FrameIcon,
			isActive: false,
			items: projects.map((p) => ({
				title: p.name,
				url: `/app/projects/${p.id}`,
			})),
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
		{#each items as item (item.title)}
			<Collapsible.Root open={item.isActive} class="group/collapsible">
				{#snippet child({ props })}
					<Sidebar.MenuItem {...props}>
						<Collapsible.Trigger>
							{#snippet child({ props })}
								<Sidebar.MenuButton {...props} tooltipContent={item.title}>
									{#if item.icon}
										<item.icon />
									{/if}
									<a href={item.url}><span>{item.title}</span></a>
									<ChevronRightIcon
										class="ml-auto transition-transform duration-200 group-data-[state=open]/collapsible:rotate-90"
									/>
								</Sidebar.MenuButton>
							{/snippet}
						</Collapsible.Trigger>
						<Collapsible.Content>
							<Sidebar.MenuSub>
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
							</Sidebar.MenuSub>
						</Collapsible.Content>
					</Sidebar.MenuItem>
				{/snippet}
			</Collapsible.Root>
		{/each}
	</Sidebar.Menu>
</Sidebar.Group>
