<script lang="ts">
	import {
		Card,
		CardHeader,
		CardTitle,
		CardContent,
	} from "$lib/components/ui/card";
	import type { Task } from "$lib/types";
	import TaskCard from "./task-card.svelte";

	export let tasks: Task[];

	const planned = tasks?.filter((task) => task.board_id === 0) || [];
	const inProgress = tasks?.filter((task) => task.board_id === 1) || [];
	const done = tasks?.filter((task) => task.board_id === 2) || [];
</script>

<div class="grid gap-4 md:grid-cols-3">
	<Card>
		<CardHeader>
			<CardTitle class="text-lg">Planned</CardTitle>
		</CardHeader>
		<CardContent>
			{#if planned.length === 0}
				<p class="text-muted-foreground text-sm">No tasks</p>
			{:else}
				{#each planned as task}
					<TaskCard {task} />
				{/each}
			{/if}
		</CardContent>
	</Card>
	<Card>
		<CardHeader>
			<CardTitle class="text-lg">In Progress</CardTitle>
		</CardHeader>
		<CardContent>
			{#if inProgress.length === 0}
				<p class="text-muted-foreground text-sm">No tasks</p>
			{:else}
				{#each inProgress as task}
					<TaskCard {task} />
				{/each}
			{/if}
		</CardContent>
	</Card>
	<Card>
		<CardHeader>
			<CardTitle class="text-lg">Done</CardTitle>
		</CardHeader>
		<CardContent>
			{#if done.length === 0}
				<p class="text-muted-foreground text-sm">No tasks</p>
			{:else}
				{#each done as task}
					<TaskCard {task} />
				{/each}
			{/if}
		</CardContent>
	</Card>
</div>
