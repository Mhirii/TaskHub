<script lang="ts">
	import { Button } from "$lib/components/ui/button";
	import {
		Card,
		CardContent,
		CardDescription,
		CardFooter,
		CardHeader,
		CardTitle,
	} from "$lib/components/ui/card";
	import {
		Popover,
		PopoverContent,
		PopoverTrigger,
	} from "$lib/components/ui/popover";
	import { Calendar, Flag, Pen, User } from "@lucide/svelte";
	import type { Task } from "$lib/types";
	import {
		Tooltip,
		TooltipContent,
		TooltipTrigger,
	} from "$lib/components/ui/tooltip";
	import { cn } from "$lib/utils";

	export let task: Task;
</script>

<Card class="mb-2">
	<CardHeader class="px-2 flex flex-row items-center justify-between w-full">
		<div class="flex flex-col items-center gap-2 w-full justify-start">
			<CardTitle>
				<p class="font-medium">{task.name}</p>
			</CardTitle>
		</div>
		<div class="flex items-center gap-2 w-full justify-end">
			<Popover>
				<PopoverTrigger>
					<Button variant="outline" size="sm" class="h-8 w-8">
						<Pen />
					</Button>
				</PopoverTrigger>
				<PopoverContent>
					<p>Edit task</p>
				</PopoverContent>
			</Popover>
		</div>
	</CardHeader>
	<CardFooter
		class="px-2 flex flex-row items-center justif-between w-full gap-2"
	>
		<Tooltip>
			<TooltipTrigger>
				<Button variant="ghost" size="sm" class="h-8 w-8">
					<Calendar
						class={cn(
							"h-4 w-4",
							task.due_date ? "text-foreground" : "text-muted-foreground",
						)}
					/>
				</Button>
			</TooltipTrigger>
			<TooltipContent>
				{#if task.due_date}
					<p>Due date: {task.due_date}</p>
				{:else}
					<p>No due date</p>
				{/if}
			</TooltipContent>
		</Tooltip>
		<Tooltip>
			<TooltipTrigger>
				<Button variant="ghost" size="sm" class="h-8 w-8">
					<Flag
						class={cn(
							"h-4 w-4",
							task.priority ? "text-foreground" : "text-muted-foreground",
						)}
					/>
				</Button>
			</TooltipTrigger>
			<TooltipContent>
				{#if task.priority}
					<p>Priority: {task.priority}</p>
				{:else}
					<p>No priority</p>
				{/if}
			</TooltipContent>
		</Tooltip>
		<Tooltip>
			<TooltipTrigger>
				<Button variant="ghost" size="sm" class="h-8 w-8">
					<User
						class={cn(
							"h-4 w-4",
							task.assignee_id ? "text-foreground" : "text-muted-foreground",
						)}
					/>
				</Button>
			</TooltipTrigger>
			<TooltipContent>
				{#if task.assignee_id}
					<p>Assigned to: {task.assignee_id}</p>
				{:else}
					<p>No assignee</p>
				{/if}
			</TooltipContent>
		</Tooltip>
	</CardFooter>
</Card>
