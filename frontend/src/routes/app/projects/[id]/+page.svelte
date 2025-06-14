<script lang="ts">
	import { page } from "$app/state";
	import type { PageData } from "./$types";
	import ProjectCardSkeleton from "./project-card-skeleton.svelte";
	import ProjectColumnsSkeleton from "./project-column-skeleton.svelte";
	import NoTasksFound from "./no-tasks-found.svelte";
	import ProjectColumns from "./project-column.svelte";
	import type { Project, Task } from "$lib/types";
	import {
		Card,
		CardHeader,
		CardTitle,
		CardDescription,
	} from "$lib/components/ui/card";

	let params = page.params;
	const id = params.id;

	export let data: PageData;

	let projectLoading = true;
	export let project: Project = {
		id: 0,
		name: "",
		description: "",
		ownerID: 0,
		task_count: 0,
		open_task_count: 0,
		members_count: 0,
		created_at: "",
		updated_at: "",
	};

	data.project.then((p) => {
		console.log(p);
		projectLoading = false;
		project = p;
	});

	let tasksLoading = true;
	export let tasks: Task[];
	data.tasks.then((t) => {
		console.log(t);
		tasksLoading = false;
		tasks = t;
	});
</script>

{#if id}
	<h1>Project {id}</h1>
{:else}
	<h1>Project</h1>
{/if}

<div class="flex justify-center">
	<div class="w-full xl:max-w-[1200px] py-10">
		{#if projectLoading}
			<ProjectCardSkeleton />
			<ProjectColumnsSkeleton />
		{:else}
			<Card class="mb-8">
				<CardHeader>
					<CardTitle class="text-3xl font-bold">{project.name}</CardTitle>
					<CardDescription>{project.description}</CardDescription>
				</CardHeader>
			</Card>
			{#if tasksLoading}
				<div></div>
			{:else if tasks?.length === 0}
				<NoTasksFound />
			{:else}
				<ProjectColumns {tasks} />
			{/if}
		{/if}
	</div>
</div>
