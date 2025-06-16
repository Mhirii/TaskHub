<script lang="ts" module>
	import GalleryVerticalEndIcon from "@lucide/svelte/icons/gallery-vertical-end";

	const sampleData = {
		user: {
			name: "shadcn",
			email: "m@example.com",
			avatar: "/avatars/shadcn.jpg",
		},
		teams: [
			{
				name: "iTeam",
				logo: GalleryVerticalEndIcon,
				plan: "University",
			},
		],
	};
</script>

<script lang="ts">
	import NavMain from "./nav-main.svelte";
	import NavUser from "./nav-user.svelte";
	import TeamSwitcher from "./team-switcher.svelte";
	import * as Sidebar from "$lib/components/ui/sidebar/index.js";
	import type { ComponentProps } from "svelte";
	import type { Project } from "$lib/types";

	let {
		projects,
		user,
		ref = $bindable(null),
		collapsible = "icon",
		...restProps
	}: ComponentProps<typeof Sidebar.Root> & {
		projects: Project[];
		user: { username: string; email: string; avatar: string; roles: string[] };
	} = $props();
</script>

<Sidebar.Root {collapsible} {...restProps}>
	<Sidebar.Header>
		<TeamSwitcher teams={sampleData.teams} />
	</Sidebar.Header>
	<Sidebar.Content>
		<NavMain {projects} />
	</Sidebar.Content>
	<Sidebar.Footer>
		<NavUser {user} />
	</Sidebar.Footer>
	<Sidebar.Rail />
</Sidebar.Root>
