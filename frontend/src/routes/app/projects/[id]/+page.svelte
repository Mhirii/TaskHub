<script lang="ts">
	import { page } from "$app/state";
	import type { PageData } from "./$types";
	import ProjectCardSkeleton from "./project-card-skeleton.svelte";
	import ProjectColumnsSkeleton from "./project-column-skeleton.svelte";
	import NoTasksFound from "./no-tasks-found.svelte";
	import ProjectColumns from "./project-column.svelte";
	import {
		Card,
		CardHeader,
		CardTitle,
		CardDescription,
	} from "$lib/components/ui/card";
	import { Plus } from "@lucide/svelte";
	import Button from "$lib/components/ui/button/button.svelte";

	let params = page.params;
	const id = params.id;

	let projectId: number;
	try {
		projectId = Number.parseInt(id);
	} catch (e) {
		console.error(e);
		projectId = 99;
	}

	export let data: PageData;
</script>

{#if id}
	<h1>Project {id}</h1>
{:else}
	<h1>Project</h1>
{/if}

<div class="flex justify-center">
	<div class="w-full xl:max-w-[1200px] py-10">
		{#await data.project}
			<ProjectCardSkeleton />
			<ProjectColumnsSkeleton />
		{:then p}
			<Card class="mb-8">
				<CardHeader class="flex flex-row items-center justify-between">
					<div>
						<CardTitle class="text-3xl font-bold">{p.name}</CardTitle>
						<CardDescription>{p.description}</CardDescription>
					</div>
					<Button>
						<Plus />
						Add Member
					</Button>
				</CardHeader>
			</Card>
			{#await data.tasks}
				<div></div>
			{:then t}
				<ProjectColumns tasks={t} {projectId} {data} />
			{/await}
		{/await}
	</div>
</div>
