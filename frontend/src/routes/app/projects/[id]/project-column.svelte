<script lang="ts">
	import { Button } from "$lib/components/ui/button";
	import {
		Card,
		CardHeader,
		CardTitle,
		CardContent,
	} from "$lib/components/ui/card";
	import type { Task } from "$lib/types";
	import { Plus } from "@lucide/svelte";
	import TaskCard from "./task-card.svelte";
	import * as Dialog from "$lib/components/ui/dialog/index.js";
	import TaskForm from "./task-form.svelte";
	import type { Infer, SuperValidated } from "sveltekit-superforms";
	import type { TaskSchema, TaskUpdateSchema } from "./schema";
	import AddTaskButton from "./add-task-button.svelte";

	let {
		tasks,
		data,
		projectId,
	}: {
		tasks: Task[];
		data: {
			taskForm: SuperValidated<Infer<TaskSchema>>;
			taskUpdateForm: SuperValidated<Infer<TaskUpdateSchema>>;
		};
		projectId: number;
	} = $props();

	const mapTasks = (tasks: Task[], boardId: number) => {
		if (!tasks || tasks.length === 0) return [];
		let highestOrder = 0;
		const filtered = tasks.filter((task) => {
			if (task.board_id === boardId) {
				if (task.order > highestOrder) {
					highestOrder = task.order;
				}
				return true;
			} else return false;
		});
		return filtered.sort((a, b) => b.order - a.order);
	};

	const planned = mapTasks(tasks, 0);
	const inProgress = mapTasks(tasks, 1);
	const done = mapTasks(tasks, 2);
</script>

<div class="grid gap-4 lg:grid-cols-3">
	<Card>
		<CardHeader class="flex flex-row items-center justify-between">
			<CardTitle class="text-lg">Planned</CardTitle>
			<Dialog.Root>
				<Dialog.Trigger>
					<Button
						variant="outline"
						size="sm"
						class="h-8 w-8"
						onclick={() => console.log("add task")}
					>
						<Plus />
					</Button>
				</Dialog.Trigger>
				<Dialog.Content>
					<TaskForm {projectId} board={0} {data} />
				</Dialog.Content>
			</Dialog.Root>
		</CardHeader>
		<CardContent>
			{#if planned.length === 0}
				<p class="text-muted-foreground text-sm">No tasks</p>
			{:else}
				{#each planned as task}
					<TaskCard {task} {projectId} boardId={0} {data} />
				{/each}
			{/if}
		</CardContent>
	</Card>
	<Card>
		<CardHeader class="flex flex-row items-center justify-between">
			<CardTitle class="text-lg">In Progress</CardTitle>
			<AddTaskButton {projectId} {data} boardId={1} />
		</CardHeader>
		<CardContent>
			{#if inProgress.length === 0}
				<p class="text-muted-foreground text-sm">No tasks</p>
			{:else}
				{#each inProgress as task}
					<TaskCard {task} {projectId} boardId={1} {data} />
				{/each}
			{/if}
		</CardContent>
	</Card>
	<Card>
		<CardHeader class="flex flex-row items-center justify-between">
			<CardTitle class="text-lg">Done</CardTitle>
			<AddTaskButton {projectId} {data} boardId={2} />
		</CardHeader>
		<CardContent>
			{#if done.length === 0}
				<p class="text-muted-foreground text-sm">No tasks</p>
			{:else}
				{#each done as task}
					<TaskCard {task} {projectId} boardId={2} {data} />
				{/each}
			{/if}
		</CardContent>
	</Card>
</div>
