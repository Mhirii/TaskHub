<script lang="ts">
	import type { PageData } from "./$types";
	import ProjectsCardSkeleton from "./projects-card-skeleton.svelte";
	import ProjectsTableSkeleton from "./projects-table-skeleton.svelte";
	import NoProjectsFound from "./no-projects-found.svelte";
	import ProjectsCard from "./projects-card.svelte";
	import ProjectsTable from "./projects-table.svelte";
	import {
		Card,
		CardHeader,
		CardTitle,
		CardDescription,
	} from "$lib/components/ui/card";
	import type { Project } from "$lib/types";
	import { AlertCircleIcon, Plus } from "@lucide/svelte";
	import * as Alert from "$lib/components/ui/alert/index.js";

	let isLoading = true;
	let projects: Project[] = [];

	const skeletonRows = Array(5).fill(null);

	export let data: PageData;
	console.log(data);
	data.projects.then((p) => {
		console.log(p);
		projects = p;
		isLoading = false;
	});
</script>

<div class="flex justify-center">
	<div class="w-full xl:max-w-[1200px] py-10">
		<Card class="mb-8">
			<CardHeader class="flex flex-row items-center justify-between">
				<div>
					<CardTitle class="text-3xl font-bold">Projects</CardTitle>
					<CardDescription>
						Manage your projects and track progress
					</CardDescription>
				</div>
				<a
					href="/app/projects/new"
					class="inline-flex items-center justify-center rounded-md bg-primary px-4 py-2 text-sm font-medium text-primary-foreground hover:bg-primary/90 focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:pointer-events-none disabled:opacity-50"
				>
					<Plus />
					New Project
				</a>
			</CardHeader>
		</Card>
		{#if isLoading}
			<!-- Mobile view: Skeleton Cards -->
			<div class="lg:hidden space-y-4">
				{#each skeletonRows as _}
					<ProjectsCardSkeleton {data} />
				{/each}
			</div>

			<!-- Desktop view: Skeleton Table -->
			<div class="hidden lg:block">
				<ProjectsTableSkeleton />
			</div>
		{:else if projects?.length === 0}
			<NoProjectsFound />
		{:else if projects?.length > 0}
			<div class="space-y-8">
				<!-- Mobile view: Cards -->
				<div class="lg:hidden space-y-4">
					{#each projects as project}
						<ProjectsCard {project} />
					{/each}
				</div>

				<!-- Desktop view: Table -->
				<div class="hidden lg:block w-full">
					<ProjectsTable {projects} />
				</div>
			</div>
		{:else}
			<Alert.Root variant="destructive">
				<AlertCircleIcon />
				<Alert.Title>You have no projects.</Alert.Title>
				<Alert.Description>
					<p>
						Please contact your administrator if you think this is an issue.
					</p>
				</Alert.Description>
			</Alert.Root>
		{/if}
	</div>
</div>
